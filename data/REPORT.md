# AssetMantle endpoint radar

_Generated 2026-05-09T06:39:23Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22285186 | no | ✓ 68d (WE1) | 180/513 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22285186 | no | ✓ 61d (WE1) | 94/258 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22285186 | no | ✓ 43d (WE1) | 126/510 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22285186 | no | ✓ 36d (WE1) | 156/657 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22285186 | no | ✓ 85d (E8) | 198/519 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22285186 | no | ✓ 68d (WE1) | 170/491 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22285186 | no | ✓ 61d (WE1) | 51/257 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22285186 | no | ✓ 43d (WE1) | 136/426 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22285186 | no | ✓ 36d (WE1) | 236/399 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22285186 | no | ✓ 85d (E8) | 199/524 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 61d (WE1) | 99/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 85d (E7) | 313/0 ms | canonical foundation endpoint |

