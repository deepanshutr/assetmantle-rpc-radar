# AssetMantle endpoint radar

_Generated 2026-04-30T06:47:51Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22144517 | no | ✓ 77d (WE1) | 150/363 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22144517 | no | ✓ 70d (WE1) | 78/409 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22144517 | no | ✓ 52d (WE1) | 113/406 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22144517 | no | ✓ 45d (WE1) | 168/436 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 56d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22144517 | no | ✓ 87d (E7) | 209/357 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22144517 | no | ✓ 77d (WE1) | 157/423 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22144517 | no | ✓ 70d (WE1) | 82/295 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22144517 | no | ✓ 52d (WE1) | 130/386 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22144517 | no | ✓ 45d (WE1) | 151/686 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✗ 361d (localhost) | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22144517 | no | ✓ 87d (E8) | 220/372 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 70d (WE1) | 114/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 87d (E7) | 143/0 ms | canonical foundation endpoint |

