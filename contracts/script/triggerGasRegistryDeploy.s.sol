// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerGasRegistry_Migrate} from "../src/TriggerGasRegistry_Migrate.sol";
import {CREATE3} from "lib/solady/src/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasRegistryDeploy is Script {
    function run() public {
        bytes32 SALT = keccak256(abi.encodePacked("TriggerGasRegistry"));

        bytes32 implementation_salt = keccak256(abi.encodePacked("ImplementationV1"));

        bytes memory implementation_code = type(TriggerGasRegistry_Migrate).creationCode;


        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        address implementation = CREATE3.deployDeterministic(implementation_code, implementation_salt);

        bytes memory initData = abi.encodeWithSelector(TriggerGasRegistry_Migrate.initialize.selector, deployer);
        bytes memory proxy_code = abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(address(implementation), initData));

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        console.log("Proxy deployed to:", proxy);
        console.log("Implementation deployed to:", implementation);
        console.log("Deployer:", deployer);
        console.log("Proxy owner:", TriggerGasRegistry_Migrate(payable(proxy)).owner());

        vm.stopBroadcast();
    }
}