# AssetMantle endpoint radar

_Generated 2026-05-14T06:55:50Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22363707 | no | ✓ 63d (WE1) | 130/396 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22363707 | no | ✓ 56d (WE1) | 72/380 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22363707 | no | ✓ 38d (WE1) | 129/504 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22363707 | no | ✓ 31d (WE1) | 148/457 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22363707 | no | ✓ 63d (R13) | 175/264 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22363707 | no | ✓ 63d (WE1) | 129/371 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22363707 | no | ✓ 56d (WE1) | 85/396 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22363707 | no | ✓ 38d (WE1) | 130/467 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22363707 | no | ✓ 31d (WE1) | 153/394 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22363707 | no | ✓ 63d (R13) | 184/417 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 56d (WE1) | 57/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 63d (R13) | 90/0 ms | canonical foundation endpoint |

