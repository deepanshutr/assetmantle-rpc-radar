# AssetMantle endpoint radar

_Generated 2026-07-13T08:45:23Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23310010 | no | ✓ 62d (WE1) | 136/332 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 55d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23310010 | no | ✓ 37d (WE1) | 116/456 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23310010 | no | ✓ 89d (WE1) | 159/349 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23310010 | no | ✓ 63d (YR2) | 181/259 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23310010 | no | ✓ 62d (WE1) | 150/412 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 55d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23310010 | no | ✓ 37d (WE1) | 129/475 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23310010 | no | ✓ 89d (WE1) | 157/340 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23310011 | no | ✓ 63d (YR2) | 386/402 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 55d (WE1) | 16/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 63d (YR2) | 92/0 ms | canonical foundation endpoint |

