/*
    This contract is used to create, update, and delete jobs.
    TODO: Add modifiers, so that the one who created the job can only update or delete it.
    TODO: Fetch function, which shall get Keeper list working on the job
*/

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract TriggerXJobCreator {
    uint256 private _job_id_counter;
    mapping(uint256 => Job) public jobs;

    enum ArgType { None, Static, Dynamic }

    struct Job {
        uint256 jobId;  // Changed from uint32 to uint256
        string jobType;
        string status;
        uint32 timeframe;
        uint256 blockNumber;
        address contractAddress;
        string targetFunction; 
        uint256 timeInterval;
        ArgType argType;
        bytes[] arguments;
        string apiEndpoint;  
    }

    event JobCreated(uint256 indexed jobId);
    event JobDeleted(uint256 indexed jobId);
    event JobUpdated(uint256 indexed jobId);

    function createJob(
        string memory jobType,
        uint32 timeframe,
        address contractAddress,
        string memory targetFunction,
        uint256 timeInterval,
        ArgType argType,
        bytes[] memory arguments,
        string memory apiEndpoint 
    ) public returns (uint256) {
        _job_id_counter++;
        uint256 newJobId = _job_id_counter;

        jobs[newJobId] = Job({
            jobId: newJobId,
            jobType: jobType,
            status: "Created",
            timeframe: timeframe,
            blockNumber: block.number,
            contractAddress: contractAddress,
            targetFunction: targetFunction,
            timeInterval: timeInterval,
            argType: argType,
            arguments: arguments,
            apiEndpoint: apiEndpoint 
        });

        emit JobCreated(newJobId);
        return newJobId;
    }

    // Function to update a job
    function updateJob(
        uint256 jobId,
        string memory jobType,
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
        job.timeframe = timeframe;
        job.contractAddress = contractAddress;
        job.targetFunction = targetFunction;
        job.timeInterval = timeInterval;
        job.argType = argType;
        job.arguments = arguments;
        job.apiEndpoint = apiEndpoint; 

        emit JobUpdated(jobId);
    }

    // Function to delete a job
    function deleteJob(uint256 jobId) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        delete jobs[jobId];
        emit JobDeleted(jobId);
    }

    // Additional helper functions
    function getJob(uint256 jobId) public view returns (bytes memory) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        Job storage job = jobs[jobId];
        return abi.encode(
            job.jobId,
            job.jobType,
            job.status,
            job.timeframe,
            job.blockNumber,
            job.contractAddress,
            job.targetFunction,
            job.timeInterval,
            job.argType,
            job.arguments,
            job.apiEndpoint
        );
    }

    function getJobStatus(uint256 jobId) public view returns (string memory) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        return jobs[jobId].status;
    }

    function getJobArgumentCount(uint256 jobId) public view returns (uint256) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        return jobs[jobId].arguments.length;
    }

    function getJobArgument(uint256 jobId, uint256 argIndex) public view returns (bytes memory) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        require(argIndex < jobs[jobId].arguments.length, "Argument index out of bounds");
        return jobs[jobId].arguments[argIndex];
    }
}