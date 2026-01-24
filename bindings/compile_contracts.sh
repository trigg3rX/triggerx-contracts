#!/bin/bash
# Compile Solidity contracts from source and extract ABIs

set -e

script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

cd "$script_path"

OUTPUT_DIR="$script_path/abis"
mkdir -p "$OUTPUT_DIR"

# Check dependencies
if ! command -v forge &> /dev/null; then
    echo "Error: forge (Foundry) is not installed"
    exit 1
fi

if ! command -v jq &> /dev/null; then
    echo "Error: jq is not installed"
    exit 1
fi

echo "=== Compiling contracts from source ==="

# Navigate to contracts root and build
CONTRACTS_ROOT="$script_path/../contracts"
echo ""
echo "Compiling from contracts root: $CONTRACTS_ROOT"
echo "-------------------------------------------"

cd "$CONTRACTS_ROOT"

# Clean and build
forge build

# Function to extract ABI for a contract
extract_abi() {
    local contract=$1
    
    echo "Extracting ABI for $contract..."
    
    # Find the compiled artifact (Foundry outputs to out/<ContractName>.sol/<ContractName>.json)
    # Try the standard Foundry output path first
    artifact_file="out/${contract}.sol/${contract}.json"
    
    if [ ! -f "$artifact_file" ]; then
        echo "Warning: Standard path not found, searching..."
        # Search for the artifact in out directory
        artifact_file=$(find out -type f -name "${contract}.json" -path "*/${contract}.sol/${contract}.json" 2>/dev/null | head -1)
    fi
    
    if [ -z "$artifact_file" ] || [ ! -f "$artifact_file" ]; then
        echo "Warning: Exact path not found, trying broader search..."
        artifact_file=$(find out -type f -name "${contract}.json" 2>/dev/null | head -1)
    fi
    
    if [ -z "$artifact_file" ] || [ ! -f "$artifact_file" ]; then
        echo "Error: Could not find compiled artifact for $contract"
        return 1
    fi
    
    echo "Found artifact: $artifact_file"
    
    # Extract ABI using jq
    jq '.abi' "$artifact_file" > "$OUTPUT_DIR/${contract}.abi"
    
    # Also extract bytecode if available (for deployment)
    bytecode=$(jq -r '.bytecode.object // empty' "$artifact_file")
    if [ -n "$bytecode" ] && [ "$bytecode" != "null" ] && [ "$bytecode" != "0x" ]; then
        echo "$bytecode" > "$OUTPUT_DIR/${contract}.bin"
    fi
    
    echo "Successfully extracted ABI for $contract"
}

# Extract ABIs for all contracts
echo ""
echo "=== Extracting ABIs ==="

# Contracts from othentic-core-contracts
extract_abi "AttestationCenter"
extract_abi "AvsGovernance"

# Contracts from triggerx-contracts
extract_abi "JobRegistry"
extract_abi "TriggerGasRegistry"
extract_abi "TaskExecutionHub"
extract_abi "TaskExecutionSpoke"
extract_abi "AvsGovernanceLogic"
extract_abi "TriggerXSafeFactory"
extract_abi "TriggerXSafeModule"

cd "$script_path"

echo ""
echo "=== Compilation complete ==="
echo "ABIs extracted to: $OUTPUT_DIR"
ls -la "$OUTPUT_DIR"
