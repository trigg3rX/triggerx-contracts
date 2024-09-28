#!/bin/bash

cd contracts

forge script script/TriggerXDeployer.s.sol:TriggerXDeployer --rpc-url $HOLESKY_RPC_URL --broadcast