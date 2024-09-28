// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {TriggerXDeploymentLib} from "./utils/TriggerXDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";

import {
    Quorum,
    StrategyParams,
    IStrategy
} from "@eigenlayer-middleware/interfaces/IECDSAStakeRegistryEventsAndErrors.sol";

contract TriggerXDeployer is Script {
    using UpgradeableProxyLib for address;

    address private deployer;
    address proxyAdmin;
    TriggerXDeploymentLib.DeploymentData triggerXDeployment;
    Quorum internal quorum;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("AVS_OWNER_PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
        quorum.strategies.push(
            StrategyParams({strategy: IStrategy(address(420)), multiplier: 10_000})
        );
    }

    function run() external {
        vm.startBroadcast(deployer);
        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();

        triggerXDeployment =
            TriggerXDeploymentLib.deployContracts(proxyAdmin, quorum);

        vm.stopBroadcast();

        verifyDeployment();
        TriggerXDeploymentLib.writeDeploymentJson(triggerXDeployment);
    }

    function verifyDeployment() internal view {
        require(
            triggerXDeployment.stakeRegistry != address(0), "StakeRegistry address cannot be zero"
        );
        require(
            triggerXDeployment.triggerXServiceManager != address(0),
            "TriggerXServiceManager address cannot be zero"
        );
        require(proxyAdmin != address(0), "ProxyAdmin address cannot be zero");
        require(
            triggerXDeployment.triggerXTaskManager != address(0),
            "TriggerXTaskManager address cannot be zero"
        );
    }
}