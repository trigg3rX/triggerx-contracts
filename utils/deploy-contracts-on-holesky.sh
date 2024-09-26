#!/bin/bash

HOLESKY_RPC_URL=$(grep HOLESKY_RPC_URL .env | cut -d '=' -f2)
HOLESKY_PRIVATE_KEY=$(grep HOLESKY_PRIVATE_KEY .env | cut -d '=' -f2)
ETHERSCAN_API_KEY=$(grep ETHERSCAN_API_KEY .env | cut -d '=' -f2)

REGISTRY_COORDINATOR=$(jq .registryCoordinator utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)
AVS_DIRECTORY=$(jq .avsDirectory utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)
STAKE_REGISTRY=$(jq .stakeRegistry utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)

cd contracts

echo "Deploying Task Manager Interface..."
output1=$(forge create ItxTaskManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY)
ITX_TASK_MANAGER_ADDRESS=$(echo "$output1" | grep "Deployed to:" | awk '{print $NF}')

echo "Deploying Task Manager..."
output2=$(forge create txTaskManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY --constructor-args $REGISTRY_COORDINATOR 100 --verify --etherscan-api-key $ETHERSCAN_API_KEY)
TX_TASK_MANAGER_ADDRESS=$(echo "$output2" | grep "Deployed to:" | awk '{print $NF}')

echo "Deploying Service Manager..."
output3=$(forge create txServiceManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY --constructor-args $AVS_DIRECTORY $REGISTRY_COORDINATOR $STAKE_REGISTRY $TX_TASK_MANAGER_ADDRESS --verify --etherscan-api-key $ETHERSCAN_API_KEY)
TX_SERVICE_MANAGER_ADDRESS=$(echo "$output3" | grep "Deployed to:" | awk '{print $NF}')

echo "All contracts deployed successfully!"

json=$(jq -n \
  --arg itxTaskManager "$ITX_TASK_MANAGER_ADDRESS" \
  --arg txTaskManager "$TX_TASK_MANAGER_ADDRESS" \
  --arg txServiceManager "$TX_SERVICE_MANAGER_ADDRESS" \
  '{
    itxTaskManager: $itxTaskManager,
    txTaskManager: $txTaskManager,
    txServiceManager: $txServiceManager,
  }')
echo "$json" | jq '.' > ../utils/deployments/holeskyDeployments.json

echo "Addresses and transaction hashes are saved in utils/deployments/holeskyDeployments.json"