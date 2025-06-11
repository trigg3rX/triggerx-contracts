// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "forge-std/Script.sol";
import "forge-std/console.sol";

interface ITriggerXStakeRegistry {
    function getStake(address user) external view returns (uint256, uint256);
    function points(address user) external view returns (uint256);
}

interface ITriggerGasRegistry {
    function batchMigrateUsers(
        address[] calldata users,
        uint256[] calldata ethSpentAmounts,
        uint256[] calldata tgBalances,
        uint256[] calldata userPoints
    ) external;
    function getBalance(address user) external view returns (uint256, uint256);
    function points(address user) external view returns (uint256);
}

contract TriggerGasMigration is Script {
    address constant OLD_REG    = 0xF1d505d1f6df11795c77A8A1b7476609E7b6361a;
    address constant NEW_REG    = 0x85ea3eB894105bD7e7e2A8D34cf66C8E8163CD2a;

    function run() external {
        uint256 pk      = vm.envUint("PRIVATE_KEY");

        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(pk);

        ITriggerXStakeRegistry oldReg = ITriggerXStakeRegistry(OLD_REG);
        ITriggerGasRegistry      newReg = ITriggerGasRegistry(NEW_REG);

        address[] memory users = new address[](70);
        users[0] = 0xD5E9061656252a0b44D98C6944B99046FDDf49cA;
        users[1] = 0xBc0435FB4f37345a420fbD09e5700f3A72fd0534;
        users[2] = 0xC76EA60887CA82C474cf6dfc17f918DDd68D6cA2;
        users[3] = 0x751F7CB39462FFc9cbE0276D411EbA4a2a7Ebe46;
        users[4] = 0xc073A5E091DC60021058346b10cD5A9b3F0619fE;
        users[5] = 0xE3304AB782c3272D1D7964ba7043e11Fd7ecFEB8;
        users[6] = 0x28C8568C73E62E42246dFa4a6aAbeCF4Ef497bE3;
        users[7] = 0xB4e6ee231C86bBcCB35935244CBE9cE333D30Bdf;
        users[8] = 0x88826a677aDB340F0c7b8CCd6aF6aD96a40b0085;
        users[9] = 0xEd9C18fcc15c02b8B590354dd46C014c62F3DCDE;
        users[10] = 0x45E4657eE118325F852609e9dB45211fe1A0044C;
        users[11] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        users[12] = 0x93c0f6dD62D4A5DcfBB4a887048Ff563cBa9C9Fb;
        users[13] = 0x21e3dBDB71FDe76BC37fA30A0702E7CD2D984546;
        users[14] = 0x7c0B569F6575608614e0b0dd9D65359D0757a34f;
        users[15] = 0x8FCE7E1C5Ac0d40451FEd06c127f42F0fd7E8ed0;
        users[16] = 0x931628CEb6f30c580B58813e19744a0359f37bC8;
        users[17] = 0xBBE7A17CCD734Fb382de00970CB1DB113DC73Cb5;
        users[18] = 0x2465b8016dd6d573e790Bd99323Ba7e91a2868CF;
        users[19] = 0x334eBE7A724ea661fEC3246f21Ce25e3dc25f9C6;
        users[20] = 0x70D5D108fEf49be40d2c2388ae722B23F4Df2175;
        users[21] = 0x69b4fe79f7D054468B23858204978e6F677e044A;
        users[22] = 0x4d264374A1EA425BAdd7930628e3b1481b0dBb69;
        users[23] = 0x6d14CdB5aD9aBAf8D3eA1f5F9bdCE26C95c406E5;
        users[24] = 0x655558724583731AEcB88711E76Aed9081225D79;
        users[25] = 0x149fd3c6326900Bb0489095579e8e8011133d06b;
        users[26] = 0xC81354b4661C425dA082601b578c6EA9A6510Fa1;
        users[27] = 0x7C81c3d1241Cc5d6d46b53a0a67e67c9610a7653;
        users[28] = 0xf439883b6A26F5D74AC0472aa4365cc7e6C2A4Dc;
        users[29] = 0xC9dC9c361c248fFA0890d7E1a263247670914980;
        users[30] = 0x71F1DaA35BD934A412577C1e3b5FfFB3572eeBed;
        users[31] = 0xe9f3EFAC8160E4Fff14F9e7828DE69cAb15541A5;
        users[32] = 0x0FbF9c8212Ce7a85D70350f438C28284A806aD5a;
        users[33] = 0x844294eBecC14e934d9d72a319E6Fa9DbAf95Bc5;
        users[34] = 0xC05cA8C526422Ef68C0cFdFE87fd8dB577174F56;
        users[35] = 0x82Bb444dfc9B9799A59B068A1677c2fEdF032917;
        users[36] = 0x7Db951c0E6D8906687B459427eA3F3F2b456473B;
        users[37] = 0xa059C8e31dAc3Eb0cEe5C821202A3da07E876611;
        users[38] = 0xba7e4875EB063006b7C9E94EA6bE3ad8be90be4E;
        users[39] = 0x8a51c19C4B3ed916769B56DbC039f8DF876029bE;
        users[40] = 0xB5291a0Df6A91Ab9e75c4992eA3f4abeC7f548Be;
        users[41] = 0x71Cfae5b40Ce7f623208DE3A8D08162572A4AfD9;
        users[42] = 0x055e07963B956AD77087d70fE351d5B5Cf646208;
        users[43] = 0x5074411b284B10E706dc7316812D6Ed244a56085;
        users[44] = 0xe8c1A8f1c4fb011dE16F1E9BDC8bf8653fADD713;
        users[45] = 0x34Ff6F8A4C7F1673D4c046170d006C10763120A6;
        users[46] = 0xF4Fcb1A9b5D853cE1C32B5513D4D5cCeC2a3CFe0;
        users[47] = 0xDa76B028117544D91a6bC86034f3421D94F97cAA;
        users[48] = 0x998638A1Ec37784f059461B072cB9a4C2ce6E155;
        users[49] = 0x1A396b462a0701fD0C284C793f290Eb724cf2311;
        users[50] = 0xaeb1D436436D86b448a45538bc311AAdc18a6DEc;
        users[51] = 0xA4857864c95e4f6C6ca54844832bDDe036e7a3Ed;
        users[52] = 0x63Bb96dC6467240DbbDE51CCF16dE7f7C8FdC02D;
        users[53] = 0xE4346694184e76C1F2eAF1BeBbdf2f30Da5460E8;
        users[54] = 0xdD6c4B5ff53dECDd5F770840B1803F6880d5c057;
        users[55] = 0x28a8681595Ff376f12E9eEc274c01b0928b50D3a;
        users[56] = 0xBb4768E6D19f9c025fe68870B451b36D1F614b5B;
        users[57] = 0xdfC54a9f540182a01FFc2a94B2Cc62F698E7a899;
        users[58] = 0x4D85968D0507db91D5c21db19b1652af76671Fd0;
        users[59] = 0x17d3EE21a6739D9452F714C83c9122D21Af41786;
        users[60] = 0x50a5271FFf29ABAA78F3d63A72aa849B122E7aBD;
        users[61] = 0xD8a5F149F40d2C79143Ad8EA9272828F8a00d15a;
        users[62] = 0x49089A3bFc1F67E1e83Cf56eF441249a09a72c89;
        users[63] = 0x1407BA540dBB5957e3760C67A63791cD22f1F181;
        users[64] = 0x0a067a261c5F5e8C4c0b9137430b4FE1255EB62e;
        users[65] = 0x20e3Bec40047423eEeD14F88935Ed5FF83457F33;
        users[66] = 0xBd96B9cB77513293A0C7541238c88d0268Fa55c5;
        users[67] = 0x12f36dCB059B50Bb862336299571380103DCAbD4;
        users[68] = 0x0000EA905fFcb21aA98251CfA292cdeA347a6416;
        users[69] = 0xd06CBfc7684372097B80682908DEeB20e95b6307;

        console.log("total users:", users.length);

        // batch size (tweak to your gas limits)
        uint256 batch = 100;
        for (uint256 i = 0; i < users.length; i += batch) {
            uint256 end = i + batch;
            if (end > users.length) end = users.length;

            // slice users[i:end]
            address[] memory chunk = new address[](end - i);
            for (uint256 j = i; j < end; j++) {
                chunk[j - i] = users[j];
            }

            // fetch old state
            uint256[] memory ethSpent = new uint256[](chunk.length);
            uint256[] memory tgBal    = new uint256[](chunk.length);
            uint256[] memory ptsArr   = new uint256[](chunk.length);

            for (uint256 k = 0; k < chunk.length; k++) {
                (ethSpent[k], tgBal[k]) = oldReg.getStake(chunk[k]);
                ptsArr[k]               = oldReg.points(chunk[k]);
            }

            // migrate
            newReg.batchMigrateUsers(chunk, ethSpent, tgBal, ptsArr);
            console.log("migrated batch", i / batch + 1, "size", chunk.length);

            // verify
            for (uint256 k = 0; k < chunk.length; k++) {
                (uint256 e2, uint256 t2) = newReg.getBalance(chunk[k]);
                uint256 p2              = newReg.points(chunk[k]);
                require(e2 == ethSpent[k],   "ETH mismatch");
                require(t2 == tgBal[k],       "TG mismatch");
                require(p2 == ptsArr[k],      "pts mismatch");
            }
            console.log("batch", i / batch + 1, "verified");
        }

        vm.stopBroadcast();
        console.log("migration complete");
    }
}
