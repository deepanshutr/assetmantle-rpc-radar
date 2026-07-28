package radar

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// Discovery ran 13 times without landing, and its output contained
// `https://assetmantle-rpc.polkachu.com:443` alongside the already-tracked
// `https://assetmantle-rpc.polkachu.com`. Those are the same endpoint. Once
// best.json feeds real workflows, a duplicated host silently doubles that
// provider's weight in the ranking.
func TestNormalizeCollapsesDefaultPorts(t *testing.T) {
	same := [][2]string{
		{"https://a.example.com", "https://a.example.com:443"},
		{"https://a.example.com/", "https://a.example.com:443/"},
		{"http://b.example.com", "http://b.example.com:80"},
		{"https://c.example.com", "HTTPS://C.EXAMPLE.COM:443"},
	}
	for _, p := range same {
		if normalize(p[0]) != normalize(p[1]) {
			t.Errorf("%q and %q must normalize alike, got %q vs %q",
				p[0], p[1], normalize(p[0]), normalize(p[1]))
		}
	}
}

// gRPC entries are bare host:port with no scheme, and there :443 is a real
// coordinate, not a default to be dropped. polkachu serves gRPC on 14690;
// collapsing ports here would merge a working endpoint with a dead one.
func TestNormalizeKeepsPortsOnSchemelessHosts(t *testing.T) {
	a := "assetmantle-grpc.polkachu.com:443"
	b := "assetmantle-grpc.polkachu.com:14690"
	if normalize(a) == normalize(b) {
		t.Fatalf("schemeless ports are distinct endpoints, %q == %q", a, b)
	}
	if normalize(a) != "assetmantle-grpc.polkachu.com:443" {
		t.Fatalf("schemeless host must survive normalization, got %q", normalize(a))
	}
}

func bestFixture() []Result {
	mk := func(kind, url, source string, ok bool, h int64, p50 int64) Result {
		return Result{
			Endpoint:     Endpoint{Kind: kind, URL: url, Source: source},
			OK:           ok,
			BlockHeight:  h,
			Network:      "mantle-1",
			LatencyP50Ms: p50,
		}
	}
	return []Result{
		mk("rpc", "https://slow.example", "x", true, 100, 900),
		mk("rpc", "https://fast.example", "y", true, 100, 120),
		mk("rpc", "https://dead.example", "z", false, 0, 0),
		mk("rpc", "https://lagging.example", "w", true, 40, 10),
		mk("rest", "https://rest-ok.example", "y", true, 100, 200),
	}
}

func TestBestExcludesUnhealthyEndpoints(t *testing.T) {
	best := BestEndpoints(bestFixture())
	for _, u := range best["rpc"] {
		if u == "https://dead.example" {
			t.Fatal("an unhealthy endpoint must never be recommended")
		}
	}
}

func TestBestExcludesLaggingEndpointsEvenIfFast(t *testing.T) {
	// A node 60 blocks behind answers quickly and returns stale data — the
	// worst kind of endpoint to hand a workflow.
	best := BestEndpoints(bestFixture())
	for _, u := range best["rpc"] {
		if u == "https://lagging.example" {
			t.Fatal("a lagging endpoint must not be recommended despite low latency")
		}
	}
}

func TestBestRanksCurrentEndpointsByLatency(t *testing.T) {
	best := BestEndpoints(bestFixture())
	if len(best["rpc"]) < 2 {
		t.Fatalf("want at least 2 healthy rpc endpoints, got %v", best["rpc"])
	}
	if best["rpc"][0] != "https://fast.example" {
		t.Fatalf("fastest current endpoint must rank first, got %v", best["rpc"])
	}
}

func TestBestGroupsByKind(t *testing.T) {
	best := BestEndpoints(bestFixture())
	if len(best["rest"]) != 1 || best["rest"][0] != "https://rest-ok.example" {
		t.Fatalf("rest bucket wrong: %v", best["rest"])
	}
	if _, ok := best["grpc"]; ok && len(best["grpc"]) != 0 {
		t.Fatalf("grpc bucket should be absent or empty, got %v", best["grpc"])
	}
}

func TestBestRejectsWrongNetwork(t *testing.T) {
	rs := bestFixture()
	rs[1].Network = "some-testnet-1" // the otherwise-best rpc
	best := BestEndpoints(rs)
	for _, u := range best["rpc"] {
		if u == "https://fast.example" {
			t.Fatal("an endpoint serving a different chain must never be recommended")
		}
	}
}

func TestWriteBestProducesStableConsumableJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "best.json")
	if err := WriteBest(path, bestFixture()); err != nil {
		t.Fatalf("WriteBest: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var got struct {
		GeneratedAt string              `json:"generated_at"`
		Endpoints   map[string][]string `json:"endpoints"`
	}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("best.json must be valid JSON: %v", err)
	}
	if got.GeneratedAt == "" {
		t.Fatal("best.json needs a generated_at so consumers can spot a stale file")
	}
	if len(got.Endpoints["rpc"]) == 0 {
		t.Fatalf("rpc list empty: %+v", got.Endpoints)
	}
}

// The real registry must yield a usable recommendation for every protocol, or
// L3 consumers silently fall back to their pinned defaults forever.
func TestCommittedStateYieldsARecommendationPerKind(t *testing.T) {
	if _, err := os.Stat(liveStatePath); os.IsNotExist(err) {
		t.Skip("no committed state file")
	}
	state, err := ReadState(liveStatePath)
	if err != nil {
		t.Fatalf("ReadState: %v", err)
	}
	best := BestEndpoints(state)
	for _, kind := range []string{"rpc", "rest"} {
		if len(best[kind]) == 0 {
			t.Errorf("no healthy %s endpoint recommended from the committed state", kind)
		}
	}
}
