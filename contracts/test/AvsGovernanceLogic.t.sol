// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test} from "forge-std/Test.sol";
import {AvsGovernanceLogic} from "../src/AvsGovernanceLogic.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {MockProxyHub} from "./mocks/MockProxyHub.sol";
import {IAvsGovernanceLogic} from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";
import {console2} from "forge-std/console2.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract AvsGovernanceLogicTest is Test {
    AvsGovernanceLogic public avsGovernance;
    MockEndpoint public mockEndpoint;
    MockProxyHub public mockProxyHub;
    address public owner = address(0x1);
    address public operator1 = address(0x100);
    address public operator2 = address(0x101);
    uint32 public constant DST_EID = 20201;
    
    event Whitelisted(address indexed operator);
    event UnWhitelisted(address indexed operator);
    event OperatorRegistered(address indexed operator);
    event OperatorUnregistered(address indexed operator);
    event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee);
    
    function setUp() public {
        console2.log("Starting AvsGovernanceLogic test setup");
        
        vm.startPrank(owner);
        
        // Deploy mocks
        console2.log("Deploying MockEndpoint");
        mockEndpoint = new MockEndpoint();
        console2.log("MockEndpoint deployed at", address(mockEndpoint));
        
        console2.log("Deploying MockProxyHub");
        mockProxyHub = new MockProxyHub();
        console2.log("MockProxyHub deployed at", address(mockProxyHub));
        
        // Deploy AvsGovernanceLogic
        console2.log("Deploying AvsGovernanceLogic");
        try new AvsGovernanceLogic(
            address(mockEndpoint),
            address(mockProxyHub),
            DST_EID,
            owner,
            address(0x999) // placeholder for avsGovernance address
        ) returns (AvsGovernanceLogic avsGov) {
            avsGovernance = avsGov;
            console2.log("AvsGovernanceLogic deployed at", address(avsGovernance));
        } catch Error(string memory reason) {
            console2.log("AvsGovernanceLogic deployment failed with reason:", reason);
            revert(reason);
        } catch (bytes memory reason) {
            console2.log("AvsGovernanceLogic deployment failed with low-level error");
            revert("Low-level error");
        }
        
        // Fund contract for message fees
        vm.deal(address(avsGovernance), 100 ether);
        
        vm.stopPrank();
        console2.log("AvsGovernanceLogic test setup completed");
    }
    
    function test_Constructor() public {
        assertEq(avsGovernance.owner(), owner);
        assertEq(avsGovernance.proxyHub(), address(mockProxyHub));
        assertEq(avsGovernance.dstEid(), DST_EID);
    }
    
    function test_SetProxyHub() public {
        address newProxyHub = address(0x202);
        
        vm.prank(owner);
        avsGovernance.setProxyHub(newProxyHub);
        
        assertEq(avsGovernance.proxyHub(), newProxyHub);
    }
    
    function test_SetProxyHub_OnlyOwner() public {
        address newProxyHub = address(0x202);
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, operator1));
        vm.prank(operator1);
        avsGovernance.setProxyHub(newProxyHub);
    }
    
    function test_Withdraw() public {
        uint256 initialBalance = 10 ether;
        address payable recipient = payable(address(0x300));
        
        // Fund contract
        vm.deal(address(avsGovernance), initialBalance);
        
        // Withdraw funds
        vm.prank(owner);
        avsGovernance.withdraw(recipient, 5 ether);
        
        assertEq(address(avsGovernance).balance, 5 ether);
        assertEq(recipient.balance, 5 ether);
    }
    
    function test_Withdraw_OnlyOwner() public {
        address payable recipient = payable(address(0x300));
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, operator1));
        vm.prank(operator1);
        avsGovernance.withdraw(recipient, 1 ether);
    }
    
    function test_Withdraw_InsufficientBalance() public {
        uint256 initialBalance = 1 ether;
        address payable recipient = payable(address(0x300));
        
        // Fund contract
        vm.deal(address(avsGovernance), initialBalance);
        
        vm.expectRevert("Insufficient balance");
        vm.prank(owner);
        avsGovernance.withdraw(recipient, 2 ether);
    }
    
    function test_AddToWhitelist() public {
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;
        
        vm.expectEmit(true, false, false, false);
        emit Whitelisted(operator1);
        
        vm.expectEmit(true, false, false, false);
        emit Whitelisted(operator2);
        
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        assertTrue(avsGovernance.isWhitelisted(operator1));
        assertTrue(avsGovernance.isWhitelisted(operator2));
    }
    
    function test_AddToWhitelist_AlreadyWhitelisted() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        // First add to whitelist
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Try to add again
        vm.expectRevert("Already whitelisted");
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
    }
    
    function test_RemoveFromWhitelist() public {
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;
        
        // First add to whitelist
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Now remove
        vm.expectEmit(true, false, false, false);
        emit UnWhitelisted(operator1);
        
        vm.expectEmit(true, false, false, false);
        emit UnWhitelisted(operator2);
        
        vm.prank(owner);
        avsGovernance.removeFromWhitelist(operators);
        
        assertFalse(avsGovernance.isWhitelisted(operator1));
        assertFalse(avsGovernance.isWhitelisted(operator2));
    }
    
    function test_RemoveFromWhitelist_NotWhitelisted() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        vm.expectRevert("Not whitelisted");
        vm.prank(owner);
        avsGovernance.removeFromWhitelist(operators);
    }
    
    function test_BeforeOperatorRegistered() public {
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        // Add to whitelist
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // This should succeed because operator1 is whitelisted
        vm.prank(address(0x999)); // Some external contract
        avsGovernance.beforeOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }
    
    function test_BeforeOperatorRegistered_NotWhitelisted() public {
        vm.expectRevert("Operator is not whitelisted");
        vm.prank(address(0x999)); // Some external contract
        avsGovernance.beforeOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }
    
    function test_AfterOperatorRegistered() public {
        // Add operator to whitelist
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        
        vm.prank(owner);
        avsGovernance.addToWhitelist(operators);
        
        // Set up mock endpoint
        uint256 fee = 1 ether;
        mockEndpoint.setQuotedFee(fee);
        mockEndpoint.setMsgGuid(bytes32(uint256(123)));
        
        // Fund contract for message fees
        vm.deal(address(avsGovernance), 2 ether);
        
        // Expect events in the correct order
        vm.expectEmit(true, true, false, true);
        emit MessageSent(DST_EID, bytes32(uint256(123)), fee);
        
        vm.expectEmit(true, false, false, false);
        emit OperatorRegistered(operator1);
        
        // Register operator
        vm.prank(address(0x999));
        avsGovernance.afterOperatorRegistered(operator1, 0, [uint256(0), 0, 0, 0], address(0));
    }
    
    function test_AfterOperatorUnregistered() public {
        // Set up mock endpoint
        uint256 fee = 1 ether;
        mockEndpoint.setQuotedFee(fee);
        mockEndpoint.setMsgGuid(bytes32(uint256(456)));
        
        // Fund contract for message fees
        vm.deal(address(avsGovernance), 2 ether);
        
        // Expect events in the correct order
        vm.expectEmit(true, true, false, true);
        emit MessageSent(DST_EID, bytes32(uint256(456)), fee);
        
        vm.expectEmit(true, false, false, false);
        emit OperatorUnregistered(operator1);
        
        // Unregister operator
        vm.prank(address(0x999));
        avsGovernance.afterOperatorUnregistered(operator1);
    }
    
    function test_ReceiveEther() public {
        uint256 amount = 5 ether;
        address sender = address(0x123);
        
        // Fund the sender
        vm.deal(sender, amount);
        
        uint256 initialBalance = address(avsGovernance).balance;
        
        vm.prank(sender);
        (bool success,) = address(avsGovernance).call{value: amount}("");
        
        assertTrue(success);
        assertEq(address(avsGovernance).balance, initialBalance + amount);
    }
    
    // function test_BuildExecutorOptions() public {
    //     // This is an internal function, so we'll try to call it through a public function
    //     // For this test to work, we'd need to expose a public testing function in the contract
    //     // or use other advanced testing techniques
    //     // This is just a placeholder test
    // }
}  