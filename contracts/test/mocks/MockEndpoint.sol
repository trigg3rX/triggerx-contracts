// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {MessagingParams, MessagingFee, MessagingReceipt} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OAppReceiver.sol";

// Mock implementation with only the functions we need for testing
contract MockEndpoint {
    uint256 private quotedFee;
    bytes32 private msgGuid;
    
    function setQuotedFee(uint256 _fee) external {
        quotedFee = _fee;
    }
    
    function setMsgGuid(bytes32 _guid) external {
        msgGuid = _guid;
    }
    
    function quote(MessagingParams calldata, address) external view returns (MessagingFee memory) {
        return MessagingFee({
            nativeFee: quotedFee,
            lzTokenFee: 0
        });
    }
    
    function send(
        MessagingParams calldata,
        address
    ) external payable returns (MessagingReceipt memory) {
        // Ensure we have enough funds
        require(msg.value >= quotedFee, "Insufficient fee");
        
        // Return the configured guid instead of reverting
        return MessagingReceipt({
            guid: msgGuid,
            nonce: 1,
            fee: MessagingFee({
                nativeFee: msg.value,
                lzTokenFee: 0
            })
        });
    }
    
    function setPeer(uint32, bytes32) external {}
    
    function setDelegate(address) external {}
    
    // Only include the minimal set of functions needed for your tests
    function getSendLibrary(address, uint32) external pure returns (address) {
        return address(0);
    }
    
    function isValidSendLibrary(address, address, uint32) external pure returns (bool) {
        return true;
    }
    
    // Add other functions as needed for your tests
} 