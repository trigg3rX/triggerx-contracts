package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ContractABI represents the structure for storing ABI data
type ContractABI struct {
	ContractName string      `json:"contractName"`
	ABI          interface{} `json:"abi"`
	Address      string      `json:"address"`
	Network      string      `json:"network"`
}

// EtherscanResponse represents the Etherscan API response
type EtherscanResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// NetworkConfig holds network-specific configuration
type NetworkConfig struct {
	Name     string
	ChainID  int
	APIKey   string
	BaseURL  string
	Explorer string
}

// Supported networks
var networks = map[string]NetworkConfig{
	"mainnet": {
		Name:     "mainnet",
		ChainID:  1,
		BaseURL:  "https://api.etherscan.io/api",
		Explorer: "https://etherscan.io",
	},
	"sepolia": {
		Name:     "sepolia",
		ChainID:  11155111,
		BaseURL:  "https://api-sepolia.etherscan.io/api",
		Explorer: "https://sepolia.etherscan.io",
	},
	"holesky": {
		Name:     "holesky",
		ChainID:  17000,
		BaseURL:  "https://api-holesky.etherscan.io/api",
		Explorer: "https://holesky.etherscan.io",
	},
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run fetch-abi.go <network> <contract_address> <contract_name> [api_key]")
		fmt.Println("Example: go run fetch-abi.go mainnet 0x1234...5678 MyContract YOUR_API_KEY")
		fmt.Println("Supported networks: mainnet, sepolia, holesky")
		os.Exit(1)
	}

	networkName := strings.ToLower(os.Args[1])
	contractAddress := os.Args[2]
	contractName := os.Args[3]
	
	var apiKey string
	if len(os.Args) > 4 {
		apiKey = os.Args[4]
	} else {
		apiKey = os.Getenv("ETHERSCAN_API_KEY")
	}

	network, exists := networks[networkName]
	if !exists {
		fmt.Printf("❌ Unsupported network: %s\n", networkName)
		fmt.Println("Supported networks: mainnet, sepolia, holesky")
		os.Exit(1)
	}

	if apiKey == "" {
		fmt.Println("⚠️  No API key provided. Rate limiting may apply.")
	}

	fmt.Printf("🔍 Fetching ABI for %s on %s...\n", contractName, network.Name)

	abi, err := fetchABI(network, contractAddress, apiKey)
	if err != nil {
		fmt.Printf("❌ Failed to fetch ABI: %v\n", err)
		os.Exit(1)
	}

	// Create abis directory if it doesn't exist
	if err := os.MkdirAll("abis", 0755); err != nil {
		fmt.Printf("❌ Failed to create abis directory: %v\n", err)
		os.Exit(1)
	}

	// Save ABI to file
	filename := filepath.Join("abis", contractName+".abi")
	if err := saveABI(abi, filename); err != nil {
		fmt.Printf("❌ Failed to save ABI: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ ABI saved to %s\n", filename)
	fmt.Printf("🔗 Explorer: %s/address/%s\n", network.Explorer, contractAddress)
}

func fetchABI(network NetworkConfig, address, apiKey string) (interface{}, error) {
	url := fmt.Sprintf("%s?module=contract&action=getabi&address=%s", network.BaseURL, address)
	if apiKey != "" {
		url += fmt.Sprintf("&apikey=%s", apiKey)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var etherscanResp EtherscanResponse
	if err := json.Unmarshal(body, &etherscanResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if etherscanResp.Status != "1" {
		return nil, fmt.Errorf("API error: %s", etherscanResp.Message)
	}

	var abi interface{}
	if err := json.Unmarshal([]byte(etherscanResp.Result), &abi); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	return abi, nil
}

func saveABI(abi interface{}, filename string) error {
	abiBytes, err := json.MarshalIndent(abi, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal ABI: %w", err)
	}

	return os.WriteFile(filename, abiBytes, 0644)
} 