# AssetMantle endpoint radar

_Generated 2026-07-22T08:06:26Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23454648 | no | ✓ 53d (WE1) | 136/182 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 46d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23454648 | no | ✓ 87d (WE1) | 103/401 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23454648 | no | ✓ 80d (WE1) | 161/300 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23454648 | no | ✓ 54d (YR2) | 141/226 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23454648 | no | ✓ 53d (WE1) | 136/302 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 46d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23454648 | no | ✓ 87d (WE1) | 118/410 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23454648 | no | ✓ 80d (WE1) | 164/331 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23454648 | no | ✓ 54d (YR2) | 158/319 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 46d (WE1) | 14/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 54d (YR2) | 22/0 ms | canonical foundation endpoint |

