# Changelog

All notable changes to TriggerX Contracts will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-27

### Added
- **Dispatcher Signature Verification**: Added signature verification to `TaskExecutionHub` and `TaskExecutionSpoke` for enhanced security
- **Deadline & Nonce Parameters**: Added `deadline` and `nonce` parameters to `executeFunction` to prevent replay attacks
- **JobId Parameter**: Added `jobId` parameter to `TriggerXSafeModule.execJobFromHub` for proper job tracking
- **PackedJobIdLib**: Library for efficient job ID packing/unpacking

### Changed
- Updated `executeFunction` signature across Hub and Spoke contracts
- Updated `TriggerXSafeModule` to support new execution flow

### Security
- Replay attack prevention via nonce tracking
- Deadline-based transaction expiry
- Dispatcher-only execution enforcement

### Deployments (Testnet)

#### TaskExecutionHub (Base Sepolia - 84532)
| Contract | Type | Address |
|----------|------|---------|
| TaskExecutionHub | Proxy | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` |
| TaskExecutionHub | Implementation | `0xfe754df6a0301b2d48696E3960b8A439Cb85Ebc2` |

#### TaskExecutionSpoke
| Chain | Proxy | Implementation |
|-------|-------|----------------|
| Sepolia (11155111) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0x275bb750A2430248c5b17faE4956762b4e0D0690` |
| OP Sepolia (11155420) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0x31C937708Bb02bD02eB5AC6F5865e76f7EdB70A9` |
| Arbitrum Sepolia (421614) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0xDD8157dd23C38b594EfB21eadE73AEA76252069C` |

#### TriggerXSafeModule (Same address on all chains)
| Chain | Address |
|-------|---------|
| Base Sepolia (84532) | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| Sepolia (11155111) | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| OP Sepolia (11155420) | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| Arbitrum Sepolia (421614) | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |

## [0.x.x] - Pre-1.0 (Legacy)

Previous development iterations before formal versioning was established.
