package radar

import "testing"

// Regression guards from the adversarial review of PR #3.
//
// best.json is a recommendation other systems act on, so the admission rule
// must demand proof rather than extend benefit of the doubt. The dangerous
// shape is a 200 response carrying an unparseable body: it unmarshals cleanly,
// leaves Network empty and BlockHeight zero, and — because error pages are
// fast — sorts to the front of a latency-ranked list.

func probed(kind, url, network string, height, p50 int64) Result {
	return Result{
		Endpoint:     Endpoint{Kind: kind, URL: url, Source: "x"},
		OK:           true,
		Network:      network,
		BlockHeight:  height,
		LatencyP50Ms: p50,
	}
}

func TestFastErrorPageIsNeverRecommended(t *testing.T) {
	results := []Result{
		// A gateway error page: HTTP 200, valid JSON, no chain-id, no height,
		// and quicker than every real node.
		probed("rest", "https://gateway-error.example", "", 0, 12),
		probed("rest", "https://real-a.example", "mantle-1", 100, 210),
		probed("rest", "https://real-b.example", "mantle-1", 100, 340),
	}
	best := BestEndpoints(results)

	for _, u := range best["rest"] {
		if u == "https://gateway-error.example" {
			t.Fatalf("an endpoint that proved neither chain nor height must not "+
				"be recommended, got %v", best["rest"])
		}
	}
	if len(best["rest"]) == 0 || best["rest"][0] != "https://real-a.example" {
		t.Fatalf("fastest PROVEN endpoint must lead, got %v", best["rest"])
	}
}

func TestWrongChainCannotMoveTheMedian(t *testing.T) {
	// Three hosts serving the Ethereum L2 named "Mantle" alongside two real
	// mantle-1 nodes. If they counted, the median would jump to ~80M and every
	// real endpoint would be dropped as "lagging" while being paged as "stuck".
	results := []Result{
		probed("rpc", "https://real-a.example", "mantle-1", 23_549_342, 200),
		probed("rpc", "https://real-b.example", "mantle-1", 23_549_342, 300),
		probed("rpc", "https://l2-a.example", "mantle", 80_000_000, 50),
		probed("rpc", "https://l2-b.example", "mantle", 80_000_001, 60),
		probed("rpc", "https://l2-c.example", "mantle", 80_000_002, 70),
	}

	if got := medianBlockHeight(results); got != 23_549_342 {
		t.Fatalf("median must be computed over mantle-1 only, got %d", got)
	}
	best := BestEndpoints(results)
	if len(best["rpc"]) != 2 {
		t.Fatalf("both real endpoints must survive, got %v", best["rpc"])
	}
	for _, u := range best["rpc"] {
		if u == "https://l2-a.example" {
			t.Fatal("a wrong-chain endpoint must never be recommended")
		}
	}
}

func TestHealthyForeignChainNodeIsNotPagedAsStuck(t *testing.T) {
	// The alerting counterpart: a foreign-chain node in the set must not make a
	// healthy endpoint we own look stuck.
	prev := []Result{probed("rpc", "https://rpc.assetmantle.one", "mantle-1", 23_549_000, 200)}
	prev[0].Source = "foundation"
	cur := []Result{
		probed("rpc", "https://rpc.assetmantle.one", "mantle-1", 23_549_342, 200),
		probed("rpc", "https://l2.example", "mantle", 80_000_000, 50),
	}
	cur[0].Source = "foundation"

	if alerts := CriticalAlerts(prev, cur); len(alerts) != 0 {
		t.Fatalf("a foreign-chain node must not manufacture a stuck page, got %+v", alerts)
	}
}

func TestGrpcIsAdmittedAsReachabilityOnly(t *testing.T) {
	// Documented honestly rather than silently: the grpc probe is a bare TLS
	// dial, so these entries carry no chain or height guarantee.
	results := []Result{probed("grpc", "grpc.assetmantle.one:443", "", 0, 58)}
	if got := BestEndpoints(results)["grpc"]; len(got) != 1 {
		t.Fatalf("grpc must still be listed on reachability, got %v", got)
	}
}
