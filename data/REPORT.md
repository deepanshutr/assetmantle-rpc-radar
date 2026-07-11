# AssetMantle endpoint radar

_Generated 2026-07-11T07:38:16Z — automated, do not edit._

**8/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23277023 | no | ✓ 64d (WE1) | 133/313 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 57d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23277023 | no | ✓ 39d (WE1) | 107/359 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✗ | — | — | ✓ 32d (WE1) | — | _err: http 500_ |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23277023 | no | ✓ 65d (YR2) | 157/309 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23277023 | no | ✓ 64d (WE1) | 122/331 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 57d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23277023 | no | ✓ 39d (WE1) | 115/162 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✗ | — | — | ✓ 32d (WE1) | — | _err: http 500_ |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23277023 | no | ✓ 65d (YR2) | 142/528 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp: lookup assetmantle-grpc.polkachu.com: i…_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 57d (WE1) | 22/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 65d (YR2) | 87/0 ms | canonical foundation endpoint |

