// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {console2} from "forge-std/console2.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasRegistryDeploy is Script {
    struct UserBalance {
        address user;
        uint256 ethSpent;
        uint256 tgBalance;
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        address operator = 0xd2B4F73FE4c747716F20839c37C451f241226b03;
        address triggerGasRegistryAddress = 0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a;
        
        // Define users to check
        address[] memory users = new address[](6);
        users[0] = 0xD5E9061656252a0b44D98C6944B99046FDDf49cA;
        users[1] = 0xBc0435FB4f37345a420fbD09e5700f3A72fd0534;
        users[2] = 0xC76EA60887CA82C474cf6dfc17f918DDd68D6cA2;
        users[3] = 0x751F7CB39462FFc9cbE0276D411EbA4a2a7Ebe46;
        users[4] = 0xc073A5E091DC60021058346b10cD5A9b3F0619fE;
        users[5] = 0xE3304AB782c3272D1D7964ba7043e11Fd7ecFEB8;

        address payable proxy = payable(triggerGasRegistryAddress);

        // STEP 1: Capture BEFORE state
        console.log("\n=== BEFORE UPGRADE ===");
        UserBalance[] memory beforeBalances = new UserBalance[](users.length);
        address oldImplementation = getImplementation(proxy);
        console.log("Current implementation:", oldImplementation);
        console.log("Current owner:", TriggerGasRegistry(proxy).owner());
        
        for (uint256 i = 0; i < users.length; i++) {
            (uint256 ethSpent, uint256 tgBalance) = TriggerGasRegistry(proxy).balances(users[i]);
            beforeBalances[i] = UserBalance({
                user: users[i],
                ethSpent: ethSpent,
                tgBalance: tgBalance
            });
            console2.log("User %s: ETH=%s, TG=%s", users[i], ethSpent, tgBalance);
        }

        vm.startBroadcast(deployerPrivateKey);

        // STEP 2: Deploy new implementation
        TriggerGasRegistry newImplementation = new TriggerGasRegistry();
        console.log("\n=== PERFORMING UPGRADE ===");
        console.log("Deploying new implementation to:", address(newImplementation));
        
        // STEP 3: Perform upgrade
        TriggerGasRegistry(proxy).upgradeToAndCall(address(newImplementation), "");
        TriggerGasRegistry(proxy).setOperator(operator);
        TriggerGasRegistry(proxy).setTGPerETH(1000);

        vm.stopBroadcast();

        // STEP 4: Capture AFTER state
        console.log("\n=== AFTER UPGRADE ===");
        address finalImplementation = getImplementation(proxy);
        console.log("New implementation:", finalImplementation);
        console.log("New owner:", TriggerGasRegistry(proxy).owner());
        console.log("New operator:", TriggerGasRegistry(proxy).operatorRole());

        UserBalance[] memory afterBalances = new UserBalance[](users.length);
        for (uint256 i = 0; i < users.length; i++) {
            (uint256 ethSpent, uint256 tgBalance) = TriggerGasRegistry(proxy).balances(users[i]);
            afterBalances[i] = UserBalance({
                user: users[i],
                ethSpent: ethSpent,
                tgBalance: tgBalance
            });
            console2.log("User %s: ETH=%s, TG=%s", users[i], ethSpent, tgBalance);
        }

        // STEP 5: Compare BEFORE vs AFTER
        console.log("\n=== COMPARISON RESULTS ===");
        bool allDataPreserved = true;
        uint256 usersWithData = 0;
        
        for (uint256 i = 0; i < users.length; i++) {
            bool ethMatches = beforeBalances[i].ethSpent == afterBalances[i].ethSpent;
            bool tgMatches = beforeBalances[i].tgBalance == afterBalances[i].tgBalance;
            
            if (beforeBalances[i].ethSpent > 0 || beforeBalances[i].tgBalance > 0) {
                usersWithData++;
            }
            
            if (ethMatches && tgMatches) {
                console2.log("[OK] User %s: Data preserved", users[i]);
            } else {
                console2.log("[ERROR] User %s: Data changed!", users[i]);
                console2.log("  ETH: %s -> %s", beforeBalances[i].ethSpent, afterBalances[i].ethSpent);
                console2.log("  TG: %s -> %s", beforeBalances[i].tgBalance, afterBalances[i].tgBalance);
                allDataPreserved = false;
            }
        }

        console.log("\n=== FINAL VERIFICATION ===");
        console2.log("Total users checked: %s", users.length);
        console2.log("Users with balance data: %s", usersWithData);
        console2.log("Implementation changed: %s", oldImplementation != finalImplementation);
        
        if (allDataPreserved && usersWithData > 0) {
            console.log("SUCCESS: All user data preserved during upgrade!");
        } else if (allDataPreserved && usersWithData == 0) {
            console.log("WARNING: No user data found (might be expected)");
        } else {
            console.log("CRITICAL: Some user data was lost during upgrade!");
        }

        console.log("\n=== DEPLOYMENT SUMMARY ===");
        console.log("Proxy address:", proxy);
        console.log("Old implementation:", oldImplementation);
        console.log("New implementation:", finalImplementation);
        console.log("Deployer:", deployer);
    }

    function getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        bytes32 impl;
        assembly {
            impl := sload(slot)
        }
        return address(uint160(uint256(impl)));
    }
}