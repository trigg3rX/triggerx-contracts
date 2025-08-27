# TriggerX Contracts

This repository contains the contracts for the TriggerX project.

## Types of contracts

### 1. Othentic Stack Contracts

- AVS Governance Contract (deployed on Ethereum)
  - AVS Governance Logic Contract (deployed on Ethereum, as a hook for the AVS Governance Contract)
- Attestation Center Contract (deployed on Base)

### 2. Support Contracts

- TriggerXJobRegistry Contract (deployed on all chains where users can register their jobs)
- TriggerGasRegistry Contract (deployed on all chains where users buy and consume TG - TriggerGas)
- TaskExecutionHub and TaskExecutionSpoke Contracts (deployed on all chains where users need to execute their tasks)

## Folder structure

- `contracts/`: Contains the contracts for the TriggerX project.
  - `src/`: Contains the Support Contracts for the TriggerX project.
  - `script/`: Contains the scripts deploy, upgrade, verify and migrate functions for contracts.
  - `test/`: Contains the unit and security tests for the contracts.
- `bindings/`: Contains the bindings for the TriggerX project.
  - `abis/`: Contains the abis and bytecodes for Othentic Contracts.
  - `contracts/`: Contains bindings generated using geth abigen.

## Contract Deployment Commands

1. Environment variables: Copy the .env.example file and fill the values required.

   **Note**: The value of `RPC_URL` determines the chain where the contracts are deployed in scripts, so make sure to change it to the correct chain (testnet or mainnet) before deployment.

2. Deploy Othentic Contracts using `otcli`:

```bash
npm install -g @othentic/cli
```

Deploying on Ethereum Sepolia and Base Sepolia with 0.1 ETH as initial deposit for L1 and L2 MessageHandler:

```bash
otcli network deploy --name TriggerX --eth --l2-rewards --security-provider eigenlayer --l1-chain 11155111 --l2-chain 84532 --l1-initial-deposit 100000000000000000 --l2-initial-deposit 100000000000000000
```

The CLI will ask for registration details for the AVS to EigenLayer.

3. Configure the TriggerX AVS:

```bash
# Set Max Voting Power for Operators
otcli network set-max-voting-power --l1-chain 11155111

# Set Staking Contracts
otcli network set-staking-contracts --l1-chain 11155111

# Create Task Definition
otcli network create-task-definition --l2-chain 84532
```

### **Note: All commands below are meant to be executed from `./contracts/` directory.**

- User need to set the salt for contracts as required.
- The RPC Url in script determines the chain where the contracts are deployed.
- Set the `AVS_GOVERNANCE_ADDRESS` and `ATTESTATION_CENTER_ADDRESS` in the `.env` file.

4. Deploy JobRegistry Contract:

```bash
forge script script/deploy/1_deployJobRegistry.s.sol:DeployJobRegistry --verify --verifier-url "https://api.etherscan.io/v2/api" --etherscan-api-version v2 --chain 84532 --etherscan-api-key <ETHERSCAN_API_KEY> --broadcast 
```

5. Deploy TriggerGasRegistry Contract:

- User need to set the operator address (TaskExecutionAddress) to used in the script.

```bash
forge script script/deploy/2_deployTriggerGasRegistry.s.sol:DeployTriggerGasRegistry --verify --verifier-url "https://api.etherscan.io/v2/api" --etherscan-api-version v2 --chain 84532 --etherscan-api-key <ETHERSCAN_API_KEY> --broadcast 
```

- If the contracts are being deployed from scratch, then user won't have the operator address (TaskExecutionAddress) to be used in the script. In that case, user need to call the `SetOperatorOnGasRegistry` script after the deployment of the TaskExecution Contracts and set the TaskExecutionAddress in`.env`.

```bash
forge script script/deploy/2_deployTriggerGasRegistry.s.sol:SetOperatorOnGasRegistry --broadcast 
```

6. Deploy TaskExecutionHub Contract:

- Set `JOB_REGISTRY_ADDRESS` and `TRIGGER_GAS_REGISTRY_ADDRESS` in the `.env` file.
- Deploys the TaskExecutionHub Contract on the Base Chain (As we have deployed AttestationCenter Contract on Base Chain).
- 0.1 ETH is deposited to the TaskExecutionHub Contract.

```bash
forge script script/deploy/3_deployTaskExecutionHub.s.sol:DeployTaskExecutionHub --verify --verifier-url "https://api.etherscan.io/v2/api" --etherscan-api-version v2 --chain 84532 --etherscan-api-key <ETHERSCAN_API_KEY> --broadcast 
```

- If the a new Spoke is deployed, then user need to call the `AddSokesToHub` script after the deployment of the TaskExecutionHub Contract, and setting the `spokeEids` in the script function.

```bash
forge script script/deploy/3_deployTaskExecutionHub.s.sol:AddSokesToHub --broadcast 
```

- If the a new JobRegistry is deployed, then user need to call the `SetJobRegistryonHub` script after the deployment of the TaskExecutionHub Contract.

```bash
forge script script/deploy/3_deployTaskExecutionHub.s.sol:SetJobRegistryonHub --broadcast 
```

- If the a new TriggerGasRegistry is deployed, then user need to call the `SetTriggerGasRegistryonHub` script after the deployment of the TaskExecutionHub Contract.


```bash
forge script script/deploy/3_deployTaskExecutionHub.s.sol:SetTriggerGasRegistryonHub --broadcast 
```

7. Deploy TaskExecutionSpoke Contract:

```bash
forge script script/deploy/4_deployTaskExecutionSpoke.s.sol:DeployTaskExecutionSpoke --multi --chain 11155420 --broadcast
```

- To verify the TaskExecutionSpoke Contracts:

```bash
forge verify-contract --watch --compiler-version 0.8.27 --verifier-url "https://api.etherscan.io/v2/api" --chain 421614 --constructor-args $(cast abi-encode "constructor(address,address)" "<LZ_ENDPOINT_ADDRESS>" "<DEPLOYER_ADDRESS>") --etherscan-api-key <ETHERSCAN_API_KEY> <IMPLEMENTATION_ADDRESS> src/lz/TaskExecutionSpoke.sol:TaskExecutionSpoke
```

8. Deploy AVS Governance Logic Contract:

- Set `TASK_EXECUTION_ADDRESS` in the `.env` file.
- 0.1 ETH is deposited to the AvsGovernanceLogic Contract.
- Sets the AvsGovernanceLogic Contract on the AvsGovernance Contract.

```bash
forge script script/deploy/5_deployAvsGovernanceLogic.s.sol:DeployAvsGovernanceLogic --verify --verifier-url "https://api.etherscan.io/v2/api" --etherscan-api-version v2 --chain 11155111 --etherscan-api-key <ETHERSCAN_API_KEY> --broadcast 
```
