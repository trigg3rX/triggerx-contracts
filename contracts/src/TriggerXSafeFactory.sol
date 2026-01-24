// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {SafeProxyFactory} from "../lib/safe-contracts/contracts/proxies/SafeProxyFactory.sol";
import {SafeProxy} from "../lib/safe-contracts/contracts/proxies/SafeProxy.sol";

/**
 * @title TriggerXSafeFactory
 * @notice Factory contract for creating Safe wallets for TriggerX users.
 * @dev This factory deploys Safe wallets using the SafeProxyFactory, following the same
 *      pattern as the Safe Core SDK. It supports both Safe and SafeL2 singletons.
 */
contract TriggerXSafeFactory {
    // Maps a user to their Safe wallets.
    mapping(address => address[]) private userSafeWallets;
    // Maps a user to their current saltNonce for deterministic address generation.
    mapping(address => uint256) private userSaltNonces;
    
    address public immutable safeProxyFactory;
    address public immutable safeSingleton;

    event SafeWalletCreated(
        address indexed user, 
        address indexed safeWallet, 
        uint256 saltNonce
    );

    constructor(
        address _safeProxyFactory,
        address _safeSingleton
    ) {
        require(_safeProxyFactory != address(0), "Zero proxy factory");
        require(_safeSingleton != address(0), "Zero singleton");
        safeProxyFactory = _safeProxyFactory;
        safeSingleton = _safeSingleton;
    }

    /**
     * @notice Deploy a new Safe wallet for a user with the user as the sole owner.
     * @dev This function follows the same deployment pattern as the Safe Core SDK:
     *      1. Encodes the setup call with user as sole owner and threshold of 1
     *      2. Uses SafeProxyFactory.createProxyWithNonce with an incrementing saltNonce
     *      3. The saltNonce ensures deterministic addresses and prevents collisions
     * @param user The user who will own this Safe wallet.
     * @return safeAddress The address of the deployed Safe wallet.
     */
    function createSafeWallet(address user) external returns (address safeAddress) {
        require(user != address(0), "Zero user");
        
        // Get the next saltNonce for this user
        uint256 saltNonce = userSaltNonces[user];
        userSaltNonces[user] = saltNonce + 1;
        
        // Create owners array with the user as the sole owner
        address[] memory owners = new address[](1);
        owners[0] = user;
        
        // Create initializer data for Safe setup
        // This matches the Safe Core SDK's setup call signature
        bytes memory initializer = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            owners,           // _owners - list of Safe owners
            1,                // _threshold - number of required confirmations (1 for single owner)
            address(0),       // to - contract address for optional delegate call during setup
            "",               // data - data payload for optional delegate call
            address(0),       // fallbackHandler - handler for fallback calls to this contract
            address(0),       // paymentToken - token that should be used for payment (0 is ETH)
            0,                // payment - value that should be paid
            address(0)        // paymentReceiver - address that should receive the payment
        );
        
        // Create Safe proxy using the SafeProxyFactory
        // This uses CREATE2 for deterministic address generation based on:
        // - safeSingleton address
        // - initializer data
        // - saltNonce
        SafeProxyFactory factory = SafeProxyFactory(safeProxyFactory);
        SafeProxy safeProxy = factory.createProxyWithNonce(safeSingleton, initializer, saltNonce);
        
        safeAddress = address(safeProxy);
        require(safeAddress != address(0), "Zero proxy");
        userSafeWallets[user].push(safeAddress);
        emit SafeWalletCreated(user, safeAddress, saltNonce);
        return safeAddress;
    }

    /**
     * @notice Returns all Safe wallets created for a particular user.
     * @param user The owner's address.
     * @return Array of Safe wallet addresses.
     */
    function getSafeWallets(address user) external view returns (address[] memory) {
        return userSafeWallets[user];
    }

    /**
     * @notice Returns the most recently created Safe wallet for a user, or address(0) if none.
     * @param user The owner's address.
     * @return The address of the most recent Safe wallet, or address(0) if none exists.
     */
    function latestSafeWallet(address user) external view returns (address) {
        address[] storage wallets = userSafeWallets[user];
        if (wallets.length == 0) return address(0);
        return wallets[wallets.length - 1];
    }

    /**
     * @notice Returns the current saltNonce for a user.
     * @dev This can be used to predict the address of the next Safe wallet for a user.
     * @param user The owner's address.
     * @return The current saltNonce value.
     */
    function getUserSaltNonce(address user) external view returns (uint256) {
        return userSaltNonces[user];
    }

    /**
     * @notice Calculates the predicted address for a user's next Safe wallet.
     * @dev This matches the Safe Core SDK's address prediction logic using CREATE2.
     *      The address is deterministic based on:
     *      - SafeProxyFactory address
     *      - Safe singleton address
     *      - Initializer data (which includes the user's address)
     *      - saltNonce (user's current saltNonce)
     * @param user The user who will own the Safe wallet.
     * @return predictedAddress The predicted address of the next Safe wallet.
     */
    function predictSafeAddress(address user) external view returns (address predictedAddress) {
        require(user != address(0), "Zero user");
        
        uint256 saltNonce = userSaltNonces[user];
        
        // Create owners array with the user as the sole owner
        address[] memory owners = new address[](1);
        owners[0] = user;
        
        // Create the same initializer that will be used during deployment
        bytes memory initializer = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            owners,
            1,
            address(0),
            "",
            address(0),
            address(0),
            0,
            address(0)
        );
        
        // Calculate salt following SafeProxyFactory's logic:
        // salt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce))
        bytes32 salt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce));
        
        // Calculate CREATE2 address
        // Address = keccak256(0xff ++ factoryAddress ++ salt ++ keccak256(deploymentCode))[12:]
        bytes memory deploymentCode = abi.encodePacked(
            type(SafeProxy).creationCode,
            uint256(uint160(safeSingleton))
        );
        
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                safeProxyFactory,
                salt,
                keccak256(deploymentCode)
            )
        );
        
        predictedAddress = address(uint160(uint256(hash)));
    }
}