// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IAttestationCenter{
    function operatorsIdsByAddress(address) external view returns (uint256);
}