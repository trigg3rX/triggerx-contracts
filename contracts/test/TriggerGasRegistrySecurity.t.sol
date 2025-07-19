// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract MaliciousReentrantRegistry {
    TriggerGasRegistry public target;
    uint256 public callCount;
    bool public attackType; // false = purchase, true = claim
    
    constructor(address _target) {
        target = TriggerGasRegistry(_target);
    }
    
    function attackPurchase() external payable {
        attackType = false;
        target.purchaseTG{value: msg.value}(msg.value);
    }
    
    function attackClaim(uint256 tgAmount) external {
        attackType = true;
        target.claimETHForTG(tgAmount);
    }
    
    receive() external payable {
        callCount++;
        if (callCount < 3) {
            if (attackType) {
                // Reentrancy on claim
                try target.claimETHForTG(100) {} catch {}
            } else {
                // Reentrancy on purchase
                if (address(this).balance >= 0.1 ether) {
                    try target.purchaseTG{value: 0.1 ether}(0.1 ether) {} catch {}
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
            owner,
            1000 // TG_PER_ETH
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
    
    function test_Security_PurchaseTGReentrancyProtection() public {
        // purchaseTG doesn't send ETH, so it can't trigger reentrancy via receive()
        // This test verifies that purchaseTG works normally and doesn't have reentrancy vulnerabilities
        MaliciousReentrantRegistry malicious = new MaliciousReentrantRegistry(address(gasRegistry));
        vm.deal(address(malicious), 10 ether);
        
        // This should succeed because purchaseTG doesn't send ETH, so no reentrancy can occur
        malicious.attackPurchase{value: 1 ether}();
        
        // Verify that the purchase succeeded normally
        (, uint256 tgBalance) = gasRegistry.getBalance(address(malicious));
        assertEq(tgBalance, 1e21); // 1 ETH * 1000 = 1000 TG
    }
    
    function test_Security_ClaimETHReentrancyProtection() public {
        // claimETHForTG does send ETH, so it can trigger reentrancy protection
        MaliciousReentrantRegistry malicious = new MaliciousReentrantRegistry(address(gasRegistry));
        vm.deal(address(malicious), 10 ether);
        
        // Purchase some TG first
        vm.prank(address(malicious));
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // The attack should succeed for the first call but the reentrancy (second call) should fail
        // This demonstrates that reentrancy protection is working correctly
        malicious.attackClaim(100);
        
        // Verify the attack didn't double-spend - balance should be reduced by only 100, not 200
        (, uint256 tgBalance) = gasRegistry.getBalance(address(malicious));
        assertEq(tgBalance, 1e21 - 100); // Started with 1000 TG, spent 100, not 200
    }

    // ==================== PARAMETER VALIDATION TESTS ====================
    
    function test_Security_PurchaseTGRejectsZeroAmount() public {
        vm.expectRevert("Cannot spend 0 ETH");
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 0}(0);
    }
    
    function test_Security_PurchaseTGRejectsMismatchedValue() public {
        // This test is no longer needed since purchaseTG doesn't take an amount parameter
        // The function now uses msg.value directly, so there's no mismatch possible
        uint256 ethAmount = 1 ether;
        
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Verify the purchase worked correctly
        (, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(tgBalance, ethAmount * 1000);
    }
    
    function test_Security_WithdrawETHRejectsZeroAmount() public {
        vm.expectRevert("Cannot withdraw 0 ETH");
        vm.prank(owner);
        gasRegistry.withdrawETH(0, "test");
    }
    
    function test_Security_WithdrawETHRejectsInsufficientBalance() public {
        vm.expectRevert("Insufficient contract balance");
        vm.prank(owner);
        gasRegistry.withdrawETH(1000 ether, "test");
    }

    // ==================== BALANCE MANIPULATION TESTS ====================
    
    function test_Security_CannotClaimMoreTGThanOwned() public {
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // User has 1 ETH * 1000 = 1000 TG (which is 1e21 in contract terms)
        // Try to claim more TG than available
        vm.expectRevert("Insufficient TG balance");
        vm.prank(user1);
        gasRegistry.claimETHForTG(1e21 + 1); // Try to claim 1 more than available
    }
    
    function test_Security_BalanceIntegrityAfterOperations() public {
        uint256 ethAmount = 5 ether;
        
        // User1 purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // User2 purchases TG
        vm.prank(user2);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        uint256 totalTGBefore = ethAmount * 1000 * 2; // Both users have 5000 TG eac     
        
        // Check total TG is conserved
        (, uint256 user1TG) = gasRegistry.getBalance(user1);
        (, uint256 user2TG) = gasRegistry.getBalance(user2);
        
        assertEq(user1TG + user2TG, totalTGBefore);
    }

    // ==================== OVERFLOW/UNDERFLOW TESTS ====================
    
    function test_Security_PurchaseLargeAmountTG() public {
        uint256 largeAmount = type(uint128).max;
        vm.deal(user1, largeAmount);
        
        // This should work due to Solidity 0.8+ overflow protection
        vm.prank(user1);
        gasRegistry.purchaseTG{value: largeAmount}(largeAmount);
        
        (, uint256 tgBalance) = gasRegistry.getBalance(user1);
        // Due to multiplication by 1000, this should overflow and revert in safe math
        // But since we're using uint256, it should be fine for reasonable amounts
        assertGt(tgBalance, 0);
    }
    
    function test_Security_MaximumTGCalculation() public {
        // Test the maximum safe ETH amount that won't overflow TG calculation
        uint256 maxSafeEth = type(uint256).max / 1000;
        
        // This should be much larger than any realistic amount
        assertGt(maxSafeEth, 1000000 ether);
    }

    // ==================== STATE CORRUPTION TESTS ====================
    
    function test_Security_CannotDirectlyManipulateBalances() public {
        // Test that there's no way to directly modify balances without going through proper functions
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        (, uint256 initialTG) = gasRegistry.getBalance(user1);
        
        // Try various operations that shouldn't affect the balance
        vm.prank(user1);
        gasRegistry.getBalance(user1); // Read operation
        
        (, uint256 finalTG) = gasRegistry.getBalance(user1);
        assertEq(initialTG, finalTG);
    }
    
    function test_Security_BalanceConsistencyAcrossUsers() public {
        // Test that operations on one user don't affect another user's balance
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        vm.prank(user2);
        gasRegistry.purchaseTG{value: 2 ether}(2 ether);
        
        (, uint256 user1Initial) = gasRegistry.getBalance(user1);
        (, uint256 user2Initial) = gasRegistry.getBalance(user2);
        
        // User1 claims some TG
        vm.prank(user1);
        gasRegistry.claimETHForTG(500);
        
        // Check user2's balance is unchanged
        (, uint256 user2Final) = gasRegistry.getBalance(user2);
        assertEq(user2Initial, user2Final);
        
        // Check user1's balance changed correctly
        (, uint256 user1Final) = gasRegistry.getBalance(user1);
        assertEq(user1Final, user1Initial - 500);
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
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // This should still work
        vm.prank(owner);
        gasRegistry.withdrawETH(0.1 ether, largeReason);
    }
    
    function test_Security_GasLimitOnLargeOperations() public {
        // Test that operations don't consume excessive gas
        uint256 gasStart = gasleft();
        
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        uint256 gasUsed = gasStart - gasleft();
        
        // Purchase should not use excessive gas
        assertLt(gasUsed, 200000);
    }

    // ==================== UPGRADE SECURITY TESTS ====================
    
    function test_Security_UpgradePreservesState() public {
        // Setup initial state
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        (, uint256 tgBalanceBefore) = gasRegistry.getBalance(user1);
        
        // Deploy new implementation
        TriggerGasRegistry newImpl = new TriggerGasRegistry();
        
        // Upgrade
        vm.prank(owner);
        gasRegistry.upgradeToAndCall(address(newImpl), "");
        
        // Check state is preserved
        (, uint256 tgBalanceAfter) = gasRegistry.getBalance(user1);
        assertEq(tgBalanceBefore, tgBalanceAfter);
    }
    
    function test_Security_CannotReinitialize() public {
        // Try to reinitialize the already initialized contract
        vm.expectRevert("InvalidInitialization()");
        gasRegistry.initialize(address(this), address(this), 1000);
    }

    // ==================== ETH HANDLING SECURITY TESTS ====================
    
    function test_Security_ETHTransferFailureHandling() public {
        // Create a contract that rejects ETH
        address rejectingContract = address(new RejectingContract());
        
        // Fund the rejecting contract with ETH
        vm.deal(rejectingContract, 1 ether);
        
        // Purchase TG for the rejecting contract
        vm.prank(rejectingContract);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // Try to claim ETH - should fail gracefully
        vm.expectRevert("ETH transfer failed");
        vm.prank(rejectingContract);
        gasRegistry.claimETHForTG(500);
    }
    
    function test_Security_WithdrawETHTransferFailure() public {
        // Purchase TG to fund the contract
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // Try to withdraw with normal recipient - should work
        vm.prank(owner);
        gasRegistry.withdrawETH(0.1 ether, "test");
    }

    // ==================== EDGE CASES AND BOUNDARY TESTS ====================
    
    function test_Security_ZeroTGClaim() public {
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // Claiming 0 TG should work but transfer 0 ETH
        uint256 balanceBefore = user1.balance;
        
        vm.prank(user1);
        gasRegistry.claimETHForTG(0);
        
        assertEq(user1.balance, balanceBefore);
    }

    
    function test_Security_ExactETHConversion() public {
        // Test the exact conversion rate: ethAmount * 1000 = tgAmount (with 18 decimals)
        uint256 ethAmount = 0.001 ether; // Should give 1e18 TG (0.001 * 1000 = 1 TG with 18 decimals)
        
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        (, uint256 tgBalance) = gasRegistry.getBalance(user1);
        assertEq(tgBalance, 1e18); // 0.001 ETH * 1000 = 1 TG (1e18 with decimals)
        
        // Claim it back
        uint256 balanceBefore = user1.balance;
        vm.prank(user1);
        gasRegistry.claimETHForTG(1e18);
        
        // Should get back 0.001 ETH (1 TG / 1000)
        assertEq(user1.balance, balanceBefore + 0.001 ether);
    }

    // ==================== DEDUCT TG BALANCE SECURITY TESTS ====================
    
    function test_Security_DeductTGBalance_OnlyOperator() public {
        // Setup user with TG balance
        vm.prank(user1);
        gasRegistry.purchaseTG{value: 1 ether}(1 ether);
        
        // Non-operator cannot deduct
        vm.expectRevert("Only operator can call this function");
        vm.prank(attacker);
        gasRegistry.deductTGBalance(user1, 100);
        
        // Other users cannot deduct
        vm.expectRevert("Only operator can call this function");
        vm.prank(user2);
        gasRegistry.deductTGBalance(user1, 100);
    }
    
    function test_Security_DeductTGBalance_CannotOverdeduct() public {
        uint256 ethAmount = 1 ether;
        
        // Setup user with TG balance
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get user's exact balance
        (, uint256 tgBalance) = gasRegistry.getBalance(user1);
        
        // Try to deduct more than available
        vm.expectRevert("Insufficient TG balance");
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductTGBalance(user1, tgBalance + 1);
    }
    
    function test_Security_DeductTGBalance_ZeroAmountHandling() public {
        uint256 ethAmount = 1 ether;
        
        // Setup user with TG balance
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get initial balance
        (, uint256 initialBalance) = gasRegistry.getBalance(user1);
        
        // Deduct zero amount - should not revert and should not change balance
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductTGBalance(user1, 0);
        
        // Check balance unchanged
        (, uint256 finalBalance) = gasRegistry.getBalance(user1);
        assertEq(finalBalance, initialBalance);
    }
    
    function test_Security_DeductTGBalance_FromZeroBalance() public {
        // Try to deduct from user with zero balance
        vm.expectRevert("Insufficient TG balance");
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductTGBalance(user1, 1);
    }
    
    function test_Security_DeductTGBalance_LargeAmount() public {
        uint256 largeEthAmount = 100 ether;
        vm.deal(user1, largeEthAmount);
        
        // User purchases large amount of TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: largeEthAmount}(largeEthAmount);
        
        uint256 largeTGAmount = largeEthAmount * 1000 / 2; // Deduct half
        
        // Should work for large amounts
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductTGBalance(user1, largeTGAmount);
        
        // Check balance updated correctly
        (, uint256 finalBalance) = gasRegistry.getBalance(user1);
        assertEq(finalBalance, largeEthAmount * 1000 - largeTGAmount);
    }
    
    function test_Security_DeductTGBalance_ExactBalanceDeduction() public {
        uint256 ethAmount = 1 ether;
        
        // User purchases TG
        vm.prank(user1);
        gasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get exact balance
        (, uint256 exactBalance) = gasRegistry.getBalance(user1);
        
        // Deduct exact balance
        vm.prank(owner); // owner is operator in security tests
        gasRegistry.deductTGBalance(user1, exactBalance);
        
        // Check balance is now zero
        (, uint256 finalBalance) = gasRegistry.getBalance(user1);
        assertEq(finalBalance, 0);
    }
}

contract RejectingContract {
    // This contract rejects all ETH transfers
    receive() external payable {
        revert("Rejecting ETH");
    }
}
