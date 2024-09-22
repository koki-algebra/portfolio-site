.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Format code
	@go fmt ./...

.PHONY: lint
lint: ## Lint code
	@golangci-lint run --fix --allow-parallel-runners

.PHONY: generate
generate: ## Generate code
	@go generate ./...
