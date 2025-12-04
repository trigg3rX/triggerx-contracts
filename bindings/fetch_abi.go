package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// ContractsData represents the structure of the contracts.json file
type ContractsData map[string]map[string]interface{}

const (
	OUTPUT_DIR = "abis"
)

// Network configuration for different chain IDs
var networkConfigs = map[string]struct {
	networkName string
	chainID     string
}{
	"11155111": {"sepolia", "11155111"},   // Ethereum Sepolia
	"84532":    {"base-sepolia", "84532"}, // Base Sepolia
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(OUTPUT_DIR, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	// Read the testnet deployment file
	contractsDataRaw, err := os.ReadFile("./contracts.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to read contracts.json: %v", err))
	}

	var contractsData ContractsData
	if err := json.Unmarshal(contractsDataRaw, &contractsData); err != nil {
		panic(fmt.Sprintf("Failed to parse contracts.json: %v", err))
	}

	// Process each network
	for chainID, contracts := range contractsData {
		config, exists := networkConfigs[chainID]
		if !exists {
			fmt.Printf("Skipping unknown chain ID: %s\n", chainID)
			continue
		}

		fmt.Printf("\nProcessing network: %s (chain ID: %s)\n", config.networkName, config.chainID)
		fmt.Println(strings.Repeat("-", 50))

		// Process each contract in the network
		for contractName, contractData := range contracts {
			var implementationAddress string

			// Handle contract data - in contracts.json, addresses are stored as strings
			switch v := contractData.(type) {
			case string:
				// Direct address string
				implementationAddress = v
			default:
				fmt.Printf("Skipping %s (unknown format)\n", contractName)
				continue
			}

			// Skip if address is the zero address or empty
			if implementationAddress == "" || implementationAddress == "0x0000000000000000000000000000000000000000" {
				fmt.Printf("Skipping %s (zero or empty address)\n", contractName)
				continue
			}

			// Fetch ABI
			fmt.Printf("Fetching ABI for %s (address: %s)...\n", contractName, implementationAddress)

			abiPath := filepath.Join(OUTPUT_DIR, fmt.Sprintf("%s.abi", contractName))

			// Skip if already exists
			if _, err := os.Stat(abiPath); err == nil {
				fmt.Printf("ABI already exists for %s\n", contractName)
				continue
			}

			abi, err := fetchABIFromEtherscan(implementationAddress, config.chainID)
			if err != nil {
				fmt.Printf("Failed to fetch ABI for %s: %v\n", contractName, err)
				continue
			}

			// Write ABI to file
			if err := os.WriteFile(abiPath, []byte(abi), 0644); err != nil {
				fmt.Printf("Failed to write ABI for %s: %v\n", contractName, err)
				continue
			}

			fmt.Printf("Successfully fetched ABI for %s\n", contractName)
		}
	}

	fmt.Println("\nABI fetching complete!")
}

func fetchABIFromEtherscan(address string, chainID string) (string, error) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	baseURL := "https://api.etherscan.io/v2/api"

	url := fmt.Sprintf("%s?chainid=%s&module=contract&action=getabi&address=%s&apikey=%s", baseURL, chainID, address, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch ABI: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %v", err)
	}

	if result.Status != "1" {
		return "", fmt.Errorf("etherscan API error: %s", result.Message)
	}

	return result.Result, nil
}
