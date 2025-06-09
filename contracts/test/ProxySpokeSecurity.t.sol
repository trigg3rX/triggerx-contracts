// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";

contract MaliciousSpokeTarget {
    bool public shouldRevert;
    bool public shouldConsumeMuchGas;
    ProxySpoke public spoke;
    
    constructor(address _spoke) {
        spoke = ProxySpoke(_spoke);
    }
    
    function setRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }
    
    function setConsumeMuchGas(bool _shouldConsume) external {
        shouldConsumeMuchGas = _shouldConsume;
    }
    
    function maliciousFunction() external payable {
        if (shouldRevert) {
            revert("Malicious revert");
        }
        
        if (shouldConsumeMuchGas) {
            // Consume much gas
            for (uint256 i = 0; i < 10000; i++) {
                keccak256(abi.encode(i, block.timestamp));
            }
        }
    }
    
    receive() external payable {
        if (shouldRevert) {
            revert("Rejecting ETH");
        }
    }
}

contract ProxySpokeSecurityTest is Test {
    ProxySpoke public proxySpoke;
    MockEndpoint public mockEndpoint;
    
    address public owner = address(0x1);
    address public attacker = address(0x666);
    address public keeper1 = address(0x100);
    address public keeper2 = address(0x101);
    
    uint32 public constant HUB_EID = 10101;
    
    function setUp() public {
        vm.startPrank(owner);
        
        mockEndpoint = new MockEndpoint();
        
        address[] memory initialKeepers = new address[](2);
        initialKeepers[0] = keeper1;
        initialKeepers[1] = keeper2;
        
        proxySpoke = new ProxySpoke(
            address(mockEndpoint),
            owner,
            HUB_EID,
            initialKeepers
        );
        
        vm.stopPrank();
        
        vm.deal(keeper1, 10 ether);
        vm.deal(keeper2, 10 ether);
        vm.deal(attacker, 10 ether);
    }

    // ==================== ACCESS CONTROL TESTS ====================
    
    function test_Security_OnlyKeeperCanExecuteFunction() public {
        address target = address(0x999);
        bytes memory data = abi.encodeWithSignature("someFunction()");
        
        vm.expectRevert("Spoke: Keeper not registered");
        vm.prank(attacker);
        proxySpoke.executeFunction(target, data);
    }
    
    function test_Security_InitialKeepersAreSetCorrectly() public {
        assertTrue(proxySpoke.isKeeper(keeper1));
        assertTrue(proxySpoke.isKeeper(keeper2));
        assertFalse(proxySpoke.isKeeper(attacker));
    }

    // ==================== CROSS-CHAIN MESSAGING SECURITY TESTS ====================
    
    function test_Security_OnlyValidActionTypesAccepted() public {
        // Set up peer relationship first
        vm.prank(owner);
        proxySpoke.setPeer(HUB_EID, bytes32(uint256(uint160(address(this)))));
        
        // Test invalid action type
        bytes memory invalidPayload = abi.encode(uint8(99), attacker); // Invalid action type
        
        vm.expectRevert();
        vm.prank(address(mockEndpoint));
        proxySpoke.lzReceive(Origin({srcEid: HUB_EID, sender: bytes32(uint256(uint160(address(this)))), nonce: uint64(0)}), bytes32(0), invalidPayload, address(0), bytes(""));
    }
    
    function test_Security_MalformedPayloadHandling() public {
        // Test malformed payload
        bytes memory malformedPayload = abi.encode("invalid", 123, true);
        
        vm.expectRevert();
        vm.prank(address(mockEndpoint));
        proxySpoke.lzReceive(Origin({srcEid: uint32(0), sender: bytes32(0), nonce: uint64(0)}), bytes32(0), malformedPayload, address(0), bytes(""));
    }

    // ==================== FUNCTION EXECUTION SECURITY TESTS ====================
    
    function test_Security_FunctionExecutionWithMaliciousTarget() public {
        MaliciousSpokeTarget maliciousTarget = new MaliciousSpokeTarget(address(proxySpoke));
        maliciousTarget.setRevert(true);
        
        bytes memory data = abi.encodeWithSelector(MaliciousSpokeTarget.maliciousFunction.selector);
        
        vm.expectRevert("Function execution failed");
        vm.prank(keeper1);
        proxySpoke.executeFunction(address(maliciousTarget), data);
    }
    
    function test_Security_CannotExecuteOnSpokeItself() public {
        // Try to call sensitive functions on the ProxySpoke itself
        bytes memory data = abi.encodeWithSelector(ProxySpoke.executeFunction.selector, address(0), "");
        
        vm.expectRevert("Function execution failed");
        vm.prank(keeper1);
        proxySpoke.executeFunction(address(proxySpoke), data);
    }

    // ==================== KEEPER MANAGEMENT SECURITY TESTS ====================
    
    function test_Security_KeeperRegistrationSecurity() public {
        // Set up peer relationship first
        vm.prank(owner);
        proxySpoke.setPeer(HUB_EID, bytes32(uint256(uint160(address(this)))));
        
        // Initially, attacker is not a keeper
        assertFalse(proxySpoke.isKeeper(attacker));
        
        // Register attacker via LayerZero message
        bytes memory payload = abi.encode(ProxySpoke.ActionType.REGISTER, attacker);
        
        vm.prank(address(mockEndpoint));
        proxySpoke.lzReceive(Origin({srcEid: HUB_EID, sender: bytes32(uint256(uint160(address(this)))), nonce: uint64(0)}), bytes32(0), payload, address(0), bytes(""));
        
        // Now attacker should be a keeper
        assertTrue(proxySpoke.isKeeper(attacker));
    }
    
    function test_Security_KeeperUnregistrationSecurity() public {
        address testKeeper = address(0x999);
        
        // Set up peer relationship first
        vm.prank(owner);
        proxySpoke.setPeer(HUB_EID, bytes32(uint256(uint160(address(this)))));
        
        // Register a keeper first
        bytes memory registerPayload = abi.encode(ProxySpoke.ActionType.REGISTER, testKeeper);
        
        vm.prank(address(mockEndpoint));
        proxySpoke.lzReceive(Origin({srcEid: HUB_EID, sender: bytes32(uint256(uint160(address(this)))), nonce: uint64(0)}), bytes32(0), registerPayload, address(0), bytes(""));
        
        assertTrue(proxySpoke.isKeeper(testKeeper));
        
        // Unregister the keeper
        bytes memory unregisterPayload = abi.encode(ProxySpoke.ActionType.UNREGISTER, testKeeper);
        
        vm.prank(address(mockEndpoint));
        proxySpoke.lzReceive(Origin({srcEid: HUB_EID, sender: bytes32(uint256(uint160(address(this)))), nonce: uint64(0)}), bytes32(0), unregisterPayload, address(0), bytes(""));
        
        assertFalse(proxySpoke.isKeeper(testKeeper));
    }

    // ==================== HELPER FUNCTIONS ====================
    
    function dummyFunction() external pure returns (bool) {
        return true;
    }
    
    receive() external payable {
        // Allow contract to receive ETH
    }
} 