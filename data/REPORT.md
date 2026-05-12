# AssetMantle endpoint radar

_Generated 2026-05-12T06:52:55Z — automated, do not edit._

**12/15 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22332340 | no | ✓ 65d (WE1) | 144/335 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22332339 | no | ✓ 58d (WE1) | 42/89 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22332340 | no | ✓ 40d (WE1) | 122/460 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22332340 | no | ✓ 33d (WE1) | 140/470 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✓ | 22332340 | no | ✓ 65d (R13) | 179/207 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## REST (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22332339 | no | ✓ 65d (WE1) | 124/166 ms |  |
| `https://assetmantle-rest.publicnode.com` | ✓ | 22332339 | no | ✓ 58d (WE1) | 33/91 ms |  |
| `https://assetmantle-rest.stakerhouse.com` | ✓ | 22332340 | no | ✓ 40d (WE1) | 122/235 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22332340 | no | ✓ 33d (WE1) | 135/307 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | — | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✓ | 22332340 | no | ✓ 65d (R13) | 164/214 ms | canonical foundation endpoint (currently broken; this radar tracks the fix) |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 58d (WE1) | 30/0 ms |  |
| `grpc.assetmantle.one:443` | ✓ | — | no | ✓ 65d (R13) | 69/0 ms | canonical foundation endpoint |

