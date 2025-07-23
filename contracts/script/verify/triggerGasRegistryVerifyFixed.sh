#!/usr/bin/env bash
source .env

# Accept chain ID as argument, default to Base Sepolia
CHAIN_ID="${1:-84532}"

# Configure network-specific settings based on chain ID
case "$CHAIN_ID" in
  "84532")
    NETWORK_NAME="Base Sepolia"
    RPC_URL=${BASE_SEPOLIA_RPC}
    EXPLORER_URL="https://sepolia.basescan.org/address"
    ;;
  "11155420")
    NETWORK_NAME="OP Sepolia"
    RPC_URL=${OPSEPOLIA_RPC}
    EXPLORER_URL="https://sepolia-optimism.etherscan.io/address"
    ;;
  *)
    echo "Error: Unsupported chain ID: $CHAIN_ID" >&2
    echo "Supported chains: 84532 (Base Sepolia), 11155420 (OP Sepolia)" >&2
    exit 1
    ;;
esac

ETHERSCAN_API_KEY=${ETHERSCAN_API_KEY}

# TriggerGasRegistry proxy address (same on both networks)
PROXY_ADDRESS="0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a"

echo "=== Verifying TriggerGasRegistry on $NETWORK_NAME (Chain ID: $CHAIN_ID) ==="
echo "Proxy address: $PROXY_ADDRESS"

if [ -z "$ETHERSCAN_API_KEY" ]; then
  echo "Error: ETHERSCAN_API_KEY must be set" >&2
  exit 1
fi

if [ -z "$RPC_URL" ]; then
  echo "Error: RPC_URL for $NETWORK_NAME must be set in .env" >&2
  exit 1
fi

# Fetch implementation address from proxy contract using EIP-1967 implementation slot
# Implementation slot: 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc
echo "Fetching current implementation address from proxy..."
CURRENT_IMPLEMENTATION_ADDRESS=$(cast storage --rpc-url "$RPC_URL" "$PROXY_ADDRESS" 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc | sed 's/0x000000000000000000000000/0x/')

if [ -z "$CURRENT_IMPLEMENTATION_ADDRESS" ] || [ "$CURRENT_IMPLEMENTATION_ADDRESS" = "0x" ]; then
  echo "Error: Could not fetch implementation address from proxy" >&2
  exit 1
fi

echo "Current implementation address: $CURRENT_IMPLEMENTATION_ADDRESS"
echo ""

echo "=== Verifying TriggerGasRegistry Implementation ==="

# Verify the current implementation contract using Etherscan v2 API
forge verify-contract \
  --watch --compiler-version 0.8.27 \
  --verifier-url "https://api.etherscan.io/v2/api?chainid=$CHAIN_ID" \
  --etherscan-api-key "$ETHERSCAN_API_KEY" \
  "$CURRENT_IMPLEMENTATION_ADDRESS" \
  src/TriggerGasRegistry.sol:TriggerGasRegistry

if [ $? -eq 0 ]; then
  echo "\n✅ TriggerGasRegistry Implementation verified successfully on $NETWORK_NAME!\n"
  echo "🔗 View on Explorer:\n$EXPLORER_URL/$CURRENT_IMPLEMENTATION_ADDRESS\n"
else
  echo "❌ Implementation verification failed!"
  echo "This is critical - exiting."
  exit 1
fi

echo "=== Verifying TriggerGasRegistry Proxy ==="

# Get the owner address by calling the owner() function on the proxy
echo "Fetching owner address from proxy..."
OWNER_ADDRESS=$(cast call --rpc-url "$RPC_URL" "$PROXY_ADDRESS" "owner()(address)" | sed 's/0x000000000000000000000000/0x/')

if [ -z "$OWNER_ADDRESS" ] || [ "$OWNER_ADDRESS" = "0x" ]; then
  echo "Warning: Could not fetch owner address, using default"
  OWNER_ADDRESS="0x88826a677adb340f0c7b8ccd6af6ad96a40b0085"
fi

echo "Owner address: $OWNER_ADDRESS"

# For proxy verification, we need to use the ORIGINAL deployment constructor arguments
# The proxy was originally deployed with a different implementation and then upgraded
echo "Using original deployment constructor arguments..."

# Original implementation from deployment bytecode analysis
ORIGINAL_IMPLEMENTATION="0x9f452702490e11af38c213de30c000a46e4da399"

# The proxy was deployed with initialize(address) - single parameter
INIT_CALL_DATA=$(cast calldata "initialize(address)" "$OWNER_ADDRESS")
echo "Initialize call data: $INIT_CALL_DATA"

# Use cast to encode the constructor arguments with ORIGINAL implementation
CONSTRUCTOR_ARGS=$(cast abi-encode "constructor(address,bytes)" "$ORIGINAL_IMPLEMENTATION" "$INIT_CALL_DATA")
# Remove the 0x prefix for forge verify-contract
CONSTRUCTOR_ARGS=${CONSTRUCTOR_ARGS#0x}

echo "Proxy constructor args (with original implementation): $CONSTRUCTOR_ARGS"

# Verify the proxy contract using Etherscan v2 API
forge verify-contract \
  --watch --compiler-version 0.8.27 \
  --verifier-url "https://api.etherscan.io/v2/api?chainid=$CHAIN_ID" \
  --etherscan-api-key "$ETHERSCAN_API_KEY" \
  --constructor-args "$CONSTRUCTOR_ARGS" \
  "$PROXY_ADDRESS" \
  lib/othentic-core-contracts/lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol:ERC1967Proxy

if [ $? -eq 0 ]; then
  echo "\n✅ TriggerGasRegistry Proxy verified successfully on $NETWORK_NAME!\n"
  echo "🔗 View on Explorer:\n$EXPLORER_URL/$PROXY_ADDRESS\n"
  echo "📋 Contract Summary:\n- Current Implementation: $CURRENT_IMPLEMENTATION_ADDRESS ✅ VERIFIED\n- Proxy: $PROXY_ADDRESS ✅ VERIFIED\n- Original Implementation: $ORIGINAL_IMPLEMENTATION\n- Network: $NETWORK_NAME (Chain ID: $CHAIN_ID)\n- Owner: $OWNER_ADDRESS"
  echo "\n💡 Note: Proxy was verified using original deployment constructor arguments."
  echo "The proxy has been upgraded since original deployment, which is why we use the original implementation address for verification."
else
  echo "❌ Proxy verification failed!"
  echo "📋 Contract Summary:\n- Current Implementation: $CURRENT_IMPLEMENTATION_ADDRESS ✅ VERIFIED\n- Proxy: $PROXY_ADDRESS ❌ (verification failed)\n- Network: $NETWORK_NAME (Chain ID: $CHAIN_ID)\n- Owner: $OWNER_ADDRESS"
fi 