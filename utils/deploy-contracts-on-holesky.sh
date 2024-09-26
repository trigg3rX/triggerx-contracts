#!/bin/bash

HOLESKY_RPC_URL=$(grep HOLESKY_RPC_URL .env | cut -d '=' -f2)
HOLESKY_PRIVATE_KEY=$(grep HOLESKY_PRIVATE_KEY .env | cut -d '=' -f2)

# get registry coordinator address from utils/deployments.json
REGISTRY_COORDINATOR=$(jq .registryCoordinator utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)
AVS_DIRECTORY=$(jq .avsDirectory utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)
STAKE_REGISTRY=$(jq .stakeRegistry utils/deployments/eigenDeployments.json | cut -d '"' -f2 | cut -d '"' -f2)

# cd contracts 

# echo "Deploying Task Manager Interface..."
# output1=$(forge create ItxTaskManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY)
# ITX_TASK_MANAGER_ADDRESS=$(echo "$output1" | grep "Deployed to:" | awk '{print $NF}')
# ITX_TASK_MANAGER_HASH=$(echo "$output1" | grep "Transaction hash:" | awk '{print $NF}')

# echo "Deploying Task Manager..."
# output2=$(forge create txTaskManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY --constructor-args $REGISTRY_COORDINATOR 100)
# TX_TASK_MANAGER_ADDRESS=$(echo "$output2" | grep "Deployed to:" | awk '{print $NF}')
# TX_TASK_MANAGER_HASH=$(echo "$output2" | grep "Transaction hash:" | awk '{print $NF}')

# echo "Deploying Service Manager..."
# output3=$(forge create txServiceManager --rpc-url $HOLESKY_RPC_URL --private-key $HOLESKY_PRIVATE_KEY --constructor-args $AVS_DIRECTORY $REGISTRY_COORDINATOR $STAKE_REGISTRY $TX_TASK_MANAGER_ADDRESS)
# TX_SERVICE_MANAGER_ADDRESS=$(echo "$output3" | grep "Deployed to:" | awk '{print $NF}')
# TX_SERVICE_MANAGER_HASH=$(echo "$output3" | grep "Transaction hash:" | awk '{print $NF}')

# echo "All contracts deployed successfully!"

# json=$(jq -n \
#   --arg itxTaskManager "$ITX_TASK_MANAGER_ADDRESS" \
#   --arg itxTaskManagerHash "$ITX_TASK_MANAGER_HASH" \
#   --arg txTaskManager "$TX_TASK_MANAGER_ADDRESS" \
#   --arg txTaskManagerHash "$TX_TASK_MANAGER_HASH" \
#   --arg txServiceManager "$TX_SERVICE_MANAGER_ADDRESS" \
#   --arg txServiceManagerHash "$TX_SERVICE_MANAGER_HASH" \
#   '{
#     itxTaskManager: $itxTaskManager,
#     itxTaskManagerHash: $itxTaskManagerHash,
#     txTaskManager: $txTaskManager,
#     txTaskManagerHash: $txTaskManagerHash,
#     txServiceManager: $txServiceManager,
#     txServiceManagerHash: $txServiceManagerHash
#   }')
# echo "$json" | jq '.' > ../utils/holeskyDeployments.json

echo "Addresses and transaction hashes are saved in utils/holeskyDeployments.json"