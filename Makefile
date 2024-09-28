############################# HELP MESSAGE #############################
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

############################# ENVIRONMENT #############################

.PHONY: install
install:
	yarn install

.PHONY: init-submodules
init-submodules:
	git submodule update --init --recursive

.PHONY: setup-env
setup-env:
	/bin/bash -c 'source .env'

.PHONY: setup
setup: install init-submodules setup-env ## Setup the project

############################# CONTRACTS #############################

build-contracts: ## Build contracts
	./utils/build-contracts.sh

deploy-contracts-on-holesky: ## Deploy contracts on Holesky
	./utils/deploy-contracts-on-holesky.sh

deploy-contracts-on-shasta: ## Deploy contracts on Shasta
	./utils/deploy-contracts-on-shasta.sh

############################# KEEPER #############################

register-keeper-on-el-and-avs: ## Register keeper on EL and AVS
	./utils/register-keeper-on-el-and-avs.sh