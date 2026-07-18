# AssetMantle endpoint radar

_Generated 2026-07-18T07:25:39Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23389693 | no | ✓ 57d (WE1) | 251/315 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 50d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23389693 | no | ✓ 32d (WE1) | 137/531 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23389693 | no | ✓ 84d (WE1) | 244/404 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23389693 | no | ✓ 58d (YR2) | 182/461 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23389693 | no | ✓ 57d (WE1) | 142/429 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 50d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23389693 | no | ✓ 32d (WE1) | 148/525 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23389693 | no | ✓ 84d (WE1) | 170/395 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23389693 | no | ✓ 58d (YR2) | 192/311 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 50d (WE1) | 93/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 58d (YR2) | 162/0 ms | canonical foundation endpoint |

