# AssetMantle endpoint radar

_Generated 2026-07-07T09:21:39Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23214358 | no | ✓ 68d (WE1) | 134/385 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 61d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23214358 | no | ✓ 43d (WE1) | 184/1271 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23214358 | no | ✓ 36d (WE1) | 192/469 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23214358 | no | ✓ 69d (YR2) | 188/1533 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23214358 | no | ✓ 68d (WE1) | 174/447 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 61d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23214358 | no | ✓ 43d (WE1) | 126/387 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23214358 | no | ✓ 36d (WE1) | 153/374 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23214358 | no | ✓ 69d (YR2) | 172/386 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 61d (WE1) | 22/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 69d (YR2) | 53/0 ms | canonical foundation endpoint |

