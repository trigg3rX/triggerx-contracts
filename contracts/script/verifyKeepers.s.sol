// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";

contract VerifyKeepers is Script {
    // ProxyHub addresses for different chains
    address constant PROXY_ADDRESS = 0x96c5F575940DBe98fd9600F74F0c36139A7Be2A2;

    // Operators array
    address[] OPERATORS = [
        0x19aBAdfBb70e60CEa506291BeE00d8307CB62667,
        0x2dcFBc2613b60EC755A2530F818d0E1A8Ed1Db3f,
        0x95AB9BCa0eECE37468fd6038E39021e7EEd208a3,
        0xBaE997eFC6E496d8A34181C2443d97b491A3f057,
        0x45E4657eE118325F852609e9dB45211fe1A0044C,
        0xEbe4f49161eb9002eD19E06F943DCe73D24B596E,
        0x350ddC16818CDFA1E0c06c26Fd5e76360032faa4,
        0x011FCbAE5f306cd793456ab7d4c0CC86756c693D,
        0x0a067a261c5F5e8C4c0b9137430b4FE1255EB62e
    ];

    function run() external {
        // Create forks for both chains
        uint256 baseFork = vm.createFork(vm.envString("BASE_SEPOLIA_RPC"));
        uint256 opFork = vm.createFork(vm.envString("OPSEPOLIA_RPC"));

        // Verify on Base Sepolia
        vm.selectFork(baseFork);
        console2.log("\nVerifying on Base Sepolia:");
        verifyKeepersOnChain(OPERATORS, PROXY_ADDRESS);

        // Verify on OP Sepolia
        vm.selectFork(opFork);
        console2.log("\nVerifying on OP Sepolia:");
        verifyKeepersOnChain(OPERATORS, PROXY_ADDRESS);
    }

    function verifyKeepersOnChain(address[] memory operators, address proxyHubAddress) view internal {
        ProxyHub proxyHub = ProxyHub(payable(proxyHubAddress));
        
        for (uint i = 0; i < operators.length; i++) {
            bool isKeeper = proxyHub.isKeeper(operators[i]);
            console2.log("Operator %s is %s on chain %s", 
                operators[i], 
                isKeeper ? "a keeper" : "not a keeper",
                proxyHubAddress
            );
        }
    }
} 