#!/bin/bash

source .env

cd contracts

forge script script/TriggerXStakeDeployerOPS.s.sol:TriggerXStakeDeployerOPS \
        --chain-id 11155420 \
        --rpc-url $OP_SEPOLIA_RPC_URL \
        --private-key $TRIGGERX_OWNER_PRIVATE_KEY \
        --verify \
        --verifier blockscout \
        --verifier-url https://optimism-sepolia.blockscout.com/api/ \
        --etherscan-api-key $BLOCKSCOUT_API_KEY \
        --broadcast -vvvv