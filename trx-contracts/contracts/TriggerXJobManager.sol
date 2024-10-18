// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract TriggerXJobManager {
    uint32 private _job_id_counter;

    mapping(uint32 => Job) public jobs;

    mapping(address => uint32) public userJobsCount;

    mapping(address => uint32[]) public userJobs;

    mapping(address => uint256) public userTotalStake;

    enum ArgType {
        None,
        Static,
        Dynamic
    }

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

    event JobCreated(
        uint32 indexed jobId,
        address indexed creator,
        uint256 stakeAmount
    );
    event JobDeleted(
        uint32 indexed jobId,
        address indexed creator,
        uint256 stakeRefunded
    );
    event JobUpdated(uint32 indexed jobId);

    modifier onlyJobCreator(uint32 jobId) {
        require(
            jobs[jobId].jobCreator == msg.sender,
            "Only job creator can call this"
        );
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
        job.jobCreator = msg.sender;
        job.stakeAmount = stakeAmount;

        emit JobUpdated(jobId);
    }

    function deleteJob(
        uint32 jobId,
        uint256 stakeConsumed
    ) public onlyJobCreator(jobId) {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        require(
            keccak256(bytes(jobs[jobId].status)) != keccak256(bytes("Deleted")),
            "Job already deleted"
        );

        uint256 stakeToRefund = jobs[jobId].stakeAmount - stakeConsumed;

        jobs[jobId].status = "Deleted";

        userJobsCount[msg.sender]--;

        (bool sent, ) = msg.sender.call{value: stakeToRefund}("");
        require(sent, "Failed to refund stake");

        emit JobDeleted(jobId, msg.sender, stakeToRefund);
    }

    function addTaskId(uint32 jobId, uint32 taskId) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        jobs[jobId].taskIds.push(taskId);
    }

    function setJobStatus(uint32 jobId, string memory status) public {
        require(jobs[jobId].jobId != 0, "Job does not exist");
        require(
            keccak256(bytes(status)) == keccak256(bytes("Created")) ||
            keccak256(bytes(status)) == keccak256(bytes("Executing")) ||
            keccak256(bytes(status)) == keccak256(bytes("Finished")),
            "Invalid status"
        );

        jobs[jobId].status = status;
        emit JobUpdated(jobId);
    }

    function getJobArgs(uint32 jobId) public view returns (bytes[] memory) {
        return jobs[jobId].arguments;
    }

    function getTaskIds(uint32 jobId) public view returns (uint32[] memory) {
        return jobs[jobId].taskIds;
    }
}