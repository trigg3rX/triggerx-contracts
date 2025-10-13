// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerXSafeModule} from "../../src/TriggerXSafeModule.sol";

contract DeployTriggerXSafeModule is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);
    address taskExecutionHub = 0x179c62e83c3f90981B65bc12176FdFB0f2efAD54;

    function run() public {
        vm.startBroadcast(deployerPrivateKey);
        bytes32 salt = keccak256(abi.encodePacked("triggerx-safe-module", deployer));
        TriggerXSafeModule module = new TriggerXSafeModule{salt: salt}(taskExecutionHub);
        console.log("TriggerXSafeModule deployed at:", payable(address(module)));
        vm.stopBroadcast();
    }
}