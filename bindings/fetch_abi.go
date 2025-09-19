package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type ContractsData struct {
	Eth      map[string]string `json:"Eth"`
	Base     map[string]string `json:"Base"`
	Arbitrum map[string]string `json:"Arbitrum"`
}

const (
	OUTPUT_DIR = "abis"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(OUTPUT_DIR, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	ethClient, err := ethclient.Dial(fmt.Sprintf("https://eth-mainnet.g.alchemy.com/v2/%s", os.Getenv("ALCHEMY_API_KEY")))
	if err != nil {
		panic(err)
	}

	baseClient, err := ethclient.Dial(fmt.Sprintf("https://base-mainnet.g.alchemy.com/v2/%s", os.Getenv("ALCHEMY_API_KEY")))
	if err != nil {
		panic(err)
	}

	arbitrumClient, err := ethclient.Dial(fmt.Sprintf("https://arb-mainnet.g.alchemy.com/v2/%s", os.Getenv("ALCHEMY_API_KEY")))
	if err != nil {
		panic(err)
	}
	fmt.Println("Fetched clients")

	contractsDataRaw, err := os.ReadFile("./mainnet.json")
	if err != nil {
		panic(err)
	}

	var contractsData ContractsData
	if err := json.Unmarshal(contractsDataRaw, &contractsData); err != nil {
		panic(err)
	}

	fetchBytecode := func(name string, address string, client *ethclient.Client, network string) {
		binPath := filepath.Join(OUTPUT_DIR, name+"."+network+".bin")
		if _, err := os.Stat(binPath); err == nil {
			return
		}

		addr := common.HexToAddress(address)
		bytecode, err := client.CodeAt(context.Background(), addr, nil)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(
			binPath,
			[]byte(fmt.Sprintf("%x", bytecode)),
			0644,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched bytecode for %s on %s\n", name, network)
	}

	for name, addr := range contractsData.Eth {
		fetchBytecode(name, addr, ethClient, "eth")
		abi, err := fetchABIFromEtherscan(addr, "1")
		if err != nil {
			panic(err)
		}

		abiPath := filepath.Join(OUTPUT_DIR, name+".eth.abi")
		if _, err := os.Stat(abiPath); err == nil {
			continue
		}

		err = os.WriteFile(
			abiPath,
			[]byte(abi),
			0644,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched ABI for %s on Ethereum\n", name)
	}

	for name, addr := range contractsData.Base {
		fetchBytecode(name, addr, baseClient, "base")
		abi, err := fetchABIFromEtherscan(addr, "8453")
		if err != nil {
			panic(err)
		}

		abiPath := filepath.Join(OUTPUT_DIR, name+".base.abi")
		if _, err := os.Stat(abiPath); err == nil {
			continue
		}

		err = os.WriteFile(
			abiPath,
			[]byte(abi),
			0644,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched ABI for %s on Base\n", name)
	}

	for name, addr := range contractsData.Arbitrum {
		fetchBytecode(name, addr, arbitrumClient, "arbitrum")
		abi, err := fetchABIFromEtherscan(addr, "42161")
		if err != nil {
			panic(err)
		}

		abiPath := filepath.Join(OUTPUT_DIR, name+".arbitrum.abi")
		if _, err := os.Stat(abiPath); err == nil {
			continue
		}

		err = os.WriteFile(
			abiPath,
			[]byte(abi),
			0644,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched ABI for %s on Arbitrum\n", name)
	}
}

func fetchABIFromEtherscan(address string, network string) (string, error) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("ETHERSCAN_API_KEY not set in environment")
	}

	baseURL := "https://api.etherscan.io/v2/api"

	url := fmt.Sprintf("%s?chainid=%s&module=contract&action=getabi&address=%s&apikey=%s", baseURL, network, address, apiKey)

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
