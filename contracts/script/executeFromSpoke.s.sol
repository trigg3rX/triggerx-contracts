// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {Target} from "../src/Target.sol";

contract ExecuteFromSpoke is Script {
    // Deployed contract addresses
    address constant SPOKE_ADDRESS = 0xc6e9cc75D299C0a16f1c2376Dc5De7ebcFcE55E4;
    address constant TARGET_ADDRESS = 0x2C34442fD5D21C0cE207DABA53EDE46038D02C77;

    function run() external {
        uint256 keeperPrivateKey = vm.envUint("KEEPER_PRIVATE_KEY");
        address keeperAddress = vm.addr(keeperPrivateKey);
        console.log("Executing from Spoke using keeper:", keeperAddress);

        vm.startBroadcast(keeperPrivateKey);

        // Create an instance of the Spoke contract
        ProxySpoke spoke = ProxySpoke(payable(SPOKE_ADDRESS));

        // Define the value to set in the Target contract
        uint256 valueToSet = 12345;
        bytes memory targetCalldata = abi.encodeWithSignature("setValue(uint256)", valueToSet);

        // Execute the function on the Spoke contract
        spoke.executeFunction{value: 0.001 ether}(TARGET_ADDRESS, targetCalldata);
        console.log("Executed function on Spoke to set value in Target contract.");

        vm.stopBroadcast();
    }
}