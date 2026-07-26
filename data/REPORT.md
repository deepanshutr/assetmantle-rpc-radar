# AssetMantle endpoint radar

_Generated 2026-07-26T08:06:05Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23519166 | no | ✓ 49d (WE1) | 203/511 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 42d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23519166 | no | ✓ 83d (WE1) | 175/577 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23519166 | no | ✓ 76d (WE1) | 239/655 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23519166 | no | ✓ 50d (YR2) | 225/661 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23519166 | no | ✓ 49d (WE1) | 210/491 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 42d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23519166 | no | ✓ 83d (WE1) | 200/610 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23519166 | no | ✓ 76d (WE1) | 241/506 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23519166 | no | ✓ 50d (YR2) | 262/736 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 42d (WE1) | 60/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 50d (YR2) | 82/0 ms | canonical foundation endpoint |

