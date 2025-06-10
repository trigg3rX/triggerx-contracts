#!/bin/bash

# Load env vars
source .env

# Check if dry run
DRY_RUN=${1:-false}

# Chain config
RPC_URL=${BASE_SEPOLIA_RPC}     # Replace with your RPC
ETHERSCAN_API_KEY=${BASESCAN_API_KEY}  # Add your key to .env

# Deployment values
IMPLEMENTATION_SALT=$(cast keccak "test3_ImplementationV1")
PROXY_SALT=$(cast keccak "test3_TriggerGasRegistry")
CREATE3_DEPLOYER=0x88826a677aDB340F0c7b8CCd6aF6aD96a40b0085   # replace with actual address
DEPLOYER_PK=$PRIVATE_KEY

# Updated addresses from verification script
IMPLEMENTATION_ADDRESS=0xC3E2510534C9CDD08C2d95f49c2fa1f72C6E1114
PROXY_ADDRESS=0xB76fE384ff7fCDDb042980894Fd6d672a432dC4b

# Validate required env vars
if [ -z "$ETHERSCAN_API_KEY" ]; then
    echo "Error: BASESCAN_API_KEY not found in .env"
    exit 1
fi

# Get bytecode
IMPLEMENTATION_BYTECODE=$(forge inspect src/TriggerGasRegistry.sol:TriggerGasRegistry bytecode)
if [ -z "$IMPLEMENTATION_BYTECODE" ]; then
    echo "Error: Could not get implementation bytecode"
    exit 1
fi

# Derive deterministic addresses with correct initialize function call (now requires initialOwner parameter)
INIT_CALLDATA=$(cast calldata "initialize(address)" $CREATE3_DEPLOYER)

echo "Contract Addresses:"
echo "Implementation: $IMPLEMENTATION_ADDRESS"
echo "Proxy: $PROXY_ADDRESS"
echo ""

if [ "$DRY_RUN" = "true" ]; then
    echo "DRY RUN MODE - Not submitting to Etherscan"
    echo ""
    echo "Implementation verification command:"
    echo "forge verify-contract --chain-id 84532 --watch --compiler-version 0.8.27 --etherscan-api-key \$BASESCAN_API_KEY $IMPLEMENTATION_ADDRESS src/TriggerGasRegistry.sol:TriggerGasRegistry"
    echo ""
    echo "After implementation verification, manually verify proxy on BaseScan:"
    echo "1. Go to: https://sepolia.basescan.org/address/$PROXY_ADDRESS"
    echo "2. Click 'Contract' tab -> 'More Options' -> 'Is this a proxy?'"
    echo "3. Select 'Verify Proxy'"
    echo "4. Enter Implementation Address: $IMPLEMENTATION_ADDRESS"
    echo "5. Select Proxy Type: EIP-1967 (Transparent/UUPS)"
    exit 0
fi

echo "Verifying Implementation at $IMPLEMENTATION_ADDRESS"
echo "This contains your TriggerGasRegistry business logic..."
echo ""

forge verify-contract --chain-id 84532 \
  --watch \
  --compiler-version 0.8.27 \
  --etherscan-api-key $ETHERSCAN_API_KEY \
  $IMPLEMENTATION_ADDRESS \
  src/TriggerGasRegistry.sol:TriggerGasRegistry

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Implementation verified successfully!"
    echo ""
    echo "🔗 NEXT STEP: Manually verify the proxy to see TriggerGasRegistry methods:"
    echo "1. Go to: https://sepolia.basescan.org/address/$PROXY_ADDRESS"
    echo "2. Click 'Contract' tab -> 'More Options' -> 'Is this a proxy?'"
    echo "3. Select 'Verify Proxy'"
    echo "4. Enter Implementation Address: $IMPLEMENTATION_ADDRESS"
    echo "5. Select Proxy Type: EIP-1967 (Transparent/UUPS)"
    echo ""
    echo "After proxy verification, $PROXY_ADDRESS will show all TriggerGasRegistry read/write methods!"
else
    echo ""
    echo "❌ Implementation verification failed!"
    echo "Check the error above and try again."
fi
