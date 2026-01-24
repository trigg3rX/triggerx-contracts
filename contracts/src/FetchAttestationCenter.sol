// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {AttestationCenter} from "@othentic-core-contracts/NetworkManagement/L2/AttestationCenter.sol";

/**
 * @title FetchAttestationCenter
 * @dev A contract for fetching the AttestationCenter contract
 * @notice This contract's sole purpose is to include the AttestationCenter contract in the bindings
 */
contract FetchAttestationCenter {
    // This contract exists only to ensure AttestationCenter is compiled
    AttestationCenter public attestationCenter;
}