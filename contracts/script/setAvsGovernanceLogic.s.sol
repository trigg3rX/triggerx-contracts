// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.25;

import {Script} from "forge-std/Script.sol";
import {IAvsGovernanceLogic} from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";

interface IAvsGovernance {
    function setAvsGovernanceLogic(IAvsGovernanceLogic _avsGovernanceLogic) external;
}

contract SetAvsGovernanceLogic is Script {
    function run() external {
        // AVS Governance address on Holesky
        
        address avsGovernance = 0xe6010F17f9ED9B5B9535Aac79a778982Bc45Cb56;
        // AVS Governance Logic address on Holesky
      
        address avsGovernanceLogic = 0x2cFADBDd050bB83A45B80C2a045516470A67c991;

        // Start broadcasting transactions
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        // Cast to IAvsGovernance interface and call setAvsGovernanceLogic
        IAvsGovernance(avsGovernance).setAvsGovernanceLogic(IAvsGovernanceLogic(avsGovernanceLogic));

        // Stop broadcasting transactions
        vm.stopBroadcast();
    }
}
