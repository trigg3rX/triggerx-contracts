// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Vm} from "forge-std/Vm.sol";
import {ProxyAdmin} from "@openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {EmptyContract} from "@eigenlayer-contracts/test/mocks/EmptyContract.sol";

library UpgradeableProxyLib {
    // Storage slot constants for proxy contracts
    bytes32 internal constant IMPLEMENTATION_SLOT =
        bytes32(uint256(keccak256("keeperX.contracts.implementation")) - 1);
    bytes32 internal constant ADMIN_SLOT =
        bytes32(uint256(keccak256("keeperX.contracts.admin")) - 1);

    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function deployProxyAdmin() internal returns (address) {
        return address(new ProxyAdmin());
    }

    function setUpEmptyProxy(
        address admin
    ) internal returns (address) {
        address emptyContract = address(new EmptyContract());
        return address(new TransparentUpgradeableProxy(emptyContract, admin, ""));
    }

    function upgrade(address proxy, address impl) internal {
        ProxyAdmin admin = getProxyAdmin(proxy);
        admin.upgrade(TransparentUpgradeableProxy(payable(proxy)), impl);
    }

    function upgradeAndCall(address proxy, address impl, bytes memory initData) internal {
        ProxyAdmin admin = getProxyAdmin(proxy);
        admin.upgradeAndCall(TransparentUpgradeableProxy(payable(proxy)), impl, initData);
    }

    function getImplementation(
        address proxy
    ) internal view returns (address) {
        bytes32 value = vm.load(proxy, IMPLEMENTATION_SLOT);
        return address(uint160(uint256(value)));
    }

    function getProxyAdmin(
        address proxy
    ) internal view returns (ProxyAdmin) {
        bytes32 value = vm.load(proxy, ADMIN_SLOT);
        return ProxyAdmin(address(uint160(uint256(value))));
    }
}