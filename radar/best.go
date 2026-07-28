package radar

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// expectedNetwork is the chain these endpoints must actually be serving. A node
// that answers happily for a different chain-id is worse than one that is
// simply down: a consumer would read plausible-looking wrong data and never
// notice.
const expectedNetwork = "mantle-1"

// maxLagBlocks is how far behind the median a node may sit and still be
// recommended. A node that answers fast but serves stale state is the worst
// possible endpoint to hand a workflow.
const maxLagBlocks = 30

// BestEndpoints ranks the endpoints worth pointing a consumer at, per kind:
// healthy, on the right chain, current with the median height, ordered by p50
// latency.
//
// The result is a fallback chain, not a single pick. Consumers walk it in
// order, so one provider having a bad minute costs a retry rather than a run.
func BestEndpoints(results []Result) map[string][]string {
	median := medianBlockHeight(results)

	byKind := map[string][]Result{}
	for _, r := range results {
		if !r.OK {
			continue
		}
		if r.Kind == "grpc" {
			// The grpc probe is a bare TLS dial, so it can read neither
			// chain-id nor height. This list carries REACHABILITY ONLY and
			// none of the guarantees below — see the README caveat.
			byKind[r.Kind] = append(byKind[r.Kind], r)
			continue
		}
		// Require PROOF that the endpoint serves this chain and is current,
		// rather than giving unread fields the benefit of the doubt.
		//
		// A 200 carrying an unparseable body — a gateway error page, or the
		// JSON-RPC error envelope Tendermint returns with HTTP 200 — unmarshals
		// cleanly and leaves Network empty and BlockHeight zero. The old
		// "empty means grpc, so allow it" clause admitted exactly that, and the
		// height check self-disabled on BlockHeight == 0. Error pages are also
		// FAST, so the bogus entry sorted to index 0 and became the pick.
		if r.Network != expectedNetwork || r.BlockHeight <= 0 {
			continue
		}
		if median > 0 && median-r.BlockHeight > maxLagBlocks {
			continue
		}
		byKind[r.Kind] = append(byKind[r.Kind], r)
	}

	out := map[string][]string{}
	for kind, rs := range byKind {
		sort.SliceStable(rs, func(i, j int) bool {
			return rs[i].LatencyP50Ms < rs[j].LatencyP50Ms
		})
		urls := make([]string, 0, len(rs))
		for _, r := range rs {
			urls = append(urls, r.URL)
		}
		out[kind] = urls
	}
	return out
}

// WriteBest publishes the ranked recommendation consumers read instead of
// pinning endpoints of their own. Committed alongside state.json and REPORT.md,
// so it is versioned, diffable, and fetchable as a raw file with no service to
// operate.
func WriteBest(path string, results []Result) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	wrap := struct {
		GeneratedAt time.Time           `json:"generated_at"`
		Network     string              `json:"network"`
		Endpoints   map[string][]string `json:"endpoints"`
	}{time.Now().UTC(), expectedNetwork, BestEndpoints(results)}
	b, err := json.MarshalIndent(wrap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}
