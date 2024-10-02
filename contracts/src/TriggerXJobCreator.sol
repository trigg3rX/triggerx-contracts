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
        uint32[] taskIds; 
        address jobCreator;
    }

    event JobCreated(uint256 indexed jobId);
    event JobDeleted(uint256 indexed jobId);
    event JobUpdated(uint256 indexed jobId);

    modifier onlyJobCreator(uint256 jobId) {
        require(jobs[jobId].jobCreator == msg.sender, "Only the job creator can call this function");
        _;
    }

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
            apiEndpoint: apiEndpoint,
            taskIds: new uint32[](0),
            jobCreator: msg.sender
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
    ) public onlyJobCreator(jobId) {
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
        job.taskIds = new uint32[](0);
        job.jobCreator = msg.sender;
        emit JobUpdated(jobId);
    }

    function addTaskId(uint256 jobId, uint32 taskId) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        jobs[jobId].taskIds.push(taskId);
    }

    // Function to delete a job
    function deleteJob(uint256 jobId) public onlyJobCreator(jobId) {
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
            job.apiEndpoint,
            job.taskIds,
            job.jobCreator
        );
    }

    function getJobsByCreator(address jobCreator) public view returns (Job[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= _job_id_counter; i++) {
            if (jobs[i].jobCreator == jobCreator) {
                count++;
            }
        }
        Job[] memory result = new Job[](count);
        uint256 index = 0;
        for (uint256 i = 1; i <= _job_id_counter; i++) {
            if (jobs[i].jobCreator == jobCreator) {
                result[index++] = jobs[i];
            }
        }
        return result;
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