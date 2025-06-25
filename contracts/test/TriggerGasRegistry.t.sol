// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract TriggerGasRegistryTest is Test {
    TriggerGasRegistry public implementation;
    TriggerGasRegistry public gasRegistry;
    ERC1967Proxy public proxy;
    
    address public owner = address(0x1);
    address public operator = address(0x2);
    address public user1 = address(0x100);
    address public user2 = address(0x101);
    address public keeper = address(0x200);
    
    event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount);
    event TGClaimed(address indexed user, uint256 amount);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    event TGBalanceDeducted(address indexed user, uint256 amount);
    
    function setUp() public {
        vm.startPrank(owner);
        
        // Deploy implementation
        implementation = new TriggerGasRegistry();
        
        // Deploy proxy with initialize function
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector,
            owner,
            operator
        );
        proxy = new ERC1967Proxy(address(implementation), initData);
        
        // Get the proxy as TriggerGasRegistry
        gasRegistry = TriggerGasRegistry(address(proxy));
        
        vm.stopPrank();
        
        // Fund users with ETH for purchasing TG
        vm.deal(user1, 10 ether);
        vm.deal(user2, 10 ether);
    }
    
    function test_Initialize() public {
        assertEq(gasRegistry.owner(), owner);
    }
    
    function test_PurchaseTG() public {
        uint256 ethAmount = 1 ether;
        
        vm.expectEmit(true, false, false, true);
        emit TGPurchased(user1, ethAmount, ethAmount * 1000);
        
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Check balance was recorded correctly
        (uint256 ethSpent, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(ethSpent, ethAmount);
        assertEq(tgBalance, ethAmount * 1000); // 1 ETH should give 1000 TG tokens
        
        // Check contract balance increased
        assertEq(address(gasRegistry).balance, ethAmount);
    }
    
    function test_PurchaseTG_ZeroAmount() public {
        vm.expectRevert("Cannot spend 0 ETH");
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 0}(0);
    }
    
    function test_PurchaseTG_MismatchedAmount() public {
        vm.expectRevert("Sent ETH must match amount");
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 0.5 ether}(1 ether);
    }
    
    function test_GetBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Check balance info
        (uint256 ethSpent, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(ethSpent, ethAmount);
        assertEq(tgBalance, ethAmount * 1000);
        
        // User 2 has no balance
        (ethSpent, tgBalance) = gasRegistry.getBalance(user2);
        assertEq(ethSpent, 0);
        assertEq(tgBalance, 0);
    }
    
    function test_ClaimETHForTG() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Calculate expected return
        uint256 tgAmount = 500; // Claim 500 TG tokens
        uint256 expectedEthReturn = tgAmount / 1000; // 0.5 ETH
        
        uint256 userBalanceBefore = user1.balance;
        
        vm.expectEmit(true, false, false, true);
        emit TGClaimed(user1, tgAmount);
        
        // User claims ETH for TG
        vm.prank(user1);
        gasRegistry.claimETHForTG(tgAmount);
        
        // Check TG balance decreased
        (uint256 ethSpent, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(ethSpent, ethAmount); // ETH spent unchanged
        assertEq(tgBalance, ethAmount * 1000 - tgAmount); // TG balance decreased
        
        // Check user received ETH
        assertEq(user1.balance, userBalanceBefore + expectedEthReturn);
        
        // Check contract ETH balance decreased
        assertEq(address(gasRegistry).balance, ethAmount - expectedEthReturn);
    }
    
    function test_ClaimETHForTG_InsufficientTG() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Try to claim more TG than available
        uint256 tgAmount = ethAmount * 1000 + 1; // 1 more than available
        
        vm.expectRevert("Insufficient TG balance");
        vm.prank(user1);
        gasRegistry.claimETHForTG(tgAmount);
    }
    
    function test_UpdateTGBalances() public {
        uint256 ethAmount = 1 ether;
        
        // Both users purchase TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        vm.prank(user2);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Transfer TG from user1 to keeper
        uint256 tgAmount = 300;
        
        vm.expectEmit(true, true, false, true);
        emit TGTransferred(user1, keeper, tgAmount);
        
        vm.prank(owner);
        gasRegistry.updateTGBalances(keeper, user1, tgAmount);
        
        // Check user1's TG decreased
        (uint256 ethSpent, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(ethSpent, ethAmount);
        assertEq(tgBalance, ethAmount * 1000 - tgAmount);
        
        // Check keeper's TG increased
        (ethSpent, tgBalance) = gasRegistry.getBalance(keeper);
        assertEq(ethSpent, 0); // No ETH spent
        assertEq(tgBalance, tgAmount); // TG transferred
    }
    
    function test_UpdateTGBalances_OnlyOwner() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user2));
        vm.prank(user2); // Not owner
        gasRegistry.updateTGBalances(keeper, user1, 100);
    }
    
    function test_UpdateTGBalances_InsufficientTG() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Try to transfer more TG than available
        uint256 tgAmount = ethAmount * 1000 + 1; // 1 more than available
        
        vm.expectRevert("Insufficient user TG balance");
        vm.prank(owner);
        gasRegistry.updateTGBalances(keeper, user1, tgAmount);
    }
    
    function test_Upgrade() public {
        // Deploy new implementation
        TriggerGasRegistry newImplementation = new TriggerGasRegistry();
        
        // Upgrade to new implementation
        vm.prank(owner);
        gasRegistry.upgradeToAndCall(address(newImplementation), "");
        
        // Verify upgrade was successful
        // This would typically involve checking that the implementation address changed
        // and that functionality works as expected
    }
    
    function test_Upgrade_OnlyOwner() public {
        // Deploy new implementation
        TriggerGasRegistry newImplementation = new TriggerGasRegistry();
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1); // Not owner
        gasRegistry.upgradeToAndCall(address(newImplementation), "");
    }

    function test_WithdrawETH() public {
        uint256 ethAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        uint256 withdrawAmount = 0.5 ether;
        string memory reason = "Test withdrawal";
        
        uint256 ownerBalanceBefore = owner.balance;
        
        vm.expectEmit(true, false, false, true);
        emit ETHWithdrawn(owner, withdrawAmount, reason);
        
        vm.prank(owner);
        gasRegistry.withdrawETH(withdrawAmount, reason);
        
        // Check owner received ETH
        assertEq(owner.balance, ownerBalanceBefore + withdrawAmount);
        
        // Check contract balance decreased
        assertEq(address(gasRegistry).balance, ethAmount - withdrawAmount);
    }
    
    function test_WithdrawETH_OnlyOwner() public {
        uint256 ethAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1); // Not owner
        gasRegistry.withdrawETH(0.5 ether, "Test");
    }
    
    function test_WithdrawETH_ZeroAmount() public {
        vm.expectRevert("Cannot withdraw 0 ETH");
        vm.prank(owner);
        gasRegistry.withdrawETH(0, "Test");
    }
    
    function test_WithdrawETH_InsufficientBalance() public {
        uint256 ethAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        vm.expectRevert("Insufficient contract balance");
        vm.prank(owner);
        gasRegistry.withdrawETH(ethAmount + 1, "Test");
    }

    // ==================== DEDUCT TG BALANCE TESTS ====================
    
    function test_DeductTGBalance_Normal() public {
        uint256 ethAmount = 1 ether;
        uint256 deductAmount = 500;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get initial balance
        (, uint256 initialTG) = gasRegistry.getBalance(user1);
        
        // Expect event to be emitted
        vm.expectEmit(true, false, false, true);
        emit TGBalanceDeducted(user1, deductAmount);
        
        // Operator deducts TG
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, deductAmount);
        
        // Check balance decreased
        (, uint256 finalTG) = gasRegistry.getBalance(user1);
        assertEq(finalTG, initialTG - deductAmount);
    }
    
    function test_DeductTGBalance_ZeroAmount() public {
        uint256 ethAmount = 1 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get initial balance
        (, uint256 initialTG) = gasRegistry.getBalance(user1);
        
        // Deduct 0 TG - should not emit event and should not revert
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, 0);
        
        // Check balance unchanged
        (, uint256 finalTG) = gasRegistry.getBalance(user1);
        assertEq(finalTG, initialTG);
    }
    
    function test_DeductTGBalance_InsufficientBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Try to deduct more TG than available
        uint256 tgBalance = ethAmount * 1000;
        
        vm.expectRevert("Insufficient TG balance");
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, tgBalance + 1);
    }
    
    function test_DeductTGBalance_OnlyOperator() public {
        uint256 ethAmount = 1 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Try to deduct as non-operator
        vm.expectRevert("Only operator can call this function");
        vm.prank(user1);
        gasRegistry.deductTGBalance(user1, 100);
        
        // Try to deduct as owner (not operator)
        vm.expectRevert("Only operator can call this function");
        vm.prank(owner);
        gasRegistry.deductTGBalance(user1, 100);
    }
    
    function test_DeductTGBalance_ExactBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get exact balance
        (, uint256 tgBalance) = gasRegistry.getBalance(user1);
        
        // Deduct exact balance
        vm.expectEmit(true, false, false, true);
        emit TGBalanceDeducted(user1, tgBalance);
        
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, tgBalance);
        
        // Check balance is now zero
        (, uint256 finalTG) = gasRegistry.getBalance(user1);
        assertEq(finalTG, 0);
    }
    
    function test_DeductTGBalance_FromUserWithZeroBalance() public {
        // Try to deduct from user with no TG balance
        vm.expectRevert("Insufficient TG balance");
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, 1);
    }
    
    function test_DeductTGBalance_MultipleDeductions() public {
        uint256 ethAmount = 2 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        uint256 initialTG = ethAmount * 1000;
        
        // First deduction
        vm.expectEmit(true, false, false, true);
        emit TGBalanceDeducted(user1, 300);
        
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, 300);
        
        (, uint256 afterFirst) = gasRegistry.getBalance(user1);
        assertEq(afterFirst, initialTG - 300);
        
        // Second deduction
        vm.expectEmit(true, false, false, true);
        emit TGBalanceDeducted(user1, 700);
        
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, 700);
        
        (, uint256 afterSecond) = gasRegistry.getBalance(user1);
        assertEq(afterSecond, initialTG - 300 - 700);
    }
    
    function test_DeductTGBalance_DoesNotAffectOtherUsers() public {
        uint256 ethAmount = 1 ether;
        
        // Both users purchase TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        vm.prank(user2);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get initial balances
        (, uint256 user1InitialTG) = gasRegistry.getBalance(user1);
        (, uint256 user2InitialTG) = gasRegistry.getBalance(user2);
        
        // Deduct from user1
        vm.prank(operator);
        gasRegistry.deductTGBalance(user1, 500);
        
        // Check user1's balance changed
        (, uint256 user1FinalTG) = gasRegistry.getBalance(user1);
        assertEq(user1FinalTG, user1InitialTG - 500);
        
        // Check user2's balance unchanged
        (, uint256 user2FinalTG) = gasRegistry.getBalance(user2);
        assertEq(user2FinalTG, user2InitialTG);
    }
} 