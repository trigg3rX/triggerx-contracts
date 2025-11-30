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
    
    event ETHDeposited(address indexed user, uint256 ethAmount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    event ETHBalanceDeducted(address indexed user, uint256 amount);
    
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
        
        // Fund users with ETH for depositing
        vm.deal(user1, 10 ether);
        vm.deal(user2, 10 ether);
    }
    
    function test_Initialize() public view {
        assertEq(gasRegistry.owner(), owner);
    }
    
    function test_DepositETH() public {
        uint256 ethAmount = 1 ether;
        
        vm.expectEmit(true, false, false, true);
        emit ETHDeposited(user1, ethAmount);
        
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Check balance was recorded correctly
        uint256 balance = gasRegistry.getBalance(user1);
        assertEq(balance, ethAmount);
        
        // Check contract balance increased
        assertEq(address(gasRegistry).balance, ethAmount);
    }
    
    function test_DepositETH_ZeroAmount() public {
        vm.expectRevert("Cannot deposit 0 ETH");
        vm.prank(user1);
        gasRegistry.depositETH{value: 0}(0);
    }
    
    function test_DepositETH_MismatchedAmount() public {
        uint256 ethAmount = 1 ether;
        
        vm.expectRevert("Sent ETH must match amount");
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount + 1);
    }
    
    function test_GetBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Check balance info
        uint256 balance = gasRegistry.getBalance(user1);
        assertEq(balance, ethAmount);
        
        // User 2 has no balance
        balance = gasRegistry.getBalance(user2);
        assertEq(balance, 0);
    }
    
    function test_WithdrawETHBalance() public {
        uint256 ethAmount = 1 ether;
        uint256 withdrawAmount = 0.5 ether;
        
        // User 1 deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        uint256 userBalanceBefore = user1.balance;
        
        // User withdraws ETH
        vm.prank(user1);
        gasRegistry.withdrawETHBalance(withdrawAmount);
        
        // Check balance decreased
        uint256 balance = gasRegistry.getBalance(user1);
        assertEq(balance, ethAmount - withdrawAmount);
        
        // Check user received ETH
        assertEq(user1.balance, userBalanceBefore + withdrawAmount);
        
        // Check contract ETH balance decreased
        assertEq(address(gasRegistry).balance, ethAmount - withdrawAmount);
    }
    
    function test_WithdrawETHBalance_InsufficientBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User 1 deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Try to withdraw more than available
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(user1);
        gasRegistry.withdrawETHBalance(ethAmount + 1);
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
        
        // First, user deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Operator deducts some ETH (this adds to totalDeductedBalance)
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, 0.5 ether);
        
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
        
        // First, user deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Operator deducts some ETH
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, 0.5 ether);
        
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
        
        // First, user deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Operator deducts some ETH
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, 0.5 ether);
        
        vm.expectRevert("Cannot withdraw more than deducted balance");
        vm.prank(owner);
        gasRegistry.withdrawETH(0.6 ether, "Test");
    }

    // ==================== DEDUCT ETH BALANCE TESTS ====================
    
    function test_DeductETHBalance_Normal() public {
        uint256 ethAmount = 1 ether;
        uint256 deductAmount = 0.5 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get initial balance
        uint256 initialBalance = gasRegistry.getBalance(user1);
        
        // Expect event to be emitted
        vm.expectEmit(true, false, false, true);
        emit ETHBalanceDeducted(user1, deductAmount);
        
        // Operator deducts ETH
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, deductAmount);
        
        // Check balance decreased
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, initialBalance - deductAmount);
        
        // Check totalDeductedBalance increased
        assertEq(gasRegistry.totalDeductedBalance(), deductAmount);
    }
    
    function test_DeductETHBalance_ZeroAmount() public {
        uint256 ethAmount = 1 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get initial balance
        uint256 initialBalance = gasRegistry.getBalance(user1);
        
        // Deduct 0 ETH - should not emit event and should not revert
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, 0);
        
        // Check balance unchanged
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, initialBalance);
        
        // Check totalDeductedBalance unchanged
        assertEq(gasRegistry.totalDeductedBalance(), 0);
    }
    
    function test_DeductETHBalance_InsufficientBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Try to deduct more ETH than available
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, ethAmount + 1);
    }
    
    function test_DeductETHBalance_OnlyOperator() public {
        uint256 ethAmount = 1 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Try to deduct as non-operator
        vm.expectRevert("Only operator can call this function");
        vm.prank(user1);
        gasRegistry.deductETHBalance(user1, 0.1 ether);
        
        // Try to deduct as owner (not operator)
        vm.expectRevert("Only operator can call this function");
        vm.prank(owner);
        gasRegistry.deductETHBalance(user1, 0.1 ether);
    }
    
    function test_DeductETHBalance_ExactBalance() public {
        uint256 ethAmount = 1 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get exact balance
        uint256 balance = gasRegistry.getBalance(user1);
        
        // Deduct exact balance
        vm.expectEmit(true, false, false, true);
        emit ETHBalanceDeducted(user1, balance);
        
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, balance);
        
        // Check balance is now zero
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, 0);
        
        // Check totalDeductedBalance
        assertEq(gasRegistry.totalDeductedBalance(), balance);
    }
    
    function test_DeductETHBalance_FromUserWithZeroBalance() public {
        // Try to deduct from user with no ETH balance
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, 1);
    }
    
    function test_DeductETHBalance_MultipleDeductions() public {
        uint256 ethAmount = 2 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        uint256 initialBalance = ethAmount;
        
        // First deduction
        uint256 firstDeduction = 0.3 ether;
        vm.expectEmit(true, false, false, true);
        emit ETHBalanceDeducted(user1, firstDeduction);
        
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, firstDeduction);
        
        uint256 afterFirst = gasRegistry.getBalance(user1);
        assertEq(afterFirst, initialBalance - firstDeduction);
        assertEq(gasRegistry.totalDeductedBalance(), firstDeduction);
        
        // Second deduction
        uint256 secondDeduction = 0.7 ether;
        vm.expectEmit(true, false, false, true);
        emit ETHBalanceDeducted(user1, secondDeduction);
        
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, secondDeduction);
        
        uint256 afterSecond = gasRegistry.getBalance(user1);
        assertEq(afterSecond, initialBalance - firstDeduction - secondDeduction);
        assertEq(gasRegistry.totalDeductedBalance(), firstDeduction + secondDeduction);
    }
    
    function test_DeductETHBalance_DoesNotAffectOtherUsers() public {
        uint256 ethAmount = 1 ether;
        
        // Both users deposit ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        vm.prank(user2);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get initial balances
        uint256 user1InitialBalance = gasRegistry.getBalance(user1);
        uint256 user2InitialBalance = gasRegistry.getBalance(user2);
        
        // Deduct from user1
        uint256 deductAmount = 0.5 ether;
        vm.prank(operator);
        gasRegistry.deductETHBalance(user1, deductAmount);
        
        // Check user1's balance changed
        uint256 user1FinalBalance = gasRegistry.getBalance(user1);
        assertEq(user1FinalBalance, user1InitialBalance - deductAmount);
        
        // Check user2's balance unchanged
        uint256 user2FinalBalance = gasRegistry.getBalance(user2);
        assertEq(user2FinalBalance, user2InitialBalance);
        
        // Check totalDeductedBalance
        assertEq(gasRegistry.totalDeductedBalance(), deductAmount);
    }
} 