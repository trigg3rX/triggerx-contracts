// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/// @notice Migration script for TriggerGasRegistry users
/// @dev Usage:
///   1. Fetch users automatically: node script/fetchUsers.js (requires RPC URLs in env)
///   2. Run migration for a specific chain:
///      forge script script/migrate.s.sol:TriggerGasMigration --sig "run(string,string)" "base" "data/users.json" -vvvv
///   3. Or set CHAIN and USERS_FILE env vars:
///      CHAIN=optimism USERS_FILE=data/users.json forge script script/migrate.s.sol:TriggerGasMigration -vvvv
///   4. JSON file format: {"base": {"chainId": 84532, "users": [...]}, "optimism": {...}, ...}

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "forge-std/StdJson.sol";

interface IOldRegistry {
    function getBalance(address user) external view returns (uint256, uint256);

    function withdrawETH(uint256 amount, string memory reason) external;
}

interface INewRegistry {
    function batchMigrateUsers(
        address[] calldata users,
        uint256[] calldata ethAmounts
    ) external payable;

    function getBalance(address user) external view returns (uint256);
}

contract TriggerGasMigration is Script {
    using stdJson for string;

    // Update these addresses for your new deployment
    address constant OLD_REG = 0x93dDB2307F3Af5df85F361E5Cddd898Acd3d132d;
    address constant NEW_REG = 0xe2AC670F7D66c69D547A44D08F9bA1Fc0Fc0f991;

    function run() external {
        // Try to get chain from env, default to "optimism"
        try vm.envString("CHAIN") returns (string memory chain) {
            runMigration(chain, "");
        } catch {
            runMigration("arbitrum", "");
        }
    }

    function runMigration(
        string memory chainName,
        string memory usersFilePath
    ) public {
        uint256 pk = vm.envUint("PRIVATE_KEY");

        // Select fork based on chain
        if (keccak256(bytes(chainName)) == keccak256(bytes("base"))) {
            vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        } else if (
            keccak256(bytes(chainName)) == keccak256(bytes("optimism"))
        ) {
            vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        } else if (
            keccak256(bytes(chainName)) == keccak256(bytes("arbitrum"))
        ) {
            vm.createSelectFork(vm.envString("ARB_RPC"));
        } else if (
            keccak256(bytes(chainName)) == keccak256(bytes("ethereum"))
        ) {
            vm.createSelectFork(vm.envString("ETH_SEPOLIA_RPC"));
        }

        vm.startBroadcast(pk);

        IOldRegistry oldReg = IOldRegistry(OLD_REG);
        INewRegistry newReg = INewRegistry(NEW_REG);

        // Load user addresses from JSON file or use hardcoded fallback
        address[] memory users;
        string memory filePath = usersFilePath;

        if (bytes(filePath).length == 0) {
            // Try to load from environment variable
            try vm.envString("USERS_FILE") returns (string memory envFile) {
                filePath = envFile;
            } catch {
                filePath = "data/users.json"; // Default path
            }
        }

        if (bytes(filePath).length > 0) {
            users = loadUsersFromChain(filePath, chainName);
        }

        console.log("Total users to migrate:", users.length);

        // // Withdraw all ETH from old registry
        // uint256 oldBalance = address(OLD_REG).balance;
        // console.log("Old Registry Balance:", oldBalance);
        // if (oldBalance > 0) {
        //     try oldReg.withdrawETH(oldBalance, "Migration") {
        //         console.log("Withdrawn ETH from old registry");
        //     } catch Error(string memory reason) {
        //         console.log("Withdraw failed:", reason);
        //     } catch {
        //         console.log("Withdraw failed (unknown reason)");
        //     }
        // }

        // Calculate migration amounts
        uint256[] memory migrationAmounts = new uint256[](users.length);
        uint256 totalMigrationEth = 0;

        for (uint256 k = 0; k < users.length; k++) {
            (, uint256 tgBal) = oldReg.getBalance(users[k]);

            // Convert TG balance to ETH amount (1000 TG = 1 ETH)
            uint256 ethAmount = tgBal / 1000;

            migrationAmounts[k] = ethAmount;
            totalMigrationEth += ethAmount;

            // console.log("User:", users[k]);
            // console.log("TG Bal:", tgBal);
            // console.log("ETH Amt:", ethAmount);
        }

        console.log("Required ETH for this migration:", totalMigrationEth);

        uint256 newRegBalBefore = address(NEW_REG).balance;

        // Migrate batch
        newReg.batchMigrateUsers{value: totalMigrationEth}(
            users,
            migrationAmounts
        );

        // Verify migration
        console.log("\n--- Verifying Migration ---");

        // 1. Check Old Registry Balance
        uint256 oldRegBalAfter = address(OLD_REG).balance;
        if (oldRegBalAfter == 0) {
            console.log("OK: Old Registry is empty");
        } else {
            console.log(
                "Warning: Old Registry still has balance:",
                oldRegBalAfter
            );
        }

        // 2. Check New Registry Balance
        uint256 newRegBalAfter = address(NEW_REG).balance;
        if (newRegBalAfter == newRegBalBefore + totalMigrationEth) {
            console.log("OK: New Registry received correct ETH amount");
        } else {
            console.log("Error: New Registry ETH balance mismatch");
            console.log("   Expected increase:", totalMigrationEth);
            console.log(
                "   Actual increase:",
                newRegBalAfter - newRegBalBefore
            );
        }

        // 3. Verify User Balances
        uint256 verifiedCount = 0;
        for (uint256 k = 0; k < users.length; k++) {
            uint256 newEthBal = newReg.getBalance(users[k]);
            if (newEthBal == migrationAmounts[k]) {
                verifiedCount++;
            } else {
                console.log("Error: Balance mismatch for user:", users[k]);
                console.log("   Expected:", migrationAmounts[k]);
                console.log("   Actual:", newEthBal);
            }
        }

        if (verifiedCount == users.length) {
            console.log("OK: All", users.length, "users verified successfully");
        } else {
            revert("Migration verification failed for some users");
        }

        console.log("\nMigration completed successfully!");
        console.log("Total ETH migrated:", totalMigrationEth);

        vm.stopBroadcast();
    }

    /// @notice Load user addresses from a multi-chain JSON file
    /// @param filePath Path to the JSON file (relative to project root or absolute)
    /// @param chainName Chain name: "base", "optimism", "arbitrum", or "ethereum"
    /// @return Array of user addresses for the specified chain
    /// @dev JSON file should have format: {"base": {"chainId": 84532, "users": [...]}, ...}
    function loadUsersFromChain(
        string memory filePath,
        string memory chainName
    ) public view returns (address[] memory) {
        string memory json = vm.readFile(filePath);
        string memory chainPath = string.concat(".", chainName, ".users");
        address[] memory users = json.readAddressArray(chainPath);
        return users;
    }
}
