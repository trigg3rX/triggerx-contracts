// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {ProxyAdmin} from "@openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {ECDSAStakeRegistry} from "@eigenlayer-middleware/unaudited/ECDSAStakeRegistry.sol";
import {TriggerXServiceManager} from "../../src/TriggerXServiceManager.sol";
import {TriggerXTaskManager} from "../../src/TriggerXTaskManager.sol";
import {IDelegationManager} from "@eigenlayer-contracts/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer-contracts/contracts/interfaces/IAVSDirectory.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {ITriggerXTaskManager} from "../../src/ITriggerXTaskManager.sol";
import {Quorum} from "@eigenlayer-middleware/interfaces/IECDSAStakeRegistryEventsAndErrors.sol";
import {UpgradeableProxyLib} from "./UpgradeableProxyLib.sol";
import {Strings} from "@openzeppelin-contracts/contracts/utils/Strings.sol";

library TriggerXDeploymentLib {
    using stdJson for *;
    using Strings for *;
    using UpgradeableProxyLib for address;

    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct DeploymentData {
        address triggerXServiceManager;
        address triggerXTaskManager;
        address stakeRegistry;
    }

    function deployContracts(
        address proxyAdmin,
        Quorum memory quorum
    ) internal returns (DeploymentData memory) {
        DeploymentData memory result;

        // First, deploy upgradeable proxy contracts that will point to the implementations.
        result.triggerXServiceManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.triggerXTaskManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.stakeRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);

        // Read core contract addresses from JSON file
        string memory eigenJson = vm.readFile("script/utils/deployments/eigenDeployments.json");
        address delegationManager = eigenJson.readAddress(".delegationManager");
        address avsDirectory = eigenJson.readAddress(".avsDirectory");
        address registryCoordinator = eigenJson.readAddress(".registryCoordinator");
        
        // Deploy the implementation contracts, using the proxy contracts as inputs
        address stakeRegistryImpl =
            address(new ECDSAStakeRegistry(IDelegationManager(delegationManager)));

        address triggerXServiceManagerImpl = address(
            new TriggerXServiceManager(
                IAVSDirectory(avsDirectory),
                IRegistryCoordinator(registryCoordinator),
                IStakeRegistry(result.stakeRegistry),
                ITriggerXTaskManager(result.triggerXTaskManager)
            )
        );

        address triggerXTaskManagerImpl = address(
            new TriggerXTaskManager(
                IRegistryCoordinator(registryCoordinator),
                uint32(uint160(address(result.triggerXTaskManager)))
            )
        );

        bytes memory upgradeCall = abi.encodeCall(
            ECDSAStakeRegistry.initialize, (result.triggerXServiceManager, 1, quorum)
        );

        console2.log("Reaching here.....");

        UpgradeableProxyLib.upgradeAndCall(result.stakeRegistry, stakeRegistryImpl, upgradeCall);
        console2.log("StakeRegistry proxy upgraded at:", result.stakeRegistry);

        UpgradeableProxyLib.upgrade(result.triggerXServiceManager, triggerXServiceManagerImpl);
        console2.log("TriggerXServiceManager proxy upgraded at:", result.triggerXServiceManager);

        UpgradeableProxyLib.upgrade(result.triggerXTaskManager, triggerXTaskManagerImpl);
        console2.log("TriggerXTaskManager proxy upgraded at:", result.triggerXTaskManager);

        string memory holeskyJson = vm.readFile("script/utils/deployments/holeskyDeployments.json");
        holeskyJson = holeskyJson.serialize("triggerXServiceManager", result.triggerXServiceManager);
        holeskyJson = holeskyJson.serialize("triggerXTaskManager", result.triggerXTaskManager);
        holeskyJson = holeskyJson.serialize("stakeRegistry", result.stakeRegistry);
        vm.writeFile("script/utils/holeskyDeployments.json", holeskyJson);
        console2.log("Deployment addresses written to holeskyDeployments.json");
        console2.log("---------------------------------------------------------------------------------\n");

        return result;
    }
}
