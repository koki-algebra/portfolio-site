.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Format
	@npm run fmt

.PHONY: lint
lint: ## Lint
	@npm run lint

.PHONY: generate
generate: ## Generate code
	@npx buf generate --config ../proto/buf.yaml ../proto
