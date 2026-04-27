package radar

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// WriteState writes the full probe results as JSON.
func WriteState(path string, results []Result) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	wrap := struct {
		GeneratedAt time.Time `json:"generated_at"`
		Count       int       `json:"count"`
		Results     []Result  `json:"results"`
	}{time.Now().UTC(), len(results), results}
	b, err := json.MarshalIndent(wrap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// WriteReport writes a per-kind markdown summary.
func WriteReport(path string, results []Result) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# AssetMantle endpoint radar\n\n")
	fmt.Fprintf(&sb, "_Generated %s — automated, do not edit._\n\n", time.Now().UTC().Format(time.RFC3339))
	ok, total := count(results)
	fmt.Fprintf(&sb, "**%d/%d endpoints healthy.**\n\n", ok, total)
	for _, kind := range []string{"rpc", "rest", "grpc"} {
		group := filterKind(results, kind)
		if len(group) == 0 {
			continue
		}
		fmt.Fprintf(&sb, "## %s (%d)\n\n", strings.ToUpper(kind), len(group))
		fmt.Fprintln(&sb, "| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |")
		fmt.Fprintln(&sb, "|---|---|---|---|---|---|---|")
		sort.Slice(group, func(i, j int) bool { return group[i].URL < group[j].URL })
		for _, r := range group {
			renderRow(&sb, r)
		}
		fmt.Fprintln(&sb)
	}
	return os.WriteFile(path, []byte(sb.String()), 0o644)
}

func renderRow(sb *strings.Builder, r Result) {
	status := "✓"
	if !r.OK {
		status = "✗"
	}
	height := "—"
	if r.BlockHeight > 0 {
		height = fmt.Sprintf("%d", r.BlockHeight)
	}
	archive := "—"
	if r.OK {
		if r.Archive {
			archive = "yes"
		} else {
			archive = "no"
		}
	}
	tls := "—"
	if r.TLSIssuer != "" {
		trust := "✓"
		if !r.TLSTrusted {
			trust = "✗"
		}
		tls = fmt.Sprintf("%s %dd (%s)", trust, r.TLSDaysToExpiry, truncate(r.TLSIssuer, 18))
	}
	lat := "—"
	if r.LatencyP50Ms > 0 {
		lat = fmt.Sprintf("%d/%d ms", r.LatencyP50Ms, r.LatencyP95Ms)
	}
	notes := r.Notes
	if r.Error != "" {
		notes = "_err: " + truncate(r.Error, 60) + "_"
	}
	fmt.Fprintf(sb, "| `%s` | %s | %s | %s | %s | %s | %s |\n",
		r.URL, status, height, archive, tls, lat, notes)
}

func count(rs []Result) (ok, total int) {
	for _, r := range rs {
		total++
		if r.OK {
			ok++
		}
	}
	return
}

func filterKind(rs []Result, kind string) []Result {
	var out []Result
	for _, r := range rs {
		if r.Kind == kind {
			out = append(out, r)
		}
	}
	return out
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
