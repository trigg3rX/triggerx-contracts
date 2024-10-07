// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// import {ProxyAdmin} from "@openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol";
// import {TransparentUpgradeableProxy} from "@openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
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
import {Strings} from "@openzeppelin-contracts/contracts/utils/Strings.sol";

library TriggerXDeploymentLib {
    using stdJson for *;
    using Strings for *;

    Vm internal constant VM = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct DeploymentData {
        address triggerXServiceManager;
        address triggerXTaskManager;
        address stakeRegistry;
    }

    function deployContracts(
        // address proxyAdmin,
        Quorum memory quorum
    ) internal returns (DeploymentData memory) {
        DeploymentData memory result;

        // result.triggerXServiceManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        // result.triggerXTaskManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        // result.stakeRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        
        // Read core contract addresses from JSON file
        string memory eigenJson = VM.readFile("script/utils/deployments/eigenDeployments.json");
        address delegationManager = eigenJson.readAddress(".delegationManager");
        address avsDirectory = eigenJson.readAddress(".avsDirectory");
        address registryCoordinator = eigenJson.readAddress(".registryCoordinator");
        
        // Deploy contracts directly
        result.stakeRegistry = address(
            new ECDSAStakeRegistry(IDelegationManager(delegationManager))
        );

        result.triggerXServiceManager = address(
            new TriggerXServiceManager(
                IAVSDirectory(avsDirectory),
                IRegistryCoordinator(registryCoordinator),
                IStakeRegistry(result.stakeRegistry),
                ITriggerXTaskManager(result.triggerXTaskManager)
            )
        );

        // We need to deploy TaskManager before ServiceManager due to dependency
        result.triggerXTaskManager = address(
            new TriggerXTaskManager(
                IRegistryCoordinator(registryCoordinator),
                uint32(uint160(address(result.triggerXTaskManager))),
                TriggerXServiceManager(result.triggerXServiceManager)
            )
        );



        // Initialize the stake registry
        ECDSAStakeRegistry(result.stakeRegistry).initialize(
            result.triggerXServiceManager,
            1,
            quorum
        );

        // Log deployments
        console2.log("Direct contract deployments:");
        console2.log("StakeRegistry deployed at:", result.stakeRegistry);
        console2.log("TriggerXTaskManager deployed at:", result.triggerXTaskManager);
        console2.log("TriggerXServiceManager deployed at:", result.triggerXServiceManager);

        // Write deployment addresses to JSON
        string memory holeskyJson = VM.readFile("script/utils/deployments/holeskyDeployments.json");
        holeskyJson = holeskyJson.serialize("triggerXServiceManager", result.triggerXServiceManager);
        holeskyJson = holeskyJson.serialize("triggerXTaskManager", result.triggerXTaskManager);
        holeskyJson = holeskyJson.serialize("stakeRegistry", result.stakeRegistry);
        VM.writeFile("script/utils/deployments/holeskyDeployments.json", holeskyJson);
        console2.log("Deployment addresses written to holeskyDeployments.json");
        console2.log("---------------------------------------------------------------------------------\n");

        return result;
    }
}