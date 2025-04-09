// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@layerzero-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import "@layerzero-v2/contracts/interfaces/ILayerZeroReceiver.sol";
import "../interfaces/IAttestationCenter.sol";

/**
 * @title ProxySpoke
 * @dev Executes cross-chain function calls validated by a hub (Base Sepolia) via LayerZero V2
 */
contract ProxySpoke is ILayerZeroReceiver {
    // LayerZero endpoint contract
    address public immutable endpoint;

    // Hub (Base Sepolia) endpoint ID
    uint32 public immutable hubEid;

    // Events
    event CrossChainRequestSent(address indexed keeper, address target, bytes data);
    event FunctionExecuted(address indexed keeper, address target, bytes data, uint256 value);

    /**
     * @param _endpoint LayerZero V2 endpoint
     * @param _hubEid Hub endpoint ID (Base Sepolia)
     */
    constructor(address _endpoint, uint32 _hubEid) {
        endpoint = _endpoint;
        hubEid = _hubEid;
    }


    /**
     * @notice Called by a keeper to execute a cross-chain function.
     * Sends a message to the hub to verify keeper status and request execution.
     * @param target The contract to call after verification
     * @param data The callData to send to the target contract
     */
    function executeFunction(address target, bytes memory data) external payable {
        // Encode the original keeper (msg.sender), target, and call data
        bytes memory message = abi.encode(msg.sender, target, data);

        // Create standard options with gas limit
        bytes memory options = abi.encodePacked(
            uint16(1),          // type: gas limit override (0x0001)
            uint256(1000000)     // gas limit (200,000)
        );
        
        // Send message to hub with all available ETH
        ILayerZeroEndpointV2(endpoint).send{value: msg.value}(
            MessagingParams({
                dstEid: hubEid,
                receiver: bytes32(uint256(uint160(address(this)))),
                message: message,
                options: options,
                payInLzToken: false
            }),
            address(this)
        );

        emit CrossChainRequestSent(msg.sender, target, data);
    }

    /**
     * @dev Called by LayerZero to deliver a message after hub validation.
     * Executes the target function if hub confirmed the keeper.
     */
    function lzReceive(
        Origin calldata _origin,
        bytes32,            // _guid (unused)
        bytes calldata _message,
        address,            // _executor (unused)
        bytes calldata      // _extraData (unused)
    ) external payable override {
        require(msg.sender == endpoint, "Spoke: Unauthorized endpoint");
        require(_origin.srcEid == hubEid, "Spoke: Invalid source endpoint");
        require(_origin.sender == bytes32(uint256(uint160(address(this)))), "Spoke: Invalid remote sender");

        (address target, bytes memory data) = abi.decode(_message, (address, bytes));
        _executeFunction(target, data);
    }

    /**
     * @dev Executes the target function with provided calldata and ETH.
     */
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Spoke: Function execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    // === Required LayerZero interface stubs ===

    function allowInitializePath(Origin calldata) external pure override returns (bool) {
        return true;
    }

    function nextNonce(uint32, bytes32) external pure override returns (uint64) {
        return 0; // Nonce tracking disabled (can be implemented if needed)
    }
    
    // Add a receive function to accept ETH
    receive() external payable {}
}