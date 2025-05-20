// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import "forge-std/console.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {ITriggerGasRegistry} from "../src/interfaces/ITriggerGasRegistry.sol";

import "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasDeployerOPS is Script {
    function writeDeploymentToJson(address implementation, address proxy) internal {
        string memory jsonString = string.concat(
            "{",
            '"triggerGasRegistry": {',
            '"proxy": "', vm.toString(proxy), '",',
            '"implementation": "', vm.toString(implementation), '"',
            '}',
            "}"
        );

        vm.writeFile("./script/output/gas.opsepolia.json", jsonString);
    }

    function run() external returns (address) {
        string memory config = vm.readFile("./script/input/parameters.holesky.json");
        address triggerXOwner = abi.decode(vm.parseJson(config, ".triggerXOwner"), (address));
        
        vm.startBroadcast();

        TriggerGasRegistry implementation = new TriggerGasRegistry();
        
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector,
            triggerXOwner
        );

        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            initData
        );

        vm.stopBroadcast();

        console.log("TriggerGasRegistry implementation deployed to:", address(implementation));
        console.log("Proxy contract deployed to:", address(proxy));

        writeDeploymentToJson(address(implementation), address(proxy));

        return address(proxy);
    }
}
