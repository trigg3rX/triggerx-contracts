// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ProxyHub} from "./lz/ProxyHub.sol";

// Test contract that exposes _batchBroadcast functionality
contract TestableProxyHub is ProxyHub {
    constructor(
        address _endpoint,
        address _ownerAddress,
        uint32 _srcEid,
        uint32 _thisChainEid,
        address[] memory _initialKeepers
    ) ProxyHub(_endpoint, _ownerAddress, _srcEid, _thisChainEid, _initialKeepers) {}
    
    function testRegisterBroadcast(address keeper) external {
        isKeeper[keeper] = true;
        emit KeeperRegistered(keeper);
        _batchBroadcast(ActionType.REGISTER, keeper);
    }
    
    function testUnregisterBroadcast(address keeper) external {
        isKeeper[keeper] = false;
        emit KeeperUnregistered(keeper);
        _batchBroadcast(ActionType.UNREGISTER, keeper);
    }
} 