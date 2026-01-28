# Changelog

All notable changes to TriggerX Contracts will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-27

### Security Enhancements

#### TaskDispatcher Signature Verification
- Added `taskDispatcher` address state variable for authorized signers
- Added `_verifyTaskDispatcherSignature()` using ECDSA + EIP-191 (`toEthSignedMessageHash`)
- Signature hash includes: `jobId`, `target`, `deadline`, `msg.sender`, `block.chainid`

#### Safe Module Job Owner Validation
- Added `triggerXSafeModule` address state variable
- Added `_validateSafeModuleCalldata()` for on-chain jobOwner validation
- Uses assembly to extract `jobOwner` from calldata at offset 164 (4 + 5*32)
- Reverts with `JobOwnerMismatch` if passed jobOwner doesn't match registry

#### Transaction Security
- `deadline` parameter prevents stale transaction execution
- Chain ID included in signature to prevent cross-chain replay
- Custom errors for gas-efficient reverts: `SignatureExpired`, `JobNotFound`, `TaskDispatcherNotSet`, `InvalidSignature`, `InvalidSafeModuleCalldata`, `JobOwnerMismatch`

### Contract Changes

#### TaskExecutionHub
- Added `taskDispatcher` state variable
- Added `triggerXSafeModule` state variable
- Updated `executeFunction(jobId, ethAmount, target, data, deadline, signature)`
- Added `setTaskDispatcher(address)` with `TaskDispatcherUpdated` event
- Added `setTriggerXSafeModule(address)` with `TriggerXSafeModuleUpdated` event
- Added `_validateSafeModuleCalldata()` internal function

#### TaskExecutionSpoke
- Mirror security changes from Hub
- Added `taskDispatcher` and `triggerXSafeModule` state variables
- Updated `executeFunction` with same signature as Hub
- Cross-chain execution now requires dispatcher signature

#### TriggerXSafeModule
- Updated `execJobFromHub(safeAddress, actionTarget, actionValue, actionData, operation, jobOwner)`
- Validates `safe.isOwner(jobOwner)` before execution
- Reverts with `SafeNotOwnedByJobOwner` if validation fails
- Immutable `taskExecutionHub` address for authorization

#### JobRegistry
- Added `getJobOwner(uint256 jobId)` for job ownership queries
- Added `unpackJobId(uint256 jobId)` for packed job ID decoding

### Deployments (Testnet)

#### TaskExecutionHub (Base Sepolia - 84532)
| Type | Address |
|------|---------|
| Proxy | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` |
| Implementation | `0xfe754df6a0301b2d48696E3960b8A439Cb85Ebc2` |

#### TaskExecutionSpoke
| Chain | Proxy | Implementation |
|-------|-------|----------------|
| Sepolia (11155111) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0x275bb750A2430248c5b17faE4956762b4e0D0690` |
| OP Sepolia (11155420) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0x31C937708Bb02bD02eB5AC6F5865e76f7EdB70A9` |
| Arbitrum Sepolia (421614) | `0x179c62e83c3f90981B65bc12176FdFB0f2efAD54` | `0xDD8157dd23C38b594EfB21eadE73AEA76252069C` |

#### TriggerXSafeModule (Deterministic - same address all chains)
| Chain | Address |
|-------|---------|
| Base Sepolia | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| Sepolia | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| OP Sepolia | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |
| Arbitrum Sepolia | `0xE4EF5835ceE541255B3B9BC3E82F11496145BB35` |

#### TriggerXSafeFactory
| Chain | Address |
|-------|---------|
| Base Sepolia | `0x04359eDC46Cd6C6BD7F6359512984222BE10F8Be` |
| Sepolia | `0xdf76E2A796a206D877086c717979054544B1D9Bc` |
| OP Sepolia | `0x04359eDC46Cd6C6BD7F6359512984222BE10F8Be` |
| Arbitrum Sepolia | `0x04359eDC46Cd6C6BD7F6359512984222BE10F8Be` |

---

## [0.x.x] - Pre-1.0 (Legacy)

Previous development iterations before formal versioning was established.
