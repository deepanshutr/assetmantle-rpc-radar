# AssetMantle endpoint radar

_Generated 2026-05-18T07:15:19Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22426305 | no | ✓ 59d (WE1) | 186/417 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22426305 | no | ✓ 52d (WE1) | 120/534 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22426305 | no | ✓ 34d (WE1) | 174/566 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22426305 | no | ✓ 86d (WE1) | 211/343 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22426305 | no | ✓ 59d (R13) | 227/534 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22426305 | no | ✓ 59d (WE1) | 195/435 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22426305 | no | ✓ 52d (WE1) | 86/532 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22426305 | no | ✓ 34d (WE1) | 213/602 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22426305 | no | ✓ 86d (WE1) | 190/452 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22426305 | no | ✓ 59d (R13) | 230/463 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 52d (WE1) | 86/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 59d (R13) | 24/0 ms | canonical foundation endpoint |

