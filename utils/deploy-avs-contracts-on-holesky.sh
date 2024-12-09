#!/bin/bash

source .env

cd contracts

forge script script/TriggerXDeployerHolesky.s.sol:TriggerXDeployerHolesky \
        --chain-id 17000 \
        --rpc-url $HOLESKY_RPC_URL \
        --private-key $TRIGGERX_OWNER_PRIVATE_KEY \
        --verify \
        --verifier blockscout \
        --verifier-url https://eth-holesky.blockscout.com/api/ \
        --etherscan-api-key $BLOCKSCOUT_API_KEY \
        --broadcast -vvvv
