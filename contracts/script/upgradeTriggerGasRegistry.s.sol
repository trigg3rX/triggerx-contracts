// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerGasRegistry_Migrate} from "../src/TriggerGasRegistry_Migrate.sol";
import {CREATE3} from "lib/solady/src/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TriggerGasRegistryDeploy is Script {
    function run() public {

        bytes32 implementation_salt = keccak256(abi.encodePacked("ImplementationV3"));

        bytes memory implementation_code = type(TriggerGasRegistry_Migrate).creationCode;


        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        address implementation = CREATE3.deployDeterministic(implementation_code, implementation_salt);

        // TODO: upgrade the implementation
        address payable proxy = payable(0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a);
        TriggerGasRegistry_Migrate(proxy).upgradeToAndCall(
            implementation,""
        );

        vm.stopBroadcast();

        // also Verify balances
        address[] memory users = new address[](6);
        users[0] = 0xD5E9061656252a0b44D98C6944B99046FDDf49cA;
        users[1] = 0xBc0435FB4f37345a420fbD09e5700f3A72fd0534;
        users[2] = 0xC76EA60887CA82C474cf6dfc17f918DDd68D6cA2;
        users[3] = 0x751F7CB39462FFc9cbE0276D411EbA4a2a7Ebe46;
        users[4] = 0xc073A5E091DC60021058346b10cD5A9b3F0619fE;
        users[5] = 0xE3304AB782c3272D1D7964ba7043e11Fd7ecFEB8;
        uint256[] memory ethSpentAmounts = new uint256[](6);
        uint256[] memory tgBalances = new uint256[](6);
        uint256[] memory userPoints = new uint256[](6);

        // TODO: console balances
        for (uint256 i = 0; i < users.length; i++) {
            (uint256 ethSpent, uint256 tgBalance) = TriggerGasRegistry_Migrate(proxy).balances(users[i]);
            console.log("User:", users[i]);
            console.log("ETH spent:", ethSpent);
            console.log("TG balance:", tgBalance);
            console.log("Points:", TriggerGasRegistry_Migrate(proxy).points(users[i]));
        }

        console.log("Proxy upgraded to:", proxy);
        console.log("Implementation deployed to:", implementation);
        console.log("Deployer:", deployer);
        console.log("Proxy owner:", TriggerGasRegistry_Migrate(proxy).owner());
    }
}