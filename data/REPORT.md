# AssetMantle endpoint radar

_Generated 2026-04-28T06:50:07Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22113185 | no | ✓ 79d (WE1) | 140/317 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22113185 | no | ✓ 72d (WE1) | 28/340 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22113185 | no | ✓ 54d (WE1) | 103/403 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22113185 | no | ✓ 47d (WE1) | 391/883 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 58d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22113185 | no | ✓ 89d (E7) | 198/280 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22113185 | no | ✓ 79d (WE1) | 136/183 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22113185 | no | ✓ 72d (WE1) | 25/334 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22113185 | no | ✓ 54d (WE1) | 119/358 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22113185 | no | ✓ 47d (WE1) | 137/337 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✗ 363d (localhost) | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22113185 | no | ✓ 89d (E8) | 198/281 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 72d (WE1) | 19/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 89d (E7) | 81/0 ms | canonical foundation endpoint |

