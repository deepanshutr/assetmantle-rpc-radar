# AssetMantle endpoint radar

_Generated 2026-05-16T06:43:30Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22394880 | no | ✓ 61d (WE1) | 188/417 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22394880 | no | ✓ 54d (WE1) | 215/306 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22394880 | no | ✓ 36d (WE1) | 172/553 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22394880 | no | ✓ 88d (WE1) | 200/428 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22394880 | no | ✓ 61d (R13) | 212/493 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22394880 | no | ✓ 61d (WE1) | 210/460 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22394880 | no | ✓ 54d (WE1) | 81/525 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22394880 | no | ✓ 36d (WE1) | 197/551 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22394880 | no | ✓ 88d (WE1) | 219/426 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22394880 | no | ✓ 61d (R13) | 202/415 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 54d (WE1) | 21/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 61d (R13) | 20/0 ms | canonical foundation endpoint |

