#!/usr/bin/env node

/**
 * Script to fetch all users who have interacted with the old TriggerGasRegistry contract
 * 
 * Usage:
 *   node script/fetchUsers.js
 * 
 * Environment variables required:
 *   - ETHERSCAN_API_KEY: API key for Etherscan (works for all chains)
 *   - OLD_REGISTRY_ADDRESS: Address of the old registry contract (default: 0x664CB20BCEEc9416D290AC820e5446e61a5c75e4)
 */

const axios = require("axios");
const fs = require("fs");
const path = require("path");
const dotenv = require("dotenv");
const { ethers } = require("ethers");
dotenv.config();

// Old registry contract address
// NOTE: Make sure this contract is deployed on the chains you're querying!
// The default address may only exist on mainnet chains, not testnets
const OLD_REGISTRY_ADDRESS = process.env.OLD_REGISTRY_ADDRESS || "0x93dDB2307F3Af5df85F361E5Cddd898Acd3d132d";

// Chain configurations for Etherscan API v2
// API v2 uses unified endpoint: https://api.etherscan.io/v2/api
// Specify chain using chainid parameter
const ETHERSCAN_API_V2_URL = "https://api.etherscan.io/v2/api";

const CHAINS = {
  // base: {
  //   name: "Base Sepolia",
  //   chainId: 84532,
  //   apiKey: process.env.ETHERSCAN_API_KEY,
  // },
  // optimism: {
  //   name: "Optimism Sepolia",
  //   chainId: 11155420,
  //   apiKey: process.env.ETHERSCAN_API_KEY,
  // },
  arbitrum: {
    name: "Arbitrum One",
    chainId: 42161,
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  // ethereum: {
  //   name: "Ethereum Sepolia",
  //   chainId: 11155111,
  //   apiKey: process.env.ETHERSCAN_API_KEY,
  // },
  
};

// Event signature for TriggerGasRegistry
// Only fetching TGPurchased events which have 'user' as indexed parameter
const EVENT_SIGNATURES = [
  "TGPurchased(address,uint256,uint256)",
  "ETHDeposited(address,uint256)",
];

/**
 * Sleep function for rate limiting
 */
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

/**
 * Fetch logs from Etherscan API v2
 */
async function fetchLogsFromEtherscan(chainConfig, topic0, fromBlock = 0, toBlock = "latest") {
  const params = {
    chainid: chainConfig.chainId, // Required: specify chain using chainid parameter
    module: "logs",
    action: "getLogs",
    address: OLD_REGISTRY_ADDRESS,
    // topic0: topic0, // COMMENTED OUT to fetch ALL events
    fromBlock: fromBlock,
    toBlock: toBlock,
    apikey: chainConfig.apiKey,
  };

  // If topic0 is provided, add it
  if (topic0) {
    params.topic0 = topic0;
  }

  try {
    // Debug: Log the request URL (without API key)
    const debugParams = { ...params, apikey: '***' };
    console.log(`   Request: ${ETHERSCAN_API_V2_URL}?${new URLSearchParams(debugParams).toString()}`);

    const response = await axios.get(ETHERSCAN_API_V2_URL, { params });

    // Debug: log the full response for debugging
    // console.log(`   API Response Status: ${response.data.status}`);

    if (response.data.status === "1" && response.data.result) {
      const logs = Array.isArray(response.data.result) ? response.data.result : [];
      // console.log(`   ✅ API returned ${logs.length} logs`);
      return logs;
    } else if (response.data.message === "No records found" || response.data.result === "No logs found" || response.data.result === "No records found") {
      // console.log(`   ℹ️  No records found`);
      return [];
    } else {
      // API returned an error
      const errorMsg = response.data.message || response.data.result || "Unknown error";
      console.error(`   ❌ API Error: ${errorMsg}`);
      return [];
    }
  } catch (error) {
    console.error(`   HTTP Error:`, error.message);
    console.error(`   Code:`, error.code);
    if (error.response) {
      console.error(`   Status:`, error.response.status);
      console.error(`   Data:`, JSON.stringify(error.response.data));
    } else {
      console.error(`   Stack:`, error.stack);
    }
    return [];
  }
}

/**
 * Fetch all unique users from events on a specific chain using Etherscan API v2
 */
async function fetchUsersFromChain(chainConfig) {
  if (!chainConfig.apiKey) {
    console.warn(`⚠️  Skipping ${chainConfig.name}: ETHERSCAN_API_KEY not set`);
    return new Set();
  }

  console.log(`\n📡 Fetching users from ${chainConfig.name} (chainId: ${chainConfig.chainId}) via Etherscan API v2...`);

  try {
    const users = new Set();

    // DEBUG: First fetch ALL events to see what's happening
    console.log(`   🔎 Debugging: Fetching ALL events for contract ${OLD_REGISTRY_ADDRESS}...`);
    const allLogs = await fetchLogsFromEtherscan(chainConfig, null); // Pass null for topic0
    console.log(`   found ${allLogs.length} total events on this chain`);

    if (allLogs.length > 0) {
      console.log(`   Events found on chain:`);
      const uniqueTopics = new Set();
      allLogs.forEach(log => {
        if (log.topics && log.topics.length > 0) {
          uniqueTopics.add(log.topics[0]);
        }
      });

      uniqueTopics.forEach(topic => {
        let eventName = "Unknown";
        if (topic === "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b") eventName = "Upgraded(address)";
        if (topic === "0x7f26b83ff96e1f2b6a682f133852f6798a09c0040d2842aa36ee80b797d74b86") eventName = "OwnershipTransferred(address,address)";
        if (topic === "0x4cc662b651de5212ae4a670ed3fab666cbe3f524d278d775c6e42685132bcf24") eventName = "TGPurchased";
        if (topic === "0x6e411cb5d9291a2cae763a3f0c9822e358795cba820d85ff477ab658c071c6b0") eventName = "ETHDeposited";

        console.log(`   - ${topic} (${eventName})`);
      });
    }

    // Fetch events for each event signature
    for (const eventSig of EVENT_SIGNATURES) {
      try {
        const eventName = eventSig.split("(")[0];
        // Generate topic0 (event signature hash) - Etherscan expects it with 0x prefix
        const topic0 = ethers.id(eventSig);

        // Use client-side filtering on allLogs instead of making new requests
        const logs = allLogs.filter(log => log.topics[0] && log.topics[0].toLowerCase() === topic0.toLowerCase());
        console.log(`   Found ${logs.length} logs for ${eventName}`);

        // Extract user addresses from logs
        // The first indexed parameter (user) is in topics[1]
        for (const log of logs) {
          if (log.topics && Array.isArray(log.topics) && log.topics.length > 1) {
            // topics[1] is the indexed user address (padded to 32 bytes)
            let topicHex = log.topics[1];
            if (!topicHex.startsWith("0x")) {
              topicHex = "0x" + topicHex;
            }
            // Extract last 40 characters (20 bytes = address length)
            // Remove 0x prefix, take last 40 chars, then add 0x back
            const addressHex = topicHex.replace("0x", "").slice(-40);
            const userAddress = ethers.getAddress("0x" + addressHex);
            users.add(userAddress);
          }
        }
      } catch (error) {
        console.error(`   Error fetching ${eventSig.split("(")[0]}:`, error.message);
      }
    }

    console.log(`   ✅ Found ${users.size} unique users on ${chainConfig.name}`);
    return users;
  } catch (error) {
    console.error(`   ❌ Error fetching from ${chainConfig.name}:`, error.message);
    return new Set();
  }
}

/**
 * Main function
 */
async function main() {
  console.log("🚀 Starting user fetch from TriggerGasRegistry events via Etherscan API v2");
  console.log(`📋 Old Registry Address: ${OLD_REGISTRY_ADDRESS}`);
  console.log(`📋 API Endpoint: ${ETHERSCAN_API_V2_URL}\n`);

  if (!process.env.ETHERSCAN_API_KEY) {
    console.error("❌ ETHERSCAN_API_KEY environment variable is required");
    console.error("   Get your API key from: https://etherscan.io/myapikey");
    process.exit(1);
  }

  const results = {};

  // Fetch users from each chain
  for (const [chainKey, chainConfig] of Object.entries(CHAINS)) {
    const users = await fetchUsersFromChain(chainConfig);
    results[chainKey] = {
      chainId: chainConfig.chainId,
      users: Array.from(users).sort(), // Sort for consistency
    };
  }

  // Write results to JSON file
  const outputPath = path.join(__dirname, "../data/users.json");
  const outputData = {
    base: results.base || { chainId: 84532, users: [] },
    optimism: results.optimism || { chainId: 11155420, users: [] },
    arbitrum: results.arbitrum || { chainId: 421614, users: [] },
    ethereum: results.ethereum || { chainId: 11155111, users: [] },
  };

  // Ensure data directory exists
  const dataDir = path.dirname(outputPath);
  if (!fs.existsSync(dataDir)) {
    fs.mkdirSync(dataDir, { recursive: true });
  }

  fs.writeFileSync(outputPath, JSON.stringify(outputData, null, 2));

  console.log("\n✅ Results saved to data/users.json");
  console.log("\n📊 Summary:");
  for (const [chainKey, data] of Object.entries(outputData)) {
    console.log(`   ${CHAINS[chainKey].name}: ${data.users.length} users`);
  }
}

// Run the script
main().catch((error) => {
  console.error("❌ Fatal error:", error);
  process.exit(1);
});
