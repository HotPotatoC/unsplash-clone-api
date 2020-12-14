# APP Path
APP_PATH := cmd/api/main.go

# Output File
BIN_FILE := .bin/main
# Output file extension
BIN_EXTENSION :=

.PHONY: help build run

# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Runs the api server
	@go run cmd/api/main.go

build: ## Compiles the app
	@echo ">> Building go file..."
	@go build -ldflags="-s -w" -o $(BIN_FILE)$(BIN_EXTENSION) $(APP_PATH)

build-and-run: ## Compiles the app then runs it
	@$(MAKE) build
	@.bin/main


.DEFAULT_GOAL := help