# AssetMantle endpoint radar

_Generated 2026-07-09T09:18:30Z — automated, do not edit._

**8/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23246681 | no | ✓ 66d (WE1) | 203/457 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 59d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23246681 | no | ✓ 41d (WE1) | 180/588 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✗ | — | — | ✓ 34d (WE1) | — | _err: http 500_ |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23246682 | no | ✓ 67d (YR2) | 217/1911 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23246681 | no | ✓ 66d (WE1) | 191/432 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 59d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23246681 | no | ✓ 41d (WE1) | 184/519 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✗ | — | — | ✓ 34d (WE1) | — | _err: http 500_ |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23246682 | no | ✓ 67d (YR2) | 218/1665 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 59d (WE1) | 32/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 67d (YR2) | 24/0 ms | canonical foundation endpoint |

