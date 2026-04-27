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

## See also

- [`deepanshutr/assetmantle-infra`](https://github.com/deepanshutr/assetmantle-infra) — the production aggregator (Caddy on Akash) that consumes the endpoints this radar tracks.
