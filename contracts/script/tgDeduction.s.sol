// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Script} from "forge-std/Script.sol";
import {TriggerGasRegistry} from "../src/TriggerGasRegistry.sol";
import {JobRegistry} from "../src/JobRegistry.sol";
import {TaskExecutionHub} from "../src/lz/TaskExecutionHub.sol";
import {CREATE3} from "lib/solady/src/utils/CREATE3.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {console2} from "forge-std/console2.sol";

contract TGDeduction is Script {
    TriggerGasRegistry public triggerGasRegistry;
    JobRegistry public jobRegistry;
    TaskExecutionHub public taskExecutionHub;
    
    address public deployer;
    address public jobOwner;
    address public keeper;
    
    uint32 public constant SRC_EID = 40259; // L1 chain ID
    uint32 public constant THIS_EID = 40245; // This L2 chain ID

    function run() public {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        deployer = vm.addr(privateKey);
        
        // Create additional addresses for testing
        jobOwner = deployer; // Same as deployer for this test
        keeper = deployer;   // Same as deployer for this test

        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(privateKey);

        console2.log("=== Starting TG Deduction Test ===");
        console2.log("Deployer:");
        console2.logAddress(deployer);
        console2.log("Job Owner:");
        console2.logAddress(jobOwner);
        console2.log("Keeper:");
        console2.logAddress(keeper);

        // Deploy contracts
        deployTriggerGasRegistry();
        deployJobRegistry();
        deployTaskExecutionHub();

        // Set TaskExecutionHub as operator of TriggerGasRegistry now that it's deployed
        triggerGasRegistry.setOperator(address(taskExecutionHub));

        // Test the complete flow
        completeFlowTest();

        vm.stopBroadcast();
    }

    function deployTriggerGasRegistry() public {
        console2.log("\n--- Deploying TriggerGasRegistry ---");

        bytes32 SALT = keccak256(abi.encodePacked("TriggerGasRegistrysfsdf"));
        bytes32 implementation_salt = keccak256(abi.encodePacked("triggerGasRegistrysdsrs"));

        bytes memory implementation_code = type(TriggerGasRegistry).creationCode;
        address implementation = CREATE3.deployDeterministic(implementation_code, implementation_salt);

        bytes memory initData = abi.encodeWithSelector(TriggerGasRegistry.initialize.selector, deployer, address(0), 1000);

        bytes memory proxy_code = abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(address(implementation), initData));

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        triggerGasRegistry = TriggerGasRegistry(proxy);
        
        console2.log("TriggerGasRegistry implementation deployed to:", implementation);
        console2.log("TriggerGasRegistry proxy deployed to:", address(triggerGasRegistry));
    }

    function deployJobRegistry() public {
        console2.log("\n--- Deploying JobRegistry ---");
        
        bytes32 SALT = keccak256(abi.encodePacked("JobRegistrysfsdf"));
        bytes32 implementation_salt = keccak256(abi.encodePacked("jobRegistrysfsdf"));

        bytes memory implementation_code = type(JobRegistry).creationCode;
        address implementation = CREATE3.deployDeterministic(implementation_code, implementation_salt);

        bytes memory initData = abi.encodeWithSelector(JobRegistry.initialize.selector, deployer);

        bytes memory proxy_code = abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(address(implementation), initData));

        address proxy = CREATE3.deployDeterministic(proxy_code, SALT);

        jobRegistry = JobRegistry(proxy);
        
        console2.log("JobRegistry implementation deployed to:", implementation);
        console2.log("JobRegistry proxy deployed to:", address(jobRegistry));
    }

    function deployTaskExecutionHub() public {
        console2.log("\n--- Deploying TaskExecutionHub ---");
        
        // Setup initial keepers array
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = keeper;
        
        // Deploy TaskExecutionHub
        taskExecutionHub = new TaskExecutionHub(
            address(0x6EDCE65403992e310A62460808c4b910D972f10f),
            deployer
        );
        
        // Initialize the contract
        taskExecutionHub.initialize(
            deployer,
            SRC_EID,
            THIS_EID,
            initialKeepers,
            address(jobRegistry),
            address(triggerGasRegistry)
        );
        
        console2.log("TaskExecutionHub deployed to:", address(taskExecutionHub));
        
        // Fund the TaskExecutionHub for message fees
        vm.deal(address(taskExecutionHub), 10 ether);
    }

    function completeFlowTest() public {
        console2.log("\n=== Testing Complete Flow ===");
        
        // Step 1: Purchase TG
        purchaseTG();
        
        // Step 2: Create a job
        uint256 jobId = createJob();
        
        // Step 3: Verify keeper registration
        keeperRegistration();
        
        // Step 4: Execute function and verify TG deduction
        executeFunction(jobId);
        
        // Step 5: Test different jobTypes with proper data
        differentJobTypes();
        
        console2.log("\n=== Test Completed Successfully ===");
    }

    function purchaseTG() public {
        console2.log("\n--- Step 1: Purchasing TG ---");
        
        uint256 ethAmount = 1 ether;
        uint256 initialBalance = address(jobOwner).balance;
        
        console2.log("Initial ETH balance: ", vm.toString(initialBalance));
        console2.log("Purchasing TG with ", vm.toString(ethAmount), " ETH");
        
        // Get initial TG balance
        (, uint256 initialTGBalance) = triggerGasRegistry.getBalance(jobOwner);
        console2.log("Initial TG balance: ", vm.toString(initialTGBalance));
        
        // Purchase TG
        triggerGasRegistry.purchaseTG{value: ethAmount}(ethAmount);
        
        // Get final TG balance
        (, uint256 finalTGBalance) = triggerGasRegistry.getBalance(jobOwner);
        console2.log("Final TG balance: ", vm.toString(finalTGBalance));
        console2.log("TG purchased: ", vm.toString(finalTGBalance - initialTGBalance));
        
        // Verify the purchase
        require(finalTGBalance > initialTGBalance, "TG balance should increase");
        console2.log("TG purchase successful");
    }

    function createJob() public returns (uint256 jobId) {
        console2.log("\n--- Step 2: Creating Job ---");
        
        string memory jobName = "Test Job";
        uint8 jobType = 1;
        uint256 timeFrame = 3600; // 1 hour
        address targetContract = address(0x123); // Mock target contract
        bytes memory data = encodeJobData(jobType, 300, bytes32(0)); // 5 minute interval for jobType 1
        
        console2.log("Creating job with name:");
        console2.logString(jobName);
        console2.log("Job owner:");
        console2.logAddress(jobOwner);
        
        // Create the job
        jobId = jobRegistry.createJob(jobName, jobType, timeFrame, targetContract, data);
        
        console2.log("Job created with ID: ", vm.toString(jobId));
        
        // Verify the job was created correctly
        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        require(job.jobOwner == jobOwner, "Job owner should match");
        require(job.isActive, "Job should be active");
        
        console2.log("Job creation successful");
        return jobId;
    }

    /**
     * @dev Helper function to encode job data based on jobType requirements
     * @param jobType The type of the job
     * @param timeInterval The time interval (for jobType 1 or 2)
     * @param ipfsHash The IPFS hash (for jobType 2, 4, or 6)
     * @return data The encoded data
     */
    function encodeJobData(uint8 jobType, uint256 timeInterval, bytes32 ipfsHash) internal pure returns (bytes memory data) {
        if (jobType == 1 || jobType == 2) {
            // For jobType 1 or 2, include timeInterval
            if (jobType == 2 || jobType == 4 || jobType == 6) {
                // For jobType 2, 4, or 6, also include ipfsHash
                data = abi.encode(timeInterval, ipfsHash);
            } else {
                // For jobType 1, only timeInterval
                data = abi.encode(timeInterval);
            }
        } else {
            // For other jobTypes, include recurringJob (bool)
            bool recurringJob = true; // Default to true for testing
            if (jobType == 4 || jobType == 6) {
                // For jobType 4 or 6, also include ipfsHash
                data = abi.encode(recurringJob, ipfsHash);
            } else {
                // For other jobTypes, only recurringJob
                data = abi.encode(recurringJob);
            }
        }
    }

    function keeperRegistration() public view {
        console2.log("\n--- Step 3: Verifying Keeper Registration ---");
        
        bool isRegistered = taskExecutionHub.isKeeper(keeper);
        console2.log("Keeper registration status: ", vm.toString(isRegistered));
        
        require(isRegistered, "Keeper should be registered");
        console2.log("Keeper registration verified");
    }

    function executeFunction(uint256 jobId) public {
        console2.log("\n--- Step 4: Testing Function Execution and TG Deduction ---");
        
        uint256 tgAmountToDeduct = 100; // Amount of TG to deduct
        address targetContract = address(0x123); // Mock target contract
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Get initial TG balance
        (, uint256 initialTGBalance) = triggerGasRegistry.getBalance(jobOwner);
        console2.log("Initial TG balance before execution: ", vm.toString(initialTGBalance));
        console2.log("TG amount to deduct: ", vm.toString(tgAmountToDeduct));
        
        // Create a simple mock contract that will always succeed
        bytes memory mockBytecode = hex"600180600c6000396000f3006000fd"; // Simple bytecode that returns true
        vm.etch(targetContract, mockBytecode);
        
        console2.log("Executing function through TaskExecutionHub...");
        
        // Execute the function
        taskExecutionHub.executeFunction(jobId, tgAmountToDeduct, targetContract, data);
        
        // Get final TG balance
        (, uint256 finalTGBalance) = triggerGasRegistry.getBalance(jobOwner);
        console2.log("Final TG balance after execution: ", vm.toString(finalTGBalance));
        console2.log("TG deducted: ", vm.toString(initialTGBalance - finalTGBalance));
        
        // Verify the deduction
        require(finalTGBalance == initialTGBalance - tgAmountToDeduct, "TG should be deducted correctly");
        console2.log("Function execution and TG deduction successful");
    }

    function differentJobTypes() public {
        console2.log("\n--- Step 5: Testing Different Job Types ---");
        
        // Test jobType 2 (requires timeInterval and ipfsHash)
        console2.log("Testing jobType 2...");
        bytes32 testIpfsHash = keccak256(abi.encodePacked("test-ipfs-hash"));
        bytes memory data2 = encodeJobData(2, 600, testIpfsHash); // 10 minute interval
        uint256 jobId2 = jobRegistry.createJob("Test Job Type 2", 2, 7200, address(0x456), data2);
        console2.log("JobType 2 created with ID: ", vm.toString(jobId2));
        
        // Test jobType 3 (requires recurringJob only)
        console2.log("Testing jobType 3...");
        bytes memory data3 = encodeJobData(3, 0, bytes32(0));
        uint256 jobId3 = jobRegistry.createJob("Test Job Type 3", 3, 1800, address(0x789), data3);
        console2.log("JobType 3 created with ID: ", vm.toString(jobId3));
        
        // Test jobType 4 (requires recurringJob and ipfsHash)
        console2.log("Testing jobType 4...");
        bytes memory data4 = encodeJobData(4, 0, testIpfsHash);
        uint256 jobId4 = jobRegistry.createJob("Test Job Type 4", 4, 900, address(0xabc), data4);
        console2.log("JobType 4 created with ID: ", vm.toString(jobId4));
        
        console2.log("All job types tested successfully");
    }
}