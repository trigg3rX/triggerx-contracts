// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";
import {ERC1967Proxy} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";
import {TaskExecutionSpoke} from "../../src/lz/TaskExecutionSpoke.sol";
import {IAttestationCenter} from "../../src/interfaces/IAttestationCenter.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";

contract DeployAll is Script {
    // --- Configuration (Update if needed) ---
    // LayerZero Endpoints
    // address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x1a44076050125825900e736c501f859c50fE728c;
    address constant LZ_ENDPOINT_HOLESKY = 0x1a44076050125825900e736c501f859c50fE728c;
    address constant LZ_ENDPOINT_ARBITRUM_SEPOLIA = 0x1a44076050125825900e736c501f859c50fE728c;


    // Endpoint IDs (EIDs for LayerZero)
    // uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 30184; // Base Sepolia Endpoint ID
    uint32 constant HOLESKY_EID = 30101; // Holesky Endpoint ID
    uint32 constant ARBITRUM_SEPOLIA_EID = 30110; // Arbitrum Sepolia Endpoint ID
    // uint32 constant POLYGON_ARBITRUM_EID = 40267; // Polygon Amoy Endpoint ID
    // uint32 constant AVALANCHE_FUJI_EID = 40106; // Avalanche Fuji Endpoint ID
    // uint32 constant BNB_TESTNET_EID = 40102; // BNB Testnet Endpoint ID

    address constant JOB_REGISTRY_ADDRESS = 0xAf1189aFd1F1880F09AeC3Cbc32cf415c735C710; // proxy address on arbitrum
    address constant TRIGGER_GAS_REGISTRY_ADDRESS = 0x93dDB2307F3Af5df85F361E5Cddd898Acd3d132d; // proxy address on arbitrum
    address constant ATTESATION_CENTER_ADDRESS = 0x6DFee10D13d5B43AaF97bDA908C1D76d4313aF5f;
    address constant AVS_GOVERNANCE_LOGIC_ADDRESS = 0xc1F545Fe807b72429952dbFEFE8702658e4C5875;
    address constant HUB_IMPLEMENTATION_ADDRESS = 0x005a2Da42955E3837587D9cD87A19511F8041263;
    address constant SPOKE_IMPLEMENTATION_ADDRESS = 0x04730cEa1156832Ad18f4dA0c674EcC88d745eE8;

    bytes32 SALT = bytes32(keccak256(abi.encodePacked("triggerX_taskExecutionProxy_V1")));

    // Struct to hold network deployment information
    struct NetworkInfo {
        string name;
        string rpcEnvVar;
        address endpoint;
        uint32 eid;
    }

    // State variables to avoid stack depth issues
    address[] private operators;
    address private hubAddress;
    address[] private spokeAddresses;

    function deployHub(uint256 deployerPrivateKey, address deployerAddress) internal {
        console.log("\n=== Deploying TaskExecutionHub on Base Sepolia ===");

        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy the implementation contract
        // TaskExecutionHub hubImpl = new TaskExecutionHub(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress);
        // console.log("TaskExecutionHub implementation deployed at:", address(hubImpl));

        // 2. Prepare the initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionHub.initialize.selector,
            deployerAddress,                   // _ownerAddress
            HOLESKY_EID,                       // _srcEid
            BASE_SEPOLIA_EID,                  // _originEid
            operators,                         // _initialKeepers
            JOB_REGISTRY_ADDRESS,              // _jobRegistryAddress (random)
            TRIGGER_GAS_REGISTRY_ADDRESS       // _triggerGasRegistryAddress (random)
        );

        // 3. Prepare the proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(HUB_IMPLEMENTATION_ADDRESS, initData)
        );
        
        // 4. Deploy proxy using CREATE3
        hubAddress = CREATE3.deployDeterministic(proxyBytecode, SALT);
        console.log("TaskExecutionHub proxy deployed at:", hubAddress);

        // Note: setPeer and setOperator calls are moved to configureHub function
        // where we can handle cross-chain calls properly

        vm.stopBroadcast();
    }

    function configureHub(uint256 deployerPrivateKey) internal {
        vm.startBroadcast(deployerPrivateKey);

        TaskExecutionHub hub = TaskExecutionHub(payable(hubAddress));

        // Create spoke EIDs array
        uint32[] memory spokeEids = new uint32[](1);
        // spokeEids[0] = OP_SEPOLIA_EID;
        spokeEids[0] = ARBITRUM_SEPOLIA_EID;

        // Add all spoke endpoints to Hub
        hub.addSpokes(spokeEids);
        // console.log("Added spoke endpoint:", OP_SEPOLIA_EID, "(OP Sepolia)");
        console.log("Added spoke endpoint:", ARBITRUM_SEPOLIA_EID, "(Arbitrum Sepolia)");

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 0.0001 ether);
        // console.log("Sent 1 ETH to Hub contract at:", address(hub));

        // Note: Cross-chain configuration calls (setPeer, setOperator) should be handled separately
        // as they require interacting with contracts on different networks

        vm.stopBroadcast();
    }

    function deploySpoke(
        string memory networkName,
        string memory rpcEnvVar,
        address endpoint,
        uint256 deployerPrivateKey,
        address deployerAddress
    ) internal returns (address) {
        console.log(string.concat("\n=== Deploying TaskExecutionSpoke on ", networkName, " ==="));
        
        // Create fork for the network
        vm.createSelectFork(vm.envString(rpcEnvVar));
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy implementation
        // TaskExecutionSpoke spokeImpl = new TaskExecutionSpoke(endpoint, deployerAddress);
        // console.log(string.concat("TaskExecutionSpoke implementation on ", networkName, " deployed at:"), address(spokeImpl));

        // 2. Prepare initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionSpoke.initialize.selector,
            deployerAddress,    // _ownerAddress
            BASE_SEPOLIA_EID,   // _hubEid
            operators,          // _initialKeepers
            JOB_REGISTRY_ADDRESS,       // _jobRegistryAddress (random)
            TRIGGER_GAS_REGISTRY_ADDRESS        // _triggerGasRegistryAddress (random)
        );

        // 3. Prepare proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(SPOKE_IMPLEMENTATION_ADDRESS, initData)
        );
        
        // 4. Deploy proxy using CREATE3
        address spokeAddr = CREATE3.deployDeterministic(proxyBytecode, SALT);
        console.log(string.concat("TaskExecutionSpoke proxy deployed on ", networkName, " at:"), spokeAddr);

        // Note: setOperator call is commented out as it requires cross-chain interaction
        // This should be handled separately after deployment

        vm.stopBroadcast();
        
        return spokeAddr;
    }

    function deployAllSpokes(uint256 deployerPrivateKey, address deployerAddress) internal {
        spokeAddresses = new address[](1);
        
        // Deploy OP Sepolia spoke
        // spokeAddresses[0] = deploySpoke(
        //     "OP Sepolia",
        //     "OPSEPOLIA_RPC", 
        //     LZ_ENDPOINT_OP_SEPOLIA,
        //     deployerPrivateKey,
        //     deployerAddress
        // );

        // Deploy Arbitrum Sepolia spoke
        spokeAddresses[0] = deploySpoke(
            "Arbitrum",
            "ARBITRUM_RPC",
            LZ_ENDPOINT_ARBITRUM_SEPOLIA, 
            deployerPrivateKey,
            deployerAddress
        );
    }

    function printDeploymentSummary() internal view {
        console.log("\n--- Deployment Complete ---");
        console.log("Hub Address:", hubAddress);
        
        TaskExecutionHub hub = TaskExecutionHub(payable(hubAddress));
        console.log("Hub Owner:", hub.owner());
        
        // OP Sepolia spoke
        // TaskExecutionSpoke opSpoke = TaskExecutionSpoke(payable(spokeAddresses[0]));
        // console.log("OP Sepolia Spoke Address:", spokeAddresses[0]);
        // console.log("OP Sepolia Spoke Owner:", opSpoke.owner());
        
        // Arbitrum Sepolia spoke
        TaskExecutionSpoke arbSpoke = TaskExecutionSpoke(payable(spokeAddresses[0]));
        console.log("Arbitrum Sepolia Spoke Address:", spokeAddresses[0]);
        console.log("Arbitrum Sepolia Spoke Owner:", arbSpoke.owner());
        
        console.log("---------------------------");
    }

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // Create a fork using the Base Sepolia RPC to fetch operators
        vm.createSelectFork(vm.envString("BASE_RPC"));
        
        // Fetch operators from AttestationCenter
        // operators = fetchOperatorsFromAttestationCenter();
        // console.log("Successfully fetched", operators.length, "operators");

        // Deploy Hub on Base Sepolia
        deployHub(deployerPrivateKey, deployerAddress);
        
        // Configure Hub with spoke endpoints
        configureHub(deployerPrivateKey);

        // Deploy all spokes
        deployAllSpokes(deployerPrivateKey, deployerAddress);

        // Print final deployment summary
        printDeploymentSummary();
    }
} 