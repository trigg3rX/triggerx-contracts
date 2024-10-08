############################# HELP MESSAGE #############################
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

############################# ENVIRONMENT #############################

.PHONY: setup
setup:  ## Setup the project
	./utils/setup.sh

############################# CONTRACTS #############################

build-contracts: ## Build contracts
	./utils/build-contracts.sh

deploy-contracts-on-holesky: ## Deploy contracts on Holesky
	./utils/deploy-contracts-on-holesky.sh

deploy-contracts-on-nile: ## Deploy contracts on Nile
	./utils/deploy-contracts-on-nile.sh