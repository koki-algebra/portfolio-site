.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run Application
	@direnv allow .
	@docker compose up -d

.PHONY: fmt
fmt: ## Format
	@make -C ./backend fmt
	@make -C ./proto fmt

.PHONY: lint
lint: ## Lint code
	@make -C ./backend lint
	@make -C ./proto lint

.PHONY: test
test: ## Test code
	@make -C ./backend test

.PHONY: generate
generate: ## Generate code
	@make -C ./backend generate
	@make -C ./proto generate

.PHONY: clear
clear: ## Clear Application
	@docker compose down --volumes

.PHONY: logs
logs: ## Show API Server log
	@docker compose logs -f api_server

.PHONY: psql
psql: ## Login PostgreSQL
	@export PGPASSWORD=$(DB_PASSWORD); \
	psql --host localhost --port $(DB_PORT) --username $(DB_USER) --dbname $(DB_DATABASE)
