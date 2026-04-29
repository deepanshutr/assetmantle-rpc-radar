# AssetMantle endpoint radar

_Generated 2026-04-29T06:45:53Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22128811 | no | ✓ 78d (WE1) | 146/390 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22128811 | no | ✓ 71d (WE1) | 30/158 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22128811 | no | ✓ 53d (WE1) | 106/419 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22128811 | no | ✓ 46d (WE1) | 118/231 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 57d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22128811 | no | ✓ 88d (E7) | 213/318 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22128811 | no | ✓ 78d (WE1) | 139/380 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22128811 | no | ✓ 71d (WE1) | 33/106 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22128811 | no | ✓ 53d (WE1) | 135/461 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22128811 | no | ✓ 46d (WE1) | 139/361 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✗ 362d (localhost) | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22128811 | no | ✓ 88d (E8) | 213/280 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 71d (WE1) | 49/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 88d (E7) | 84/0 ms | canonical foundation endpoint |

