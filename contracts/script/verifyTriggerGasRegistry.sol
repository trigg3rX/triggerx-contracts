// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

contract VerifyTriggerGasRegistry is Script {
    function run() public {
        address proxy = 0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a;
        address implementation = 0x9F452702490e11Af38C213De30c000A46E4dA399;

        TriggerGasRegistry triggerGasRegistry = TriggerGasRegistry(proxy);

        // Get deployer private key
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        // Start broadcasting as deployer
        vm.startBroadcast(deployerPrivateKey);

        // Test 1: Verify owner
        console2.log("Owner:", triggerGasRegistry.owner());
        console2.log("Deployer:", deployer);
        require(triggerGasRegistry.owner() == deployer, "Owner mismatch");

        // Test 2: Purchase TG tokens
        uint256 purchaseAmount = 0.1 ether;
        triggerGasRegistry.purchaseTG{value: purchaseAmount}(purchaseAmount);
        
        // Verify balance after purchase
        (uint256 ethSpent, uint256 tgBalance) = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("ETH spent: ", Strings.toString(ethSpent)));
        console2.log(string.concat("TG balance: ", Strings.toString(tgBalance)));
        require(ethSpent == purchaseAmount, "ETH spent mismatch");
        require(tgBalance == purchaseAmount * 1000, "TG balance mismatch");

        // Test 3: Claim ETH for TG
        uint256 claimAmount = 200; // 200 TG tokens
        uint256 initialBalance = address(deployer).balance;
        console2.log(string.concat("Initial ETH balance: ", Strings.toString(initialBalance)));
        
        // Get TG balance before claim
        (, uint256 beforeTgBalance) = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("TG balance before claim: ", Strings.toString(beforeTgBalance)));
        
        // Calculate expected ETH to receive (200 TG = 0.0002 ETH)
        uint256 expectedEth = claimAmount / 1000;
        console2.log(string.concat("Expected ETH to receive: ", Strings.toString(expectedEth)));
        
        triggerGasRegistry.claimETHForTG(claimAmount);
        
        // Get TG balance after claim
        (, uint256 afterTgBalance) = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("TG balance after claim: ", Strings.toString(afterTgBalance)));
        
        // Verify ETH balance after claim
        uint256 finalBalance = address(deployer).balance;
        console2.log(string.concat("Final ETH balance: ", Strings.toString(finalBalance)));
        uint256 ethReceived = finalBalance - initialBalance;
        console2.log(string.concat("ETH received from claim: ", Strings.toString(ethReceived)));
        
        require(afterTgBalance == beforeTgBalance - claimAmount, "TG balance not reduced correctly");
        require(ethReceived >= expectedEth, "ETH claim amount mismatch");

        // Test 4: Withdraw ETH (as owner)
        uint256 withdrawAmount = 0.05 ether;
        uint256 contractBalance = address(triggerGasRegistry).balance;
        triggerGasRegistry.withdrawETH(withdrawAmount, "Test withdrawal");
        
        // Verify contract balance after withdrawal
        uint256 newContractBalance = address(triggerGasRegistry).balance;
        console2.log(string.concat("Contract balance after withdrawal: ", Strings.toString(newContractBalance)));
        require(newContractBalance == contractBalance - withdrawAmount, "Withdrawal amount mismatch");

        vm.stopBroadcast();

        console2.log("All tests passed successfully!");
    }
}