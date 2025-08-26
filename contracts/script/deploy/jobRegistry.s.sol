// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {JobRegistry} from "../../src/JobRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";

contract DeployJobRegistry is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        // Create fork for Arbitrum Sepolia
        string memory rpcUrl = vm.envString("ARBITRUM_SEPOLIA_RPC");
        require(bytes(rpcUrl).length > 0, "ARBITRUM_SEPOLIA_RPC not set");
        vm.createSelectFork(rpcUrl);
        
        vm.startBroadcast(deployerPrivateKey);

        // Validate salt is set
        string memory saltString = vm.envString("JR_SALT");
        require(bytes(saltString).length > 0, "JR_SALT not set");
        bytes32 salt = keccak256(abi.encode(saltString));

        // Deploy the implementation contract
        JobRegistry implementation = new JobRegistry();
        console.log("JobRegistry implementation deployed at:", address(implementation));

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector,
            deployer // Set deployer as initial owner
        );

        bytes memory proxyCode = abi.encodePacked(
            type(ERC1967Proxy).creationCode, 
            abi.encode(address(implementation), initData)
        );
        
        address proxy = CREATE3.deployDeterministic(proxyCode, salt);
        
        console.log("JobRegistry proxy deployed at:", proxy);
        console.log("Initial owner:", deployer);

        vm.stopBroadcast();

        // Verify the deployment
        JobRegistry proxyContract = JobRegistry(proxy);
        require(proxyContract.owner() == deployer, "Owner not set correctly");
        require(proxyContract.getJobCounter() == 0, "Job counter not initialized");

        // Log deployment info
        console.log("=== JobRegistry Deployment Summary ===");
        console.log("Network: Arbitrum Sepolia");
        console.log("RPC URL:", rpcUrl);
        console.log("Implementation:", address(implementation));
        console.log("Proxy:", proxy);
        console.log("Owner:", deployer);
        console.log("Salt used:", saltString);
        console.log("========================================");
    }

    function upgradeJobRegistry(address proxyAddress, address newImplementation) external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        // Create fork for Arbitrum Sepolia
        string memory rpcUrl = vm.envString("ARBITRUM_SEPOLIA_RPC");
        require(bytes(rpcUrl).length > 0, "ARBITRUM_SEPOLIA_RPC not set");
        vm.createSelectFork(rpcUrl);
        
        vm.startBroadcast(deployerPrivateKey);

        // Verify the proxy address is valid
        require(proxyAddress != address(0), "Invalid proxy address");
        require(newImplementation != address(0), "Invalid implementation address");

        JobRegistry jobRegistry = JobRegistry(proxyAddress);
        
        // Verify the caller is the owner
        require(jobRegistry.owner() == deployer, "Caller is not the owner");

        jobRegistry.upgradeToAndCall(newImplementation, "");

        console.log("JobRegistry upgraded to:", newImplementation);
        console.log("Proxy address:", proxyAddress);
        console.log("Upgrade performed by:", deployer);

        vm.stopBroadcast();
    }
}