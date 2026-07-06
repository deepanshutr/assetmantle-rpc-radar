# AssetMantle endpoint radar

_Generated 2026-07-06T09:53:19Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23198528 | no | ✓ 69d (WE1) | 121/272 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 62d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23198528 | no | ✓ 44d (WE1) | 104/405 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23198528 | no | ✓ 37d (WE1) | 128/278 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23198528 | no | ✓ 70d (YR2) | 158/354 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23198528 | no | ✓ 69d (WE1) | 134/357 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 62d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23198528 | no | ✓ 44d (WE1) | 117/369 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23198528 | no | ✓ 37d (WE1) | 127/216 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23198528 | no | ✓ 70d (YR2) | 136/244 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 62d (WE1) | 18/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 70d (YR2) | 20/0 ms | canonical foundation endpoint |

