// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployTriggerGasRegistry is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);
    address operator = vm.envAddress("TASK_EXECUTION_ADDRESS");
    uint256 tgPerEth = vm.envUint("TG_PER_ETH");

    bytes32 SALT = keccak256(abi.encodePacked(vm.envString("GAS_REGISTRY_SALT")));
    bytes32 IMPL_SALT = keccak256(abi.encodePacked(vm.envString("GAS_REGISTRY_IMPL_SALT")));

    function run() public {
        // Create fork for this chain using environment variable
        string memory rpcUrl = vm.envString("RPC_URL");
        vm.createSelectFork(rpcUrl);

        bytes memory implementation_code = type(TriggerGasRegistry).creationCode;

        vm.startBroadcast(deployerPrivateKey);

        // Deploy implementation
        address implementation = CREATE3.deployDeterministic(implementation_code, IMPL_SALT);

        // Deploy proxy
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector, 
            deployer, 
            operator, 
            tgPerEth
        );
        bytes memory proxy_code = abi.encodePacked(
            type(ERC1967Proxy).creationCode, 
            abi.encode(address(implementation), initData)
        );

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        // Verify deployment
        TriggerGasRegistry registry = TriggerGasRegistry(payable(proxy));
        require(registry.owner() == deployer, "Owner mismatch");
        require(registry.operatorRole() == operator, "Operator mismatch");
        require(registry.TG_PER_ETH() == tgPerEth, "TG_PER_ETH mismatch");

        vm.stopBroadcast();

        // Print summary
        console.log("");
        console.log("=== TRIGGER GAS REGISTRY DEPLOYMENT SUMMARY ===");
        console.log("Proxy:", proxy);
        console.log("Implementation:", implementation);
        console.log("Deployer:", deployer);
        console.log("");
    }
}

contract SetOperatorOnGasRegistry is Script {
    function run() public {
        string memory rpcUrl = vm.envString("RPC_URL");
        vm.createSelectFork(rpcUrl);
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        TriggerGasRegistry registry = TriggerGasRegistry(vm.envAddress("GAS_REGISTRY_ADDRESS"));
        registry.setOperator(vm.envAddress("TASK_EXECUTION_ADDRESS"));
    }
}
