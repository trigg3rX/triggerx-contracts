// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract TriggerXJobCreator {
    uint32 private _job_id_counter;
    mapping(uint32 => Job) public jobs;

    enum ArgType { None, Static, Dynamic }

    struct Job {
        uint256 jobId;
        string jobType;
        //string jobDescription;
        string status;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
        uint32 timeframe;
        uint256 blockNumber;
        address contractAddress;
        string targetFunction; 
        uint256 timeInterval;
        ArgType argType;
        bytes[] arguments;
        string apiEndpoint;  
    }

    event JobCreated(uint32 indexed jobId, string jobType, address contractAddress);
    event JobDeleted(uint32 indexed jobId);
    event JobUpdated(uint32 indexed jobId);

    // Function to create a job with dynamic arguments using API
    function createJob(
        string memory jobType,
        //string memory jobDescription,
        bytes memory quorumNumbers,
        uint32 quorumThresholdPercentage,
        uint32 timeframe,
        address contractAddress,
        string memory targetFunction,
        uint256 timeInterval,
        ArgType argType,
        bytes[] memory arguments,
        string memory apiEndpoint 
    ) public returns (uint32) {
        _job_id_counter++;
        uint32 newJobId = _job_id_counter;

        jobs[newJobId] = Job({
            jobId: newJobId,
            jobType: jobType,
            //jobDescription: jobDescription,
            status: "Created",
            quorumNumbers: quorumNumbers,
            quorumThresholdPercentage: quorumThresholdPercentage,
            timeframe: timeframe,
            blockNumber: block.number,
            contractAddress: contractAddress,
            targetFunction: targetFunction,
            timeInterval: timeInterval,
            argType: argType,
            arguments: arguments,
            apiEndpoint: apiEndpoint 
        });

        emit JobCreated(newJobId, jobType, contractAddress);
        return newJobId;
    }

    // Function to update a job
    function updateJob(
        uint32 jobId,
        string memory jobType,
        //string memory jobDescription,
        bytes memory quorumNumbers,
        uint32 quorumThresholdPercentage,
        uint32 timeframe,
        address contractAddress,
        string memory targetFunction,
        uint256 timeInterval,
        ArgType argType,
        bytes[] memory arguments,
        string memory apiEndpoint 
    ) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        Job storage job = jobs[jobId];

        job.jobType = jobType;
        //job.jobDescription = jobDescription;
        job.quorumNumbers = quorumNumbers;
        job.quorumThresholdPercentage = quorumThresholdPercentage;
        job.timeframe = timeframe;
        job.contractAddress = contractAddress;
        job.targetFunction = targetFunction;
        job.timeInterval = timeInterval;
        job.argType = argType;
        job.arguments = arguments;
        job.apiEndpoint = apiEndpoint; 

        emit JobUpdated(jobId);
    }

    // Additional helper functions remain the same
    function getJob(uint32 jobId) public view returns (Job memory) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        return jobs[jobId];
    }

    function getJobArgumentCount(uint32 jobId) public view returns (uint256) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        return jobs[jobId].arguments.length;
    }

    function getJobArgument(uint32 jobId, uint256 argIndex) public view returns (bytes memory) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        require(argIndex < jobs[jobId].arguments.length, "Argument index out of bounds");
        return jobs[jobId].arguments[argIndex];
    }
}