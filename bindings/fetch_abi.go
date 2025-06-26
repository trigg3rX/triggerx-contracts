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
	Eth  map[string]string `json:"Eth"`
	Base map[string]string `json:"Base"`
}

const (
	OUTPUT_DIR = "bindings/abis"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	ethClient, err := ethclient.Dial(os.Getenv("ETH_RPC_URL"))
	if err != nil {
		panic(err)
	}

	baseClient, err := ethclient.Dial(os.Getenv("BASE_RPC_URL"))
	if err != nil {
		panic(err)
	}

	contractsDataRaw, err := os.ReadFile("./bindings/triggerx.prod.json")
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
		abi, err := fetchABIFromEtherscan(addr, "17000")
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
		fmt.Printf("Fetched ABI for %s on holesky\n", name)
	}

	for name, addr := range contractsData.Base {
		fetchBytecode(name, addr, baseClient, "base")
		abi, err := fetchABIFromEtherscan(addr, "84532")
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
		fmt.Printf("Fetched ABI for %s on holesky\n", name)
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
