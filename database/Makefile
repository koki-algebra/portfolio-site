.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Format
	@sqlfluff fix

.PHONY: lint
lint: ## Lint
	@sqlfluff lint

.PHONY: new
new: ## Add new migration file. e.g. $ make new NAME=add_user_table
	@goose -dir migrations -s create $(NAME) sql

.PHONY: status
status: ## Show migration status
	@goose -dir migrations status

.PHONY: up
up: ## Migration up
	@goose -dir migrations up
