# AssetMantle endpoint radar

_Generated 2026-05-15T07:01:01Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22379438 | no | ✓ 62d (WE1) | 116/215 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22379438 | no | ✓ 55d (WE1) | 37/178 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22379438 | no | ✓ 37d (WE1) | 103/198 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22379438 | no | ✓ 89d (WE1) | 188/365 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22379438 | no | ✓ 62d (R13) | 153/210 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22379438 | no | ✓ 62d (WE1) | 135/297 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22379438 | no | ✓ 55d (WE1) | 39/344 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22379438 | no | ✓ 37d (WE1) | 136/451 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22379438 | no | ✓ 89d (WE1) | 129/281 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22379438 | no | ✓ 62d (R13) | 174/204 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 55d (WE1) | 43/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 62d (R13) | 20/0 ms | canonical foundation endpoint |

