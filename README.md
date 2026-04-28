# assetmantle-rpc-radar

Health + discovery watcher for AssetMantle public endpoints (RPC / REST / gRPC).

- **Daily** at 06:00 UTC: probe every known endpoint, commit `data/state.json` + regenerated `data/REPORT.md`.
- **Weekly** Mon 06:00 UTC: scrape discovery sources, append new endpoints to `known.yaml` if any.
- **Critical alerts** (host down, height stuck > 30 min, cert expiring < 14 days) are pushed to Telegram (`@x_anyhow_x_bot`) via secrets.

## Files

```
known.yaml          # tracked endpoints (kind, url, notes)
data/state.json     # last full probe snapshot (committed)
data/REPORT.md      # human-readable per-endpoint table
radar/              # Go package
  probe.go          # /status + /block?height=1 + TLS + latency
  discover.go       # chain-registry / cosmos.directory / polkachu peers
  report.go         # markdown renderer
  alert.go          # Telegram critical alerter
main.go             # CLI: `radar {health|discover|all}`
```

## Local dev

```
go run . health             # probe all known endpoints, regenerate state + report
go run . discover           # scrape sources, update known.yaml
TELEGRAM_BOT_TOKEN=… TELEGRAM_CHAT_ID=… go run . health   # with alerts
```

## Discovery sources

1. `cosmos/chain-registry` → `assetmantle/chain.json` `apis.{rpc,rest,grpc}[]`
2. `cosmos.directory/assetmantle` JSON
3. `polkachu.com/api/v2/chains/assetmantle/live_peers`

## URL convention quirks (verify before adding to `known.yaml`)

Discovery sources sometimes list URLs that look right but aren't. Verify each new entry actually responds before relying on it.

- **PublicNode REST is `*-rest.publicnode.com`, NOT `*-api.publicnode.com`.** The `-api` pattern (e.g. `assetmantle-api.publicnode.com`) returns 404 from Cloudflare for every path. Use `assetmantle-rest.publicnode.com` instead. This was a real bug — the original `known.yaml` had the wrong URL and the radar reported "publicnode REST down" indefinitely until the URL was corrected.
- **PublicNode RPC is `*-rpc.publicnode.com`** — the consistent pattern across RPC/REST/gRPC is `<chain>-{rpc,rest,grpc}.publicnode.com`.
- **Polkachu gRPC is on `:14690` plaintext h2c**, NOT `:443` TLS. The `:443` port is connection-refused. Probe with `openssl s_client -alpn h2 -connect host:port` before assuming `:443` works.
- **Notional URL pattern can flip.** `cosmos.directory` shows `https://api-assetmantle-ia.cosmosia.notional.ventures` for REST while older lists may have `https://rest-assetmantle-ia.cosmosia.notional.ventures` — the old one is wrong. Always cross-check against `cosmos.directory` for the canonical URL.
- **Stakewolle has URL paths** (e.g. `https://public.stakewolle.com/cosmos/assetmantle/rpc`). Caddy's `reverse_proxy` directive can use it as a probe target (radar) but cannot use it as a load-balanced upstream (the path component breaks the pool member URL parser). Track in radar; don't add to the aggregator pool.
- **Polkachu/Cloudflare quirk**: Polkachu sits behind Cloudflare; you'll see `server: cloudflare` and `cf-ray` headers on every response. Don't read those as "Polkachu is down" or "we're talking to Cloudflare directly" — it's CF-as-CDN in front of Polkachu's origin.

## See also

- [`deepanshutr/assetmantle-infra`](https://github.com/deepanshutr/assetmantle-infra) — the production aggregator (Caddy on Akash) that consumes the endpoints this radar tracks. Its `akash/README.md` has the full Caddy-gotcha catalog (CORS double-header, `tls_server_name` placeholder, `lb_policy first` + no health_uri = sticky broken primary).
