# AssetMantle endpoint radar

_Generated 2026-07-28T06:21:31Z — automated, do not edit._

**10/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 23550259 | no | ✓ 48d (WE1) | 193/456 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✗ | — | — | ✓ 40d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 23550259 | no | ✓ 82d (WE1) | 184/612 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 23550259 | no | ✓ 74d (WE1) | 241/507 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 23550259 | no | ✓ 48d (YR2) | 257/1617 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 23550259 | no | ✓ 48d (WE1) | 185/405 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✗ | — | — | ✓ 40d (WE1) | — | _err: http 403_ |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 23550259 | no | ✓ 82d (WE1) | 212/632 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 23550259 | no | ✓ 74d (WE1) | 234/449 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 23550260 | no | ✓ 48d (YR2) | 258/1607 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 40d (WE1) | 28/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 48d (YR2) | 81/0 ms | canonical foundation endpoint |

