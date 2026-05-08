# AssetMantle endpoint radar

_Generated 2026-05-08T06:32:00Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22269447 | no | ✓ 69d (WE1) | 174/516 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22269446 | no | ✓ 62d (WE1) | 47/212 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22269447 | no | ✓ 44d (WE1) | 123/500 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22269447 | no | ✓ 37d (WE1) | 173/509 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 48d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22269447 | no | ✓ 86d (E8) | 199/536 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22269447 | no | ✓ 69d (WE1) | 163/473 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22269446 | no | ✓ 62d (WE1) | 86/282 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22269447 | no | ✓ 44d (WE1) | 136/431 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22269447 | no | ✓ 37d (WE1) | 169/350 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22269447 | no | ✓ 86d (E8) | 199/523 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 62d (WE1) | 92/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 86d (E7) | 332/0 ms | canonical foundation endpoint |

