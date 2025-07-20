// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {JobRegistry} from "../../src/JobRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployJobRegistry is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        vm.startBroadcast(deployerPrivateKey);

        // Deploy the implementation contract
        JobRegistry implementation = new JobRegistry();
        console.log("JobRegistry implementation deployed at:", address(implementation));

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector,
            deployer // Set deployer as initial owner
        );

        // Deploy the proxy contract
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            initData
        );
        
        console.log("JobRegistry proxy deployed at:", address(proxy));
        console.log("Initial owner:", deployer);

        vm.stopBroadcast();

        // Log deployment info
        console.log("=== JobRegistry Deployment Summary ===");
        console.log("Implementation:", address(implementation));
        console.log("Proxy:", address(proxy));
        console.log("Owner:", deployer);
        console.log("========================================");
    }

    function upgradeJobRegistry(address proxyAddress, address newImplementation) external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);

        JobRegistry jobRegistry = JobRegistry(proxyAddress);
        jobRegistry.upgradeToAndCall(newImplementation, "");

        console.log("JobRegistry upgraded to:", newImplementation);
        console.log("Proxy address:", proxyAddress);

        vm.stopBroadcast();
    }
} 