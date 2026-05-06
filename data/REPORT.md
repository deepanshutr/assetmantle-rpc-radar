# AssetMantle endpoint radar

_Generated 2026-05-06T06:50:23Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22238410 | no | ✓ 71d (WE1) | 130/377 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22238410 | no | ✓ 64d (WE1) | 39/205 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22238410 | no | ✓ 46d (WE1) | 113/373 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22238411 | no | ✓ 39d (WE1) | 211/435 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 50d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22238411 | no | ✓ 88d (E8) | 186/534 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22238410 | no | ✓ 71d (WE1) | 132/325 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22238410 | no | ✓ 64d (WE1) | 40/123 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22238411 | no | ✓ 46d (WE1) | 134/460 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22238411 | no | ✓ 39d (WE1) | 172/331 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22238411 | no | ✓ 88d (E8) | 190/506 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 64d (WE1) | 112/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 88d (E7) | 305/0 ms | canonical foundation endpoint |

