#!/bin/bash

echo "Fetching contract ABIs and bytecode..."
go run ./bindings/fetch_abi.go

echo "Generating bindings..."

script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

if [[ "$(docker images -q abigen-with-interfaces 2> /dev/null)" == "" ]]; then
    docker build -t abigen-with-interfaces -f ./bindings/abigen-with-interfaces.Dockerfile $script_path
fi

function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    network=$4
    echo "creating bindings for $contract on $network..."
    mkdir -p $binding_dir/${contract}
    
    abi_file="$script_path/abis/${contract}.${network}.abi"
    bin_file="$script_path/abis/${contract}.${network}.bin"

    rm -f $binding_dir/${contract}/binding.go
    docker run --rm \
        --user $(id -u):$(id -g) \
        -v $(realpath $binding_dir):/home/binding_dir \
        -v $(realpath $script_path):/home/bindings \
        abigen-with-interfaces \
        --bin=/home/bindings/abis/${contract}.${network}.bin \
        --abi=/home/bindings/abis/${contract}.${network}.abi \
        --pkg=contract${contract} \
        --out=/home/binding_dir/${contract}/binding.go
}

cd $script_path

# Process Holesky contracts
for abi_file in $script_path/abis/*.eth.abi; do
    if [ -f "$abi_file" ]; then
        contract=$(basename "$abi_file" .eth.abi)
        if [ -f "$script_path/abis/$contract.eth.bin" ]; then
            create_binding . "$contract" ./contracts "eth"
            cp $script_path/abis/$contract.eth.abi $script_path/abis/$contract.eth.abi.json
        fi
    fi
done

# Process Base contracts
for abi_file in $script_path/abis/*.base.abi; do
    if [ -f "$abi_file" ]; then
        contract=$(basename "$abi_file" .base.abi)
        if [ -f "$script_path/abis/$contract.base.bin" ]; then
            create_binding . "$contract" ./contracts "base"
            cp $script_path/abis/$contract.base.abi $script_path/abis/$contract.base.abi.json
        fi
    fi
done
