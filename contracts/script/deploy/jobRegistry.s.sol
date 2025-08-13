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
        
        vm.startBroadcast(deployerPrivateKey);

        bytes32 salt = keccak256(abi.encode("triggerX_eigenlayer_jobRegistry_V1"));

        // Deploy the implementation contract
        JobRegistry implementation = new JobRegistry();
        console.log("JobRegistry implementation deployed at:", address(implementation));

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector,
            deployer // Set deployer as initial owner
        );

        bytes memory proxy_code = abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(address(implementation), initData));
        
        address proxy = CREATE3.deployDeterministic(proxy_code, salt);
        
        console.log("JobRegistry proxy deployed at:", proxy);
        console.log("Initial owner:", deployer);

        vm.stopBroadcast();

        // Log deployment info
        console.log("=== JobRegistry Deployment Summary ===");
        console.log("Implementation:", address(implementation));
        console.log("Proxy:", proxy);
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