# AssetMantle endpoint radar

_Generated 2026-07-17T07:50:54Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23373848 | no | ✓ 58d (WE1) | 163/408 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 51d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23373848 | no | ✓ 33d (WE1) | 128/497 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23373848 | no | ✓ 85d (WE1) | 239/510 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23373848 | no | ✓ 59d (YR2) | 195/445 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23373848 | no | ✓ 58d (WE1) | 170/485 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 51d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23373848 | no | ✓ 33d (WE1) | 147/456 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23373848 | no | ✓ 85d (WE1) | 188/416 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23373848 | no | ✓ 59d (YR2) | 185/300 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 51d (WE1) | 77/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 59d (YR2) | 112/0 ms | canonical foundation endpoint |

