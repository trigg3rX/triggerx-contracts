#!/bin/bash

# Set your API keys
ARBISCAN_API_KEY=C9UKWI9QU5PDF9PNB5D9SDZ7QEVXYMZE8X
ETHERSCAN_API_KEY=9DCCHJFSS3DCNMRCUZCMHZQB8I23W66JUU
BASESCAN_API_KEY=E5W5AG4YPZEH8YGQXQI33BY9GR135I66HZ

# Verify JobRegistry on Arbitrum Sepolia
forge verify-contract --rpc-url https://arb-sepolia.g.alchemy.com/v2/2bhdU9feedVXPCEcw6arkzNnqhprSsBt --etherscan-api-key $ARBISCAN_API_KEY --verifier-url https://api-sepolia.arbiscan.io/api 0x76BCd570700be5d98B8C9CbB5de47151212A9Dd0 src/JobRegistry.sol:JobRegistry

# Verify TriggerGasRegistry on Arbitrum Sepolia
forge verify-contract --rpc-url https://arb-sepolia.g.alchemy.com/v2/2bhdU9feedVXPCEcw6arkzNnqhprSsBt --etherscan-api-key $ARBISCAN_API_KEY --verifier-url https://api-sepolia.arbiscan.io/api 0x35cf6891D356c78253e4968Eb5696BDF54b43f3c src/TriggerGasRegistry.sol:TriggerGasRegistry

# Verify TaskExecutionHub on Base Sepolia
forge verify-contract --rpc-url https://base-sepolia.g.alchemy.com/v2/2bhdU9feedVXPCEcw6arkzNnqhprSsBt --etherscan-api-key $BASESCAN_API_KEY --verifier-url https://api-sepolia.basescan.org/api 0xa5a23AE80b817978EFBaDa86C835B87551B43C9a src/lz/TaskExecutionHub.sol:TaskExecutionHub

# Verify TaskExecutionSpoke on Arbitrum Sepolia
forge verify-contract --rpc-url https://arb-sepolia.g.alchemy.com/v2/2bhdU9feedVXPCEcw6arkzNnqhprSsBt --etherscan-api-key $ARBISCAN_API_KEY --verifier-url https://api-sepolia.arbiscan.io/api 0xdEE5383F801EC3985d7d058edE3EeA42FA6f5537 src/lz/TaskExecutionSpoke.sol:TaskExecutionSpoke

# Verify AvsGovernanceLogic on Ethereum Sepolia
forge verify-contract --rpc-url https://eth-sepolia.g.alchemy.com/v2/2bhdU9feedVXPCEcw6arkzNnqhprSsBt --etherscan-api-key $ETHERSCAN_API_KEY --verifier-url https://api-sepolia.etherscan.io/api 0xa5a23AE80b817978EFBaDa86C835B87551B43C9a src/AvsGovernanceLogic.sol:AvsGovernanceLogic

echo "All contracts verification submitted!"