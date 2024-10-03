// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract TriggerXJobManager {
    uint32 private _job_id_counter;

    mapping(uint32 => Job) public jobs;

    mapping(address => uint32) public userJobsCount;
    mapping(address => uint32[]) public userJobs;
    mapping(address => uint256) public userTotalStake;

    enum ArgType { None, Static, Dynamic }

    struct Job {
        uint32 jobId;
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
        uint256 stakeAmount;
    }
    
    event JobCreated(uint32 indexed jobId, address indexed creator, uint256 stakeAmount);
    event JobDeleted(uint32 indexed jobId, address indexed creator, uint256 stakeRefunded);
    event JobUpdated(uint32 indexed jobId);
    
    modifier onlyJobCreator(uint32 jobId) {
        require(jobs[jobId].jobCreator == msg.sender, "Only job creator can call this");
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
    ) public payable returns (uint32) {
        require(msg.value > 0, "Must stake some TRX to create a job");
        
        _job_id_counter++;
        uint32 newJobId = _job_id_counter;
        
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
            jobCreator: msg.sender,
            stakeAmount: msg.value
        });
        
        userJobsCount[msg.sender]++;

        userJobs[msg.sender].push(newJobId);

        userTotalStake[msg.sender] += msg.value;
        
        emit JobCreated(newJobId, msg.sender, msg.value);
        return newJobId;
    }

    // Function to update a job
    function updateJob(
        uint32 jobId,
        string memory jobType,
        uint32 timeframe,
        address contractAddress,
        string memory targetFunction,
        uint256 timeInterval,
        ArgType argType,
        bytes[] memory arguments,
        string memory apiEndpoint,
        uint256 stakeAmount
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
        job.stakeAmount = stakeAmount;

        emit JobUpdated(jobId);
    }
    
    function deleteJob(uint32 jobId, uint256 stakeConsumed) public onlyJobCreator(jobId) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        
        uint256 stakeToRefund = jobs[jobId].stakeAmount - stakeConsumed;
        
        // Remove job from userJobs array
        uint32[] storage creatorJobs = userJobs[msg.sender];
        for (uint i = 0; i < creatorJobs.length; i++) {
            if (creatorJobs[i] == jobId) {
                creatorJobs[i] = creatorJobs[creatorJobs.length - 1];
                creatorJobs.pop();
                break;
            }
        }
        
        delete jobs[jobId];
        userJobsCount[msg.sender]--;
        
        // Refund stake
        (bool sent, ) = msg.sender.call{value: stakeToRefund}("");
        require(sent, "Failed to refund stake");
        
        emit JobDeleted(jobId, msg.sender, stakeToRefund);
    }

    function addTaskId(uint32 jobId, uint32 taskId) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        jobs[jobId].taskIds.push(taskId);
    }

    function getUserJobs(address user) public view returns (Job[] memory) {
        uint32[] memory jobIds = userJobs[user];
        Job[] memory userJobsList = new Job[](jobIds.length);
        
        for (uint i = 0; i < jobIds.length; i++) {
            userJobsList[i] = jobs[jobIds[i]];
        }
        
        return userJobsList;
    }
    
    function getTotalUserStake(address user) public view returns (uint256) {
        uint32[] memory jobIds = userJobs[user];
        uint256 totalStake = 0;
        
        for (uint i = 0; i < jobIds.length; i++) {
            totalStake += jobs[jobIds[i]].stakeAmount;
        }
        
        return totalStake;
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