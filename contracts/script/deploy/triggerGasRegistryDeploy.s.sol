// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";
import {CREATE3} from "lib/solady/src/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasRegistryDeploy is Script {
    function run() public {
        // Load and validate environment variables
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        string memory rpcUrl = vm.envString("ARBITRUM_SEPOLIA_RPC");
        string memory saltString = vm.envString("TGR_SALT");
        
        require(bytes(rpcUrl).length > 0, "ARBITRUM_SEPOLIA_RPC not set");
        require(bytes(saltString).length > 0, "TGR_SALT not set");
        require(deployer != address(0), "Invalid deployer address");

        // Create fork
        vm.createSelectFork(rpcUrl);
        vm.startBroadcast(deployerPrivateKey);

        // Generate unique salts
        bytes32 implementationSalt = keccak256(abi.encodePacked(saltString, "ImplementationV1"));
        bytes32 proxySalt = keccak256(abi.encodePacked(saltString, "Proxy"));

        // 1. Deploy Implementation
        address implementation = CREATE3.deployDeterministic(
            type(TriggerGasRegistry).creationCode,
            implementationSalt
        );
        require(implementation != address(0), "Implementation deployment failed");
        console.log("Implementation deployed at:", implementation);

        // 2. Prepare Proxy
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector,
            deployer,    // owner
            deployer,    // operator
            1000        // TG_PER_ETH
        );

        // 3. Deploy Proxy
        address proxy = CREATE3.deployDeterministic(
            abi.encodePacked(
                type(ERC1967Proxy).creationCode,
                abi.encode(implementation, initData)
            ),
            proxySalt
        );
        require(proxy != address(0), "Proxy deployment failed");
        console.log("Proxy deployed at:", proxy);

        // Verify deployment
        TriggerGasRegistry registry = TriggerGasRegistry(payable(proxy));
        require(registry.owner() == deployer, "Owner not set correctly");
        require(registry.operatorRole() == deployer, "Operator not set correctly");
        require(registry.TG_PER_ETH() == 1000, "TG rate not set correctly");

        vm.stopBroadcast();

        // Deployment summary
        console.log("\n=== Deployment Successful ===");
        console.log("Network: Arbitrum Sepolia");
        console.log("Deployer:", deployer);
        console.log("Implementation:", implementation);
        console.log("Proxy:", proxy);
        console.log("Salt Used:", saltString);
        console.log("============================\n");
    }
}