# AssetMantle endpoint radar

_Generated 2026-07-14T07:44:53Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23325452 | no | ✓ 61d (WE1) | 140/394 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 54d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23325452 | no | ✓ 36d (WE1) | 117/233 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23325452 | no | ✓ 88d (WE1) | 201/430 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23325452 | no | ✓ 62d (YR2) | 156/393 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23325452 | no | ✓ 61d (WE1) | 148/268 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 54d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23325452 | no | ✓ 36d (WE1) | 135/459 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23325452 | no | ✓ 88d (WE1) | 164/401 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23325452 | no | ✓ 62d (YR2) | 168/266 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 54d (WE1) | 26/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 62d (YR2) | 41/0 ms | canonical foundation endpoint |

