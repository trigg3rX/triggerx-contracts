#!/bin/bash

# Compile the contracts first
echo "Compiling contracts..."
forge build

# ProxyHub contract address on Base Sepolia
PROXY_HUB="0x96c5F575940DBe98fd9600F74F0c36139A7Be2A2"

# Constructor arguments
LZ_ENDPOINT="0x6EDCE65403992e310A62460808c4b910D972f10f"
OWNER="0x88826a677adb340f0c7b8ccd6af6ad96a40b0085"
SRC_EID="40217"  # Holesky EID
THIS_CHAIN_EID="40245"  # Base Sepolia EID

# Initial keepers array
INITIAL_KEEPERS=(
    "0x19aBAdfBb70e60CEa506291BeE00d8307CB62667"
    "0x2dcFBc2613b60EC755A2530F818d0E1A8Ed1Db3f"
    "0x95AB9BCa0eECE37468fd6038E39021e7EEd208a3"
    "0xBaE997eFC6E496d8A34181C2443d97b491A3f057"
    "0x45E4657eE118325F852609e9dB45211fe1A0044C"
    "0xEbe4f49161eb9002eD19E06F943DCe73D24B596E"
    "0x350ddC16818CDFA1E0c06c26Fd5e76360032faa4"
    "0x011FCbAE5f306cd793456ab7d4c0CC86756c693D"
    "0x0a067a261c5F5e8C4c0b9137430b4FE1255EB62e"
)

# Format keepers array for constructor arguments
KEEPERS_ARGS="["
for keeper in "${INITIAL_KEEPERS[@]}"; do
    KEEPERS_ARGS+="$keeper,"
done
KEEPERS_ARGS=${KEEPERS_ARGS%,}  # Remove trailing comma
KEEPERS_ARGS+="]"

# Create a temporary file for the constructor arguments
TEMP_FILE=$(mktemp)
echo "constructor(address,address,uint32,uint32,address[])" > "$TEMP_FILE"
echo "$LZ_ENDPOINT" >> "$TEMP_FILE"
echo "$OWNER" >> "$TEMP_FILE"
echo "$SRC_EID" >> "$TEMP_FILE"
echo "$THIS_CHAIN_EID" >> "$TEMP_FILE"
echo "$KEEPERS_ARGS" >> "$TEMP_FILE"

echo "Verifying contract..."
# Verify the contract
forge verify-contract \
    --chain-id 84532 \
    --constructor-args-path "$TEMP_FILE" \
    --compiler-version v0.8.22 \
    --num-of-optimizations 200 \
    "$PROXY_HUB" \
    src/lz/ProxyHub.sol:ProxyHub \
    $BASESCAN_API_KEY

# Clean up
rm "$TEMP_FILE" 