# AssetMantle endpoint radar

_Generated 2026-07-20T08:36:53Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23422740 | no | ✓ 55d (WE1) | 134/324 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 48d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23422740 | no | ✓ 89d (WE1) | 103/402 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23422740 | no | ✓ 82d (WE1) | 161/359 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23422740 | no | ✓ 56d (YR2) | 153/355 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23422740 | no | ✓ 55d (WE1) | 133/359 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 48d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23422740 | no | ✓ 89d (WE1) | 116/327 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23422740 | no | ✓ 82d (WE1) | 174/374 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23422741 | no | ✓ 56d (YR2) | 180/1671 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 48d (WE1) | 93/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 56d (YR2) | 23/0 ms | canonical foundation endpoint |

