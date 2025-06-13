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

    // Deployment Salt (Must be the same for Hub and Spoke)
    // bytes32 constant SALT = "triggerXKeepers"; // Production contract salt
    bytes32 constant SALT = "test_triggerXKeepers"; 

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
        address[] memory operators = new address[](operatorCount);
        uint256 validOperators = 0;
        
        for (uint256 i = 1; i <= operatorCount; i++) {
            try attestationCenter.getOperatorPaymentDetail(i) returns (IAttestationCenter.PaymentDetails memory details) {
                if (details.operator != address(0)) {
                    operators[validOperators] = details.operator;
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
                validOperatorArray[i] = operators[i];
            }
            return validOperatorArray;
        }
        
        return operators;
    }
    

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // Create a fork using the Base Sepolia RPC to fetch operators
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        
        // Fetch operators from AttestationCenter
        address[] memory initialKeepers = fetchOperatorsFromAttestationCenter();

        console.log("Successfully fetched", initialKeepers.length, "operators");

        // --- Deploy Hub on Base Sepolia ---
        console.log("\n=== Deploying ProxyHub on Base Sepolia ===");

        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxyHub.
        // ProxyHub's constructor takes: (address _endpoint, address _owner, uint32 _srcEid, uint32 _thisChainEid, address[] memory _initialKeepers)
        bytes memory hubBytecode = abi.encodePacked(
            type(ProxyHub).creationCode,
            abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress, HOLESKY_EID, BASE_SEPOLIA_EID, initialKeepers)
        );
        
        // Deploy using CREATE3 with no ETH value.
        address hubAddr = CREATE3.deployDeterministic(hubBytecode, SALT);
        ProxyHub hub = ProxyHub(payable(hubAddr));
        console.log("ProxyHub deployed at:", hubAddr);

        // Add spoke endpoints to Hub
        uint32[] memory spokeEids = new uint32[](2);
        spokeEids[0] = OP_SEPOLIA_EID;
        spokeEids[1] = ARBITRUM_SEPOLIA_EID;
        hub.addSpokes(spokeEids);
        console.log("Added spoke endpoint:", OP_SEPOLIA_EID);
        console.log("Added spoke endpoint:", ARBITRUM_SEPOLIA_EID);

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 1 ether);
        console.log("Sent 1 ETH to Hub contract at:", address(hub));

        vm.stopBroadcast();

        // --- Deploy Spoke on OP Sepolia ---
        console.log("\n=== Deploying ProxySpoke on OP Sepolia ===");

        // Create a fork using the OP Sepolia RPC.
        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxySpoke.
        // ProxySpoke's constructor takes: (address _endpoint, address _owner, uint32 _srcEid, address[] memory _initialKeepers)
        bytes memory spokeBytecode1 = abi.encodePacked(
            type(ProxySpoke).creationCode,
            abi.encode(LZ_ENDPOINT_OP_SEPOLIA, deployerAddress, BASE_SEPOLIA_EID, initialKeepers)
        );
        
        // Deploy using CREATE3 with the same salt.
        address spokeAddr1 = CREATE3.deployDeterministic(spokeBytecode1, SALT);
        ProxySpoke spoke1 = ProxySpoke(payable(spokeAddr1));
        console.log("ProxySpoke1 deployed at:", spokeAddr1);

        vm.stopBroadcast();

        // --- Deploy Spoke on Arbitrum Sepolia ---
        console.log("\n=== Deploying ProxySpoke on Arbitrum Sepolia ===");

        // Create a fork using the Arbitrum Sepolia RPC.
        // vm.createSelectFork(vm.envString("ARBITRUM_SEPOLIA_RPC"));
        // vm.startBroadcast(deployerPrivateKey);
        
        // // Prepare the bytecode for ProxySpoke.
        // // ProxySpoke's constructor takes: (address _endpoint, address _owner, uint32 _srcEid, address[] memory _initialKeepers)
        // bytes memory spokeBytecode2 = abi.encodePacked(
        //     type(ProxySpoke).creationCode,
        //     abi.encode(LZ_ENDPOINT_ARBITRUM_SEPOLIA, deployerAddress, BASE_SEPOLIA_EID, initialKeepers)
        // );
        
        // // Deploy using CREATE3 with the same salt.
        // address spokeAddr2 = CREATE3.deployDeterministic(spokeBytecode2, SALT);
        // ProxySpoke spoke2 = ProxySpoke(payable(spokeAddr2));
        // console.log("ProxySpoke2 deployed at:", spokeAddr2);

        // vm.stopBroadcast();

        // --- Verify final state ---
        console.log("\n--- Deployment Complete ---");
  
        // console.log("Hub Address:", address(hub));
        // console.log("Hub Owner:", hub.owner());
        // console.log("Hub AVSGovernance:", deployerAddress);
        // console.log("Spoke1 Address:", address(spoke1));
        // console.log("Spoke1 Owner:", spoke1.owner());
        // console.log("Spoke2 Address:", address(spoke2));
        // console.log("Spoke2 Owner:", spoke2.owner());
        console.log("---------------------------");
    }
} 