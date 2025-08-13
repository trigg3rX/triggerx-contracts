#!/usr/bin/env bash
source .env

# Accept chain ID as argument, default to Base Sepolia
CHAIN_ID="42161"

# Configure network-specific settings based on chain ID
case "$CHAIN_ID" in
  "42161") # Arbitrum
    NETWORK_NAME="Arbitrum"
    RPC_URL=${ARBITRUM_RPC}
    EXPLORER_URL="https://arbiscan.io/address"
    VERIFIER_URL="https://api.etherscan.io/v2/api"
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

# JobRegistry proxy address (same on both networks)
PROXY_ADDRESS="0xaf1189afd1f1880f09aec3cbc32cf415c735c710"

echo "=== Verifying JobRegistry on $NETWORK_NAME (Chain ID: $CHAIN_ID) ==="
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
echo "Fetching implementation address from proxy..."
IMPLEMENTATION_ADDRESS=$(cast storage --rpc-url "$RPC_URL" "$PROXY_ADDRESS" 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc | sed 's/0x000000000000000000000000/0x/')

if [ -z "$IMPLEMENTATION_ADDRESS" ] || [ "$IMPLEMENTATION_ADDRESS" = "0x" ]; then
  echo "Error: Could not fetch implementation address from proxy" >&2
  exit 1
fi

echo "Implementation address: $IMPLEMENTATION_ADDRESS"
echo ""

echo "=== Verifying JobRegistry Implementation ==="

# Verify the implementation contract using Etherscan v2 API
forge verify-contract \
  --watch --compiler-version 0.8.27 \
  --verifier-url "$VERIFIER_URL" \
  --verifier etherscan \
  --etherscan-api-key "$ETHERSCAN_API_KEY" \
  --etherscan-api-version v2 \
  --chain "$CHAIN_ID" \
  "$IMPLEMENTATION_ADDRESS" \
  src/JobRegistry.sol:JobRegistry

if [ $? -eq 0 ]; then
  echo ""
  echo "✅ JobRegistry Implementation verified successfully on $NETWORK_NAME!"
  echo ""
  echo "🔗 View on Explorer:"
  echo "$EXPLORER_URL/$IMPLEMENTATION_ADDRESS"
  echo ""
else
  echo "❌ Implementation verification failed!"
  exit 1
fi

echo "=== Verifying JobRegistry Proxy ==="

# Get the owner address by calling the owner() function on the proxy
echo "Fetching owner address from proxy..."
OWNER_ADDRESS=$(cast call --rpc-url "$RPC_URL" "$PROXY_ADDRESS" "owner()(address)" | sed 's/0x000000000000000000000000/0x/')

if [ -z "$OWNER_ADDRESS" ] || [ "$OWNER_ADDRESS" = "0x" ]; then
  echo "Warning: Could not fetch owner address, using default"
  OWNER_ADDRESS="0x88826a677adb340f0c7b8ccd6af6ad96a40b0085"
fi

echo "Owner address: $OWNER_ADDRESS"

# Prepare constructor args for ERC1967Proxy
# ERC1967Proxy constructor: constructor(address implementation, bytes memory _data)
# The _data is the encoded initialize call: initialize(address initialOwner)

# Use cast to properly encode the initialize function call
INIT_CALL_DATA=$(cast calldata "initialize(address)" "$OWNER_ADDRESS")
echo "Initialize call data: $INIT_CALL_DATA"

# Use cast to encode the constructor arguments
CONSTRUCTOR_ARGS=$(cast abi-encode "constructor(address,bytes)" "$IMPLEMENTATION_ADDRESS" "$INIT_CALL_DATA")
# Remove the 0x prefix for forge verify-contract
CONSTRUCTOR_ARGS=${CONSTRUCTOR_ARGS#0x}

echo "Proxy constructor args: $CONSTRUCTOR_ARGS"

# Verify the proxy contract using Etherscan v2 API
forge verify-contract \
  --watch --compiler-version 0.8.27 \
  --verifier-url "$VERIFIER_URL" \
  --verifier etherscan \
  --etherscan-api-key "$ETHERSCAN_API_KEY" \
  --etherscan-api-version v2 \
  --chain "$CHAIN_ID" \
  --constructor-args "$CONSTRUCTOR_ARGS" \
  "$PROXY_ADDRESS" \
  lib/othentic-core-contracts/lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol:ERC1967Proxy

if [ $? -eq 0 ]; then
  echo ""
  echo "✅ JobRegistry Proxy verified successfully on $NETWORK_NAME!"
  echo ""
  echo "🔗 View on Explorer:"
  echo "$EXPLORER_URL/$PROXY_ADDRESS"
  echo ""
  echo "📋 Contract Summary:"
  echo "- Implementation: $IMPLEMENTATION_ADDRESS"
  echo "- Proxy: $PROXY_ADDRESS"
  echo "- Network: $NETWORK_NAME (Chain ID: $CHAIN_ID)"
  echo "- Owner: $OWNER_ADDRESS"
else
  echo "❌ Proxy verification failed!"
fi 