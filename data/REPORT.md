# AssetMantle endpoint radar

_Generated 2026-07-16T07:53:26Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23357753 | no | ✓ 59d (WE1) | 136/317 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 52d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23357753 | no | ✓ 34d (WE1) | 104/355 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23357753 | no | ✓ 86d (WE1) | 131/357 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23357753 | no | ✓ 60d (YR2) | 142/322 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23357753 | no | ✓ 59d (WE1) | 138/236 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 52d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23357753 | no | ✓ 34d (WE1) | 116/353 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23357753 | no | ✓ 86d (WE1) | 131/333 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23357753 | no | ✓ 60d (YR2) | 158/203 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 52d (WE1) | 44/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 60d (YR2) | 24/0 ms | canonical foundation endpoint |

