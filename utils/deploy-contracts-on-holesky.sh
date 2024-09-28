#!/bin/bash

cd contracts

forge script script/TriggerXDeployer.s.sol:TriggerXDeployer \
        --rpc-url $HOLESKY_RPC_URL \
        --private-key $AVS_OWNER_PRIVATE_KEY \
        --verify \
        --etherscan-api-key $ETHERSCAN_API_KEY \
        --broadcast
