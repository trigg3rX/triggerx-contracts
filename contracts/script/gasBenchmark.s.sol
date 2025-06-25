// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";

contract GasBenchmark is Script {
    // Base Sepolia LayerZero Endpoint
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    
    uint32 constant BASE_SEPOLIA_EID = 40245;
    uint32 constant HOLESKY_EID = 40217;
    
    bytes32 constant BENCHMARK_SALT = bytes32(keccak256("GAS_BENCHMARK_SALT"));
    
    struct GasMetrics {
        uint256 spokeCount;
        uint256 deploymentGas;
        uint256 addSpokesGas;
        uint256 registrationGas;
        uint256 unregistrationGas;
        uint256 totalGas;
        uint256 ethConsumed;
        uint256 gasPerSpoke;
    }
    
    GasMetrics[] public benchmarkResults;
    
    // ActionType enum matching ProxyHub
    enum ActionType {
        REGISTER,
        UNREGISTER
    }
    
    /**
     * @notice Simulates _lzReceive by calling the LayerZero endpoint directly
     * @param hub The ProxyHub contract
     * @param action 0 for REGISTER, 1 for UNREGISTER
     * @param keeper The keeper address
     */
    function simulateLzReceive(ProxyHub hub, uint8 action, address keeper) internal {
        // Create the Origin struct
        Origin memory origin = Origin({
            srcEid: HOLESKY_EID,
            sender: bytes32(uint256(uint160(address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8)))), // Mock sender from Holesky
            nonce: 1
        });
        
        // Create the payload (ActionType, address)
        ActionType actionType = action == 0 ? ActionType.REGISTER : ActionType.UNREGISTER;
        bytes memory payload = abi.encode(actionType, keeper);
        
        // Use vm.prank to simulate the call coming from the LayerZero endpoint
        // This bypasses the OnlyEndpoint access control
        vm.prank(LZ_ENDPOINT_BASE_SEPOLIA);
        try hub.lzReceive(
            origin,
            bytes32(uint256(1)), // guid
            payload,
            address(this), // executor
            "" // extraData
        ) {
            console.log("LzReceive simulation successful");
        } catch (bytes memory reason) {
            console.log("LzReceive simulation failed:");
            console.logBytes(reason);
            
            // Let's also try to decode the error if it's a standard revert
            if (reason.length >= 4) {
                bytes4 selector = bytes4(reason);
                console.log("Error selector:", vm.toString(uint256(uint32(selector))));
            }
        }
    }
    
    function generateMockEids(uint256 count) internal pure returns (uint32[] memory) {
        uint32[] memory eids = new uint32[](count);
        uint32 startEid = 60000; // Start from a high number to avoid conflicts
        
        for (uint256 i = 0; i < count; i++) {
            eids[i] = startEid + uint32(i);
        }
        
        return eids;
    }
    
    function benchmarkDeployment(
        uint256 deployerPrivateKey,
        address deployerAddress,
        uint256 spokeCount
    ) internal returns (ProxyHub, uint256) {
        console.log("Benchmarking deployment for", spokeCount, "spokes");
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Create initial keepers
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = deployerAddress;
        
        // Measure deployment gas
        uint256 gasStart = gasleft();
        
        bytes memory hubBytecode = abi.encodePacked(
            type(ProxyHub).creationCode,
            abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress, HOLESKY_EID, BASE_SEPOLIA_EID, initialKeepers)
        );
        
        // Use unique salt for each deployment
        bytes32 uniqueSalt = keccak256(abi.encodePacked(BENCHMARK_SALT, spokeCount));
        address hubAddr = CREATE3.deployDeterministic(hubBytecode, uniqueSalt);
        
        uint256 deploymentGas = gasStart - gasleft();
        
        ProxyHub hub = ProxyHub(payable(hubAddr));
        
        // Fund the hub
        vm.deal(address(hub), 1000 ether);
        
        vm.stopBroadcast();
        
        return (hub, deploymentGas);
    }
    
    function benchmarkAddSpokes(
        ProxyHub hub,
        uint256 deployerPrivateKey,
        uint256 spokeCount
    ) internal returns (uint256) {
        console.log("Benchmarking addSpokes for", spokeCount, "spokes");
        
        vm.startBroadcast(deployerPrivateKey);
        
        uint32[] memory spokeEids = generateMockEids(spokeCount);
        
        uint256 gasStart = gasleft();
        hub.addSpokes(spokeEids);
        uint256 addSpokesGas = gasStart - gasleft();
        
        vm.stopBroadcast();
        
        return addSpokesGas;
    }
    
    function benchmarkBroadcastOperations(
        ProxyHub hub,
        uint256 deployerPrivateKey
    ) internal returns (uint256 registrationGas, uint256 unregistrationGas, uint256 ethConsumed) {
        console.log("Benchmarking broadcast operations");
        
        address testKeeper = address(0x1111111111111111111111111111111111111111);
        uint256 initialBalance = address(hub).balance;
        
        // Benchmark registration via _lzReceive simulation
        uint256 gasStart = gasleft();
        simulateLzReceive(hub, 0, testKeeper); // 0 = REGISTER
        registrationGas = gasStart - gasleft();
        
        // Benchmark unregistration via _lzReceive simulation
        gasStart = gasleft();
        simulateLzReceive(hub, 1, testKeeper); // 1 = UNREGISTER  
        unregistrationGas = gasStart - gasleft();
        
        uint256 finalBalance = address(hub).balance;
        ethConsumed = initialBalance - finalBalance;
        
        return (registrationGas, unregistrationGas, ethConsumed);
    }
    
    function runBenchmark(uint256 spokeCount) internal {
        console.log(string.concat("\n=== Benchmarking ", vm.toString(spokeCount), " Spokes ==="));
        
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        
        // Create fresh fork for each benchmark
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        
        // Benchmark deployment
        (ProxyHub hub, uint256 deploymentGas) = benchmarkDeployment(deployerPrivateKey, deployerAddress, spokeCount);
        
        // Benchmark adding spokes
        uint256 addSpokesGas = benchmarkAddSpokes(hub, deployerPrivateKey, spokeCount);
        
        // Benchmark broadcast operations
        (uint256 registrationGas, uint256 unregistrationGas, uint256 ethConsumed) = 
            benchmarkBroadcastOperations(hub, deployerPrivateKey);
        
        // Calculate metrics
        uint256 totalGas = deploymentGas + addSpokesGas + registrationGas + unregistrationGas;
        uint256 gasPerSpoke = spokeCount > 0 ? (registrationGas + unregistrationGas) / spokeCount : 0;
        
        // Store results
        GasMetrics memory metrics = GasMetrics({
            spokeCount: spokeCount,
            deploymentGas: deploymentGas,
            addSpokesGas: addSpokesGas,
            registrationGas: registrationGas,
            unregistrationGas: unregistrationGas,
            totalGas: totalGas,
            ethConsumed: ethConsumed,
            gasPerSpoke: gasPerSpoke
        });
        
        benchmarkResults.push(metrics);
        
        // Log detailed results
        console.log("Deployment Gas:", deploymentGas);
        console.log("Add Spokes Gas:", addSpokesGas);
        console.log("Registration Gas:", registrationGas);
        console.log("Unregistration Gas:", unregistrationGas);
        console.log("Total Gas:", totalGas);
        console.log("ETH Consumed:", ethConsumed);
        console.log("Gas per Spoke:", gasPerSpoke);
        console.log("=== End Benchmark ===\n");
    }
    
    function generateReport() internal view {
        console.log("\n========================= GAS BENCHMARK REPORT =========================");
        console.log("| Spokes | Deployment | AddSpokes | Registration | Unregistration | Total Gas | Gas/Spoke |");
        console.log("|--------|------------|-----------|--------------|----------------|-----------|-----------|");
        
        for (uint256 i = 0; i < benchmarkResults.length; i++) {
            GasMetrics memory metrics = benchmarkResults[i];
            console.log(
                string.concat(
                    "| ", _padNumber(metrics.spokeCount, 6),
                    " | ", _padNumber(metrics.deploymentGas, 10),
                    " | ", _padNumber(metrics.addSpokesGas, 9),
                    " | ", _padNumber(metrics.registrationGas, 12),
                    " | ", _padNumber(metrics.unregistrationGas, 14),
                    " | ", _padNumber(metrics.totalGas, 9),
                    " | ", _padNumber(metrics.gasPerSpoke, 9), " |"
                )
            );
        }
        
        console.log("=========================================================================");
        
        // Calculate scaling analysis
        if (benchmarkResults.length > 1) {
            console.log("\n=== SCALING ANALYSIS ===");
            
            GasMetrics memory first = benchmarkResults[0];
            GasMetrics memory last = benchmarkResults[benchmarkResults.length - 1];
            
            uint256 spokeMultiplier = last.spokeCount / first.spokeCount;
            uint256 gasMultiplier = last.totalGas / first.totalGas;
            
            console.log("Spoke count multiplier:", spokeMultiplier);
            console.log("Gas multiplier:", gasMultiplier);
            
            if (gasMultiplier <= spokeMultiplier) {
                console.log("Gas scaling is LINEAR or BETTER");
            } else {
                console.log("Gas scaling is WORSE than linear");
            }
            
            console.log("Average gas per broadcast operation:", (last.registrationGas + last.unregistrationGas) / 2);
        }
    }
    
    function _padNumber(uint256 number, uint256 width) internal pure returns (string memory) {
        string memory str = vm.toString(number);
        bytes memory strBytes = bytes(str);
        
        if (strBytes.length >= width) {
            return str;
        }
        
        bytes memory padded = new bytes(width);
        uint256 spaces = width - strBytes.length;
        
        for (uint256 i = 0; i < spaces; i++) {
            padded[i] = " ";
        }
        
        for (uint256 i = 0; i < strBytes.length; i++) {
            padded[spaces + i] = strBytes[i];
        }
        
        return string(padded);
    }
    
    function run() external {
        console.log("Starting Comprehensive Gas Benchmark for ProxyHub Broadcasting");
        console.log("This will test broadcasting to various numbers of L2 spokes\n");
        
        // Test different spoke counts
        uint256[] memory spokeCounts = new uint256[](8);
        spokeCounts[0] = 1;
        spokeCounts[1] = 10;
        spokeCounts[2] = 25;
        spokeCounts[3] = 50;
        spokeCounts[4] = 100;
        spokeCounts[5] = 200;
        spokeCounts[6] = 300;
        spokeCounts[7] = 500;
        
        for (uint256 i = 0; i < spokeCounts.length; i++) {
            runBenchmark(spokeCounts[i]);
        }
        
        // Generate comprehensive report
        generateReport();
        
        console.log("\nBenchmark Complete!");
        console.log("Use this data to estimate gas costs for scaling to hundreds of L2s");
    }
} 