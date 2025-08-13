#!/usr/bin/env bash
source .env
# Holesky configuration
CHAIN_ID="1"
# CHAIN_ID_HEX="0x4268"
RPC_URL=${ETHEREUM_RPC}
ETHERSCAN_API_KEY=${ETHERSCAN_API_KEY}  # Try Holesky specific key first, fallback to generic

# Static contract address
CONTRACT_ADDRESS="0xc1F545Fe807b72429952dbFEFE8702658e4C5875"

# Constructor args (pre-defined)
ENDPOINT_ADDRESS="0x1a44076050125825900e736c501f859c50fE728c"
TASK_EXECUTION_HUB="0x0E258AC902F116589c86eb3D83253d7223b9c12A"
DEST_CHAIN_ID="30184"
OWNER_ADDRESS="0xdB43967857F70F04C0ffe62ed240DB74CE4E3C87"
AVS_GOVERNANCE="0x875B5ff698B74B26f39C223c4996871F28AcDdea"

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
