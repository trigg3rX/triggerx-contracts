// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {JobRegistry} from "../../src/JobRegistry.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployJobRegistry is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    bytes32 SALT = keccak256(abi.encodePacked("put_salt_here"));
    bytes32 IMPL_SALT = keccak256(abi.encodePacked("put_salt_here"));

    function run() public {
        // Create fork for this chain
        // vm.createSelectFork(vm.envString("BASE_RPC"));
        // vm.createSelectFork(vm.envString("OP_RPC"));
        vm.createSelectFork(vm.envString("ARB_RPC"));

        bytes memory implementation_code = type(JobRegistry).creationCode;

        vm.startBroadcast(deployerPrivateKey);

        // Deploy implementation
        address implementation = CREATE3.deployDeterministic(implementation_code, IMPL_SALT);

        // Deploy proxy
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector, 
            deployer
        );
        bytes memory proxy_code = abi.encodePacked(
            type(ERC1967Proxy).creationCode, 
            abi.encode(address(implementation), initData)
        );

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        // Verify deployment
        JobRegistry registry = JobRegistry(payable(proxy));
        require(registry.owner() == deployer, "Owner mismatch");

        vm.stopBroadcast();

        // Print summary
        console.log("");
        console.log("=== JOBREGISTRY DEPLOYMENT SUMMARY ===");
        console.log("Proxy:", proxy);
        console.log("Implementation:", implementation);
        console.log("Deployer:", deployer);
        console.log("");
    }
}
