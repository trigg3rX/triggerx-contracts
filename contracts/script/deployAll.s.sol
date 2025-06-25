// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {IAttestationCenter} from "../src/interfaces/IAttestationCenter.sol";

contract DeployAll is Script {
    // --- Configuration (Update if needed) ---
    // LayerZero Endpoints
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_HOLESKY = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_ARBITRUM_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;

    address constant ATTESATION_CENTER_ADDRESS = 0x9725fB95B5ec36c062A49ca2712b3B1ff66F04eD;

    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID
    uint32 constant HOLESKY_EID = 40217; // Holesky Endpoint ID
    uint32 constant ARBITRUM_SEPOLIA_EID = 40231; // Arbitrum Sepolia Endpoint ID
    uint32 constant POLYGON_ARBITRUM_EID = 40267; // Polygon Amoy Endpoint ID
    uint32 constant AVALANCHE_FUJI_EID = 40106; // Avalanche Fuji Endpoint ID
    uint32 constant BNB_TESTNET_EID = 40102; // BNB Testnet Endpoint ID
    


    bytes32 SALT = bytes32(keccak256(abi.encodePacked(vm.envString("TESTNET_PRODUCTION_SALT"))));  // Production contract salt

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

    function fetchOperatorsFromAttestationCenter() internal returns (address[] memory) {
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
        console.log("\n=== Deploying ProxyHub on Base Sepolia ===");

        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxyHub.
        bytes memory hubBytecode = abi.encodePacked(
            type(ProxyHub).creationCode,
            abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress, HOLESKY_EID, BASE_SEPOLIA_EID, operators)
        );
        
        // Deploy using CREATE3 with no ETH value.
        hubAddress = CREATE3.deployDeterministic(hubBytecode, SALT);
        console.log("ProxyHub deployed at:", hubAddress);

        vm.stopBroadcast();
    }

    function configureHub(uint256 deployerPrivateKey) internal {
        vm.startBroadcast(deployerPrivateKey);

        ProxyHub hub = ProxyHub(payable(hubAddress));

        // Create spoke EIDs array
        uint32[] memory spokeEids = new uint32[](2);
        spokeEids[0] = OP_SEPOLIA_EID;
        spokeEids[1] = ARBITRUM_SEPOLIA_EID;

        // Add all spoke endpoints to Hub
        hub.addSpokes(spokeEids);
        console.log("Added spoke endpoint:", OP_SEPOLIA_EID, "(OP Sepolia)");
        console.log("Added spoke endpoint:", ARBITRUM_SEPOLIA_EID, "(Arbitrum Sepolia)");

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 1 ether);
        console.log("Sent 1 ETH to Hub contract at:", address(hub));

        vm.stopBroadcast();
    }

    function deploySpoke(
        string memory networkName,
        string memory rpcEnvVar,
        address endpoint,
        uint256 deployerPrivateKey,
        address deployerAddress
    ) internal returns (address) {
        console.log(string.concat("\n=== Deploying ProxySpoke on ", networkName, " ==="));
        
        // Create fork for the network
        vm.createSelectFork(vm.envString(rpcEnvVar));
        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxySpoke
        bytes memory spokeBytecode = abi.encodePacked(
            type(ProxySpoke).creationCode,
            abi.encode(endpoint, deployerAddress, BASE_SEPOLIA_EID, operators)
        );
        
        // Deploy using CREATE3 with the same salt
        address spokeAddr = CREATE3.deployDeterministic(spokeBytecode, SALT);
        console.log(string.concat("ProxySpoke deployed on ", networkName, " at:"), spokeAddr);

        vm.stopBroadcast();
        
        return spokeAddr;
    }

    function deployAllSpokes(uint256 deployerPrivateKey, address deployerAddress) internal {
        spokeAddresses = new address[](2);
        
        // Deploy OP Sepolia spoke
        spokeAddresses[0] = deploySpoke(
            "OP Sepolia",
            "OPSEPOLIA_RPC", 
            LZ_ENDPOINT_OP_SEPOLIA,
            deployerPrivateKey,
            deployerAddress
        );

        // Deploy Arbitrum Sepolia spoke
        spokeAddresses[1] = deploySpoke(
            "Arbitrum Sepolia",
            "ARBITRUM_SEPOLIA_RPC",
            LZ_ENDPOINT_ARBITRUM_SEPOLIA, 
            deployerPrivateKey,
            deployerAddress
        );
    }

    function printDeploymentSummary() internal view {
        console.log("\n--- Deployment Complete ---");
        console.log("Hub Address:", hubAddress);
        
        ProxyHub hub = ProxyHub(payable(hubAddress));
        console.log("Hub Owner:", hub.owner());
        
        // OP Sepolia spoke
        ProxySpoke opSpoke = ProxySpoke(payable(spokeAddresses[0]));
        console.log("OP Sepolia Spoke Address:", spokeAddresses[0]);
        console.log("OP Sepolia Spoke Owner:", opSpoke.owner());
        
        // Arbitrum Sepolia spoke
        ProxySpoke arbSpoke = ProxySpoke(payable(spokeAddresses[1]));
        console.log("Arbitrum Sepolia Spoke Address:", spokeAddresses[1]);
        console.log("Arbitrum Sepolia Spoke Owner:", arbSpoke.owner());
        
        console.log("---------------------------");
    }

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // Create a fork using the Base Sepolia RPC to fetch operators
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        
        // Fetch operators from AttestationCenter
        operators = fetchOperatorsFromAttestationCenter();
        console.log("Successfully fetched", operators.length, "operators");

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