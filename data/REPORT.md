# AssetMantle endpoint radar

_Generated 2026-07-23T08:09:29Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23470790 | no | ✓ 52d (WE1) | 195/462 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 45d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23470790 | no | ✓ 86d (WE1) | 187/291 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23470790 | no | ✓ 79d (WE1) | 381/518 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23470791 | no | ✓ 53d (YR2) | 240/1588 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23470790 | no | ✓ 52d (WE1) | 205/480 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 45d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23470790 | no | ✓ 86d (WE1) | 203/697 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23470790 | no | ✓ 79d (WE1) | 258/461 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23470790 | no | ✓ 53d (YR2) | 238/537 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 45d (WE1) | 37/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 53d (YR2) | 71/0 ms | canonical foundation endpoint |

