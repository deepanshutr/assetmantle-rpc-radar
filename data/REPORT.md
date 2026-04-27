# AssetMantle endpoint radar

_Generated 2026-04-27T13:58:52Z — automated, do not edit._

**7/14 endpoints healthy.**

## RPC (6)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-rpc.polkachu.com` | ✓ | 22102228 | no | ✓ 80d (WE1) | 172/426 ms |  |
| `https://assetmantle-rpc.publicnode.com` | ✓ | 22102228 | no | ✓ 73d (WE1) | 98/295 ms |  |
| `https://assetmantle-rpc.stakerhouse.com` | ✓ | 22102228 | no | ✓ 55d (WE1) | 159/532 ms |  |
| `https://public.stakewolle.com/cosmos/assetmantle/rpc` | ✓ | 22102228 | no | ✓ 47d (WE1) | 161/494 ms |  |
| `https://rpc-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✓ 58d (R12) | — | _err: Get "https://rpc-assetmantle-ia.cosmosia.notional.ventures/…_ |
| `https://rpc.assetmantle.one` | ✗ | — | — | ✗ 350d (Kubernetes Ingres…) | — | _err: Get "https://rpc.assetmantle.one/status": tls: failed to ve…_ |

## REST (5)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `https://assetmantle-api.polkachu.com` | ✓ | 22102228 | no | ✓ 80d (WE1) | 174/446 ms |  |
| `https://assetmantle-api.publicnode.com` | ✗ | — | — | ✓ 73d (WE1) | — | _err: http 404_ |
| `https://public.stakewolle.com/cosmos/assetmantle/rest` | ✓ | 22102228 | no | ✓ 47d (WE1) | 182/340 ms |  |
| `https://rest-assetmantle-ia.cosmosia.notional.ventures` | ✗ | — | — | ✗ 364d (localhost) | — | _err: Get "https://rest-assetmantle-ia.cosmosia.notional.ventures…_ |
| `https://rest.assetmantle.one` | ✗ | — | — | ✗ 350d (Kubernetes Ingres…) | — | _err: Get "https://rest.assetmantle.one/cosmos/base/tendermint/v1…_ |

## GRPC (3)

| URL | Status | Height | Archive | TLS | Latency p50/p95 | Notes |
|---|---|---|---|---|---|---|
| `assetmantle-grpc.polkachu.com:443` | ✗ | — | — | — | — | _err: tls dial: dial tcp 65.108.131.174:443: i/o timeout_ |
| `assetmantle-grpc.publicnode.com:443` | ✓ | — | no | ✓ 73d (WE1) | 63/0 ms |  |
| `grpc.assetmantle.one:443` | ✗ | — | — | ✗ 243d (Amazon RSA 2048 M…) | — | _err: tls dial: tls: failed to verify certificate: x509: certific…_ |

