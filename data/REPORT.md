# AssetMantle endpoint radar

_Generated 2026-07-15T07:48:44Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23341582 | no | ✓ 60d (WE1) | 195/470 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 53d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23341582 | no | ✓ 35d (WE1) | 175/589 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23341582 | no | ✓ 87d (WE1) | 219/622 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23341582 | no | ✓ 61d (YR2) | 252/1728 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23341582 | no | ✓ 60d (WE1) | 196/553 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 53d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23341582 | no | ✓ 35d (WE1) | 202/704 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23341582 | no | ✓ 87d (WE1) | 217/441 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23341583 | no | ✓ 61d (YR2) | 253/1513 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 53d (WE1) | 194/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 61d (YR2) | 80/0 ms | canonical foundation endpoint |

