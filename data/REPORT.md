# AssetMantle endpoint radar

_Generated 2026-07-21T08:05:46Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23438517 | no | ✓ 54d (WE1) | 156/449 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 47d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23438517 | no | ✓ 88d (WE1) | 129/515 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23438517 | no | ✓ 81d (WE1) | 163/417 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23438517 | no | ✓ 55d (YR2) | 191/454 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23438517 | no | ✓ 54d (WE1) | 152/390 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 47d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23438517 | no | ✓ 88d (WE1) | 144/446 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23438517 | no | ✓ 81d (WE1) | 165/326 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23438517 | no | ✓ 55d (YR2) | 188/279 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 47d (WE1) | 101/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 55d (YR2) | 78/0 ms | canonical foundation endpoint |

