# AssetMantle endpoint radar

_Generated 2026-05-05T06:41:25Z — automated, do not edit._

**11/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22222662 | no | ✓ 72d (WE1) | 141/311 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22222662 | no | ✓ 65d (WE1) | 43/78 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22222662 | no | ✓ 47d (WE1) | 103/398 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22222662 | no | ✓ 40d (WE1) | 153/410 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 51d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✗ | — | — | ✓ 89d (E8) | — | _err: Get "https://rpc.assetmantle.one/status": context deadline …_ |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22222662 | no | ✓ 72d (WE1) | 146/389 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22222662 | no | ✓ 65d (WE1) | 22/137 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22222662 | no | ✓ 47d (WE1) | 121/341 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22222662 | no | ✓ 40d (WE1) | 134/234 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22222664 | no | — | 168/5985 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 65d (WE1) | 94/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 89d (E7) | 255/0 ms | canonical foundation endpoint |

