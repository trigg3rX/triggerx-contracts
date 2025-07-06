// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "forge-std/Script.sol";
import "forge-std/console2.sol";

import {TriggerXAvs} from "src/TriggerXAvs.sol";

/// @title UpgradeTriggerXAvs
/// @notice Deploys a new `TriggerXAvs` implementation and upgrades an existing
///         ERC-1967/UUPS proxy to point at it.
///
/// Environment variables expected:
/// - PRIVATE_KEY : private key of the owner (must match `TriggerXAvs.owner()`)
/// - PROXY       : address of the existing proxy to upgrade
contract UpgradeTriggerXAvs is Script {
    /// @dev ERC-1967 implementation slot constant for quick verification
    bytes32 internal constant _IMPLEMENTATION_SLOT =
        0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;

    function run() external returns (address newImplementation) {
        uint256 pk = vm.envUint("DEV0_PRIVATE_KEY");
        address proxyAddr = 0x4729BC58ADC71E5386b995f99402176757d75940;

        vm.startBroadcast(pk);

        // 1. Deploy the new implementation.
        TriggerXAvs impl = new TriggerXAvs();
        newImplementation = address(impl);
        console2.log("New TriggerXAvs implementation deployed at", newImplementation);

        // 2. Upgrade the proxy (through delegatecall) via upgradeToAndCall with empty data.
        TriggerXAvs proxy = TriggerXAvs(payable(proxyAddr));
        proxy.upgradeToAndCall(newImplementation, bytes(""));
        console2.log("Proxy upgraded to", newImplementation);

        // 3. Verify the implementation slot now points to the new implementation.
        bytes32 data = vm.load(proxyAddr, _IMPLEMENTATION_SLOT);
        address recordedImpl = address(uint160(uint256(data)));
        console2.log("Implementation recorded in proxy:", recordedImpl);
        require(recordedImpl == newImplementation, "Upgrade failed: slot mismatch");

        vm.stopBroadcast();
    }
} 