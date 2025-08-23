// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";
import {CREATE3} from "lib/solady/src/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasRegistryDeploy is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);
    address operator = 0x2469e89386947535A350EEBccC5F2754fd35F474; // Task Execution Address
    uint256 tgPerEth = 1000;

    bytes32 SALT = keccak256(abi.encodePacked("TriggerGasRegistry_v1"));
    bytes32 IMPL_SALT = keccak256(abi.encodePacked("TriggerGasRegistryImpl_v1"));

    struct DeploymentInfo {
        string chainName;
        string rpcUrl;
        address proxy;
        address implementation;
    }

    function run() public {
        // Define deployment configurations for each chain
        DeploymentInfo[] memory deployments = new DeploymentInfo[](4);

        deployments[0] = DeploymentInfo({
            chainName: "Ethereum",
            rpcUrl: vm.envString("ETH_RPC"),
            proxy: address(0),
            implementation: address(0)
        });

        deployments[1] = DeploymentInfo({
            chainName: "Base",
            rpcUrl: vm.envString("BASE_RPC"),
            proxy: address(0),
            implementation: address(0)
        });
        
        deployments[2] = DeploymentInfo({
            chainName: "Optimism",
            rpcUrl: vm.envString("OPTIMISM_RPC"),
            proxy: address(0),
            implementation: address(0)
        });
        
        deployments[3] = DeploymentInfo({
            chainName: "Arbitrum",
            rpcUrl: vm.envString("ARBITRUM_RPC"),
            proxy: address(0),
            implementation: address(0)
        });

        console.log("Starting multi-chain deployment...");
        console.log("Deployer:", deployer);
        console.log("Operator:", operator);
        console.log("TG_PER_ETH:", tgPerEth);
        console.log("");

        // Deploy on each chain
        for (uint256 i = 0; i < deployments.length; i++) {
            deployOnChain(deployments[i]);
        }

        // Print summary
        console.log("");
        console.log("=== DEPLOYMENT SUMMARY ===");
        for (uint256 i = 0; i < deployments.length; i++) {
            console.log("Chain Name:", deployments[i].chainName);
            console.log("  Proxy:", deployments[i].proxy);
            console.log("  Implementation:", deployments[i].implementation);
            console.log("");
        }
    }

    function deployOnChain(
        DeploymentInfo memory deployment
    ) internal {
        console.log("Deploying on", deployment.chainName, "...");
        
        // Create fork for this chain
        vm.createSelectFork(deployment.rpcUrl);

        bytes memory implementation_code = type(TriggerGasRegistry).creationCode;

        vm.startBroadcast(deployerPrivateKey);

        // Deploy implementation
        address implementation = CREATE3.deployDeterministic(implementation_code, IMPL_SALT);

        // Deploy proxy
        bytes memory initData = abi.encodeWithSelector(
            TriggerGasRegistry.initialize.selector, 
            deployer, 
            operator, 
            tgPerEth
        );
        bytes memory proxy_code = abi.encodePacked(
            type(ERC1967Proxy).creationCode, 
            abi.encode(address(implementation), initData)
        );

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        // Store deployment info
        deployment.proxy = proxy;
        deployment.implementation = implementation;

        // Verify deployment
        TriggerGasRegistry registry = TriggerGasRegistry(payable(proxy));
        require(registry.owner() == deployer, "Owner mismatch");
        require(registry.operatorRole() == operator, "Operator mismatch");
        require(registry.TG_PER_ETH() == tgPerEth, "TG_PER_ETH mismatch");

        vm.stopBroadcast();
        console.log("Deployment on", deployment.chainName, "completed successfully");
        console.log("");
    }
}