package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

// Alert categorises one state CHANGE worth notifying the operator about.
//
// Most endpoints in known.yaml belong to third parties, and their mirrors stay
// broken for weeks at a time. Alerting on current state re-paged the operator
// every morning with the same list of other people's dead nodes, so alerts now
// fire only on a transition, and only for endpoints we can actually fix. The
// one exception is a whole-protocol blackout, which is actionable regardless of
// who owns the boxes.
type Alert struct {
	Endpoint string
	Kind     string // down | stuck | cert | recovered | blackout
	Detail   string
}

// certWarnDays is how close to TLS expiry we start caring.
const certWarnDays = 14

// alertTextMax keeps a page glanceable on a phone lock screen.
const alertTextMax = 600

// Owned reports whether we operate this endpoint, and can therefore fix it.
func (e Endpoint) Owned() bool { return e.Source == "foundation" }

// condition names the worst active problem for a result, "" when healthy.
func condition(r Result, medianHeight int64) string {
	if !r.OK {
		return "down"
	}
	if r.Kind != "grpc" && medianHeight > 0 && r.BlockHeight > 0 && medianHeight-r.BlockHeight > 100 {
		return "stuck"
	}
	if r.TLSDaysToExpiry > 0 && r.TLSDaysToExpiry < certWarnDays {
		return "cert"
	}
	return ""
}

func conditionDetail(r Result, cond string, medianHeight int64) string {
	switch cond {
	case "down":
		return r.Error
	case "stuck":
		return fmt.Sprintf("height %d, median %d, lag %d",
			r.BlockHeight, medianHeight, medianHeight-r.BlockHeight)
	case "cert":
		return fmt.Sprintf("expires in %dd (%s)", r.TLSDaysToExpiry, r.TLSIssuer)
	}
	return ""
}

// CriticalAlerts returns alerts for transitions between the previous probe and
// this one. An empty prev (first run, or a missing state file) is silent by
// design: bootstrap is not news, and an unreadable baseline must never page.
func CriticalAlerts(prev, cur []Result) []Alert {
	if len(prev) == 0 {
		return nil
	}
	prevByURL := make(map[string]Result, len(prev))
	for _, r := range prev {
		prevByURL[r.URL] = r
	}
	prevMedian, curMedian := medianBlockHeight(prev), medianBlockHeight(cur)

	var out []Alert
	for _, r := range cur {
		if !r.Owned() {
			continue
		}
		was := ""
		if p, seen := prevByURL[r.URL]; seen {
			was = condition(p, prevMedian)
		}
		now := condition(r, curMedian)
		switch {
		case now == was:
			// Unchanged since the last tick: already reported, stay silent.
		case now == "":
			out = append(out, Alert{Endpoint: r.URL, Kind: "recovered"})
		default:
			out = append(out, Alert{
				Endpoint: r.URL,
				Kind:     now,
				Detail:   conditionDetail(r, now, curMedian),
			})
		}
	}
	// Blackouts first: "the chain is unreachable over this protocol" is the most
	// severe line the system emits, and a tail-truncated message would drop it
	// precisely when everything is down.
	return append(blackoutAlerts(prev, cur), out...)
}

// blackoutAlerts fires when a protocol loses its last healthy endpoint. This is
// the only third-party condition worth paging on: it means the chain is
// unreachable over that protocol no matter whose node you point at.
func blackoutAlerts(prev, cur []Result) []Alert {
	healthyByKind := func(rs []Result) map[string]int {
		m := map[string]int{}
		for _, r := range rs {
			if _, ok := m[r.Kind]; !ok {
				m[r.Kind] = 0
			}
			if r.OK {
				m[r.Kind]++
			}
		}
		return m
	}
	before, after := healthyByKind(prev), healthyByKind(cur)

	kinds := make([]string, 0, len(after))
	for k := range after {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)

	var out []Alert
	for _, k := range kinds {
		if after[k] == 0 && before[k] > 0 {
			out = append(out, Alert{
				Endpoint: "all " + k + " endpoints",
				Kind:     "blackout",
				Detail:   fmt.Sprintf("0 healthy (was %d)", before[k]),
			})
		}
	}
	return out
}

func medianBlockHeight(results []Result) int64 {
	var hs []int64
	for _, r := range results {
		if r.OK && r.BlockHeight > 0 && r.Kind != "grpc" {
			hs = append(hs, r.BlockHeight)
		}
	}
	if len(hs) == 0 {
		return 0
	}
	sort.Slice(hs, func(i, j int) bool { return hs[i] < hs[j] })
	return hs[len(hs)/2]
}

// hostOf trims the scheme so a line stays readable inside the length cap.
func hostOf(u string) string {
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	return strings.TrimSuffix(u, "/")
}

// FormatAlerts renders the Telegram body: plain ASCII, one line per change,
// ACTION for what we must fix and AUTO for what fixed itself. The full endpoint
// table lives in data/REPORT.md and is deliberately never pasted here.
func FormatAlerts(alerts []Alert, runURL string) string {
	if len(alerts) == 0 {
		return ""
	}
	lines := make([]string, 0, len(alerts))
	for _, a := range alerts {
		tier := "ACTION"
		if a.Kind == "recovered" {
			tier = "AUTO"
		}
		line := fmt.Sprintf("[radar] %s: %s %s", tier, hostOf(a.Endpoint), a.Kind)
		if a.Detail != "" {
			line += " - " + a.Detail
		}
		lines = append(lines, line)
	}
	// The run link is what makes a page actionable, so it must survive
	// truncation. Reserve room for it and cut the alert lines instead — the
	// old order appended it last and dropped it exactly when the incident was
	// large enough to need it.
	suffix := ""
	if runURL != "" {
		suffix = "\n" + runURL
	}
	body := strings.Join(lines, "\n")
	if len(body)+len(suffix) > alertTextMax {
		keep := alertTextMax - len(suffix) - 4
		if keep < 0 {
			keep = 0
		}
		body = body[:keep] + "\n..."
	}
	return body + suffix
}

// SendAlerts pushes one grouped Telegram message via TELEGRAM_BOT_TOKEN +
// TELEGRAM_CHAT_ID env vars. No-op if either is unset (local runs).
func SendAlerts(alerts []Alert, runURL string) error {
	text := FormatAlerts(alerts, runURL)
	if text == "" {
		return nil
	}
	tok := os.Getenv("TELEGRAM_BOT_TOKEN")
	chat := os.Getenv("TELEGRAM_CHAT_ID")
	if tok == "" || chat == "" {
		// Locally this is expected and harmless. In CI it means a real page is
		// being dropped, and the caller MUST NOT advance the baseline over it —
		// returning nil here would consume the transition silently, which is
		// the exact failure this ordering exists to prevent.
		if os.Getenv("GITHUB_ACTIONS") == "true" {
			return fmt.Errorf("telegram credentials unset; %d alert(s) undelivered", len(alerts))
		}
		fmt.Fprintln(os.Stderr, "alerts: telegram creds unset, skipping")
		return nil
	}
	body := url.Values{}
	body.Set("chat_id", chat)
	body.Set("text", text)
	body.Set("disable_web_page_preview", "true")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	apiURL := "https://api.telegram.org/bot" + tok + "/sendMessage"
	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("telegram api: http %d", resp.StatusCode)
	}
	return nil
}
