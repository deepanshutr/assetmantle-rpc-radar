package radar

import (
	"strings"
	"testing"
)

// res builds a probe result for a given endpoint identity + health.
func res(kind, url, source string, ok bool) Result {
	return Result{
		Endpoint: Endpoint{Kind: kind, URL: url, Source: source},
		OK:       ok,
	}
}

// The live registry shape as of 2026-07-27: three foundation endpoints we
// operate, plus third-party mirrors we do not.
func liveState(foundationOK bool) []Result {
	return []Result{
		res("rpc", "https://rpc.assetmantle.one", "foundation", foundationOK),
		res("rpc", "https://assetmantle-rpc.polkachu.com", "polkachu", true),
		res("rpc", "https://assetmantle-rpc.publicnode.com", "publicnode", false),
		res("rpc", "https://rpc-assetmantle-ia.cosmosia.notional.ventures", "notional", false),
		res("rest", "https://rest.assetmantle.one", "foundation", foundationOK),
		res("rest", "https://assetmantle-api.polkachu.com", "polkachu", true),
		res("rest", "https://assetmantle-rest.publicnode.com", "publicnode", false),
		res("grpc", "grpc.assetmantle.one:443", "foundation", foundationOK),
		res("grpc", "assetmantle-grpc.publicnode.com:443", "publicnode", true),
		res("grpc", "assetmantle-grpc.polkachu.com:443", "polkachu", false),
	}
}

// THE REGRESSION THIS FIXES: five third-party endpoints have been down for
// 12+ consecutive days and re-paged the operator every single morning.
// An unchanged tick must be completely silent.
func TestUnchangedTickIsSilent(t *testing.T) {
	prev := liveState(true)
	cur := liveState(true)
	if got := CriticalAlerts(prev, cur); len(got) != 0 {
		t.Fatalf("unchanged tick must be silent, got %d alerts: %+v", len(got), got)
	}
}

func TestBootstrapIsSilent(t *testing.T) {
	// No previous state on disk (first run / fresh clone): never page.
	if got := CriticalAlerts(nil, liveState(false)); len(got) != 0 {
		t.Fatalf("bootstrap must be silent, got %d alerts: %+v", len(got), got)
	}
}

func TestThirdPartyBreakageNeverAlerts(t *testing.T) {
	prev := liveState(true)
	cur := liveState(true)
	// polkachu rpc (third-party) breaks. Not ours, not actionable.
	cur[1].OK = false
	if got := CriticalAlerts(prev, cur); len(got) != 0 {
		t.Fatalf("third-party breakage must not alert, got %+v", got)
	}
}

func TestOwnedEndpointDownAlertsOnce(t *testing.T) {
	prev := liveState(true)
	cur := liveState(true)
	cur[0].OK = false
	cur[0].Error = "http 502"

	got := CriticalAlerts(prev, cur)
	if len(got) != 1 {
		t.Fatalf("want exactly 1 alert, got %d: %+v", len(got), got)
	}
	if got[0].Endpoint != "https://rpc.assetmantle.one" || got[0].Kind != "down" {
		t.Fatalf("wrong alert: %+v", got[0])
	}
	// Second tick with the same condition must be silent.
	if again := CriticalAlerts(cur, cur); len(again) != 0 {
		t.Fatalf("repeat of same condition must be silent, got %+v", again)
	}
}

func TestOwnedEndpointRecoveryAlerts(t *testing.T) {
	prev := liveState(true)
	prev[0].OK = false
	cur := liveState(true)

	got := CriticalAlerts(prev, cur)
	if len(got) != 1 || got[0].Kind != "recovered" {
		t.Fatalf("want 1 recovered alert, got %+v", got)
	}
}

func TestOwnedCertExpiryAlertsOnlyOnCrossing(t *testing.T) {
	prev := liveState(true)
	prev[0].TLSDaysToExpiry = 20
	cur := liveState(true)
	cur[0].TLSDaysToExpiry = 13 // crosses the 14d threshold

	got := CriticalAlerts(prev, cur)
	if len(got) != 1 || got[0].Kind != "cert" {
		t.Fatalf("want 1 cert alert on crossing, got %+v", got)
	}
	// Still under threshold the next day -> silent, not a daily countdown.
	next := liveState(true)
	next[0].TLSDaysToExpiry = 12
	if again := CriticalAlerts(cur, next); len(again) != 0 {
		t.Fatalf("cert must not re-alert daily, got %+v", again)
	}
}

func TestKindBlackoutAlertsWhenLastHealthyEndpointDies(t *testing.T) {
	prev := liveState(true)
	cur := liveState(true)
	// Kill every rpc endpoint: the chain becomes unreachable over rpc.
	for i := range cur {
		if cur[i].Kind == "rpc" {
			cur[i].OK = false
		}
	}
	got := CriticalAlerts(prev, cur)

	var blackout *Alert
	for i := range got {
		if got[i].Kind == "blackout" {
			blackout = &got[i]
		}
	}
	if blackout == nil {
		t.Fatalf("want a blackout alert when all rpc endpoints die, got %+v", got)
	}
	if !strings.Contains(blackout.Endpoint, "rpc") {
		t.Fatalf("blackout should name the kind, got %q", blackout.Endpoint)
	}
	// Still blacked out next tick -> silent.
	if again := CriticalAlerts(cur, cur); len(again) != 0 {
		t.Fatalf("sustained blackout must not re-alert, got %+v", again)
	}
}

func TestFormatIsPlainAsciiActionableAndCapped(t *testing.T) {
	alerts := []Alert{
		{Endpoint: "https://rpc.assetmantle.one", Kind: "down", Detail: "http 502"},
		{Endpoint: "https://rest.assetmantle.one", Kind: "recovered", Detail: ""},
	}
	got := FormatAlerts(alerts, "https://github.com/x/y/actions/runs/1")

	if strings.ContainsRune(got, '\U0001F6A8') || strings.ContainsRune(got, '\U0001F4B0') {
		t.Fatalf("message must be emoji-free, got %q", got)
	}
	for _, r := range got {
		if r > 127 {
			t.Fatalf("message must be plain ASCII, found %q in %q", r, got)
		}
	}
	if !strings.Contains(got, "[radar]") {
		t.Fatalf("message needs the [radar] prefix, got %q", got)
	}
	if !strings.Contains(got, "ACTION") {
		t.Fatalf("a down endpoint we own is an ACTION, got %q", got)
	}
	if !strings.Contains(got, "AUTO") {
		t.Fatalf("a recovery is an AUTO line, got %q", got)
	}
	if len(got) > 600 {
		t.Fatalf("message must be capped at 600 chars, got %d", len(got))
	}
}

func TestFormatEmptyAlertsIsEmpty(t *testing.T) {
	if got := FormatAlerts(nil, "https://run"); got != "" {
		t.Fatalf("no alerts must render nothing, got %q", got)
	}
}

func TestOwnedReportsOperatorship(t *testing.T) {
	if !(Endpoint{Source: "foundation"}).Owned() {
		t.Fatal("foundation endpoints are ours")
	}
	if (Endpoint{Source: "polkachu"}).Owned() {
		t.Fatal("polkachu is not ours")
	}
}
