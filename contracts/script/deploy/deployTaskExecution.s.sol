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
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_ARBITRUM_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;

    address constant ATTESATION_CENTER_ADDRESS = 0xC2F5cC5417330907c0Fa46ED94d87DED90508aBd;

    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID
    uint32 constant SEPOLIA_EID = 40161; // Sepolia Endpoint ID
    uint32 constant ARBITRUM_SEPOLIA_EID = 40231; // Arbitrum Sepolia Endpoint ID

    // These addresses will be updated after deployment
    address constant JOB_REGISTRY_ADDRESS = 0x309Bb6871d548E25e6051074A1bcE73199d8B647; // Will be updated
    address constant TRIGGER_GAS_REGISTRY_ADDRESS = 0xfe68C5213EfAF66AD620A4b5c48683cB5C876a09; // Will be updated
    // address constant AVS_GOVERNANCE_LOGIC_ADDRESS = 0x0000000000000000000000000000000000000000; // Will be updated
    
    bytes32 SALT = bytes32(keccak256(abi.encodePacked(vm.envString("TASK_EXECUTION_SALT"))));

    // Struct to hold network deployment information
    struct NetworkInfo {
        string name;
        string rpcEnvVar;
        address endpoint;
        uint32 eid;
    }

    // State variables to avoid stack depth issues
    address[] private operators;
    address[] private spokeAddresses;
    address private hubAddress;

    function fetchOperatorsFromAttestationCenter() internal view returns (address[] memory) {
        console.log("Attempting to fetch operators from AttestationCenter...");
        
        IAttestationCenter attestationCenter = IAttestationCenter(ATTESATION_CENTER_ADDRESS);
        
        // Use numOfOperators() which is defined in the interface
        try attestationCenter.numOfTotalOperators() returns (uint256 operatorCount) {
            console.log("Number of operators found:", operatorCount);
            
            if (operatorCount == 0) {
                revert("No operators found in AttestationCenter");
            }
            
            return fetchOperatorsByCount(attestationCenter, operatorCount);
            
        } catch Error(string memory reason) {
            console.log("numOfOperators failed:", reason);
        } catch {
            console.log("numOfOperators failed with unknown error");
        }
        
        // If all methods fail, revert with error
        revert("Unable to fetch operators from AttestationCenter. Please check the contract address and network.");
    }
    
    function fetchOperatorsByCount(IAttestationCenter attestationCenter, uint256 operatorCount) internal view returns (address[] memory) {
        address[] memory tempOperators = new address[](operatorCount);
        uint256 validOperators = 0;
        
        for (uint256 i = 1; i <= operatorCount; i++) {
            try attestationCenter.getOperatorPaymentDetail(i) returns (IAttestationCenter.PaymentDetails memory details) {
                if (details.operator != address(0)) {
                    tempOperators[validOperators] = details.operator;
                    console.log("Operator", i, ":", details.operator);
                    validOperators++;
                }
            } catch {
                console.log("Failed to fetch operator", i, "- skipping");
            }
        }
        
        // Resize array to only include valid operators
        if (validOperators < operatorCount) {
            address[] memory validOperatorArray = new address[](validOperators);
            for (uint256 i = 0; i < validOperators; i++) {
                validOperatorArray[i] = tempOperators[i];
            }
            return validOperatorArray;
        }
        
        return tempOperators;
    }

    function deployHub(uint256 deployerPrivateKey, address deployerAddress) internal {
        console.log("\n=== Deploying TaskExecutionHub on Base Sepolia ===");

        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy the implementation contract
        TaskExecutionHub hubImpl = new TaskExecutionHub(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress);
        console.log("TaskExecutionHub implementation deployed at:", address(hubImpl));

        // 2. Prepare the initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionHub.initialize.selector,
            deployerAddress,                   // _ownerAddress
            SEPOLIA_EID,                       // _srcEid (Sepolia)
            BASE_SEPOLIA_EID,                  // _originEid
           operators,                         // _initialKeepers
            JOB_REGISTRY_ADDRESS,              // _jobRegistryAddress (will be updated)
            TRIGGER_GAS_REGISTRY_ADDRESS       // _triggerGasRegistryAddress (will be updated)
        );

        // 3. Prepare the proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(address(hubImpl), initData)
        );
        
        // 4. Deploy proxy using CREATE3
        hubAddress = CREATE3.deployDeterministic(proxyBytecode, SALT);
        console.log("TaskExecutionHub proxy deployed at:", hubAddress);

        vm.stopBroadcast();
    }

    function configureHub(uint256 deployerPrivateKey) internal {
        vm.startBroadcast(deployerPrivateKey);

        TaskExecutionHub hub = TaskExecutionHub(payable(hubAddress));

        // Create spoke EIDs array
        uint32[] memory spokeEids = new uint32[](1);
        spokeEids[0] = ARBITRUM_SEPOLIA_EID;

        // Add Arbitrum Sepolia spoke endpoint to Hub
        hub.addSpokes(spokeEids);
        console.log("Added spoke endpoint:", ARBITRUM_SEPOLIA_EID, "(Arbitrum Sepolia)");

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
        TaskExecutionSpoke spokeImpl = new TaskExecutionSpoke(endpoint, deployerAddress);
        console.log(string.concat("TaskExecutionSpoke implementation on ", networkName, " deployed at:"), address(spokeImpl));

        // 2. Prepare initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionSpoke.initialize.selector,
            deployerAddress,    // _ownerAddress
            BASE_SEPOLIA_EID,   // _hubEid
            operators,          // _initialKeepers
            JOB_REGISTRY_ADDRESS,       // _jobRegistryAddress (will be updated)
            TRIGGER_GAS_REGISTRY_ADDRESS        // _triggerGasRegistryAddress (will be updated)
        );

        // 3. Prepare proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(address(spokeImpl), initData)
        );
        
        // 4. Deploy proxy using CREATE3
        address spokeAddr = CREATE3.deployDeterministic(proxyBytecode, SALT);
        console.log(string.concat("TaskExecutionSpoke proxy deployed on ", networkName, " at:"), spokeAddr);

        vm.stopBroadcast();
        
        return spokeAddr;
    }

    function deployAllSpokes(uint256 deployerPrivateKey, address deployerAddress) internal {
        spokeAddresses = new address[](1);
        
        // Deploy Arbitrum Sepolia spoke
        spokeAddresses[0] = deploySpoke(
            "Arbitrum Sepolia",
            "ARBITRUM_SEPOLIA_RPC",
            LZ_ENDPOINT_ARBITRUM_SEPOLIA, 
            deployerPrivateKey,
            deployerAddress
        );
    }

    function printDeploymentSummary() internal view {
        console.log("\n--- Deployment Complete ---");
        // console.log("Hub Address:", hubAddress);
        
        // TaskExecutionHub hub = TaskExecutionHub(payable(hubAddress));
        // console.log("Hub Owner:", hub.owner());
        
        // Arbitrum Sepolia spoke
        TaskExecutionSpoke arbSpoke = TaskExecutionSpoke(payable(spokeAddresses[0]));
        console.log("Arbitrum Sepolia Spoke Address:", spokeAddresses[0]);
        console.log("Arbitrum Sepolia Spoke Owner:", arbSpoke.owner());
        
        console.log("---------------------------");
        console.log("IMPORTANT: Update the following addresses in your scripts:");
        console.log("JOB_REGISTRY_ADDRESS:", JOB_REGISTRY_ADDRESS);
        console.log("TRIGGER_GAS_REGISTRY_ADDRESS:", TRIGGER_GAS_REGISTRY_ADDRESS);
        //console.log("AVS_GOVERNANCE_LOGIC_ADDRESS:", AVS_GOVERNANCE_LOGIC_ADDRESS);
    }

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // Create a fork using the Base Sepolia RPC to fetch operators
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        
        // Fetch operators from AttestationCenter
        operators = [0x011FCbAE5f306cd793456ab7d4c0CC86756c693D];
        console.log("Successfully fetched", operators.length, "operators");

        // Deploy Hub on Base Sepolia
        // deployHub(deployerPrivateKey, deployerAddress);
        
        // // Configure Hub with spoke endpoints
        // configureHub(deployerPrivateKey);

        // Deploy all spokes
        deployAllSpokes(deployerPrivateKey, deployerAddress);

        // Print final deployment summary
        printDeploymentSummary();
    }
} 