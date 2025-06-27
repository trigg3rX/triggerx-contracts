// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "./interfaces/IAVSManager.sol" as avs;
import "@openzeppelin/contracts/access/Ownable.sol";

contract Avs is Ownable {

    event TaskResolved(uint64 taskId, address taskAddress);

    // task submitter decides on the criteria for a task to be completed
    // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
    //      this is for the challenge logic to verify
    // task is completed (and contract will accept its TaskResponse)
    // are signed by at least  ThresholdPercentage of the operators
    struct Task {
        string name;
        uint64 taskId;
        uint64 numberToBeSquared;
        uint64 taskResponsePeriod;
        uint64 taskChallengePeriod;
        uint8 thresholdPercentage;
        uint64 taskStatisticalPeriod;
    }

    event TaskCreated(
        uint256 taskId,
        address issuer,
        string name,
        uint64 numberToBeSquared,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint8 thresholdPercentage,
        uint64 taskStatisticalPeriod
    );
    struct ChallengeReq {
        uint64 taskId;
        address taskAddress;
        uint64 numberToBeSquared;
        avs.OperatorResInfo[]  infos;
        address[]  signedOperators ;
        address[]  noSignedOperators;
        string  taskTotalPower;
    }

    constructor() Ownable(msg.sender) {
        // constructor
    }

    function registerAVS(
        avs.AVSParams calldata params
    ) public onlyOwner returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.registerAVS(
            params
        );
        return success;
    }


    function updateAVS(
        avs.AVSParams calldata params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.updateAVS(
            params
        );
        return success;
    }

    function registerOperatorToAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.registerOperatorToAVS(
            msg.sender
        );
        return success;
    }

    function deregisterOperatorFromAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.deregisterOperatorFromAVS(
            msg.sender
        );
        return success;
    }

    function registerBLSPublicKey(
        address  avsAddr,
        bytes memory pubKey,
        bytes memory pubKeyRegistrationSignature
    ) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.registerBLSPublicKey(
            msg.sender,
            avsAddr,
            pubKey,
            pubKeyRegistrationSignature
        );
        return success;
    }

    function operatorSubmitTask(
        uint64 taskID,
        bytes calldata taskResponse,
        bytes calldata blsSignature,
        address taskContractAddress,
        uint8  phase
    ) public returns (bool) {

        bool success  = avs.AVSMANAGER_CONTRACT.operatorSubmitTask(
            msg.sender,
            taskID,
            taskResponse,
            blsSignature,
            taskContractAddress,
            phase
        );
        return success;
    }



    function createNewTask(
        string memory name,
        uint64 numberToBeSquared,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint8 thresholdPercentage,
        uint64 taskStatisticalPeriod
    ) public returns (uint64) {
        // create a new task struct
        Task memory newTask;
        newTask.name = name;
        newTask.numberToBeSquared = numberToBeSquared;
        newTask.taskResponsePeriod = taskResponsePeriod;
        newTask.taskChallengePeriod = taskChallengePeriod;
        newTask.taskStatisticalPeriod = taskStatisticalPeriod;
        newTask.thresholdPercentage = thresholdPercentage;
        require(
            thresholdPercentage<=100,
            "The threshold cannot be greater than 100"
        );
        uint64 taskID = avs.AVSMANAGER_CONTRACT.createTask(
            msg.sender,
            name,
            abi.encodePacked(keccak256(abi.encode(newTask))),
            taskResponsePeriod,
            taskChallengePeriod,
            thresholdPercentage,
            taskStatisticalPeriod
        );
        newTask.taskId = taskID;
        emit TaskCreated(newTask.taskId, msg.sender,newTask.name,newTask.numberToBeSquared, newTask.taskResponsePeriod, newTask.taskChallengePeriod,newTask.thresholdPercentage,newTask.taskStatisticalPeriod);
        return taskID;
    }

    function raiseAndResolveChallenge(
        ChallengeReq memory req
    ) public returns (bool) {

       
    }


    //query
    function getOptInOperators(address avsAddress) public view returns (address[] memory) {
        return avs.AVSMANAGER_CONTRACT.getOptInOperators(
            avsAddress
        );

    }


    function getRegisteredPubkey(address operator,address avsAddr) public view returns (bytes memory) {


        return avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(
            operator,
            avsAddr
        );
    }

    function getAVSUSDValue(address avsAddr) external view returns (uint256){
        uint256  amount = avs.AVSMANAGER_CONTRACT.getAVSUSDValue(
            avsAddr
        );
        return amount;
    }
    function getOperatorOptedUSDValue(address avsAddr,address operatorAddr) external view returns (uint256){
        uint256  amount = avs.AVSMANAGER_CONTRACT.getOperatorOptedUSDValue(
            avsAddr,
            operatorAddr
        );
        return amount;
    }

    function getAVSEpochIdentifier(address avsAddr) external view returns (string memory){
        string memory epochIdentifier = avs.AVSMANAGER_CONTRACT.getAVSEpochIdentifier(
            avsAddr
        );
        return epochIdentifier;
    }
    function getTaskInfo(address taskAddress,uint64 taskID) external view returns (avs.TaskInfo memory){
        avs.TaskInfo memory info = avs.AVSMANAGER_CONTRACT.getTaskInfo(
            taskAddress,
            taskID
        );
        return info;
    }

    function isOperator(address operator) public view returns (bool) {

        bool flag = avs.AVSMANAGER_CONTRACT.isOperator(
            operator
        );
        return flag;
    }

    function getCurrentEpoch(string memory epochIdentifier) public view returns (int64) {

        int64 currentEpoch = avs.AVSMANAGER_CONTRACT.getCurrentEpoch(
            epochIdentifier
        );
        return currentEpoch;
    }


    function getChallengeInfo(address taskAddress, uint64 taskID)  public view returns (address challenger) {

        return avs.AVSMANAGER_CONTRACT.getChallengeInfo(
            taskAddress,
            taskID
        );
    }


    function getOperatorTaskResponse(address taskAddress, address operator, uint64 taskID) public view returns (avs.TaskResultInfo memory taskResultInfo) {

        return avs.AVSMANAGER_CONTRACT.getOperatorTaskResponse(
            taskAddress,
            operator,
            taskID
        );
    }

    function getOperatorTaskResponseList(address taskAddress, uint64 taskID) public view returns (avs.OperatorResInfo[] memory operatorResInfo) {

        return avs.AVSMANAGER_CONTRACT.getOperatorTaskResponseList(
            taskAddress,
            taskID
        );
    }
}