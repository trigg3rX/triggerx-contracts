package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "context"
    "path/filepath"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/joho/godotenv"
)

type ContractAddresses struct {
    Proxy          string `json:"proxy"`
    Implementation string `json:"implementation"`
}

type TriggerXDeploymentConfig struct {
    RegistryCoordinator   ContractAddresses `json:"registryCoordinator"`
    OperatorStateRetriever string            `json:"operatorStateRetriever"`
    TriggerXServiceManager ContractAddresses `json:"triggerXServiceManager"`
    TriggerXTaskManager    ContractAddresses `json:"triggerXTaskManager"`
}

type EigenlayerDeploymentConfig struct {
    DelegationManager ContractAddresses `json:"delegationManager"`
    AvsDirectory ContractAddresses `json:"avsDirectory"`
}

type StakeDeploymentConfig struct {
    TriggerXStakeRegistry ContractAddresses `json:"triggerXStakeRegistry"`
}

const (
    OUTPUT_DIR = "bindings/abis"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }

	holeskyClient, err := ethclient.Dial(os.Getenv("HOLESKY_RPC_URL"))
    if err != nil {
        panic(err)
    }

    opSepoliaClient, err := ethclient.Dial(os.Getenv("OP_SEPOLIA_RPC_URL"))
    if err != nil {
        panic(err)
    }

    triggerxHoleskyData, err := ioutil.ReadFile("./bindings/deployment.holesky.json")
    if err != nil {
        panic(err)
    }

	eigenlayerHoleskyData, err := ioutil.ReadFile("./bindings/eigenlayer.holesky.json")
	if err != nil {
		panic(err)
	}

	stakeOPSData, err := ioutil.ReadFile("./bindings/stake.opsepolia.json")
	if err != nil {
		panic(err)
	}

    var txDeployments TriggerXDeploymentConfig
    if err := json.Unmarshal(triggerxHoleskyData, &txDeployments); err != nil {
        panic(err)
    }

	var eigenlayerDeployments EigenlayerDeploymentConfig
	if err := json.Unmarshal(eigenlayerHoleskyData, &eigenlayerDeployments); err != nil {
		panic(err)
	}

    var stakeDeployments StakeDeploymentConfig
    if err := json.Unmarshal(stakeOPSData, &stakeDeployments); err != nil {
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

        err = ioutil.WriteFile(
            binPath,
            []byte(fmt.Sprintf("%x", bytecode)),
            0644,
        )
        if err != nil {
            panic(err)
        }
        fmt.Printf("Fetched bytecode for %s on %s\n", name, network)
    }

    triggerxContracts := map[string]string{
        "RegistryCoordinator":    txDeployments.RegistryCoordinator.Implementation,
        "OperatorStateRetriever": txDeployments.OperatorStateRetriever,
        "TriggerXServiceManager": txDeployments.TriggerXServiceManager.Implementation,
        "TriggerXTaskManager":    txDeployments.TriggerXTaskManager.Implementation,
    }

	eigenlayerContracts := map[string]string{
		"DelegationManager": eigenlayerDeployments.DelegationManager.Implementation,
		"AvsDirectory": eigenlayerDeployments.AvsDirectory.Implementation,
	}

    opSepoliaContracts := map[string]string{
        "TriggerXStakeRegistry": stakeDeployments.TriggerXStakeRegistry.Implementation,
    }

    for name, addr := range triggerxContracts {
        fetchBytecode(name, addr, holeskyClient, "holesky")
        abi, err := fetchABIFromBlockscout(addr, "holesky")
        if err != nil {
            panic(err)
        }

        abiPath := filepath.Join(OUTPUT_DIR, name+".holesky.abi")
        if _, err := os.Stat(abiPath); err == nil {
            continue
        }

        err = ioutil.WriteFile(
            abiPath,
            []byte(abi),
            0644,
        )
        if err != nil {
            panic(err)
        }
        fmt.Printf("Fetched ABI for %s on holesky\n", name)
    }

	for name, addr := range eigenlayerContracts {
		fetchBytecode(name, addr, holeskyClient, "holesky")
		abi, err := fetchABIFromBlockscout(addr, "holesky")
		if err != nil {
			panic(err)
		}

		abiPath := filepath.Join(OUTPUT_DIR, name+".holesky.abi")
		if _, err := os.Stat(abiPath); err == nil {
			continue
		}

		err = ioutil.WriteFile(
			abiPath,
			[]byte(abi),
			0644,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched ABI for %s on holesky\n", name)
	}

    for name, addr := range opSepoliaContracts {
        fetchBytecode(name, addr, opSepoliaClient, "opsepolia")
        abi, err := fetchABIFromBlockscout(addr, "opsepolia")
        if err != nil {
            panic(err)
        }

        abiPath := filepath.Join(OUTPUT_DIR, name+".opsepolia.abi")
        if _, err := os.Stat(abiPath); err == nil {
            continue
        }

        err = ioutil.WriteFile(
            abiPath,
            []byte(abi),
            0644,
        )
        if err != nil {
            panic(err)
        }
        fmt.Printf("Fetched ABI for %s on opsepolia\n", name)
    }
}

func fetchABIFromBlockscout(address string, network string) (string, error) {
    var baseURL string
    switch network {
    case "holesky":
        baseURL = "https://eth-holesky.blockscout.com/api"
    case "opsepolia":
        baseURL = "https://optimism-sepolia.blockscout.com/api"
    default:
        return "", fmt.Errorf("unsupported network: %s", network)
    }

    url := fmt.Sprintf("%s?module=contract&action=getabi&address=%s", baseURL, address)
    
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("failed to fetch ABI: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
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
        return "", fmt.Errorf("blockscout API error: %s ||| %s" , result.Message, address)
    }

    return result.Result, nil
}