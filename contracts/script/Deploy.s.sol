// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Script} from "forge-std/Script.sol";
import {CounterContract} from "../src/test.sol";

contract Deploy is Script {
    function run() external returns (CounterContract) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        CounterContract counter = new CounterContract();
        
        vm.stopBroadcast();
        
        return counter;
    }
} 