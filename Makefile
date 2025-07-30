############################# HELP MESSAGE #############################
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

############################# CONTRACTS #############################
----------------------------CONTRACTS----------------------------: ## 

build-contracts: ## Build contracts
	./contracts/utils/build-contracts.sh

############################# BINDINGS ##############################
-----------------------------BINDINGS----------------------------: ## 

generate-bindings: ## Generate bindings
	./bindings/generate-bindings.sh