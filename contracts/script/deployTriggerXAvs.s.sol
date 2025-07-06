// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "forge-std/Script.sol";
import "forge-std/console2.sol";

import "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {TriggerXAvs} from "src/TriggerXAvs.sol";

/// @title DeployTriggerXAvs
/// @notice Foundry script that deploys the TriggerXAvs implementation contract and
///         an ERC-1967 proxy pointing at it. The proxy is initialised via
///         `TriggerXAvs.initialize(address initialOwner)`.
///
/// Environment variables expected:
/// - PRIVATE_KEY : the deployer's private key used for broadcasting transactions
/// - OWNER       : the address that should become the owner of the newly deployed AVS
contract DeployTriggerXAvs is Script {
    function run() external returns (TriggerXAvs proxy) {
        // --- Load required env vars ---
        uint256 deployerPk = vm.envUint("DEV0_PRIVATE_KEY");
        address owner = 0x9c0E7ECE2749091b47620b79fb43bf81923D48C7;

        // --- Start broadcasting ---
        vm.startBroadcast(deployerPk);

        // 1. Deploy implementation
        TriggerXAvs implementation = new TriggerXAvs();
        console2.log("TriggerXAvs implementation deployed at", address(implementation));

        // 2. Prepare initializer calldata
        bytes memory initData = abi.encodeWithSelector(TriggerXAvs.initialize.selector, owner);

        // 3. Deploy ERC-1967 proxy pointing to the implementation and executing initializer
        ERC1967Proxy erc1967Proxy = new ERC1967Proxy(address(implementation), initData);
        proxy = TriggerXAvs(address(erc1967Proxy));

        console2.log("TriggerXAvs proxy deployed at", address(proxy));
        console2.log("Proxy owner is", proxy.owner());

        // I want to confirm the implementation address is 0x0bA3E5336C139Bb5e94Bdc887E42557B6F6b8fFD
        // Read the ERC-1967 implementation slot directly from storage.
        bytes32 implSlot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        bytes32 raw = vm.load(address(proxy), implSlot);
        address fetchimplementation = address(uint160(uint256(raw)));
        console2.log("Implementation:", fetchimplementation);

        // --- Stop broadcasting ---
        vm.stopBroadcast();
    }
} 