# AssetMantle endpoint radar

_Generated 2026-05-13T06:56:51Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22348049 | no | ✓ 64d (WE1) | 183/449 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22348049 | no | ✓ 57d (WE1) | 213/558 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22348049 | no | ✓ 39d (WE1) | 186/628 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22348049 | no | ✓ 32d (WE1) | 193/536 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22348049 | no | ✓ 64d (R13) | 220/543 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22348049 | no | ✓ 64d (WE1) | 198/441 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22348049 | no | ✓ 57d (WE1) | 83/549 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22348049 | no | ✓ 39d (WE1) | 203/621 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22348049 | no | ✓ 32d (WE1) | 287/431 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22348049 | no | ✓ 64d (R13) | 206/617 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 57d (WE1) | 24/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 64d (R13) | 40/0 ms | canonical foundation endpoint |

