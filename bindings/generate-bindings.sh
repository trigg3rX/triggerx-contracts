#!/bin/bash

# Script to generate Go bindings for TriggerX contracts
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Generating TriggerX Contract Bindings${NC}"

# Check if abigen is installed
if ! command -v abigen &> /dev/null; then
    echo -e "${RED}❌ abigen not found. Installing...${NC}"
    go install github.com/ethereum/go-ethereum/cmd/abigen@latest
    export PATH=$PATH:$(go env GOPATH)/bin
fi

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo -e "${RED}❌ jq not found. Please install jq to extract ABIs${NC}"
    exit 1
fi

# Create directories if they don't exist
mkdir -p abis contracts

echo -e "${YELLOW}📄 Extracting ABIs...${NC}"

# Extract ABIs from compiled contracts
CONTRACT_OUT_DIR="../contracts/out"

# Extract TriggerXAvs ABI
if [ -f "$CONTRACT_OUT_DIR/TriggerXAvs.sol/TriggerXAvs.json" ]; then
    cat "$CONTRACT_OUT_DIR/TriggerXAvs.sol/TriggerXAvs.json" | jq '.abi' > abis/TriggerXAvs.abi
    echo "✅ Extracted TriggerXAvs ABI"
else
    echo -e "${RED}❌ TriggerXAvs.json not found. Run 'forge build' first.${NC}"
    exit 1
fi



# Extract TaskDefinitionLibrary ABI
if [ -f "$CONTRACT_OUT_DIR/TaskDefinitionLibrary.sol/TaskDefinitionLibrary.json" ]; then
    cat "$CONTRACT_OUT_DIR/TaskDefinitionLibrary.sol/TaskDefinitionLibrary.json" | jq '.abi' > abis/TaskDefinitionLibrary.abi
    echo "✅ Extracted TaskDefinitionLibrary ABI"
else
    echo -e "${RED}❌ TaskDefinitionLibrary.json not found. Run 'forge build' first.${NC}"
    exit 1
fi

echo -e "${YELLOW}⚙️  Generating Go bindings...${NC}"

# Generate Go bindings with separate packages to avoid conflicts
abigen --abi abis/TriggerXAvs.abi --pkg triggerxavs --type TriggerXAvs --out contracts/TriggerXAvs.go
echo "✅ Generated TriggerXAvs.go"



abigen --abi abis/TaskDefinitionLibrary.abi --pkg taskdefinition --type TaskDefinitionLibrary --out contracts/TaskDefinitionLibrary.go
echo "✅ Generated TaskDefinitionLibrary.go"

echo -e "${GREEN}✨ All bindings generated successfully!${NC}"
echo ""
echo -e "${YELLOW}📦 To use in your Go project:${NC}"
echo "  go get github.com/triggerx/triggerx-contracts/bindings"
echo ""
echo -e "${YELLOW}📖 Example usage:${NC}"
echo '  import triggerxavs "github.com/triggerx/triggerx-contracts/bindings/triggerxavs"'
echo '  import taskdefinition "github.com/triggerx/triggerx-contracts/bindings/taskdefinition"'
echo '  contract, err := triggerxavs.NewTriggerXAvs(address, client)' 
echo '  taskDefinition, err := taskdefinition.NewTaskDefinitionLibrary(address, client)' 