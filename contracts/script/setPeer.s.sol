// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/lz/ProxyHub.sol";

contract SetPeerScript is Script {
    function run() external {

        uint32 BASE_SEPOLIA_EID = 40245;
        // Load environment variables
        string memory baseSepoliaRpc = vm.envString("BASE_SEPOLIA_RPC");
        
        // Fork Base Sepolia
        vm.createSelectFork(baseSepoliaRpc);
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        
        // ProxyHub address on Base Sepolia
        address proxyHub = 0x96c5F575940DBe98fd9600F74F0c36139A7Be2A2;
        
        // Governance Logic address
        address governanceLogic = 0x2cFADBDd050bB83A45B80C2a045516470A67c991;
        
        // Start broadcasting transactions
        // vm.startBroadcast();
        
        // Set peer on ProxyHub
        ProxyHub(payable(proxyHub)).setPeer(BASE_SEPOLIA_EID, governanceLogic);
        
        // Stop broadcasting
        vm.stopBroadcast();
    }
}
