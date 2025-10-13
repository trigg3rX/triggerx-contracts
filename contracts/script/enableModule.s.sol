// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";

interface ISafeWalletFactory {
    function createSafeWallet(address user) external returns (address safeAddress);
    function latestSafeWallet(address user) external view returns (address);
    function getSafeWallets(address user) external view returns (address[] memory);
}

interface ISafe {
    function nonce() external view returns (uint256);
    function domainSeparator() external view returns (bytes32);
    function getChainId() external view returns (uint256);
    
    function execTransaction(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address payable refundReceiver,
        bytes memory signatures
    ) external payable returns (bool success);
    
    function enableModule(address module) external;
    function isModuleEnabled(address module) external view returns (bool);
}

/**
 * @title EnableSafeModuleScript
 * @notice Script to create a Safe wallet and enable TriggerXSafeModule
 * @dev This demonstrates the proper way to enable a module on a Safe wallet
 */
contract EnableSafeModuleScript is Script {
    // ⚠️ IMPORTANT: Update these addresses with your actual deployed contracts!
    // The SAFE_WALLET_FACTORY address below is currently WRONG - it points to a Safe wallet, not the factory!
    // 
    // To deploy the factory first, run:
    //   forge script script/deploy/6_deploySafeFactory.s.sol:DeploySafeFactoryL2 --rpc-url $ARB_SEPOLIA_RPC --broadcast
    // Then update the address below with the deployed factory address
    
    address SAFE_WALLET_FACTORY = vm.envOr("SAFE_WALLET_FACTORY", 0x383D4a61D0B069D02cA2Db5A82003b9561d56e19);
    address constant TRIGGERX_SAFE_MODULE = 0xca3a0c43Ac9E4FcB76C774F330fD31D4098EEacD;

    // EIP-712 TypeHash for Safe transactions
    bytes32 constant SAFE_TX_TYPEHASH = 0xbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d8;

    function run() public {
        uint256 userPrivateKey = vm.envUint("PRIVATE_KEY");
        address user = vm.addr(userPrivateKey);
        
        console.log("=== Safe Module Enablement Script ===");
        console.log("User address:", user);
        console.log("Factory address:", SAFE_WALLET_FACTORY);
        console.log("Module address:", TRIGGERX_SAFE_MODULE);
        console.log("");

        vm.startBroadcast(userPrivateKey);

        // Check if user already has a Safe wallet
        ISafeWalletFactory factory = ISafeWalletFactory(SAFE_WALLET_FACTORY);
        
        // First check if user has existing Safes
        address safe;
        try factory.latestSafeWallet(user) returns (address existingSafe) {
            safe = existingSafe;
            if (safe != address(0)) {
                console.log("Found existing Safe at:", safe);
            }
        } catch {
            console.log("Warning: Could not query existing Safes from factory");
        }
        
        if (safe == address(0)) {
            console.log("Creating new Safe wallet...");
            safe = factory.createSafeWallet(user);
            console.log("Safe deployed at:", safe);
        }

        ISafe safeProxy = ISafe(safe);
        
        // Get Safe's current nonce
        uint256 safeNonce = safeProxy.nonce();
        console.log("Safe nonce:", safeNonce);
        
        // Encode the enableModule call data
        bytes memory data = abi.encodeWithSignature("enableModule(address)", TRIGGERX_SAFE_MODULE);
        
        // Build Safe transaction parameters
        address to = safe;              // self-call to enable module
        uint256 value = 0;              // no ETH transfer
        uint8 operation = 0;            // Call operation
        uint256 safeTxGas = 0;          // use all available gas
        uint256 baseGas = 0;            // no base gas
        uint256 gasPrice = 0;           // no gas refund
        address gasToken = address(0);  // ETH
        address refundReceiver = address(0); // no refund
        
        // Calculate Safe transaction hash using EIP-712
        bytes32 safeTxHash = keccak256(
            abi.encode(
                SAFE_TX_TYPEHASH,
                to,
                value,
                keccak256(data),
                operation,
                safeTxGas,
                baseGas,
                gasPrice,
                gasToken,
                refundReceiver,
                safeNonce
            )
        );
        
        // Get domain separator from Safe
        bytes32 domainSeparator = safeProxy.domainSeparator();
        
        // Create EIP-712 hash
        bytes32 txHash = keccak256(
            abi.encodePacked(
                bytes1(0x19),
                bytes1(0x01),
                domainSeparator,
                safeTxHash
            )
        );
        
        console.log("Transaction hash:");
        console.logBytes32(txHash);
        
        // Sign the transaction hash
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userPrivateKey, txHash);
        
        // Encode the signature (65 bytes: r, s, v)
        bytes memory signature = abi.encodePacked(r, s, v);
        
        console.log("Signature generated, length:", signature.length);
        
        // Execute the transaction through Safe's execTransaction
        bool success = safeProxy.execTransaction(
            to,
            value,
            data,
            operation,
            safeTxGas,
            baseGas,
            gasPrice,
            gasToken,
            payable(refundReceiver),
            signature
        );
        
        require(success, "Failed to enable module");
        console.log("Module enabled successfully!");
        
        // Verify module is enabled
        bool isEnabled = safeProxy.isModuleEnabled(TRIGGERX_SAFE_MODULE);
        console.log("Module enabled:", isEnabled);
        require(isEnabled, "Module verification failed");

        vm.stopBroadcast();
    }
}
