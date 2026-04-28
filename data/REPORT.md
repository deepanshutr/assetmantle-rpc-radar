# AssetMantle endpoint radar

_Generated 2026-04-28T02:28:38Z — automated, do not edit._

**10/14 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22110351 | no | ✓ 80d (WE1) | 208/460 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22110351 | no | ✓ 72d (WE1) | 223/527 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22110351 | no | ✓ 54d (WE1) | 160/215 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22110351 | no | ✓ 47d (WE1) | 192/465 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 58d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22110351 | no | ✓ 89d (E7) | 235/398 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (5)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22110351 | no | ✓ 80d (WE1) | 190/422 ms |  |
| `https://assetmantle-api.publicnode.com` | ✗ | — | — | ✓ 72d (WE1) | — | _err: http 404_ |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22110351 | no | ✓ 47d (WE1) | 199/395 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✗ 363d (localhost) | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22110352 | no | ✓ 89d (E8) | 252/409 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 72d (WE1) | 13/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 89d (E7) | 143/0 ms | canonical foundation endpoint |

