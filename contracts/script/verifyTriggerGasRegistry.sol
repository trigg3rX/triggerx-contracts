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

        // Test 2: Deposit ETH
        uint256 depositAmount = 0.1 ether;
        triggerGasRegistry.depositETH{value: depositAmount}(depositAmount);
        
        // Verify balance after deposit
        uint256 ethBalance = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("ETH balance: ", Strings.toString(ethBalance)));
        require(ethBalance == depositAmount, "ETH balance mismatch");

        // Test 3: Withdraw ETH balance
        uint256 withdrawAmount = 0.05 ether;
        uint256 initialBalance = address(deployer).balance;
        console2.log(string.concat("Initial ETH balance: ", Strings.toString(initialBalance)));
        
        // Get ETH balance before withdraw
        uint256 beforeETHBalance = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("ETH balance before withdraw: ", Strings.toString(beforeETHBalance)));
        
        triggerGasRegistry.withdrawETHBalance(withdrawAmount);
        
        // Get ETH balance after withdraw
        uint256 afterETHBalance = triggerGasRegistry.getBalance(deployer);
        console2.log(string.concat("ETH balance after withdraw: ", Strings.toString(afterETHBalance)));
        
        // Verify ETH balance after withdraw
        uint256 finalBalance = address(deployer).balance;
        console2.log(string.concat("Final ETH balance: ", Strings.toString(finalBalance)));
        uint256 ethReceived = finalBalance - initialBalance;
        console2.log(string.concat("ETH received from withdraw: ", Strings.toString(ethReceived)));
        
        require(afterETHBalance == beforeETHBalance - withdrawAmount, "ETH balance not reduced correctly");
        require(ethReceived == withdrawAmount, "ETH withdraw amount mismatch");

        // Test 4: Withdraw ETH (as owner) - requires deducted balance first
        // First, deduct some balance to make it available for owner withdrawal
        uint256 deductAmount = 0.02 ether;
        triggerGasRegistry.deductETHBalance(deployer, deductAmount);
        
        uint256 ownerWithdrawAmount = 0.02 ether;
        uint256 contractBalance = address(triggerGasRegistry).balance;
        triggerGasRegistry.withdrawETH(ownerWithdrawAmount, "Test withdrawal");
        
        // Verify contract balance after withdrawal
        uint256 newContractBalance = address(triggerGasRegistry).balance;
        console2.log(string.concat("Contract balance after withdrawal: ", Strings.toString(newContractBalance)));
        require(newContractBalance == contractBalance - ownerWithdrawAmount, "Withdrawal amount mismatch");

        vm.stopBroadcast();

        console2.log("All tests passed successfully!");
    }
}