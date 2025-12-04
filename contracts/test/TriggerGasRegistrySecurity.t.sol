// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract MaliciousReentrantRegistry {
    TriggerGasRegistry public target;
    uint256 public callCount;
    bool public attackType; // false = deposit, true = withdraw
    
    constructor(address _target) {
        target = TriggerGasRegistry(_target);
    }
    
    function attackDeposit() external payable {
        attackType = false;
        target.depositETH{value: msg.value}(msg.value);
    }
    
    function attackWithdraw(uint256 ethAmount) external {
        attackType = true;
        target.withdrawETHBalance(ethAmount);
    }
    
    receive() external payable {
        callCount++;
        if (callCount < 3) {
            if (attackType) {
                // Reentrancy on withdraw
                try target.withdrawETHBalance(0.1 ether) {} catch {}
            } else {
                // Reentrancy on deposit
                if (address(this).balance >= 0.1 ether) {
                    try target.depositETH{value: 0.1 ether}(0.1 ether) {} catch {}
                }
            }
        }
    }
}

contract MaliciousUpgrade is TriggerGasRegistry {
    function maliciousFunction() external {
        // Malicious function that could drain funds
        payable(msg.sender).transfer(address(this).balance);
    }
}

contract TriggerGasRegistrySecurityTest is Test {
    TriggerGasRegistry public implementation;
    TriggerGasRegistry public gasRegistry;
    ERC1967Proxy public proxy;
    
    address public owner = address(0x1);
        
    address public attacker = address(0x666);
    address public user1 = address(0x100);
    address public user2 = address(0x101);
    address public keeper = address(0x200);
    
    function setUp() public {
        vm.startPrank(owner);
        
        implementation = new TriggerGasRegistry();
        
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector,
            owner,
            owner
        );
        proxy = new ERC1967Proxy(address(implementation), initData);
        gasRegistry = TriggerGasRegistry(address(proxy));
        
        vm.stopPrank();
        
        vm.deal(user1, 100 ether);
        vm.deal(user2, 100 ether);
        vm.deal(attacker, 100 ether);
    }

    // ==================== ACCESS CONTROL TESTS ====================
    function test_Security_OnlyOwnerCanWithdrawETH() public {
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        gasRegistry.withdrawETH(1 ether, "malicious withdrawal");
    }
    
    function test_Security_OnlyOwnerCanUpgrade() public {
        MaliciousUpgrade maliciousImpl = new MaliciousUpgrade();
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        gasRegistry.upgradeToAndCall(address(maliciousImpl), "");
    }

    // ==================== REENTRANCY TESTS ====================
    
    function test_Security_DepositETHReentrancyProtection() public {
        // depositETH doesn't send ETH, so it can't trigger reentrancy via receive()
        // This test verifies that depositETH works normally and doesn't have reentrancy vulnerabilities
        MaliciousReentrantRegistry malicious = new MaliciousReentrantRegistry(address(gasRegistry));
        vm.deal(address(malicious), 10 ether);
        
        // This should succeed because depositETH doesn't send ETH, so no reentrancy can occur
        malicious.attackDeposit{value: 1 ether}();
        
        // Verify that the deposit succeeded normally
        uint256 balance = gasRegistry.getBalance(address(malicious));
        assertEq(balance, 1 ether);
    }
    
    function test_Security_WithdrawETHReentrancyProtection() public {
        // withdrawETHBalance does send ETH, so it can trigger reentrancy protection
        MaliciousReentrantRegistry malicious = new MaliciousReentrantRegistry(address(gasRegistry));
        vm.deal(address(malicious), 10 ether);
        
        // Deposit some ETH first
        vm.prank(address(malicious));
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // The attack should succeed for the first call but the reentrancy (second call) should fail
        // This demonstrates that reentrancy protection is working correctly
        malicious.attackWithdraw(0.1 ether);
        
        // Verify the attack didn't double-spend - balance should be reduced by only 0.1 ether, not 0.2 ether
        uint256 balance = gasRegistry.getBalance(address(malicious));
        assertEq(balance, 0.9 ether); // Started with 1 ether, spent 0.1 ether, not 0.2 ether
    }

    // ==================== PARAMETER VALIDATION TESTS ====================
    
    function test_Security_DepositETHRejectsZeroAmount() public {
        vm.expectRevert("Cannot deposit 0 ETH");
        vm.prank(user1);
        gasRegistry.depositETH{value: 0}(0);
    }
    
    function test_Security_DepositETHRejectsMismatchedValue() public {
        uint256 ethAmount = 1 ether;
        
        vm.expectRevert("Sent ETH must match amount");
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount + 1);
    }
    
    function test_Security_WithdrawETHRejectsZeroAmount() public {
        vm.expectRevert("Cannot withdraw 0 ETH");
        vm.prank(owner);
        gasRegistry.withdrawETH(0, "test");
    }
    
    function test_Security_WithdrawETHRejectsInsufficientBalance() public {
        // First need to have some deducted balance
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        vm.prank(owner); // owner is operator
        gasRegistry.deductETHBalance(user1, 0.5 ether);
        
        vm.expectRevert("Cannot withdraw more than deducted balance");
        vm.prank(owner);
        gasRegistry.withdrawETH(1 ether, "test");
    }

    // ==================== BALANCE MANIPULATION TESTS ====================
    
    function test_Security_CannotWithdrawMoreETHThanOwned() public {
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Try to withdraw more ETH than available
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(user1);
        gasRegistry.withdrawETHBalance(1 ether + 1); // Try to withdraw 1 wei more than available
    }
    
    function test_Security_BalanceIntegrityAfterOperations() public {
        uint256 ethAmount = 5 ether;
        
        // User1 deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // User2 deposits ETH
        vm.prank(user2);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        uint256 totalETHBefore = ethAmount * 2; // Both users have 5 ETH each
        
        // Check total ETH is conserved
        uint256 user1ETH = gasRegistry.getBalance(user1);
        uint256 user2ETH = gasRegistry.getBalance(user2);
        
        assertEq(user1ETH + user2ETH, totalETHBefore);
    }

    // ==================== OVERFLOW/UNDERFLOW TESTS ====================
    
    function test_Security_DepositLargeAmountETH() public {
        uint256 largeAmount = type(uint128).max;
        vm.deal(user1, largeAmount);
        
        // This should work due to Solidity 0.8+ overflow protection
        vm.prank(user1);
        gasRegistry.depositETH{value: largeAmount}(largeAmount);
        
        uint256 balance = gasRegistry.getBalance(user1);
        assertEq(balance, largeAmount);
    }

    // ==================== STATE CORRUPTION TESTS ====================
    
    function test_Security_CannotDirectlyManipulateBalances() public {
        // Test that there's no way to directly modify balances without going through proper functions
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        uint256 initialBalance = gasRegistry.getBalance(user1);
        
        // Try various operations that shouldn't affect the balance
        vm.prank(user1);
        gasRegistry.getBalance(user1); // Read operation
        
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(initialBalance, finalBalance);
    }
    
    function test_Security_BalanceConsistencyAcrossUsers() public {
        // Test that operations on one user don't affect another user's balance
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        vm.prank(user2);
        gasRegistry.depositETH{value: 2 ether}(2 ether);
        
        uint256 user1Initial = gasRegistry.getBalance(user1);
        uint256 user2Initial = gasRegistry.getBalance(user2);
        
        // User1 withdraws some ETH
        vm.prank(user1);
        gasRegistry.withdrawETHBalance(0.5 ether);
        
        // Check user2's balance is unchanged
        uint256 user2Final = gasRegistry.getBalance(user2);
        assertEq(user2Initial, user2Final);
        
        // Check user1's balance changed correctly
        uint256 user1Final = gasRegistry.getBalance(user1);
        assertEq(user1Final, user1Initial - 0.5 ether);
    }

    // ==================== DENIAL OF SERVICE TESTS ====================
    
    function test_Security_CannotDOSWithLargeStringReason() public {
        // Create a very large string for withdrawal reason
        string memory largeReason = "";
        for (uint i = 0; i < 100; i++) {
            largeReason = string(abi.encodePacked(largeReason, "This is a very long reason string that might cause issues"));
        }
        
        // Fund contract first
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Deduct some balance first
        vm.prank(owner); // owner is operator
        gasRegistry.deductETHBalance(user1, 0.1 ether);
        
        // This should still work
        vm.prank(owner);
        gasRegistry.withdrawETH(0.1 ether, largeReason);
    }
    
    function test_Security_GasLimitOnLargeOperations() public {
        // Test that operations don't consume excessive gas
        uint256 gasStart = gasleft();
        
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        uint256 gasUsed = gasStart - gasleft();
        
        // Deposit should not use excessive gas
        assertLt(gasUsed, 200000);
    }

    // ==================== UPGRADE SECURITY TESTS ====================
    
    function test_Security_UpgradePreservesState() public {
        // Setup initial state
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        uint256 balanceBefore = gasRegistry.getBalance(user1);
        
        // Deploy new implementation
        TriggerGasRegistry newImpl = new TriggerGasRegistry();
        
        // Upgrade
        vm.prank(owner);
        gasRegistry.upgradeToAndCall(address(newImpl), "");
        
        // Check state is preserved
        uint256 balanceAfter = gasRegistry.getBalance(user1);
        assertEq(balanceBefore, balanceAfter);
    }
    
    function test_Security_CannotReinitialize() public {
        // Try to reinitialize the already initialized contract
        vm.expectRevert("InvalidInitialization()");
        gasRegistry.initialize(address(this), address(this));
    }

    // ==================== ETH HANDLING SECURITY TESTS ====================
    
    function test_Security_ETHTransferFailureHandling() public {
        // Create a contract that rejects ETH
        address rejectingContract = address(new RejectingContract());
        
        // Fund the rejecting contract with ETH
        vm.deal(rejectingContract, 1 ether);
        
        // Deposit ETH for the rejecting contract
        vm.prank(rejectingContract);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Try to withdraw ETH - should fail gracefully
        vm.expectRevert("ETH transfer failed");
        vm.prank(rejectingContract);
        gasRegistry.withdrawETHBalance(0.5 ether);
    }
    
    function test_Security_WithdrawETHTransferFailure() public {
        // Deposit ETH to fund the contract
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Deduct some balance first
        vm.prank(owner); // owner is operator
        gasRegistry.deductETHBalance(user1, 0.1 ether);
        
        // Try to withdraw with normal recipient - should work
        vm.prank(owner);
        gasRegistry.withdrawETH(0.1 ether, "test");
    }

    // ==================== EDGE CASES AND BOUNDARY TESTS ====================
    
    function test_Security_ZeroETHWithdraw() public {
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Withdrawing 0 ETH should revert
        vm.expectRevert("Cannot withdraw 0 ETH");
        vm.prank(user1);
        gasRegistry.withdrawETHBalance(0);
    }

    // ==================== DEDUCT ETH BALANCE SECURITY TESTS ====================
    
    function test_Security_DeductETHBalance_OnlyOperator() public {
        // Setup user with ETH balance
        vm.prank(user1);
        gasRegistry.depositETH{value: 1 ether}(1 ether);
        
        // Non-operator cannot deduct
        vm.expectRevert("Only operator can call this function");
        vm.prank(attacker);
        gasRegistry.deductETHBalance(user1, 0.1 ether);
        
        // Other users cannot deduct
        vm.expectRevert("Only operator can call this function");
        vm.prank(user2);
        gasRegistry.deductETHBalance(user1, 0.1 ether);
    }
    
    function test_Security_DeductETHBalance_CannotOverdeduct() public {
        uint256 ethAmount = 1 ether;
        
        // Setup user with ETH balance
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get user's exact balance
        uint256 balance = gasRegistry.getBalance(user1);
        
        // Try to deduct more than available
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductETHBalance(user1, balance + 1);
    }
    
    function test_Security_DeductETHBalance_ZeroAmountHandling() public {
        uint256 ethAmount = 1 ether;
        
        // Setup user with ETH balance
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get initial balance
        uint256 initialBalance = gasRegistry.getBalance(user1);
        
        // Deduct zero amount - should not revert and should not change balance
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductETHBalance(user1, 0);
        
        // Check balance unchanged
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, initialBalance);
        
        // Check totalDeductedBalance unchanged
        assertEq(gasRegistry.totalDeductedBalance(), 0);
    }
    
    function test_Security_DeductETHBalance_FromZeroBalance() public {
        // Try to deduct from user with zero balance
        vm.expectRevert("Insufficient ETH balance");
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductETHBalance(user1, 1);
    }
    
    function test_Security_DeductETHBalance_LargeAmount() public {
        uint256 largeEthAmount = 100 ether;
        vm.deal(user1, largeEthAmount);
        
        // User deposits large amount of ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: largeEthAmount}(largeEthAmount);
        
        uint256 largeDeduction = largeEthAmount / 2; // Deduct half
        
        // Should work for large amounts
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductETHBalance(user1, largeDeduction);
        
        // Check balance updated correctly
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, largeEthAmount - largeDeduction);
        
        // Check totalDeductedBalance
        assertEq(gasRegistry.totalDeductedBalance(), largeDeduction);
    }
    
    function test_Security_DeductETHBalance_ExactBalanceDeduction() public {
        uint256 ethAmount = 1 ether;
        
        // User deposits ETH
        vm.prank(user1);
        gasRegistry.depositETH{value: ethAmount}(ethAmount);
        
        // Get exact balance
        uint256 exactBalance = gasRegistry.getBalance(user1);
        
        // Deduct exact balance
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductETHBalance(user1, exactBalance);
        
        // Check balance is now zero
        uint256 finalBalance = gasRegistry.getBalance(user1);
        assertEq(finalBalance, 0);
        
        // Check totalDeductedBalance
        assertEq(gasRegistry.totalDeductedBalance(), exactBalance);
    }
}

contract RejectingContract {
    // This contract rejects all ETH transfers
    receive() external payable {
        revert("Rejecting ETH");
    }
}
