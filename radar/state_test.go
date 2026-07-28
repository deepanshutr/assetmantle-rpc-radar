package radar

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const liveStatePath = "../data/state.json"

// The committed state file is the real baseline the workflow compares against.
// Guard the round-trip so a schema drift cannot silently reduce every tick to
// "no baseline" (which would look healthy while alerting nothing at all).
func TestReadStateParsesTheCommittedStateFile(t *testing.T) {
	if _, err := os.Stat(liveStatePath); os.IsNotExist(err) {
		t.Skip("no committed state file")
	}
	got, err := ReadState(liveStatePath)
	if err != nil {
		t.Fatalf("ReadState: %v", err)
	}
	if len(got) == 0 {
		t.Fatal("committed state parsed to zero results — schema drift")
	}
	var owned int
	for _, r := range got {
		if r.URL == "" || r.Kind == "" {
			t.Fatalf("result missing identity fields: %+v", r)
		}
		if r.Owned() {
			owned++
		}
	}
	if owned == 0 {
		t.Fatal("no foundation-owned endpoints found — Owned() would silence everything")
	}
}

// THE REPORTED BUG, on production data: the same five third-party endpoints
// have been down for 12+ consecutive days and paged every morning. Replaying
// the committed state against itself must produce complete silence.
func TestCommittedStateAgainstItselfIsSilent(t *testing.T) {
	if _, err := os.Stat(liveStatePath); os.IsNotExist(err) {
		t.Skip("no committed state file")
	}
	state, err := ReadState(liveStatePath)
	if err != nil {
		t.Fatalf("ReadState: %v", err)
	}
	if alerts := CriticalAlerts(state, state); len(alerts) != 0 {
		t.Fatalf("standing state must not alert, got %d: %+v", len(alerts), alerts)
	}
}

// The counterpart to the silence tests: silencing the noise must not silence
// the alarm. Flip a foundation endpoint in the real committed state and the
// page must still fire, with an actionable one-liner.
func TestFoundationOutageOnRealStateStillPages(t *testing.T) {
	if _, err := os.Stat(liveStatePath); os.IsNotExist(err) {
		t.Skip("no committed state file")
	}
	prev, err := ReadState(liveStatePath)
	if err != nil {
		t.Fatalf("ReadState: %v", err)
	}
	cur, err := ReadState(liveStatePath)
	if err != nil {
		t.Fatalf("ReadState: %v", err)
	}
	var flipped string
	for i := range cur {
		if cur[i].Owned() && cur[i].OK {
			cur[i].OK = false
			cur[i].Error = "http 502"
			flipped = cur[i].URL
			break
		}
	}
	if flipped == "" {
		t.Skip("no healthy foundation endpoint in committed state to flip")
	}

	alerts := CriticalAlerts(prev, cur)
	if len(alerts) == 0 {
		t.Fatal("a foundation endpoint going down must still page")
	}
	text := FormatAlerts(alerts, "https://github.com/x/y/actions/runs/1")
	if !strings.Contains(text, "ACTION") || !strings.Contains(text, "down") {
		t.Fatalf("page must be an actionable one-liner, got %q", text)
	}
	if !strings.Contains(text, hostOf(flipped)) {
		t.Fatalf("page must name the endpoint %q, got %q", flipped, text)
	}
}

func TestReadStateMissingFileIsSilentNotAnError(t *testing.T) {
	got, err := ReadState(filepath.Join(t.TempDir(), "absent.json"))
	if err != nil {
		t.Fatalf("missing state must not error, got %v", err)
	}
	if got != nil {
		t.Fatalf("missing state must yield nil, got %+v", got)
	}
}
