# AssetMantle endpoint radar

_Generated 2026-05-07T06:55:05Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22254073 | no | ✓ 70d (WE1) | 180/446 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22254073 | no | ✓ 63d (WE1) | 97/321 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22254073 | no | ✓ 45d (WE1) | 175/589 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22254073 | no | ✓ 38d (WE1) | 270/514 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 49d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22254073 | no | ✓ 87d (E8) | 215/560 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22254073 | no | ✓ 70d (WE1) | 178/429 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22254073 | no | ✓ 63d (WE1) | 227/598 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22254074 | no | ✓ 45d (WE1) | 199/635 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22254074 | no | ✓ 38d (WE1) | 270/484 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22254074 | no | ✓ 87d (E8) | 214/559 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 63d (WE1) | 74/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 87d (E7) | 351/0 ms | canonical foundation endpoint |

