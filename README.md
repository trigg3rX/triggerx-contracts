# TriggerX AVS Contract

## Overview

The TriggerX AVS (Actively Validated Service) contract is an **upgradeable proxy** that serves as TriggerX's interface to the Imua AVS-Manager precompile. It provides a modular architecture for managing operators, tasks, and validation logic while maintaining compatibility with the Imua ecosystem.

## Contract Architecture

### Core Design Principles

- **Upgradeable**: Uses OpenZeppelin's UUPS (Universal Upgradeable Proxy Standard) pattern
- **Modular**: Pluggable reward and slashing modules for future extensibility
- **Proxy Pattern**: Acts as a thin wrapper around Imua's AVS-Manager precompile
- **Access Control**: Owner-based permissions for administrative functions

### Inheritance Structure

```solidity
TriggerXAvs
├── OwnableUpgradeable (Access control)
└── UUPSUpgradeable (Upgrade mechanism)
```

## Contract Components

### 1. Storage Variables

| Variable | Type | Description |
|----------|------|-------------|
| `rewardManager` | `address` | Future on-chain reward handler (initially 0x0) |
| `slasher` | `address` | Future on-chain slashing handler (initially 0x0) |
| `_taskDefs` | `TaskDefinitions` | Local storage for task definitions |

### 2. Core Functionality

#### Initialization & Upgrades
- **`initialize(address initialOwner)`**: Upgradeable initializer replacing constructor
- **`_authorizeUpgrade(address)`**: Owner-only upgrade authorization

#### AVS Management (Owner Only)
- **`registerAVS(AVSParams calldata params)`**: Register the AVS with Imua
- **`updateAVS(AVSParams calldata params)`**: Update AVS configuration
- **`setRewardManager(address _rewardManager)`**: Set reward handler module
- **`setSlasher(address _slasher)`**: Set slashing handler module

#### Operator Management
- **`registerOperatorToAVS()`**: Operators can opt into the AVS
- **`deregisterOperatorFromAVS()`**: Operators can opt out of the AVS
- **`registerBLSPublicKey(...)`**: Register BLS public key for cryptographic operations

#### Task Management
- **`createTask(...)`**: Create new validation tasks
- **`operatorSubmitTask(...)`**: Operators submit task responses
- **`challenge(...)`**: Initiate challenges for task validation

#### Task Definitions
- **`createTaskDefinition(...)`**: Define reusable task templates (Owner only)
- **`getTaskDefinition(uint8 id)`**: Retrieve task definition by ID

### 3. Events

#### Operator Events
```solidity
event OperatorOptedIn(address indexed operator);
event OperatorOptedOut(address indexed operator);
event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey);
```

#### Task Events
```solidity
event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase);
event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected);
event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name);
event TaskCreated(uint64 indexed taskID, bytes32 definitionHash, uint8 kind);
```

#### Administrative Events
```solidity
event RewardManagerUpdated(address indexed newRewardManager);
event SlasherUpdated(address indexed newSlasher);
```

### 4. View Functions

The contract provides comprehensive read-only access to AVS state:

| Function | Description |
|----------|-------------|
| `getOptInOperators(address avsAddress)` | Get all opted-in operators |
| `getRegisteredPubkey(address operator, address avsAddr)` | Get operator's BLS public key |
| `getAVSUSDValue(address avsAddr)` | Get total USD value of AVS |
| `getOperatorOptedUSDValue(...)` | Get operator's staked USD value |
| `getTaskInfo(address taskAddress, uint64 taskID)` | Get detailed task information |
| `isOperator(address operator)` | Check if address is a registered operator |
| `getCurrentEpoch(string calldata epochIdentifier)` | Get current epoch number |
| `getChallengeInfo(...)` | Get challenge details |
| `getOperatorTaskResponse(...)` | Get operator's task response |
| `getOperatorTaskResponseList(...)` | Get all responses for a task |

## Data Structures

### AVSParams
Configuration parameters for AVS registration:
```solidity
struct AVSParams {
    address sender;
    string avsName;
    uint64 minStakeAmount;
    address taskAddress;
    address slashAddress;
    address rewardAddress;
    address[] avsOwnerAddresses;
    address[] whitelistAddresses;
    string[] assetIDs;
    uint64 avsUnbondingPeriod;
    uint64 minSelfDelegation;
    string epochIdentifier;
    uint64 miniOptInOperators;
    uint64 minTotalStakeAmount;
    uint64 avsRewardProportion;
    uint64 avsSlashProportion;
}
```

### TaskDefinition
Template for creating tasks:
```solidity
struct TaskDefinition {
    uint8 taskDefinitionId;
    string name;
    uint256 baseRewardFeeForAttesters;
    uint256 baseRewardFeeForPerformer;
    uint256 baseRewardFeeForAggregator;
    uint256 minimumVotingPower;
    uint256[] restrictedOperatorIds;
    uint256 maximumNumberOfAttesters;
}
```

### TaskInfo
Comprehensive task state information:
```solidity
struct TaskInfo {
    address taskContractAddress;
    string name;
    bytes hash;
    uint64 taskID;
    uint64 taskResponsePeriod;
    uint64 taskStatisticalPeriod;
    uint64 taskChallengePeriod;
    uint8 thresholdPercentage;
    uint64 startingEpoch;
    string actualThreshold;
    address[] optInOperators;
    address[] signedOperators;
    address[] noSignedOperators;
    address[] errSignedOperators;
    string taskTotalPower;
    OperatorActivePower[] operatorActivePower;
    bool isExpected;
    address[] eligibleRewardOperators;
    address[] eligibleSlashOperators;
}
```

## Usage Flow

### 1. AVS Setup (Owner)
1. Deploy and initialize the contract
2. Register AVS with Imua using `registerAVS()`
3. Create task definitions using `createTaskDefinition()`
4. Optionally set reward/slashing modules

### 2. Operator Onboarding
1. Operators call `registerOperatorToAVS()` to opt in
2. Register BLS public key using `registerBLSPublicKey()`
3. Begin participating in task validation

### 3. Task Lifecycle
1. **Creation**: Call `createTask()` with task parameters
2. **Submission**: Operators submit responses via `operatorSubmitTask()`
3. **Challenge**: Validators can challenge results using `challenge()`
4. **Resolution**: Rewards/slashing handled by configured modules

## Integration with Imua

The contract interfaces with Imua's AVS-Manager precompile at address `0x0000000000000000000000000000000000000901`. All core AVS operations are delegated to this precompile while maintaining local state for task definitions and future extensibility.

## Security Features

- **Access Control**: Owner-only administrative functions
- **Upgradeable**: UUPS pattern with owner authorization
- **Validation**: Input validation (e.g., threshold percentage ≤ 100)
- **Events**: Comprehensive logging for auditability
- **Modular**: Pluggable reward/slashing logic for security isolation

## Development

### Dependencies
- OpenZeppelin Contracts Upgradeable
- Foundry for testing and deployment
- Custom Imua AVS-Manager interface

### File Structure
```
contracts/src/
├── TriggerXAvs.sol              # Main contract
├── interfaces/
│   └── IAVSManager.sol          # Imua precompile interface
├── libraries/
│   └── TaskDefinitionLibrary.sol # Task definition management
└── precompiles/
    └── ImuaAVSManager.sol       # Precompile constants
```

## License

MIT License - See LICENSE file for details.
