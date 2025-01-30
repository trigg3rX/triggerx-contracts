#!/bin/bash

source .env

cd contracts

forge script script/TriggerXUpdaterHolesky.s.sol:TriggerXUpdaterHolesky \
        --chain-id 17000 \
        --rpc-url $HOLESKY_RPC_URL \
        --private-key $TRIGGERX_OWNER_PRIVATE_KEY \
        --verify \
        --etherscan-api-key $ETHERSCAN_API_KEY \
        --broadcast -vvvv