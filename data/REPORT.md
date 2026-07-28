# AssetMantle endpoint radar

_Generated 2026-07-28T08:12:44Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23551512 | no | ✓ 47d (WE1) | 191/423 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 40d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23551512 | no | ✓ 81d (WE1) | 174/573 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23551512 | no | ✓ 74d (WE1) | 207/505 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23551513 | no | ✓ 48d (YR2) | 222/1599 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23551512 | no | ✓ 47d (WE1) | 194/425 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 40d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23551512 | no | ✓ 81d (WE1) | 204/585 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23551512 | no | ✓ 74d (WE1) | 207/455 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23551513 | no | ✓ 48d (YR2) | 217/1398 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 40d (WE1) | 13/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 48d (YR2) | 26/0 ms | canonical foundation endpoint |

