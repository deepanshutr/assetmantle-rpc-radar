# AssetMantle endpoint radar

_Generated 2026-07-12T07:57:10Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23293356 | no | ✓ 63d (WE1) | 123/350 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 56d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23293356 | no | ✓ 38d (WE1) | 107/415 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23293356 | no | ✓ 31d (WE1) | 140/417 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23293356 | no | ✓ 64d (YR2) | 174/393 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23293356 | no | ✓ 63d (WE1) | 141/248 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 56d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23293356 | no | ✓ 38d (WE1) | 118/363 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23293356 | no | ✓ 31d (WE1) | 143/337 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23293357 | no | ✓ 64d (YR2) | 167/1792 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 56d (WE1) | 31/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 64d (YR2) | 42/0 ms | canonical foundation endpoint |

