// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@eigenlayer-contracts/test/mocks/EmptyContract.sol";
import "forge-std/Script.sol";
import "forge-std/console.sol";
import "@openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
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
// import {ISocketRegistry} from "@eigenlayer-middleware/interfaces/ISocketRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/RegistryCoordinator.sol";
import {IndexRegistry} from "@eigenlayer-middleware/IndexRegistry.sol";
import {StakeRegistry, IStrategy} from "@eigenlayer-middleware/StakeRegistry.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/BLSApkRegistry.sol";
// import {SocketRegistry} from "@eigenlayer-middleware/SocketRegistry.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/OperatorStateRetriever.sol";
import {VetoableSlasher} from "@eigenlayer-middleware/slashers/VetoableSlasher.sol";

import {ITriggerXTaskManager} from "../src/interfaces/ITriggerXTaskManager.sol";
import {ITriggerXServiceManager} from "../src/interfaces/ITriggerXServiceManager.sol";

import {TriggerXTaskManager} from "../src/TriggerXTaskManager.sol";
import {TriggerXServiceManager} from "../src/TriggerXServiceManager.sol";

import {IServiceManager} from "@eigenlayer-middleware/interfaces/IServiceManager.sol";

contract TriggerXDeployerHolesky is Script {

    struct TriggerXContracts {
        ITriggerXTaskManager triggerXTaskManager;
        TriggerXTaskManager triggerXTaskManagerImplementation;
        ITriggerXServiceManager triggerXServiceManager;
        TriggerXServiceManager triggerXServiceManagerImplementation;
        RegistryCoordinator registryCoordinator;
        IRegistryCoordinator registryCoordinatorImplementation;
        IIndexRegistry indexRegistry;
        IIndexRegistry indexRegistryImplementation;
        IStakeRegistry stakeRegistry;
        IStakeRegistry stakeRegistryImplementation;
        BLSApkRegistry apkRegistry;
        BLSApkRegistry apkRegistryImplementation;
        VetoableSlasher vetoableSlasher;
        VetoableSlasher vetoableSlasherImplementation;
        // ISocketRegistry socketRegistry;
        // ISocketRegistry socketRegistryImplementation;
        OperatorStateRetriever operatorStateRetriever;
        ProxyAdmin proxyAdmin;
        PauserRegistry pauserRegistry;
    }

    struct EigenLayerContracts {
        AVSDirectory avsDirectory;
        DelegationManager delegationManager;
        StrategyManager strategyManager;
        IRewardsCoordinator rewardsCoordinator;
        IAllocationManager allocationManager;
        IPermissionController permissionController;
        
        address wethStrategy;
        address stETHStrategy;
        address rETHStrategy;
        address cbETHStrategy;
        address reALTStrategy;
        address eigenStrategy;
        address beaconChainETHStrategy;
        
        uint96 wethMultiplier;
        uint96 stETHMultiplier;
        uint96 rETHMultiplier;
        uint96 cbETHMultiplier;
        uint96 reALTMultiplier;
        uint96 eigenMultiplier;
        uint96 beaconChainETHMultiplier;
    }

    struct TokenAndWeight {
        address token;
        uint96 weight;
    }

    struct DeploymentConfig {
        address triggerXOwner;
        address rewardsInitiator;
        address slasher;
        address vetoCommittee;
        address taskManager;
        address taskValidator;
        address quorumManager;
        
        uint256 numStrategies;
        uint256 numQuorum;
        uint256 maxOperatorCount;
        uint96 minimumStake;
        uint256 initalPausedStatus;
        uint32 taskResponseWindow;

        string stakeType;
        uint256 lookAheadPeriod;

        address avsDirectory;
        address delegationManager;
    }

    function writeDeploymentToJson(TriggerXContracts memory triggerXContracts) internal {
        string memory jsonString = string.concat(
            "{\n",
            '  "proxyAdmin": "', vm.toString(address(triggerXContracts.proxyAdmin)), '",\n',
            '  "pauserRegistry": "', vm.toString(address(triggerXContracts.pauserRegistry)), '",\n',
            '  "indexRegistry": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.indexRegistry)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.indexRegistryImplementation)), '"\n',
            '  },\n',
            '  "stakeRegistry": {\n', 
            '    "proxy": "', vm.toString(address(triggerXContracts.stakeRegistry)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.stakeRegistryImplementation)), '"\n',
            '  },\n',
            '  "apkRegistry": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.apkRegistry)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.apkRegistryImplementation)), '"\n',
            '  },\n',
            '  "vetoableSlasher": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.vetoableSlasher)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.vetoableSlasherImplementation)), '"\n',
            '  },\n',
            // '  "socketRegistry": {\n',
            // '    "proxy": "', vm.toString(address(triggerXContracts.socketRegistry)), '",\n',
            // '    "implementation": "', vm.toString(address(triggerXContracts.socketRegistryImplementation)), '"\n',
            // '  },\n',
            '  "registryCoordinator": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.registryCoordinator)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.registryCoordinatorImplementation)), '"\n',
            '  },\n',
            '  "operatorStateRetriever": "', vm.toString(address(triggerXContracts.operatorStateRetriever)), '",\n',
            '  "triggerXServiceManager": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.triggerXServiceManager)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.triggerXServiceManagerImplementation)), '"\n',
            '  },\n',
            '  "triggerXTaskManager": {\n',
            '    "proxy": "', vm.toString(address(triggerXContracts.triggerXTaskManager)), '",\n',
            '    "implementation": "', vm.toString(address(triggerXContracts.triggerXTaskManagerImplementation)), '"\n',
            '  }\n',
            "}"
        );

        vm.writeFile("./script/output/deployment.holesky.json", jsonString);
    }

    function run() external {
        EigenLayerContracts memory eigenLayerContracts;
        DeploymentConfig memory deploymentConfig;

        {
            string memory config = vm.readFile("./script/input/parameters.holesky.json");

            deploymentConfig.numStrategies = abi.decode(vm.parseJson(config, ".numStrategies"), (uint256));
            deploymentConfig.numQuorum = abi.decode(vm.parseJson(config, ".numQuorum"), (uint256));
            deploymentConfig.maxOperatorCount = abi.decode(vm.parseJson(config, ".maxOperatorCount"), (uint256));
            deploymentConfig.minimumStake = abi.decode(vm.parseJson(config, ".minimumStake"), (uint96));
            deploymentConfig.initalPausedStatus = abi.decode(vm.parseJson(config, ".initalPausedStatus"), (uint256));
            deploymentConfig.taskResponseWindow = abi.decode(vm.parseJson(config, ".taskResponseWindow"), (uint32));

            deploymentConfig.stakeType = abi.decode(vm.parseJson(config, ".stakeType"), (string));
            deploymentConfig.lookAheadPeriod = abi.decode(vm.parseJson(config, ".lookAheadPeriod"), (uint256));

            address deployedStrategyManager = abi.decode(vm.parseJson(config, ".strategyManager"), (address));
            address deployedAvsDirectory = abi.decode(vm.parseJson(config, ".avsDirectory"), (address));
            address deployedDelegationManager = abi.decode(vm.parseJson(config, ".delegationManager"), (address));
            address deployedRewardsCoordinator = abi.decode(vm.parseJson(config, ".rewardsCoordinator"), (address));
            address deployedAllocationManager = abi.decode(vm.parseJson(config, ".allocationManager"), (address));
            address deployedPermissionController = abi.decode(vm.parseJson(config, ".permissionController"), (address));

            eigenLayerContracts.strategyManager = StrategyManager(deployedStrategyManager);
            eigenLayerContracts.avsDirectory = AVSDirectory(deployedAvsDirectory);
            eigenLayerContracts.delegationManager = DelegationManager(deployedDelegationManager);
            eigenLayerContracts.rewardsCoordinator = IRewardsCoordinator(deployedRewardsCoordinator);
            eigenLayerContracts.allocationManager = IAllocationManager(deployedAllocationManager);
            eigenLayerContracts.permissionController = IPermissionController(deployedPermissionController);

            eigenLayerContracts.wethStrategy = abi.decode(vm.parseJson(config, ".strategies.wethStrategy"), (address));
            eigenLayerContracts.wethMultiplier = abi.decode(vm.parseJson(config, ".strategies.wethMultiplier"), (uint96));
            eigenLayerContracts.stETHStrategy = abi.decode(vm.parseJson(config, ".strategies.stETHStrategy"), (address));
            eigenLayerContracts.stETHMultiplier = abi.decode(vm.parseJson(config, ".strategies.stETHMultiplier"), (uint96));
            eigenLayerContracts.rETHStrategy = abi.decode(vm.parseJson(config, ".strategies.rETHStrategy"), (address));
            eigenLayerContracts.rETHMultiplier = abi.decode(vm.parseJson(config, ".strategies.rETHMultiplier"), (uint96));
            eigenLayerContracts.cbETHStrategy = abi.decode(vm.parseJson(config, ".strategies.cbETHStrategy"), (address));
            eigenLayerContracts.cbETHMultiplier = abi.decode(vm.parseJson(config, ".strategies.cbETHMultiplier"), (uint96));
            eigenLayerContracts.reALTStrategy = abi.decode(vm.parseJson(config, ".strategies.reALTStrategy"), (address));
            eigenLayerContracts.reALTMultiplier = abi.decode(vm.parseJson(config, ".strategies.reALTMultiplier"), (uint96));
            eigenLayerContracts.eigenStrategy = abi.decode(vm.parseJson(config, ".strategies.eigenStrategy"), (address));
            eigenLayerContracts.eigenMultiplier = abi.decode(vm.parseJson(config, ".strategies.eigenMultiplier"), (uint96));
            eigenLayerContracts.beaconChainETHStrategy = abi.decode(vm.parseJson(config, ".strategies.beaconChainETHStrategy"), (address));
            eigenLayerContracts.beaconChainETHMultiplier = abi.decode(vm.parseJson(config, ".strategies.beaconChainETHMultiplier"), (uint96));

            deploymentConfig.triggerXOwner = abi.decode(vm.parseJson(config, ".triggerXOwner"), (address));
            deploymentConfig.rewardsInitiator = abi.decode(vm.parseJson(config, ".rewardsInitiator"), (address));
            deploymentConfig.slasher = abi.decode(vm.parseJson(config, ".slasher"), (address));
            deploymentConfig.vetoCommittee = abi.decode(vm.parseJson(config, ".vetoCommittee"), (address));
            deploymentConfig.taskManager = abi.decode(vm.parseJson(config, ".taskManager"), (address));
            deploymentConfig.taskValidator = abi.decode(vm.parseJson(config, ".taskValidator"), (address));
            deploymentConfig.quorumManager = abi.decode(vm.parseJson(config, ".quorumManager"), (address));
        }

        deploymentConfig.avsDirectory = address(eigenLayerContracts.avsDirectory);
        deploymentConfig.delegationManager = address(eigenLayerContracts.delegationManager);

        TokenAndWeight[] memory deployedStrategyArray = new TokenAndWeight[](deploymentConfig.numStrategies);
        {
            deployedStrategyArray[0] = TokenAndWeight({
                token: eigenLayerContracts.wethStrategy,
                weight: eigenLayerContracts.wethMultiplier
            });
            deployedStrategyArray[1] = TokenAndWeight({
                token: eigenLayerContracts.stETHStrategy,
                weight: eigenLayerContracts.stETHMultiplier
            });
            deployedStrategyArray[2] = TokenAndWeight({
                token: eigenLayerContracts.rETHStrategy,
                weight: eigenLayerContracts.rETHMultiplier
            });
            deployedStrategyArray[3] = TokenAndWeight({
                token: eigenLayerContracts.cbETHStrategy,
                weight: eigenLayerContracts.cbETHMultiplier
            });
            deployedStrategyArray[4] = TokenAndWeight({
                token: eigenLayerContracts.reALTStrategy,
                weight: eigenLayerContracts.reALTMultiplier
            });
            deployedStrategyArray[5] = TokenAndWeight({
                token: eigenLayerContracts.eigenStrategy,
                weight: eigenLayerContracts.eigenMultiplier
            });
            deployedStrategyArray[6] = TokenAndWeight({
                token: eigenLayerContracts.beaconChainETHStrategy,
                weight: eigenLayerContracts.beaconChainETHMultiplier
            });
        }

        vm.startBroadcast();
        
        TriggerXContracts memory triggerXContracts;
        triggerXContracts.proxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();

        {
            address[] memory pausers = new address[](1);
            pausers[0] = deploymentConfig.triggerXOwner;
            triggerXContracts.pauserRegistry = new PauserRegistry(pausers, deploymentConfig.triggerXOwner);
        }

        triggerXContracts.indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.apkRegistry = BLSApkRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.registryCoordinator = RegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.triggerXServiceManager = ITriggerXServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.triggerXTaskManager = ITriggerXTaskManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );
        triggerXContracts.vetoableSlasher = VetoableSlasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(triggerXContracts.proxyAdmin), ""))
        );

        triggerXContracts.indexRegistryImplementation = new IndexRegistry(triggerXContracts.registryCoordinator);
        triggerXContracts.proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(triggerXContracts.indexRegistry))),
            address(triggerXContracts.indexRegistryImplementation)
        );

        triggerXContracts.stakeRegistryImplementation = new StakeRegistry(
            triggerXContracts.registryCoordinator, 
            IDelegationManager(eigenLayerContracts.delegationManager),
            IAVSDirectory(eigenLayerContracts.avsDirectory),
            IAllocationManager(eigenLayerContracts.allocationManager),
            ITriggerXServiceManager(address(triggerXContracts.triggerXServiceManager))
        );
        triggerXContracts.proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(triggerXContracts.stakeRegistry))),
            address(triggerXContracts.stakeRegistryImplementation)
        );

        triggerXContracts.apkRegistryImplementation = new BLSApkRegistry(triggerXContracts.registryCoordinator);
        triggerXContracts.proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(triggerXContracts.apkRegistry))),
            address(triggerXContracts.apkRegistryImplementation)
        );

        triggerXContracts.vetoableSlasherImplementation = new VetoableSlasher(
            IAllocationManager(eigenLayerContracts.allocationManager),
            IServiceManager(address(triggerXContracts.triggerXServiceManager)),
            deploymentConfig.slasher
        );

        bytes memory vetoableSlasherInitCode = abi.encodeWithSelector(
            VetoableSlasher.initialize.selector,
            deploymentConfig.vetoCommittee,
            deploymentConfig.slasher
        );

        triggerXContracts.proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(triggerXContracts.vetoableSlasher))),
            address(triggerXContracts.vetoableSlasherImplementation),
            vetoableSlasherInitCode
        );

        triggerXContracts.registryCoordinatorImplementation = new RegistryCoordinator(
            ITriggerXServiceManager(address(triggerXContracts.triggerXServiceManager)),
            triggerXContracts.stakeRegistry,
            triggerXContracts.apkRegistry,
            triggerXContracts.indexRegistry,
            IAllocationManager(eigenLayerContracts.allocationManager),
            triggerXContracts.pauserRegistry
        );
        triggerXContracts.operatorStateRetriever = new OperatorStateRetriever();

        {
            IRegistryCoordinator.OperatorSetParam[] memory operatorSetParams =
                new IRegistryCoordinator.OperatorSetParam[](deploymentConfig.numQuorum);

            for (uint256 i = 0; i < deploymentConfig.numQuorum; i++) {
                operatorSetParams[i] = IRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: uint32(deploymentConfig.maxOperatorCount),
                    kickBIPsOfOperatorStake: 11000,
                    kickBIPsOfTotalStake: 1001
                });
            }

            uint96[] memory minimumStakeForQuourm = new uint96[](deploymentConfig.numQuorum);
            for (uint256 i = 0; i < deploymentConfig.numQuorum; i++) {
                minimumStakeForQuourm[i] = deploymentConfig.minimumStake;
            }
            
            IStakeRegistry.StrategyParams[][] memory strategyParams =
                new IStakeRegistry.StrategyParams[][](deploymentConfig.numQuorum);
            for (uint256 i = 0; i < deploymentConfig.numQuorum; i++) {
                IStakeRegistry.StrategyParams[] memory params = new IStakeRegistry.StrategyParams[](1);
                params[0] = IStakeRegistry.StrategyParams({
                    strategy: IStrategy(deployedStrategyArray[i].token),
                    multiplier: deployedStrategyArray[i].weight
                });
                strategyParams[i] = params;
            }

            StakeType[] memory stakeTypes = new StakeType[](deploymentConfig.numQuorum);
            for (uint256 i = 0; i < deploymentConfig.numQuorum; i++) {
                if (keccak256(bytes(deploymentConfig.stakeType)) == keccak256(bytes("TOTAL_DELEGATED"))) {
                    stakeTypes[i] = StakeType.TOTAL_DELEGATED;
                } else if (keccak256(bytes(deploymentConfig.stakeType)) == keccak256(bytes("TOTAL_SLASHABLE"))) {
                    stakeTypes[i] = StakeType.TOTAL_SLASHABLE;
                }
            }

            uint256[] memory lookAheadPeriods = new uint256[](deploymentConfig.numQuorum);
            for (uint256 i = 0; i < deploymentConfig.numQuorum; i++) {
                lookAheadPeriods[i] = deploymentConfig.lookAheadPeriod;
            }

            triggerXContracts.proxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(triggerXContracts.registryCoordinator))),
                address(triggerXContracts.registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    RegistryCoordinator.initialize.selector,
                    deploymentConfig.triggerXOwner,
                    deploymentConfig.quorumManager,
                    deploymentConfig.quorumManager,
                    deploymentConfig.initalPausedStatus,
                    operatorSetParams,
                    minimumStakeForQuourm,
                    strategyParams,
                    stakeTypes,
                    lookAheadPeriods
                )
            );

            triggerXContracts.triggerXServiceManagerImplementation = new TriggerXServiceManager(
                IAVSDirectory(eigenLayerContracts.avsDirectory),
                eigenLayerContracts.rewardsCoordinator,
                triggerXContracts.registryCoordinator,
                triggerXContracts.stakeRegistry,
                eigenLayerContracts.permissionController,
                triggerXContracts.pauserRegistry,
                triggerXContracts.vetoableSlasher
            );

            bytes memory initcode;
            {
                initcode = abi.encodeWithSelector(
                    TriggerXServiceManager.initialize.selector,
                    ITriggerXTaskManager(deploymentConfig.taskManager),
                    deploymentConfig.triggerXOwner,
                    deploymentConfig.rewardsInitiator,
                    deploymentConfig.slasher,
                    deploymentConfig.vetoCommittee,
                    deploymentConfig.taskManager,
                    deploymentConfig.taskValidator,
                    deploymentConfig.quorumManager
                );
            }

            triggerXContracts.proxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(triggerXContracts.triggerXServiceManager))),
                address(triggerXContracts.triggerXServiceManagerImplementation),
                initcode
            );
        }

        triggerXContracts.triggerXTaskManagerImplementation = new TriggerXTaskManager(
            triggerXContracts.registryCoordinator,
            triggerXContracts.pauserRegistry
        );

        bytes memory taskManagerInitCode = abi.encodeWithSelector(
            TriggerXTaskManager.initialize.selector,
            deploymentConfig.triggerXOwner,
            deploymentConfig.taskResponseWindow,
            ITriggerXServiceManager(address(triggerXContracts.triggerXServiceManager))
        );

        triggerXContracts.proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(triggerXContracts.triggerXTaskManager))),
            address(triggerXContracts.triggerXTaskManagerImplementation),
            taskManagerInitCode
        );

        ITriggerXServiceManager(address(triggerXContracts.triggerXServiceManager)).updateTaskManager(
            address(triggerXContracts.triggerXTaskManager)
        );

        writeDeploymentToJson(triggerXContracts);

        vm.stopBroadcast();
    }
}