// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "@openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer-contracts/contracts/core/AVSDirectory.sol";
import "@eigenlayer-contracts/contracts/core/DelegationManager.sol";
import "@eigenlayer-contracts/contracts/core/StrategyManager.sol";

import {IRewardsCoordinator} from "@eigenlayer-contracts/contracts/interfaces/IRewardsCoordinator.sol";
import {IAVSDirectory} from "@eigenlayer-contracts/contracts/interfaces/IAVSDirectory.sol";
import {PauserRegistry} from "@eigenlayer-contracts/contracts/permissions/PauserRegistry.sol";

import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry, IDelegationManager, StakeType} from "@eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "@eigenlayer-middleware/interfaces/IIndexRegistry.sol";
import {IBLSApkRegistry} from "@eigenlayer-middleware/interfaces/IBLSApkRegistry.sol";
import {ISlasher} from "@eigenlayer-middleware/interfaces/ISlasher.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/RegistryCoordinator.sol";
import {IndexRegistry} from "@eigenlayer-middleware/IndexRegistry.sol";
import {StakeRegistry, IStrategy} from "@eigenlayer-middleware/StakeRegistry.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/BLSApkRegistry.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/OperatorStateRetriever.sol";
import {VetoableSlasher} from "@eigenlayer-middleware/slashers/VetoableSlasher.sol";

import {ITriggerXTaskManager} from "../src/interfaces/ITriggerXTaskManager.sol";
import {ITriggerXServiceManager} from "../src/interfaces/ITriggerXServiceManager.sol";
import {IServiceManager} from "@eigenlayer-middleware/interfaces/IServiceManager.sol";
import {TriggerXTaskManager} from "../src/TriggerXTaskManager.sol";
import {TriggerXServiceManager} from "../src/TriggerXServiceManager.sol";

contract TriggerXUpdaterHolesky is Script {
    struct DeployedContracts {
        address proxyAdmin;
        address pauserRegistry;
        address indexRegistryProxy;
        address stakeRegistryProxy;
        address apkRegistryProxy;
        address vetoableSlasherProxy;
        address registryCoordinatorProxy;
        address triggerXServiceManagerProxy;
        address triggerXTaskManagerProxy;
    }

    function run() external {
        // Read deployment addresses from JSON
        DeployedContracts memory deployed;
        address delegationManager;
        address avsDirectory;
        address strategyManager;
        address allocationManager;
        address permissionController;
        address rewardsCoordinator;
        address slasher;
        {
            string memory deploymentJson = vm.readFile("./script/output/deployment.holesky.json");
            string memory configJson = vm.readFile("./script/input/parameters.holesky.json");
            
            // Read from deployment.holesky.json
            deployed.proxyAdmin = abi.decode(vm.parseJson(deploymentJson, ".proxyAdmin"), (address));
            deployed.pauserRegistry = abi.decode(vm.parseJson(deploymentJson, ".pauserRegistry"), (address));
            deployed.indexRegistryProxy = abi.decode(vm.parseJson(deploymentJson, ".indexRegistry.proxy"), (address));
            deployed.stakeRegistryProxy = abi.decode(vm.parseJson(deploymentJson, ".stakeRegistry.proxy"), (address));
            deployed.apkRegistryProxy = abi.decode(vm.parseJson(deploymentJson, ".apkRegistry.proxy"), (address));
            deployed.vetoableSlasherProxy = abi.decode(vm.parseJson(deploymentJson, ".vetoableSlasher.proxy"), (address));
            deployed.registryCoordinatorProxy = abi.decode(vm.parseJson(deploymentJson, ".registryCoordinator.proxy"), (address));
            deployed.triggerXServiceManagerProxy = abi.decode(vm.parseJson(deploymentJson, ".triggerXServiceManager.proxy"), (address));
            deployed.triggerXTaskManagerProxy = abi.decode(vm.parseJson(deploymentJson, ".triggerXTaskManager.proxy"), (address));

            // Read from parameters.holesky.json
            delegationManager = abi.decode(vm.parseJson(configJson, ".delegationManager"), (address));
            avsDirectory = abi.decode(vm.parseJson(configJson, ".avsDirectory"), (address));
            strategyManager = abi.decode(vm.parseJson(configJson, ".strategyManager"), (address));
            allocationManager = abi.decode(vm.parseJson(configJson, ".allocationManager"), (address));
            permissionController = abi.decode(vm.parseJson(configJson, ".permissionController"), (address));
            rewardsCoordinator = abi.decode(vm.parseJson(configJson, ".rewardsCoordinator"), (address));
            slasher = abi.decode(vm.parseJson(configJson, ".txSlasher"), (address));
        }

        ProxyAdmin proxyAdmin = ProxyAdmin(deployed.proxyAdmin);

        vm.startBroadcast();

        // Deploy and upgrade new implementations
        
        // IndexRegistry
        IndexRegistry newIndexRegistry = new IndexRegistry(
            RegistryCoordinator(deployed.registryCoordinatorProxy)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.indexRegistryProxy)),
            address(newIndexRegistry)
        );

        // StakeRegistry
        StakeRegistry newStakeRegistry = new StakeRegistry(
            RegistryCoordinator(deployed.registryCoordinatorProxy),
            IDelegationManager(delegationManager),
            IAVSDirectory(avsDirectory),
            IAllocationManager(allocationManager),
            ITriggerXServiceManager(deployed.triggerXServiceManagerProxy)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.stakeRegistryProxy)),
            address(newStakeRegistry)
        );

        // BLSApkRegistry
        BLSApkRegistry newApkRegistry = new BLSApkRegistry(
            RegistryCoordinator(deployed.registryCoordinatorProxy)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.apkRegistryProxy)),
            address(newApkRegistry)
        );

        // VetoableSlasher
        VetoableSlasher newVetoableSlasher = new VetoableSlasher(
            IAllocationManager(allocationManager),
            IServiceManager(deployed.triggerXServiceManagerProxy),
            slasher
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.vetoableSlasherProxy)),
            address(newVetoableSlasher)
        );

        // RegistryCoordinator
        RegistryCoordinator newRegistryCoordinator = new RegistryCoordinator(
            ITriggerXServiceManager(deployed.triggerXServiceManagerProxy),
            IStakeRegistry(deployed.stakeRegistryProxy),
            IBLSApkRegistry(deployed.apkRegistryProxy),
            IIndexRegistry(deployed.indexRegistryProxy),
            IAllocationManager(allocationManager),
            PauserRegistry(deployed.pauserRegistry)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.registryCoordinatorProxy)),
            address(newRegistryCoordinator)
        );

        // TriggerXServiceManager
        TriggerXServiceManager newServiceManager = new TriggerXServiceManager(
            IAVSDirectory(avsDirectory),
            IRewardsCoordinator(rewardsCoordinator),
            IRegistryCoordinator(deployed.registryCoordinatorProxy),
            IStakeRegistry(deployed.stakeRegistryProxy),
            IPermissionController(permissionController),
            PauserRegistry(deployed.pauserRegistry),
            VetoableSlasher(deployed.vetoableSlasherProxy)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.triggerXServiceManagerProxy)),
            address(newServiceManager)
        );

        // TriggerXTaskManager
        TriggerXTaskManager newTaskManager = new TriggerXTaskManager(
            IRegistryCoordinator(deployed.registryCoordinatorProxy),
            PauserRegistry(deployed.pauserRegistry)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(deployed.triggerXTaskManagerProxy)),
            address(newTaskManager)
        );

        // Update the deployment.json file with new implementation addresses
        updateDeploymentJson(DeployedImplementations({
            indexRegistry: address(newIndexRegistry),
            stakeRegistry: address(newStakeRegistry),
            apkRegistry: address(newApkRegistry),
            vetoableSlasher: address(newVetoableSlasher),
            registryCoordinator: address(newRegistryCoordinator),
            triggerXServiceManager: address(newServiceManager),
            triggerXTaskManager: address(newTaskManager)
        }));

        vm.stopBroadcast();
    }

    struct DeployedImplementations {
        address indexRegistry;
        address stakeRegistry;
        address apkRegistry;
        address vetoableSlasher;
        address registryCoordinator;
        address triggerXServiceManager;
        address triggerXTaskManager;
    }

    function updateDeploymentJson(DeployedImplementations memory implementations) internal {
        string memory deploymentJson = vm.readFile("./script/output/deployment.holesky.json");
        
        // Parse the existing JSON into a new string with updated implementation addresses
        string memory updatedJson = string.concat(
            "{\n",
            '  "proxyAdmin": ', vm.parseJson(deploymentJson, ".proxyAdmin"), ',\n',
            '  "pauserRegistry": ', vm.parseJson(deploymentJson, ".pauserRegistry"), ',\n',
            '  "indexRegistry": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".indexRegistry.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.indexRegistry), '"\n',
            '  },\n',
            '  "stakeRegistry": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".stakeRegistry.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.stakeRegistry), '"\n',
            '  },\n',
            '  "apkRegistry": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".apkRegistry.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.apkRegistry), '"\n',
            '  },\n',
            '  "vetoableSlasher": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".vetoableSlasher.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.vetoableSlasher), '"\n',
            '  },\n',
            '  "registryCoordinator": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".registryCoordinator.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.registryCoordinator), '"\n',
            '  },\n',
            '  "operatorStateRetriever": ', vm.parseJson(deploymentJson, ".operatorStateRetriever"), ',\n',
            '  "triggerXServiceManager": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".triggerXServiceManager.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.triggerXServiceManager), '"\n',
            '  },\n',
            '  "triggerXTaskManager": {\n',
            '    "proxy": ', vm.parseJson(deploymentJson, ".triggerXTaskManager.proxy"), ',\n',
            '    "implementation": "', vm.toString(implementations.triggerXTaskManager), '"\n',
            '  }\n',
            "}"
        );

        vm.writeFile("./script/output/deployment.holesky.json", updatedJson);
    }
}
