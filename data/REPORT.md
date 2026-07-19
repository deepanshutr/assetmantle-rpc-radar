# AssetMantle endpoint radar

_Generated 2026-07-19T07:56:29Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23406162 | no | ✓ 56d (WE1) | 200/495 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 49d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23406162 | no | ✓ 31d (WE1) | 189/620 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23406162 | no | ✓ 83d (WE1) | 237/618 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23406162 | no | ✓ 57d (YR2) | 229/521 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23406161 | no | ✓ 56d (WE1) | 166/386 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 49d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23406162 | no | ✓ 31d (WE1) | 202/602 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23406162 | no | ✓ 83d (WE1) | 236/529 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23406162 | no | ✓ 57d (YR2) | 262/461 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 49d (WE1) | 53/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 57d (YR2) | 54/0 ms | canonical foundation endpoint |

