// SPDX-License-Identifier: MIT
pragma solidity >=0.8.17;

/// @title PrecompileLib
/// @notice Convenience library exposing Imua precompile addresses as Solidity constants,
///         plus minimal low-level wrappers for the BLS precompile.
/// @dev Addresses are taken from the PrecompileRegistry specification shared by the Imua team.
library PrecompileLib {
    // ---------------------------------------------------------------------
    // Precompile Addresses (as per PrecompileRegistry)
    // ---------------------------------------------------------------------
    address internal constant BECH32_PRECOMPILE = 0x0000000000000000000000000000000000000400;
    address internal constant ASSETS_PRECOMPILE = 0x0000000000000000000000000000000000000804;
    address internal constant DELEGATION_PRECOMPILE = 0x0000000000000000000000000000000000000805;
    address internal constant REWARD_PRECOMPILE = 0x0000000000000000000000000000000000000806;
    address internal constant BLS_PRECOMPILE = 0x0000000000000000000000000000000000000809;
    address internal constant AVS_PRECOMPILE = 0x0000000000000000000000000000000000000901;

} 