package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Alert categorises a single regression worth notifying about.
type Alert struct {
	Endpoint string
	Kind     string // down | stuck | cert
	Detail   string
}

// CriticalAlerts returns the subset of results worth paging the operator on.
func CriticalAlerts(results []Result) []Alert {
	var out []Alert
	medianHeight := medianBlockHeight(results)
	for _, r := range results {
		if !r.OK {
			out = append(out, Alert{
				Endpoint: r.URL,
				Kind:     "down",
				Detail:   r.Error,
			})
			continue
		}
		if r.Kind != "grpc" && medianHeight > 0 && r.BlockHeight > 0 && medianHeight-r.BlockHeight > 100 {
			out = append(out, Alert{
				Endpoint: r.URL,
				Kind:     "stuck",
				Detail:   fmt.Sprintf("height %d, median %d, lag %d", r.BlockHeight, medianHeight, medianHeight-r.BlockHeight),
			})
		}
		if r.TLSDaysToExpiry > 0 && r.TLSDaysToExpiry < 14 {
			out = append(out, Alert{
				Endpoint: r.URL,
				Kind:     "cert",
				Detail:   fmt.Sprintf("expires in %d days (issuer: %s)", r.TLSDaysToExpiry, r.TLSIssuer),
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
	for i := 1; i < len(hs); i++ {
		for j := i; j > 0 && hs[j-1] > hs[j]; j-- {
			hs[j-1], hs[j] = hs[j], hs[j-1]
		}
	}
	return hs[len(hs)/2]
}

// SendAlerts pushes a single grouped Telegram message via TELEGRAM_BOT_TOKEN +
// TELEGRAM_CHAT_ID env vars. No-op if either is unset (local runs).
func SendAlerts(alerts []Alert) error {
	tok := os.Getenv("TELEGRAM_BOT_TOKEN")
	chat := os.Getenv("TELEGRAM_CHAT_ID")
	if tok == "" || chat == "" {
		fmt.Fprintln(os.Stderr, "alerts: telegram creds unset, skipping")
		return nil
	}
	var sb strings.Builder
	sb.WriteString("🚨 <b>AssetMantle radar</b>\n")
	sb.WriteString(fmt.Sprintf("%d critical regression(s) detected:\n\n", len(alerts)))
	for _, a := range alerts {
		fmt.Fprintf(&sb, "• <b>%s</b> [%s]\n  <code>%s</code>\n", a.Endpoint, a.Kind, a.Detail)
	}
	body := url.Values{}
	body.Set("chat_id", chat)
	body.Set("text", sb.String())
	body.Set("parse_mode", "HTML")
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
