// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {TriggerXStakeRegistry} from "../src/TriggerXStakeRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract TriggerXStakeRegistryTest is Test {
    TriggerXStakeRegistry public implementation;
    TriggerXStakeRegistry public stakeRegistry;
    ERC1967Proxy public proxy;
    
    address public owner = address(0x1);
    address public user1 = address(0x100);
    address public user2 = address(0x101);
    address public keeper = address(0x200);
    
    event Staked(address indexed user, uint256 amount);
    event TGClaimed(address indexed user, uint256 amount);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    
    function setUp() public {
        vm.startPrank(owner);
        
        // Deploy implementation
        implementation = new TriggerXStakeRegistry();
        
        // Deploy proxy with initialize function
        bytes memory initData = abi.encodeWithSelector(
            TriggerXStakeRegistry.initialize.selector
        );
        proxy = new ERC1967Proxy(address(implementation), initData);
        
        // Get the proxy as TriggerXStakeRegistry
        stakeRegistry = TriggerXStakeRegistry(address(proxy));
        
        vm.stopPrank();
        
        // Fund users with ETH for purchasing TG
        vm.deal(user1, 10 ether);
        vm.deal(user2, 10 ether);
    }
    
    function test_Initialize() public {
        assertEq(stakeRegistry.owner(), owner);
    }
    
    function test_Stake() public {
        uint256 stakeAmount = 1 ether;
        
        vm.expectEmit(true, false, false, true);
        emit Staked(user1, stakeAmount);
        
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Check balance was recorded correctly
        (uint256 amount, uint256 TGbalance) = stakeRegistry.getStake(user1);
        assertEq(amount, stakeAmount);
        assertEq(TGbalance, stakeAmount * 1000); // 1 ETH should give 1000 TG tokens
        
        // Check contract balance increased
        assertEq(address(stakeRegistry).balance, stakeAmount);
    }
    
    function test_Stake_ZeroAmount() public {
        vm.expectRevert("Cannot stake 0 ETH");
        vm.prank(user1);
        stakeRegistry.stake{value: 0}(0);
    }
    
    function test_Stake_MismatchedAmount() public {
        vm.expectRevert("Sent ETH must match amount");
        vm.prank(user1);
        stakeRegistry.stake{value: 0.5 ether}(1 ether);
    }
    
    function test_GetStake() public {
        uint256 stakeAmount = 1 ether;
        
        // User 1 stakes
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Check stake info
        (uint256 amount, uint256 TGbalance) = stakeRegistry.getStake(user1);
        assertEq(amount, stakeAmount);
        assertEq(TGbalance, stakeAmount * 1000);
        
        // User 2 has no stake
        (amount, TGbalance) = stakeRegistry.getStake(user2);
        assertEq(amount, 0);
        assertEq(TGbalance, 0);
    }
    
    function test_ClaimETHForTG() public {
        uint256 stakeAmount = 1 ether;
        
        // User 1 stakes
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Calculate expected return
        uint256 tgAmountToClaim = 500; // Claim 500 TG tokens
        uint256 expectedEthReturn = tgAmountToClaim / 1000; // 0.5 ETH
        
        uint256 userBalanceBefore = user1.balance;
        
        vm.expectEmit(true, false, false, true);
        emit TGClaimed(user1, tgAmountToClaim);
        
        // User claims ETH for TG
        vm.prank(user1);
        stakeRegistry.claimETHForTG(tgAmountToClaim);
        
        // Check TG balance decreased
        (uint256 amount, uint256 TGbalance) = stakeRegistry.getStake(user1);
        assertEq(amount, stakeAmount); // Staked amount unchanged
        assertEq(TGbalance, stakeAmount * 1000 - tgAmountToClaim); // TG balance decreased
        
        // Check user received ETH
        assertEq(user1.balance, userBalanceBefore + expectedEthReturn);
        
        // Check contract ETH balance decreased
        assertEq(address(stakeRegistry).balance, stakeAmount - expectedEthReturn);
    }
    
    function test_ClaimETHForTG_InsufficientTG() public {
        uint256 stakeAmount = 1 ether;
        
        // User 1 stakes
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Try to claim more TG than available
        uint256 tgAmount = stakeAmount * 1000 + 1; // 1 more than available
        
        vm.expectRevert("Insufficient TG balance");
        vm.prank(user1);
        stakeRegistry.claimETHForTG(tgAmount);
    }
    
    function test_UpdateTGBalances() public {
        uint256 stakeAmount = 1 ether;
        
        // Both users stake
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        vm.prank(user2);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Transfer TG from user1 to keeper
        uint256 tgAmount = 300;
        
        vm.expectEmit(true, true, false, true);
        emit TGTransferred(user1, keeper, tgAmount);
        
        vm.prank(owner);
        stakeRegistry.updateTGBalances(keeper, user1, tgAmount);
        
        // Check user1's TG decreased
        (uint256 amount, uint256 TGbalance) = stakeRegistry.getStake(user1);
        assertEq(amount, stakeAmount);
        assertEq(TGbalance, stakeAmount * 1000 - tgAmount);
        
        // Check keeper's TG increased
        (amount, TGbalance) = stakeRegistry.getStake(keeper);
        assertEq(amount, 0); // No ETH staked
        assertEq(TGbalance, tgAmount); // TG transferred
    }
    
    function test_UpdateTGBalances_OnlyOwner() public {
        uint256 stakeAmount = 1 ether;
        
        // User 1 stakes
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user2));
        vm.prank(user2); // Not owner
        stakeRegistry.updateTGBalances(keeper, user1, 100);
    }
    
    function test_UpdateTGBalances_InsufficientTG() public {
        uint256 stakeAmount = 1 ether;
        
        // User 1 stakes
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        // Try to transfer more TG than available
        uint256 tgAmount = stakeAmount * 1000 + 1; // 1 more than available
        
        vm.expectRevert("Insufficient user TG balance");
        vm.prank(owner);
        stakeRegistry.updateTGBalances(keeper, user1, tgAmount);
    }
    
    function test_Upgrade() public {
        // Deploy new implementation
        TriggerXStakeRegistry newImplementation = new TriggerXStakeRegistry();
        
        // Upgrade to new implementation
        vm.prank(owner);
        stakeRegistry.upgradeToAndCall(address(newImplementation), "");
        
        // Verify upgrade was successful
        // This would typically involve checking that the implementation address changed
        // and that functionality works as expected
    }
    
    function test_Upgrade_OnlyOwner() public {
        // Deploy new implementation
        TriggerXStakeRegistry newImplementation = new TriggerXStakeRegistry();
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1); // Not owner
        stakeRegistry.upgradeToAndCall(address(newImplementation), "");
    }

    function test_WithdrawETH() public {
        uint256 stakeAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        uint256 withdrawAmount = 0.5 ether;
        string memory reason = "Test withdrawal";
        
        uint256 ownerBalanceBefore = owner.balance;
        
        vm.expectEmit(true, false, false, true);
        emit ETHWithdrawn(owner, withdrawAmount, reason);
        
        vm.prank(owner);
        stakeRegistry.withdrawETH(withdrawAmount, reason);
        
        // Check owner received ETH
        assertEq(owner.balance, ownerBalanceBefore + withdrawAmount);
        
        // Check contract balance decreased
        assertEq(address(stakeRegistry).balance, stakeAmount - withdrawAmount);
    }
    
    function test_WithdrawETH_OnlyOwner() public {
        uint256 stakeAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1); // Not owner
        stakeRegistry.withdrawETH(0.5 ether, "Test");
    }
    
    function test_WithdrawETH_ZeroAmount() public {
        vm.expectRevert("Cannot withdraw 0 ETH");
        vm.prank(owner);
        stakeRegistry.withdrawETH(0, "Test");
    }
    
    function test_WithdrawETH_InsufficientBalance() public {
        uint256 stakeAmount = 1 ether;
        
        // First, add some ETH to the contract
        vm.prank(user1);
        stakeRegistry.stake{value: stakeAmount}(stakeAmount);
        
        vm.expectRevert("Insufficient contract balance");
        vm.prank(owner);
        stakeRegistry.withdrawETH(stakeAmount + 1, "Test");
    }
} 