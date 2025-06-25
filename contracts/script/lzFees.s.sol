// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {AvsGovernanceLogic} from "../src/AvsGovernanceLogic.sol";

contract LzFees is Script {
    address payable constant PROXY_HUB = payable(0x9eA6026E7f8ee7444e7D82055ebdf41Fa5424cf7);
    address payable constant AVS_GOVERNANCE = payable(0xe6010F17f9ED9B5B9535Aac79a778982Bc45Cb56);

    address payable constant AVS_GOVERNANCE_LOGIC = payable(0x63a2a4AA6784412D617bf5A9D23D89536CFE7487);
    address constant LZ_ENDPOINT_HOLESKY = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    uint32 constant HOLESKY_EID = 40217;
    uint32 constant BASE_SEPOLIA_EID = 40245;

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        AvsGovernanceLogic avsGovernanceLogic = AvsGovernanceLogic(AVS_GOVERNANCE_LOGIC);

        uint256 gasLimit = avsGovernanceLogic.gasLimit();
        uint256 callValue = avsGovernanceLogic.callValue();
        vm.prank(deployerAddress);
        vm.deal(address(avsGovernanceLogic), 1000000 ether);
        avsGovernanceLogic.setGasOptions(8000000, 2.6 ether);
        console.log("Gas limit:", gasLimit);
        console.log("Call value:", callValue);

        vm.prank(AVS_GOVERNANCE);
        avsGovernanceLogic.afterOperatorRegistered(deployerAddress, 1, [uint256(0), uint256(0), uint256(0), uint256(0)], deployerAddress);

    }
}