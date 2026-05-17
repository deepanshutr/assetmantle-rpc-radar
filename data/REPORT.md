# AssetMantle endpoint radar

_Generated 2026-05-17T06:53:49Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22410492 | no | ✓ 60d (WE1) | 206/451 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22410492 | no | ✓ 53d (WE1) | 212/256 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22410492 | no | ✓ 35d (WE1) | 173/555 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22410492 | no | ✓ 87d (WE1) | 199/519 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22410492 | no | ✓ 60d (R13) | 230/512 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22410492 | no | ✓ 60d (WE1) | 192/419 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22410492 | no | ✓ 53d (WE1) | 214/256 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22410492 | no | ✓ 35d (WE1) | 199/613 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22410492 | no | ✓ 87d (WE1) | 209/414 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22410492 | no | ✓ 60d (R13) | 246/520 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 53d (WE1) | 155/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 60d (R13) | 31/0 ms | canonical foundation endpoint |

