#!/bin/bash

script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

cd "$script_path"

echo "Fetching contract ABIs..."
go run ./fetch_abi.go

echo "Generating bindings..."

if [[ "$(docker images -q abigen-with-interfaces 2> /dev/null)" == "" ]]; then
    docker build -t abigen-with-interfaces -f ./abigen-with-interfaces.Dockerfile "$script_path"
fi

function create_binding {
    contract=$1
    binding_dir=$2
    echo "Creating bindings for $contract..."
    mkdir -p $binding_dir/${contract}
    
    abi_file="$script_path/abis/${contract}.abi"
    bin_file="$script_path/abis/${contract}.bin"

    if [ ! -f "$abi_file" ]; then
        echo "Warning: ABI file not found for $contract: $abi_file"
        return 1
    fi

    rm -f $binding_dir/${contract}/binding.go
    
    # Build abigen arguments array
    abigen_args=(
        --abi=/home/bindings/abis/${contract}.abi
        --pkg=contract${contract}
        --out=/home/binding_dir/${contract}/binding.go
    )
    
    # Add bin file if it exists
    if [ -f "$bin_file" ]; then
        abigen_args=(--bin=/home/bindings/abis/${contract}.bin "${abigen_args[@]}")
    fi
    
    docker run --rm \
        --user $(id -u):$(id -g) \
        -v $(realpath $binding_dir):/home/binding_dir \
        -v $(realpath $script_path):/home/bindings \
        abigen-with-interfaces \
        "${abigen_args[@]}"
    
    # Copy ABI as JSON for reference
    cp "$abi_file" "$script_path/abis/${contract}.abi.json"
}

# Process all .abi files found in the abis directory
for abi_file in $script_path/abis/*.abi; do
    if [ -f "$abi_file" ]; then
        # Skip .abi.json files (they're copied artifacts)
        if [[ "$abi_file" == *.abi.json ]]; then
            continue
        fi
        contract=$(basename "$abi_file" .abi)
        create_binding "$contract" ./contracts

        sleep 1
        rm "$abi_file"
    fi
done
