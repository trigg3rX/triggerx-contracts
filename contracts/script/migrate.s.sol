// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "forge-std/Script.sol";
import "forge-std/console.sol";

interface ITriggerGasRegistry {
    function batchMigrateUsers(
        address[] calldata users,
        uint256[] calldata ethSpentAmounts,
        uint256[] calldata tgBalances
    ) external payable;
    function getBalance(address user) external view returns (uint256, uint256);
}

contract TriggerGasMigration is Script {
    // Update these addresses for your new deployment
    address constant OLD_REG = 0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a;
    address constant NEW_REG = 0x204F9278D6BB7714D7A40842423dFd5A27cC1b88;

    function run() external {
        uint256 pk = vm.envUint("PRIVATE_KEY");

        // Update RPC URL as needed
        // vm.createSelectFork(vm.envString("BASE_RPC"));
        vm.createSelectFork(vm.envString("OPTIMISM_RPC"));
        vm.startBroadcast(pk);

        ITriggerGasRegistry oldReg = ITriggerGasRegistry(OLD_REG);
        ITriggerGasRegistry newReg = ITriggerGasRegistry(NEW_REG);
        
        // User addresses:
        // address[] memory users = getBaseUsers();
        address[] memory users = getOptimismUsers();

        console.log("Total users to migrate:", users.length);

        // Batch size
        uint256 totalEthSpent = 0;

        // Fetch old state
        uint256[] memory ethSpent = new uint256[](users.length);
        uint256[] memory tgBal = new uint256[](users.length);
        uint256[] memory userPoints = new uint256[](users.length); // Will be zeros since no points in current contract

        for (uint256 k = 0; k < users.length; k++) {
            (ethSpent[k], tgBal[k]) = oldReg.getBalance(users[k]);
            userPoints[k] = 0; // No points in current contract
            totalEthSpent += ethSpent[k];
        }

        // Calculate required ETH for migration
        uint256 requiredETH = 0;
        for (uint256 k = 0; k < ethSpent.length; k++) {
            requiredETH += ethSpent[k];
        }

        console.log("Required ETH for this migration:", requiredETH);

        // Migrate batch
        newReg.batchMigrateUsers{value: requiredETH}(users, ethSpent, tgBal);

        // Verify migration
        for (uint256 k = 0; k < users.length; k++) {
            (uint256 newEthSpent, uint256 newTgBal) = newReg.getBalance(users[k]);
            require(newEthSpent == ethSpent[k], "ETH spent mismatch");
            require(newTgBal == tgBal[k], "TG balance mismatch");
        }

        console.log("Migration completed successfully!");
        console.log("Total ETH migrated:", totalEthSpent);

        vm.stopBroadcast();
    }

    function getBaseUsers() public pure returns (address[] memory) {
        address[] memory users = new address[](29);
        users[0] = 0x5074411b284B10E706dc7316812D6Ed244a56085;
        users[1] = 0xc073A5E091DC60021058346b10cD5A9b3F0619fE;
        users[2] = 0x88826a677aDB340F0c7b8CCd6aF6aD96a40b0085;
        users[3] = 0x0a067a261c5F5e8C4c0b9137430b4FE1255EB62e;
        users[4] = 0xA4857864c95e4f6C6ca54844832bDDe036e7a3Ed;
        users[5] = 0x4BF8E1E54E50E3b64E9c486D4230Ee4F9e7dE792;
        users[6] = 0x1A396b462a0701fD0C284C793f290Eb724cf2311;
        users[7] = 0x334eBE7A724ea661fEC3246f21Ce25e3dc25f9C6;
        users[8] = 0xe6AF6d5C12e4790Dbf8BbAa3715146BAA485Fc3B;
        users[9] = 0x4D9e5bf28A12266BFA90BBEEAA3F6600B5561B76;
        users[10] = 0x43a4685A71FecDb47484763373F9dfeE2864c149;
        users[11] = 0xC9dC9c361c248fFA0890d7E1a263247670914980;
        users[12] = 0xee2094f32c67B540CFB55C58D38D04718c9fea5B;
        users[13] = 0x0763ec3E83d89bf431a665403226af438508E97F;
        users[14] = 0x854cCC3f4C1b4EB2030108Db3454fe8eaee9F676;
        users[15] = 0x8b5BA20Fd0d5cB69594b98614167155EE5e3a393;
        users[16] = 0x077f6F930Bc1e892Ebaf4492097161d448ce442f;
        users[17] = 0xF4f20bf9f8d7D12e2C0026A9a253C53e85FceA67;
        users[18] = 0x6d14CdB5aD9aBAf8D3eA1f5F9bdCE26C95c406E5;
        users[19] = 0xD8a5F149F40d2C79143Ad8EA9272828F8a00d15a;
        users[20] = 0xdfC54a9f540182a01FFc2a94B2Cc62F698E7a899;
        users[21] = 0x07C4b9865dC8F2970C8e8E518e5DF137fF074b42;
        users[22] = 0x62Ee26BC056D1676E4064dD637B3d317F33532cf;
        users[23] = 0x7C81c3d1241Cc5d6d46b53a0a67e67c9610a7653;
        users[24] = 0x7Db951c0E6D8906687B459427eA3F3F2b456473B;
        users[25] = 0x45E4657eE118325F852609e9dB45211fe1A0044C;
        users[26] = 0x149fd3c6326900Bb0489095579e8e8011133d06b;
        users[27] = 0xd07ba8f10a16EF1e5024CF58a84aDB6Df228Bb4A;
        users[28] = 0x17d777472885F9b84bd540C3A93FaecEA1204940;
        return users;
    }

    function getOptimismUsers() public pure returns (address[] memory) {
        address[] memory users = new address[](46);
        users[0] = 0x88826a677aDB340F0c7b8CCd6aF6aD96a40b0085;
        users[1] = 0x1A396b462a0701fD0C284C793f290Eb724cf2311;
        users[2] = 0xae229e741938fd71BE75247f6B14B280031e4777;
        users[3] = 0x2460F77c64B6CA587178c8d931994b1e35C6eA5a;
        users[4] = 0xe6c691B74E4Ad2246a40D0146A18a145487E1DFe;
        users[5] = 0x0a067a261c5F5e8C4c0b9137430b4FE1255EB62e;
        users[6] = 0x0763ec3E83d89bf431a665403226af438508E97F;
        users[7] = 0x845cd84602c7c97C57DbcC7511FEF59365fD6744;
        users[8] = 0x5F954aB624bBf9369b81C9a2076eDc32ecdEB542;
        users[9] = 0x230f0AEe934C9d1Fdf2a4cE0Fc46738786e26B1a;
        users[10] = 0x4d264374A1EA425BAdd7930628e3b1481b0dBb69;
        users[11] = 0x2b64ea2059344E7BCABf3881a51de92fBbe7eDfc;
        users[12] = 0x7C81c3d1241Cc5d6d46b53a0a67e67c9610a7653;
        users[13] = 0xdfC54a9f540182a01FFc2a94B2Cc62F698E7a899;
        users[14] = 0xBd0Bc08df8c9Cd7e70e78C00f6eb94449be20C04;
        users[15] = 0x31533eDB7740462f6A0d031Cb4CCc87E90F3Ec6d;
        users[16] = 0xC2A3d741ae143BF9f27a6021Fd0a1aCBBAfcB40b;
        users[17] = 0x998638A1Ec37784f059461B072cB9a4C2ce6E155;
        users[18] = 0xA4857864c95e4f6C6ca54844832bDDe036e7a3Ed;
        users[19] = 0xD8a5F149F40d2C79143Ad8EA9272828F8a00d15a;
        users[20] = 0xf52233666b4ff2336EA47D4509E6371D91A303C1;
        users[21] = 0xF4Fcb1A9b5D853cE1C32B5513D4D5cCeC2a3CFe0;
        users[22] = 0x8b5BA20Fd0d5cB69594b98614167155EE5e3a393;
        users[23] = 0x49089A3bFc1F67E1e83Cf56eF441249a09a72c89;
        users[24] = 0xe6AF6d5C12e4790Dbf8BbAa3715146BAA485Fc3B;
        users[25] = 0x20e3Bec40047423eEeD14F88935Ed5FF83457F33;
        users[26] = 0x618dC17De04cB1DDD07E2f945fFb8357E0cAa370;
        users[27] = 0x2871e6255d8560F2e16a03373B957F86bf27382d;
        users[28] = 0x4D9e5bf28A12266BFA90BBEEAA3F6600B5561B76;
        users[29] = 0x0DFF2137f66AAC67Adfb7ffC88a2a679B46931C0;
        users[30] = 0x62Ee26BC056D1676E4064dD637B3d317F33532cf;
        users[31] = 0xF4f20bf9f8d7D12e2C0026A9a253C53e85FceA67;
        users[32] = 0x6d14CdB5aD9aBAf8D3eA1f5F9bdCE26C95c406E5;
        users[33] = 0x17d777472885F9b84bd540C3A93FaecEA1204940;
        users[34] = 0x07C4b9865dC8F2970C8e8E518e5DF137fF074b42;
        users[35] = 0xF27F42715811394E6b42d29715100db52c9d33FB;
        users[36] = 0xE2Fdf521907F5A548D01047797D5B580a24a9822;
        users[37] = 0x63Bb96dC6467240DbbDE51CCF16dE7f7C8FdC02D;
        users[38] = 0x854cCC3f4C1b4EB2030108Db3454fe8eaee9F676;
        users[39] = 0x130b11D9CC4ACEB69b9a68031860751Cb4aF1618;
        users[40] = 0x93ED9Bbc3c7acC85045fD6C8857C8e37968eB1A8;
        users[41] = 0x8443f65FEe323c1D30C52aa95152AC89FCb0dD31;
        users[42] = 0x6c275fDff8008fCe66C779CB56BdACACbB9F6436;
        users[43] = 0x655558724583731AEcB88711E76Aed9081225D79;
        users[44] = 0x0255F7A175f73a05765719c165445F63155aF8E9;
        users[45] = 0x58693935b683F73B6Cff9390Dd4fFA707dC4F069;
        return users;
    }
}
