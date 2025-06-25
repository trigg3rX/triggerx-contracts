// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {TestableProxyHub} from "../src/TestableProxyHub.sol";

contract DirectBroadcastTest is Script {
    // Base Sepolia LayerZero Endpoint
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    
    uint32 constant BASE_SEPOLIA_EID = 40245;
    uint32 constant HOLESKY_EID = 40217;
    
    
    function generateMockEids(uint256 count) internal pure returns (uint32[] memory) {
        uint32[] memory eids = new uint32[](count);
        uint32 startEid = 70000; // Start from a high number to avoid conflicts
        
        for (uint256 i = 0; i < count; i++) {
            eids[i] = startEid + uint32(i);
        }
        
        return eids;
    }
    
    function deployTestHub(
        uint256 deployerPrivateKey,
        address deployerAddress,
        uint256 spokeCount
    ) internal returns (TestableProxyHub) {
        console.log("=== Deploying Testable ProxyHub ===");
        console.log("Spoke count:", spokeCount);
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Create initial keepers array
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = deployerAddress;
        
        // Deploy TestableProxyHub with unique salt
        TestableProxyHub hub = TestableProxyHub(payable(address(new TestableProxyHub(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress, HOLESKY_EID, BASE_SEPOLIA_EID, initialKeepers))));
        console.log("Testable ProxyHub deployed at:", address(hub));
        
        // Generate mock spoke EIDs
        uint32[] memory spokeEids = generateMockEids(spokeCount);
        
        // Add all spoke endpoints to Hub
        hub.addSpokes(spokeEids);
        console.log("Added spoke endpoints:", spokeCount);
        
        // Fund the hub with enough ETH for testing
        vm.deal(address(hub), 100 ether);
        console.log("Funded hub with 100 ETH for testing");
        
        vm.stopBroadcast();
        
        return hub;
    }
    
    function testBroadcastOperations(TestableProxyHub hub, uint256 spokeCount) internal {
        console.log("\n=== Testing Direct Broadcast Operations ===");
        
        address testKeeper = address(0x2222222222222222222222222222222222222222);
        
        vm.startBroadcast();
        
        console.log("Executing keeper registration broadcast...");
        // This will show up in gas reporting
        hub.testRegisterBroadcast(testKeeper);
        
        console.log("Executing keeper unregistration broadcast...");
        // This will show up in gas reporting  
        // hub.testUnregisterBroadcast(testKeeper);
        
        vm.stopBroadcast();
        
        console.log("Broadcast operations completed for", spokeCount, "spokes");
        console.log("Check the transaction gas usage above for actual costs");
    }
    
    function runBroadcastTest(uint256 spokeCount, string memory testName) internal {
        console.log(string.concat("\n==================== ", testName, " ===================="));
        console.log("Testing direct broadcast to spokes:", spokeCount);
        
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        
        // Deploy testable hub (with broadcast for actual deployment)
        TestableProxyHub hub = deployTestHub(deployerPrivateKey, deployerAddress, spokeCount);
        
        // Test broadcast operations (with broadcast to see actual gas costs)
        testBroadcastOperations(hub, spokeCount);
        
        console.log("==================== End", testName, "====================\n");
    }
    
    function run() external {
        console.log("Starting Direct Broadcast Gas Test");
        console.log("This will directly test _batchBroadcast functionality\n");
        
        // Test different spoke counts
        runBroadcastTest(10, "Direct Test (10 spokes)");
        runBroadcastTest(50, "Direct Test (50 spokes)");
        runBroadcastTest(100, "Direct Test (100 spokes)");
        runBroadcastTest(200, "Direct Test (200 spokes)");
        runBroadcastTest(500, "Direct Test (500 spokes)");
        
        console.log("\n=== Direct Broadcast Test Complete ===");
        console.log("These results show the actual gas costs of broadcasting to hundreds of L2s");
    }
    
    /// @notice Public function to test broadcasting to a specific number of spokes
    /// @dev Can be called with: forge script --sig "runBroadcastOnce(uint256)" <script> <spokeCount>
    function runBroadcastOnce(uint256 spokeCount) external {
        require(spokeCount > 0 && spokeCount <= 1000, "Invalid spoke count: must be 1-1000");
        
        string memory testName = string.concat("Single Test (", vm.toString(spokeCount), " spokes)");
        runBroadcastTest(spokeCount, testName);
        
        console.log("=== Single Broadcast Test Complete ===");
    }
} 