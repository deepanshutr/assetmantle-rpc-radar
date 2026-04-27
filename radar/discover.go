package radar

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Discover scrapes known sources for endpoints not already in k. Returns the
// number of new entries appended.
func Discover(k *Known) (int, error) {
	existing := map[string]bool{}
	for _, e := range k.Endpoints {
		existing[normalize(e.URL)] = true
	}

	collected, err := collectFromChainRegistry()
	if err != nil {
		return 0, fmt.Errorf("chain-registry: %w", err)
	}
	if extra, err := collectFromCosmosDirectory(); err == nil {
		collected = append(collected, extra...)
	}

	added := 0
	for _, ep := range collected {
		key := normalize(ep.URL)
		if existing[key] {
			continue
		}
		existing[key] = true
		k.Endpoints = append(k.Endpoints, ep)
		added++
	}
	return added, nil
}

func normalize(u string) string {
	u = strings.TrimRight(u, "/")
	return strings.ToLower(u)
}

// chain-registry shape (subset)
type chainRegistry struct {
	APIs struct {
		RPC  []chainRegistryEndpoint `json:"rpc"`
		REST []chainRegistryEndpoint `json:"rest"`
		GRPC []chainRegistryEndpoint `json:"grpc"`
	} `json:"apis"`
}
type chainRegistryEndpoint struct {
	Address  string `json:"address"`
	Provider string `json:"provider"`
}

func collectFromChainRegistry() ([]Endpoint, error) {
	const url = "https://raw.githubusercontent.com/cosmos/chain-registry/master/assetmantle/chain.json"
	body, err := getURL(url)
	if err != nil {
		return nil, err
	}
	var c chainRegistry
	if err := json.Unmarshal(body, &c); err != nil {
		return nil, err
	}
	var out []Endpoint
	for _, e := range c.APIs.RPC {
		out = append(out, Endpoint{Kind: "rpc", URL: e.Address, Source: "chain-registry:" + e.Provider})
	}
	for _, e := range c.APIs.REST {
		out = append(out, Endpoint{Kind: "rest", URL: e.Address, Source: "chain-registry:" + e.Provider})
	}
	for _, e := range c.APIs.GRPC {
		out = append(out, Endpoint{Kind: "grpc", URL: e.Address, Source: "chain-registry:" + e.Provider})
	}
	return out, nil
}

// cosmos.directory aggregates registry data.
type cosmosDirChain struct {
	APIs struct {
		RPC  []chainRegistryEndpoint `json:"rpc"`
		REST []chainRegistryEndpoint `json:"rest"`
		GRPC []chainRegistryEndpoint `json:"grpc"`
	} `json:"apis"`
}

func collectFromCosmosDirectory() ([]Endpoint, error) {
	const url = "https://chains.cosmos.directory/assetmantle"
	body, err := getURL(url)
	if err != nil {
		return nil, err
	}
	var w struct {
		Chain cosmosDirChain `json:"chain"`
	}
	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}
	var out []Endpoint
	for _, e := range w.Chain.APIs.RPC {
		out = append(out, Endpoint{Kind: "rpc", URL: e.Address, Source: "cosmos-directory:" + e.Provider})
	}
	for _, e := range w.Chain.APIs.REST {
		out = append(out, Endpoint{Kind: "rest", URL: e.Address, Source: "cosmos-directory:" + e.Provider})
	}
	for _, e := range w.Chain.APIs.GRPC {
		out = append(out, Endpoint{Kind: "grpc", URL: e.Address, Source: "cosmos-directory:" + e.Provider})
	}
	return out, nil
}

func getURL(u string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
