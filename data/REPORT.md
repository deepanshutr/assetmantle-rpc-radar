# AssetMantle endpoint radar

_Generated 2026-04-28T05:31:46Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22112335 | no | ✓ 79d (WE1) | 167/433 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22112335 | no | ✓ 72d (WE1) | 92/448 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22112335 | no | ✓ 54d (WE1) | 123/442 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22112335 | no | ✓ 47d (WE1) | 168/509 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22112335 | no | ✓ 89d (E7) | 214/398 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22112335 | no | ✓ 79d (WE1) | 172/494 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22112335 | no | ✓ 72d (WE1) | 155/549 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22112335 | no | ✓ 54d (WE1) | 169/528 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22112335 | no | ✓ 47d (WE1) | 174/742 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22112335 | no | ✓ 89d (E8) | 219/398 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 72d (WE1) | 91/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 89d (E7) | 166/0 ms | canonical foundation endpoint |

