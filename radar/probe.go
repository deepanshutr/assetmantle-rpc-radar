package radar

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// Result is the snapshot of a single probe attempt.
type Result struct {
	Endpoint
	ProbedAt        time.Time `json:"probed_at"`
	OK              bool      `json:"ok"`
	Error           string    `json:"error,omitempty"`
	BlockHeight     int64     `json:"block_height,omitempty"`
	CatchingUp      bool      `json:"catching_up,omitempty"`
	Network         string    `json:"network,omitempty"`
	NodeMoniker     string    `json:"node_moniker,omitempty"`
	Archive         bool      `json:"archive"`
	LatencyP50Ms    int64     `json:"latency_p50_ms,omitempty"`
	LatencyP95Ms    int64     `json:"latency_p95_ms,omitempty"`
	TLSIssuer       string    `json:"tls_issuer,omitempty"`
	TLSDaysToExpiry int       `json:"tls_days_to_expiry,omitempty"`
	TLSTrusted      bool      `json:"tls_trusted"`
}

// ProbeAll probes endpoints concurrently with a small worker pool.
func ProbeAll(eps []Endpoint) []Result {
	results := make([]Result, len(eps))
	const workers = 8
	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup
	for i, ep := range eps {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, ep Endpoint) {
			defer wg.Done()
			defer func() { <-sem }()
			results[i] = ProbeOne(ep)
		}(i, ep)
	}
	wg.Wait()
	return results
}

// ProbeOne dispatches based on kind.
func ProbeOne(ep Endpoint) Result {
	r := Result{Endpoint: ep, ProbedAt: time.Now().UTC()}
	switch ep.Kind {
	case "rpc":
		probeRPC(&r)
	case "rest":
		probeREST(&r)
	case "grpc":
		probeGRPC(&r)
	default:
		r.Error = "unknown kind: " + ep.Kind
	}
	if strings.HasPrefix(ep.URL, "https://") {
		probeTLS(&r, ep.URL)
	} else if ep.Kind == "grpc" {
		probeTLS(&r, "https://"+ep.URL)
	}
	return r
}

func probeRPC(r *Result) {
	statusURL := strings.TrimRight(r.URL, "/") + "/status"
	body, lats, err := getJSONWithLatency(statusURL, 5)
	if err != nil {
		r.Error = err.Error()
		return
	}
	var s struct {
		Result struct {
			NodeInfo struct {
				Network string `json:"network"`
				Moniker string `json:"moniker"`
			} `json:"node_info"`
			SyncInfo struct {
				LatestBlockHeight   string `json:"latest_block_height"`
				EarliestBlockHeight string `json:"earliest_block_height"`
				CatchingUp          bool   `json:"catching_up"`
			} `json:"sync_info"`
		} `json:"result"`
	}
	if err := json.Unmarshal(body, &s); err != nil {
		r.Error = "parse: " + err.Error()
		return
	}
	r.OK = true
	r.Network = s.Result.NodeInfo.Network
	r.NodeMoniker = s.Result.NodeInfo.Moniker
	r.CatchingUp = s.Result.SyncInfo.CatchingUp
	fmt.Sscanf(s.Result.SyncInfo.LatestBlockHeight, "%d", &r.BlockHeight)
	r.LatencyP50Ms, r.LatencyP95Ms = pctiles(lats)

	// archive check: /block?height=1
	blockURL := strings.TrimRight(r.URL, "/") + "/block?height=1"
	if body, _, err := getJSONWithLatency(blockURL, 1); err == nil {
		r.Archive = !strings.Contains(string(body), "is not available") &&
			!strings.Contains(string(body), "lowest height")
	}
}

func probeREST(r *Result) {
	statusURL := strings.TrimRight(r.URL, "/") + "/cosmos/base/tendermint/v1beta1/blocks/latest"
	body, lats, err := getJSONWithLatency(statusURL, 5)
	if err != nil {
		r.Error = err.Error()
		return
	}
	var s struct {
		Block struct {
			Header struct {
				Height  string `json:"height"`
				ChainID string `json:"chain_id"`
			} `json:"header"`
		} `json:"block"`
	}
	if err := json.Unmarshal(body, &s); err != nil {
		r.Error = "parse: " + err.Error()
		return
	}
	r.OK = true
	r.Network = s.Block.Header.ChainID
	fmt.Sscanf(s.Block.Header.Height, "%d", &r.BlockHeight)
	r.LatencyP50Ms, r.LatencyP95Ms = pctiles(lats)

	// archive check: blocks/1
	blockURL := strings.TrimRight(r.URL, "/") + "/cosmos/base/tendermint/v1beta1/blocks/1"
	if body, _, err := getJSONWithLatency(blockURL, 1); err == nil {
		r.Archive = !strings.Contains(string(body), "not available") &&
			!strings.Contains(string(body), "code\":3")
	}
}

func probeGRPC(r *Result) {
	host := r.URL
	if u, err := url.Parse("https://" + host); err == nil {
		host = u.Host
	}
	t0 := time.Now()
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", host, &tls.Config{NextProtos: []string{"h2"}})
	if err != nil {
		r.Error = "tls dial: " + err.Error()
		return
	}
	defer conn.Close()
	r.OK = true
	r.LatencyP50Ms = time.Since(t0).Milliseconds()
}

func probeTLS(r *Result, fullURL string) {
	u, err := url.Parse(fullURL)
	if err != nil || u.Host == "" {
		return
	}
	host := u.Host
	if !strings.Contains(host, ":") {
		host += ":443"
	}
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", host, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return
	}
	defer conn.Close()
	chain := conn.ConnectionState().PeerCertificates
	if len(chain) == 0 {
		return
	}
	leaf := chain[0]
	r.TLSIssuer = leaf.Issuer.CommonName
	r.TLSDaysToExpiry = int(time.Until(leaf.NotAfter).Hours() / 24)

	// trust check: verify against system roots
	roots, _ := x509.SystemCertPool()
	intermediates := x509.NewCertPool()
	for _, c := range chain[1:] {
		intermediates.AddCert(c)
	}
	hostname, _, _ := net.SplitHostPort(host)
	if hostname == "" {
		hostname = host
	}
	if _, err := leaf.Verify(x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
		DNSName:       hostname,
	}); err == nil {
		r.TLSTrusted = true
	}
}

func getJSONWithLatency(url string, samples int) ([]byte, []time.Duration, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	var body []byte
	lats := make([]time.Duration, 0, samples)
	for i := 0; i < samples; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		t0 := time.Now()
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			cancel()
			return nil, lats, err
		}
		resp, err := client.Do(req)
		if err != nil {
			cancel()
			return nil, lats, err
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		cancel()
		if err != nil {
			return nil, lats, err
		}
		if resp.StatusCode != 200 {
			return b, lats, fmt.Errorf("http %d", resp.StatusCode)
		}
		body = b
		lats = append(lats, time.Since(t0))
	}
	return body, lats, nil
}

func pctiles(lats []time.Duration) (p50, p95 int64) {
	if len(lats) == 0 {
		return 0, 0
	}
	sorted := append([]time.Duration{}, lats...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	p50i := len(sorted) / 2
	p95i := len(sorted) * 95 / 100
	if p95i >= len(sorted) {
		p95i = len(sorted) - 1
	}
	return sorted[p50i].Milliseconds(), sorted[p95i].Milliseconds()
}
