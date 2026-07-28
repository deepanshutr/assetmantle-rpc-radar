package radar

import (
	"strings"
	"testing"
)

// Regression guards for the review's CRITICAL finding: a change-detector whose
// baseline advances independently of delivery converts one dropped page into
// permanent silence. These cover the delivery half; main.go's runHealth owns
// the ordering (send, then persist).

func TestSendAlertsFailsLoudlyInCIWhenCredentialsAreMissing(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	err := SendAlerts([]Alert{{Endpoint: "https://rpc.assetmantle.one", Kind: "down"}}, "")
	if err == nil {
		t.Fatal("in CI, undelivered alerts must return an error so the caller " +
			"does not advance the baseline over a dropped page")
	}
}

func TestSendAlertsStaysQuietLocallyWhenCredentialsAreMissing(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	if err := SendAlerts([]Alert{{Endpoint: "https://x", Kind: "down"}}, ""); err != nil {
		t.Fatalf("a local run without credentials must not error, got %v", err)
	}
}

func TestNoAlertsNeverAttemptsDelivery(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	if err := SendAlerts(nil, ""); err != nil {
		t.Fatalf("a silent tick must not error even in CI, got %v", err)
	}
}

// The run link is the only part of the page that makes it actionable.
func TestRunURLSurvivesTruncation(t *testing.T) {
	runURL := "https://github.com/deepanshutr/assetmantle-rpc-radar/actions/runs/123456789"
	var many []Alert
	for i := 0; i < 40; i++ {
		many = append(many, Alert{
			Endpoint: "https://a-very-long-endpoint-hostname-for-padding.example.com",
			Kind:     "down",
			Detail:   "dial tcp: i/o timeout after a rather verbose Go error string",
		})
	}
	got := FormatAlerts(many, runURL)

	if len(got) > alertTextMax {
		t.Fatalf("cap breached: %d chars", len(got))
	}
	if !strings.HasSuffix(got, runURL) {
		t.Fatalf("run URL must survive truncation; message ended:\n%q", got[max(0, len(got)-120):])
	}
}

// Worst-case input must not produce the least information.
func TestBlackoutSurvivesTruncationAndLeads(t *testing.T) {
	prev := liveState(true)
	cur := liveState(true)
	for i := range cur {
		cur[i].OK = false
		cur[i].Error = "dial tcp: i/o timeout after a rather verbose Go error string"
	}
	alerts := CriticalAlerts(prev, cur)
	if len(alerts) == 0 {
		t.Fatal("a total outage must alert")
	}
	if alerts[0].Kind != "blackout" {
		t.Fatalf("blackout must lead so truncation cannot drop it, got %q", alerts[0].Kind)
	}

	text := FormatAlerts(alerts, "https://github.com/x/y/actions/runs/1")
	if !strings.Contains(text, "blackout") {
		t.Fatalf("blackout line truncated away:\n%s", text)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
