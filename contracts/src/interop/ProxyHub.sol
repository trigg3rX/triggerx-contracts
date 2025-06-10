// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

// import "@interop/interfaces/IL2ToL2CrossDomainMessenger.sol";

// error CallerNotL2ToL2CrossDomainMessenger();
// error InvalidCrossDomainSender();
// error InvalidTargetContract();
// error KeeperNotRegistered();

// contract HubContract {
//     // Predeployed L2ToL2CrossDomainMessenger address on Base Sepolia
//     address public constant MESSENGER = 0x4200000000000000000000000000000000000023;
//     IL2ToL2CrossDomainMessenger messenger = IL2ToL2CrossDomainMessenger(MESSENGER);
    
//     // Mapping to track registered keepers
//     mapping(address => bool) public isKeeper;
    
//     event FunctionExecuted(address indexed keeper, address target, bytes data);
//     event CrossChainRequestReceived(address indexed keeper, address target, uint256 sourceChainId);
//     event CrossChainExecutionSent(address indexed spokeContract, uint256 destinationChainId);

//     // Owner of the contract for restricted functions
//     address public avsGovernance;

//     struct KeeperMessage {
//         address target;
//         address data;
//     }

//     constructor(address _avsGovernance) {
//         avsGovernance = _avsGovernance;
//     }


//     // Execute function locally or relay to spoke
//     function executeFunction(address target, bytes memory data) external {

//         if(msg.sender == MESSENGER) {
//             if (messenger.crossDomainMessageSender() != address(this)) revert InvalidCrossDomainSender();

//             KeeperMessage memory message = KeeperMessage({
//                 target: target,
//                 data: data
//             });

            
//         }
//         else {
//             if(!_checkKeeper(msg.sender)) revert KeeperNotRegistered();

//             _executeFunction(target,data);
//         }

//     }

//     function _checkKeeper(address keeper) internal view returns (bool) {
        
//     }


//     function _executeFunction(address target,bytes memory callData) internal payable returns (bytes memory) {

//         (bool success, bytes memory result) = target.call{value: msg.value}(callData);
//         require(success, "Function call failed");
        
//         emit FunctionExecuted(msg.sender, target,callData, msg.value);
//         return result;
//     }
// }