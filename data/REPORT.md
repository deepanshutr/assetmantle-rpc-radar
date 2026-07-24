# AssetMantle endpoint radar

_Generated 2026-07-24T08:05:21Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23486852 | no | ✓ 51d (WE1) | 137/365 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 44d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23486852 | no | ✓ 85d (WE1) | 109/361 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23486852 | no | ✓ 78d (WE1) | 130/377 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23486852 | no | ✓ 52d (YR2) | 148/331 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23486852 | no | ✓ 51d (WE1) | 119/325 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 44d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23486852 | no | ✓ 85d (WE1) | 115/398 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23486852 | no | ✓ 78d (WE1) | 132/337 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23486852 | no | ✓ 52d (YR2) | 154/340 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 44d (WE1) | 90/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 52d (YR2) | 107/0 ms | canonical foundation endpoint |

