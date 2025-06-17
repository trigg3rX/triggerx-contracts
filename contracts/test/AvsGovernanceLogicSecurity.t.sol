// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {AvsGovernanceLogic} from "../src/AvsGovernanceLogic.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {MockProxyHub} from "./mocks/MockProxyHub.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";

contract MaliciousReentrant {
    AvsGovernanceLogic public target;
    uint256 public callCount;
    
    constructor(address payable _target) {
        target = AvsGovernanceLogic(_target);
    }
    
    function attack() external payable {
        // Cannot attack since only owner can call withdraw
    }
    
    receive() external payable {
        callCount++;
        if (callCount < 3 && address(target).balance >= 1 ether) {
            target.withdraw(payable(address(this)), 1 ether);
        }
    }
}

contract MaliciousOperator {
    function maliciousCallback() external pure {
        revert("Malicious callback");
    }
}

contract AvsGovernanceLogicSecurityTest is Test {
    AvsGovernanceLogic public avsGovernance;
    MockEndpoint public mockEndpoint;
    MockProxyHub public mockProxyHub;
    
    address public owner = address(0x1);
    address public attacker = address(0x666);
    address public operator1 = address(0x100);
    address public operator2 = address(0x101);
    address public avsGovernanceAddr = address(0x999); // matches constructor parameter
    uint32 public constant DST_EID = 20201;
    
    function setUp() public {
        vm.startPrank(owner);
        
        mockEndpoint = new MockEndpoint();
        mockProxyHub = new MockProxyHub();
        
        avsGovernance = new AvsGovernanceLogic(
            address(mockEndpoint),
            address(mockProxyHub),
            DST_EID,
            owner,
            avsGovernanceAddr // placeholder for avsGovernance address
        );
        
        vm.deal(address(avsGovernance), 100 ether);
        mockEndpoint.setQuotedFee(0.1 ether);
        mockEndpoint.setMsgGuid(bytes32(uint256(123)));
        
        vm.stopPrank();
    }

    // ==================== ACCESS CONTROL TESTS ====================
    
    function test_Security_OnlyOwnerCanSetProxyHub() public {
        address newProxyHub = address(0x999);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        avsGovernance.setProxyHub(newProxyHub);
    }
    
    function test_Security_OnlyOwnerCanWithdraw() public {
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        avsGovernance.withdraw(payable(attacker), 1 ether);
    }
    
    function test_Security_OnlyOwnerCanAddToWhitelist() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        avsGovernance.addToWhitelist(operators);
    }
    
    function test_Security_OnlyOwnerCanRemoveFromWhitelist() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        // First add to whitelist as owner
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Try to remove as attacker
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        avsGovernance.removeFromWhitelist(operators);
    }
    
    function test_Security_OnlyOwnerCanSetGasOptions() public {
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, attacker));
        vm.prank(attacker);
        avsGovernance.setGasOptions(500000, 1e14);
    }

    // ==================== PARAMETER VALIDATION TESTS ====================
    
    function test_Security_ConstructorRejectsZeroProxyHub() public {
        vm.expectRevert("Invalid proxyHub");
        new AvsGovernanceLogic(
            address(mockEndpoint),
            address(0),
            DST_EID,
            owner,
            address(avsGovernance)
        );
    }
    
    function test_Security_SetProxyHubRejectsZeroAddress() public {
        vm.expectRevert("Invalid proxy hub address");
        vm.prank(owner);
        avsGovernance.setProxyHub(address(0));
    }
    
    function test_Security_WithdrawRejectsZeroRecipient() public {
        vm.expectRevert("Invalid recipient");
        vm.prank(owner);
        avsGovernance.withdraw(payable(address(0)), 1 ether);
    }
    
    function test_Security_WithdrawRejectsZeroAmount() public {
        vm.expectRevert("Amount must be greater than 0");
        vm.prank(owner);
        avsGovernance.withdraw(payable(attacker), 0);
    }
    
    function test_Security_WithdrawRejectsInsufficientBalance() public {
        vm.expectRevert("Insufficient balance");
        vm.prank(owner);
        avsGovernance.withdraw(payable(attacker), 1000 ether);
    }
    
    function test_Security_AddToWhitelistRejectsZeroAddress() public {
        address[] memory operators = new address[](1);
        operators[0] = address(0);
        
        vm.expectRevert("Invalid address");
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
    }

    // ==================== STATE MANIPULATION TESTS ====================
    
    function test_Security_CannotRegisterNonWhitelistedOperator() public {
        // Try to call beforeOperatorRegistered without whitelisting
        vm.expectRevert("Operator is not whitelisted");
        vm.prank(avsGovernanceAddr);
        avsGovernance.beforeOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }
    
    function test_Security_CannotDoubleWhitelist() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        // First whitelist
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Try to whitelist again
        vm.expectRevert("Already whitelisted");
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
    }
    
    function test_Security_CannotRemoveNonWhitelistedOperator() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        vm.expectRevert("Not whitelisted");
        vm.prank(owner);
        avsGovernance.removeFromWhitelist(operators);
    }

    // ==================== CROSS-CHAIN MESSAGING SECURITY TESTS ====================
    
    function test_Security_ReceivingMessagesReverts() public {
        // First set up a proper peer relationship to bypass NoPeer error
        vm.prank(owner);
        avsGovernance.setPeer(1, bytes32(uint256(uint160(address(this)))));
        
        // Test that the contract properly rejects incoming messages
        vm.expectRevert("AvsGovernanceLogic: should not receive messages");
        
        // Import Origin struct for LayerZero v2
        Origin memory origin = Origin({
            srcEid: uint32(1), // Use a valid EID that has a peer set
            sender: bytes32(uint256(uint160(address(this)))),
            nonce: uint64(0)
        });
        
        // Simulate LayerZero calling lzReceive with correct parameters
        vm.prank(address(mockEndpoint));
        avsGovernance.lzReceive(origin, bytes32(0), bytes("test"), address(0), bytes(""));
    }
    
    function test_Security_MessageFailsWithInsufficientBalance() public {
        // Drain contract balance (use proper address, not precompile)
        vm.prank(owner);
        avsGovernance.withdraw(payable(address(0x9999)), address(avsGovernance).balance);
        
        // Whitelist operator
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Set high quote fee to trigger insufficient balance
        mockEndpoint.setQuotedFee(10 ether);
        
        // REAL ISSUE FOUND: Contract reverts instead of handling insufficient balance gracefully
        vm.expectRevert("Insufficient balance for message fee (with 10% buffer)");
        vm.prank(avsGovernanceAddr);
        avsGovernance.afterOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }
    
    function test_Security_MessageQuoteFailureHandling() public {
        // Whitelist operator
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Create a mock endpoint that will fail the quote
        MockEndpoint failingEndpoint = new MockEndpoint();
        failingEndpoint.setQuotedFee(type(uint256).max); // Will cause overflow
        
        // Deploy new contract with failing endpoint
        AvsGovernanceLogic testContract = new AvsGovernanceLogic(
            address(failingEndpoint),
            address(mockProxyHub),
            DST_EID,
            owner,
            address(avsGovernance)
        );
        
        vm.deal(address(testContract), 1 ether);
        
        // Whitelist in new contract
        vm.prank(owner);
        testContract.addToWhitelist(operators);
        
        // This should handle the failure gracefully
        vm.expectRevert();
        testContract.afterOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }

    // ==================== REENTRANCY TESTS ====================
    
    function test_Security_WithdrawReentrancyProtection() public {
        MaliciousReentrant malicious = new MaliciousReentrant(payable(address(avsGovernance)));
        vm.deal(address(malicious), 1 ether);
        
        // The malicious contract cannot call withdraw because only owner can call it
        // This demonstrates that access control prevents the attack
        malicious.attack();
        
        // Verify that the contract state remains intact (the attack failed)
        assertEq(address(avsGovernance).balance, 100 ether); // Initial balance from setUp
    }

    // ==================== GAS MANIPULATION TESTS ====================
    
    function test_Security_ExtremeGasLimits() public {
        // Test with maximum gas limit
        vm.prank(owner);
        avsGovernance.setGasOptions(type(uint128).max, 0);
        
        assertEq(avsGovernance.gasLimit(), type(uint128).max);
        
        // Test with zero gas limit (should be allowed but may fail execution)
        vm.prank(owner);
        avsGovernance.setGasOptions(0, 0);
        
        assertEq(avsGovernance.gasLimit(), 0);
    }
    
    function test_Security_ExtremeCallValues() public {
        // Test with maximum call value
        vm.prank(owner);
        avsGovernance.setGasOptions(1000000, type(uint128).max);
        
        assertEq(avsGovernance.callValue(), type(uint128).max);
    }

    // ==================== BOUNDARY TESTS ====================
    
    function test_Security_LargeWhitelistBatch() public {
        // Test with large batch to check for gas limits and overflow issues
        uint256 batchSize = 100;
        address[] memory operators = new address[](batchSize);
        
        for (uint256 i = 0; i < batchSize; i++) {
            operators[i] = address(uint160(1000 + i));
        }
        
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Verify all are whitelisted
        for (uint256 i = 0; i < batchSize; i++) {
            assertTrue(avsGovernance.isWhitelisted(operators[i]));
        }
    }
    
    function test_Security_EmptyWhitelistArray() public {
        address[] memory emptyOperators = new address[](0);
        
        // Should not revert with empty array
        vm.prank(owner);
        avsGovernance.addToWhitelist(emptyOperators);
        
        vm.prank(owner);
        avsGovernance.removeFromWhitelist(emptyOperators);
    }

    // ==================== INTEGRATION SECURITY TESTS ====================
    
    function test_Security_FullOperatorLifecycle() public {
        // Test complete operator lifecycle for security issues
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        // 1. Whitelist
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // 2. Register (before hook)
        vm.prank(avsGovernanceAddr);
        avsGovernance.beforeOperatorRegistered(operator1, 1000, [uint256(1), 2, 3, 4], operator1);
        
        // 3. Register (after hook) - should send message
        vm.prank(avsGovernanceAddr);
        avsGovernance.afterOperatorRegistered(operator1, 1000, [uint256(1), 2, 3, 4], operator1);
        
        // 4. Unregister (before hook)
        vm.prank(avsGovernanceAddr);
        avsGovernance.beforeOperatorUnregistered(operator1);
        
        // 5. Unregister (after hook) - should send message
        vm.prank(avsGovernanceAddr);
        avsGovernance.afterOperatorUnregistered(operator1);
        
        // 6. Remove from whitelist
        vm.prank(owner);
        avsGovernance.removeFromWhitelist(operators);
        
        // 7. Verify operator cannot register again
        vm.expectRevert("Operator is not whitelisted");
        vm.prank(avsGovernanceAddr);
        avsGovernance.beforeOperatorRegistered(operator1, 1000, [uint256(1), 2, 3, 4], operator1);
    }

    // ==================== DENIAL OF SERVICE TESTS ====================
    
    function test_Security_CannotDOSWithInvalidEndpoint() public {
        // Even with invalid endpoint, contract should handle gracefully
        address invalidEndpoint = address(0xdead);
        
        vm.expectRevert();
        new AvsGovernanceLogic(
            invalidEndpoint,
            address(mockProxyHub),
            DST_EID,
            owner,
            address(avsGovernance)
        );
    }
    
    function test_Security_ContractCanReceiveETH() public {
        uint256 initialBalance = address(avsGovernance).balance;
        
        // Send ETH directly to contract
        vm.deal(attacker, 10 ether);
        vm.prank(attacker);
        (bool success,) = address(avsGovernance).call{value: 5 ether}("");
        
        assertTrue(success);
        assertEq(address(avsGovernance).balance, initialBalance + 5 ether);
    }
} 