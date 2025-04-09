// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.0;

// import "@interop/interfaces/IL2ToL2CrossDomainMessenger.sol";

// error CallerNotL2ToL2CrossDomainMessenger();
// error InvalidCrossDomainSender();
// error InvalidTargetContract();


// contract SpokeContract {
//     // Predeployed L2ToL2CrossDomainMessenger address on OP Sepolia
//     address public immutable MESSENGER = 0x4200000000000000000000000000000000000023;

//     uint256 public immutable hubChainId;

//     constructor(uint256 _hubChainId) {
//         hubChainId = _hubChainId;
//     }

//     struct KeeperMessage {
//         address keeper;
//         address target;
//         bytes data;
//     }
    
//     // Events
//     event CrossChainRequestSent(address indexed keeper, address target, bytes data);
//     event FunctionExecuted(address indexed keeper, address target, bytes data);

//     // Execute function or forward to hub
//     function executeFunction(address target, bytes memory data) external {
//         IL2ToL2CrossDomainMessenger messenger = IL2ToL2CrossDomainMessenger(MESSENGER);

//         if (msg.sender != MESSENGER) {
//             // Direct call from keeper on OP Sepolia
//             KeeperMessage memory message = KeeperMessage({
//                 keeper: msg.sender,
//                 target: target,
//                 data: data
//             });

//             messenger.sendMessage(
//                 hubChainId,
//                 address(this),
//                 abi.encodeCall(this.executeFunction, message);
//             );

//             emit CrossChainRequestSent(msg.sender, target, data);
            
//         } else {

//             if (messenger.crossDomainMessageSender() != address(this)) revert InvalidCrossDomainSender();

//             executeFunction(target, data);
//         }
//     }

//     function _executeFunction(address target,bytes memory callData) internal payable returns (bytes memory) {

//         (bool success, bytes memory result) = target.call{value: msg.value}(callData);
//         require(success, "Function call failed");
        
//         emit FunctionExecuted(msg.sender, target,callData, msg.value);
//         return result;
//     }
// }