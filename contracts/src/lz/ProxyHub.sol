// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@layerzero-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import "@layerzero-v2/contracts/interfaces/ILayerZeroReceiver.sol";
import "../interfaces/IAttestationCenter.sol";
import "forge-std/console.sol";

/**
 * @title ProxyHub
 * @dev Validates keeper registration and authorizes execution requests from spoke chains
 */
contract ProxyHub is ILayerZeroReceiver {
    // LayerZero V2 endpoint contract
    address public immutable endpoint;

    // Attestation center contract for keeper validation
    IAttestationCenter public immutable attestationCenter;

    // Keeper registry
    mapping(address => bool) public isKeeper;

    // Events
    event CrossChainRequestReceived(address indexed keeper, address target, uint32 sourceChainId);
    event CrossChainExecutionSent(address indexed toSpoke, address target, bytes data, uint32 dstChainId);
    event FunctionExecuted(address indexed keeper, address target, bytes data, uint256 value);

    /**
     * @param _endpoint LayerZero V2 endpoint on Base
     * @param _attestationCenter Address of the IAttestationCenter contract
     */
    constructor(address _endpoint, address _attestationCenter) {
        endpoint = _endpoint;
        attestationCenter = IAttestationCenter(_attestationCenter);
    }

    /**
     * @dev Handles incoming execution requests from a spoke chain.
     * Verifies keeper status, then returns the execution command back to the spoke.
     */
    function lzReceive(
        Origin calldata _origin,
        bytes32,            // _guid (unused)
        bytes calldata _message,
        address,            // _executor (unused)
        bytes calldata      // _extraData (unused)
    ) external payable override {
        require(msg.sender == endpoint, "Hub: Unauthorized endpoint");

        (address keeper, address target, bytes memory data) = abi.decode(_message, (address, address, bytes));

        require(_checkKeeper(keeper), "Hub: Keeper not registered");

        emit CrossChainRequestReceived(keeper, target, _origin.srcEid);

        bytes memory message = abi.encode(target, data);
        bytes memory options = abi.encodePacked(uint16(1), uint256(1000000)); // Lower gas limit

        // Get fee estimate
        MessagingFee memory fee = ILayerZeroEndpointV2(endpoint).quote(
           MessagingParams({
                dstEid: _origin.srcEid,
                receiver: _origin.sender,
                message: message,
                options: options,
                payInLzToken: false
            }),
            address(this)
        );

        console.log("Native Fee:", fee.nativeFee);
        console.log("LZ Token Fee:", fee.lzTokenFee);
        console.log(address(this).balance);

        // Add a buffer
        uint256 feeWithBuffer = fee.nativeFee * 120 / 100; // 20% buffer

        // Check if we have enough
        require(msg.value >= feeWithBuffer, "Insufficient fee provided");


        // Send with exact fee
        ILayerZeroEndpointV2(endpoint).send{value: feeWithBuffer}(
            MessagingParams({
                dstEid: _origin.srcEid,
                receiver: _origin.sender,
                message: message,
                options: options,
                payInLzToken: false
            }),
            address(this)
        );
        
        emit CrossChainExecutionSent(address(uint160(uint256(_origin.sender))), target, data, _origin.srcEid);
    }

    /**
     * @notice Executes a function locally if the caller is a registered keeper
     * @param target Target contract address
     * @param data Calldata for the function call
     */
    function executeFunction(address target, bytes memory data) external payable {
        require(isKeeper[msg.sender], "Hub: Keeper not registered");
        _executeFunction(target, data);
    }

    /**
     * @dev Internal function execution logic
     */
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Hub: Function execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    /**
     * @dev Checks if an address is a registered operator in the AttestationCenter.
     * @param keeper The address to check.
     * @return bool True if the call to operatorsIdsByAddress succeeds, false if it reverts.
     */
    function _checkKeeper(address keeper) internal view returns (bool) {
        // Use try/catch to handle potential reverts from the external call
        try attestationCenter.operatorsIdsByAddress(keeper) returns (uint256 /* operatorId */) {
            // If the call succeeds (returns uint256), the keeper is considered valid.
            return true;
        } catch {
            // If the call reverts (e.g., keeper not found), it returns false.
            return false;
        }
    }

    // === Required LayerZero interface stubs ===

    function allowInitializePath(Origin calldata) external pure override returns (bool) {
        return true;
    }

    function nextNonce(uint32, bytes32) external pure override returns (uint64) {
        return 0;
    }
    
    receive() external payable {}
    
}