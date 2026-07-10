# AssetMantle endpoint radar

_Generated 2026-07-10T08:59:12Z — automated, do not edit._

**8/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23261818 | no | ✓ 65d (WE1) | 123/346 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 58d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23261818 | no | ✓ 40d (WE1) | 103/416 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✗ | — | — | ✓ 33d (WE1) | — | _err: http 500_ |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23261818 | no | ✓ 66d (YR2) | 144/370 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23261818 | no | ✓ 65d (WE1) | 132/363 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 58d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23261818 | no | ✓ 40d (WE1) | 117/356 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✗ | — | — | ✓ 33d (WE1) | — | _err: http 500_ |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23261818 | no | ✓ 66d (YR2) | 153/208 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 58d (WE1) | 14/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 66d (YR2) | 19/0 ms | canonical foundation endpoint |

