#!/usr/bin/env bash
source .env
# Holesky configuration
CHAIN_ID="17000"
CHAIN_ID_HEX="0x4268"
RPC_URL=${HOLESKY_RPC}
ETHERSCAN_API_KEY=${ETHERSCAN_API_KEY}  # Try Holesky specific key first, fallback to generic

# Static contract address
CONTRACT_ADDRESS="0x4EbE2f2b7db5B48559167f2be7d760B23b00B427"

# Constructor args (pre-defined)
ENDPOINT_ADDRESS="0x6EDCE65403992e310A62460808c4b910D972f10f"
TASK_EXECUTION_HUB="0x2469e89386947535A350EEBccC5F2754fd35F474"
DEST_CHAIN_ID="40245"
OWNER_ADDRESS="0x88826a677aDB340F0c7b8CCd6aF6aD96a40b0085"
AVS_GOVERNANCE="0x12f45551f11Df20b3EcBDf329138Bdc65cc58Ec0"

CONSTRUCTOR_ARGS=$(cast abi-encode "constructor(address,address,uint32,address,address)" \
  "$ENDPOINT_ADDRESS" "$TASK_EXECUTION_HUB" "$DEST_CHAIN_ID" "$OWNER_ADDRESS" "$AVS_GOVERNANCE")

echo "AvsGovernanceLogic address: $CONTRACT_ADDRESS (chain $CHAIN_ID)"

echo "Constructor args: $CONSTRUCTOR_ARGS"

if [ -z "$ETHERSCAN_API_KEY" ]; then
  echo "Error: HOLESKY_ETHERSCAN_API_KEY or ETHERSCAN_API_KEY must be set" >&2
  exit 1
fi

# Actual verification
forge verify-contract \
  --watch --compiler-version 0.8.27 \
  --chain-id "$CHAIN_ID" \
  --etherscan-api-key "$ETHERSCAN_API_KEY" \
  --constructor-args "$CONSTRUCTOR_ARGS" \
  "$CONTRACT_ADDRESS" \
  src/AvsGovernanceLogic.sol:AvsGovernanceLogic

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ AvsGovernanceLogic verified successfully on Holesky!"
    echo ""
    echo "🔗 View on Holesky Etherscan:"
    echo "https://holesky.etherscan.io/address/$CONTRACT_ADDRESS"
    echo ""
    echo "Contract Details:"
    echo "- Address: $CONTRACT_ADDRESS"
    echo "- Network: Holesky Testnet (Chain ID: $CHAIN_ID)"
    echo "- LayerZero Endpoint: $ENDPOINT_ADDRESS"
    echo "- TaskExecutionHub: $TASK_EXECUTION_HUB"
    echo "- Destination Chain ID: $DEST_CHAIN_ID (Base Sepolia)"
    echo "- Owner: $OWNER_ADDRESS"
    echo "- AVS Governance: $AVS_GOVERNANCE"
else
    echo ""
    echo "❌ Verification failed!"
    echo "Check the error above and try again."
fi
