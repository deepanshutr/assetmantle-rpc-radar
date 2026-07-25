# AssetMantle endpoint radar

_Generated 2026-07-25T07:48:47Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23502845 | no | ✓ 50d (WE1) | 121/279 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 43d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23502845 | no | ✓ 84d (WE1) | 105/349 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23502845 | no | ✓ 77d (WE1) | 135/332 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23502845 | no | ✓ 51d (YR2) | 153/367 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23502845 | no | ✓ 50d (WE1) | 135/188 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 43d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23502845 | no | ✓ 84d (WE1) | 115/404 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23502845 | no | ✓ 77d (WE1) | 135/325 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23502845 | no | ✓ 51d (YR2) | 157/190 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 43d (WE1) | 90/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 51d (YR2) | 28/0 ms | canonical foundation endpoint |

