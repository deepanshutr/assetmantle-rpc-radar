# AssetMantle endpoint radar

_Generated 2026-07-08T08:06:02Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23229685 | no | ✓ 67d (WE1) | 668/747 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 60d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23229685 | no | ✓ 42d (WE1) | 187/590 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23229685 | no | ✓ 35d (WE1) | 207/303 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23229685 | no | ✓ 68d (YR2) | 233/508 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23229685 | no | ✓ 67d (WE1) | 404/685 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 60d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23229685 | no | ✓ 42d (WE1) | 193/568 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23229685 | no | ✓ 35d (WE1) | 209/394 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23229685 | no | ✓ 68d (YR2) | 242/1491 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 60d (WE1) | 155/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 68d (YR2) | 29/0 ms | canonical foundation endpoint |

