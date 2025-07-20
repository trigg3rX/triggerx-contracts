// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.26;

// import {Script} from "forge-std/Script.sol";
// import {console} from "forge-std/console.sol";
// import {console2} from "forge-std/console2.sol";
// import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";

// contract TriggerGasRegistryRollback is Script {
//     function run() public {
//         uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
//         address deployer = vm.addr(deployerPrivateKey);
        
//         address proxyAddress = vm.envAddress("PROXY_ADDRESS");
//         address previousImplementation = vm.envAddress("PREVIOUS_IMPLEMENTATION");
//         address operator = vm.envAddress("OPERATOR_ADDRESS");
//         uint256 tgPerEth = vm.envUint("TG_PER_ETH");
        
//         address payable proxy = payable(proxyAddress);
        
//         console.log("\n=== ROLLBACK TRIGGER GAS REGISTRY ===");
//         console.log("Proxy address:", proxyAddress);
//         console.log("Current implementation:", getImplementation(proxy));
//         console.log("Rolling back to:", previousImplementation);
        
//         // Verify the previous implementation exists
//         require(previousImplementation != address(0), "Previous implementation cannot be zero");
        
//         // Check current state before rollback
//         console.log("\n=== BEFORE ROLLBACK ===");
//         TriggerGasRegistry registry = TriggerGasRegistry(proxy);
//         console.log("Current owner:", registry.owner());
//         console.log("Current operator:", registry.operatorRole());
//         console.log("Current TG_PER_ETH:", registry.TG_PER_ETH());
//         console.log("Contract balance:", address(proxy).balance);
        
//         vm.startBroadcast(deployerPrivateKey);
        
//         // Perform rollback
//         console.log("\n=== PERFORMING ROLLBACK ===");
//         registry.upgradeToAndCall(previousImplementation, "");
        
//         // Re-initialize critical parameters
//         registry.setOperator(operator);
//         registry.setTGPerETH(tgPerEth);
        
//         vm.stopBroadcast();
        
//         // Verify rollback
//         console.log("\n=== AFTER ROLLBACK ===");
//         address newImplementation = getImplementation(proxy);
//         console.log("New implementation:", newImplementation);
//         console.log("New owner:", registry.owner());
//         console.log("New operator:", registry.operatorRole());
//         console.log("New TG_PER_ETH:", registry.TG_PER_ETH());
        
//         if (newImplementation == previousImplementation) {
//             console.log("✅ ROLLBACK SUCCESSFUL: Implementation reverted to previous version");
//         } else {
//             console.log("❌ ROLLBACK FAILED: Implementation not updated correctly");
//             revert("Rollback failed");
//         }
        
//         console.log("\n=== ROLLBACK SUMMARY ===");
//         console.log("Proxy address:", proxyAddress);
//         console.log("Previous implementation:", previousImplementation);
//         console.log("Current implementation:", newImplementation);
//         console.log("Deployer:", deployer);
//     }
    
//     function getImplementation(address proxy) internal view returns (address) {
//         bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
//         bytes32 impl;
//         assembly {
//             impl := sload(slot)
//         }
//         return address(uint160(uint256(impl)));
//     }
// } 