# AssetMantle endpoint radar

_Generated 2026-07-27T09:27:05Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23536133 | no | ✓ 48d (WE1) | 200/464 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 41d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23536133 | no | ✓ 82d (WE1) | 226/858 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23536133 | no | ✓ 75d (WE1) | 215/672 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23536133 | no | ✓ 49d (YR2) | 258/1798 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23536133 | no | ✓ 48d (WE1) | 218/531 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 41d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23536133 | no | ✓ 82d (WE1) | 191/600 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23536133 | no | ✓ 75d (WE1) | 219/455 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23536133 | no | ✓ 49d (YR2) | 247/545 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 41d (WE1) | 25/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 49d (YR2) | 58/0 ms | canonical foundation endpoint |

