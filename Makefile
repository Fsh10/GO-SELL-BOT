LOCAL_BIN=$(CURDIR)/bin
APP_NAME=bot
DOCKER_IMAGE=go-sell-bot
DOCKER_TAG=latest

.PHONY: help
help: ## Show help for commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build binary file
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(LOCAL_BIN)
	@go build -ldflags="-s -w" -o $(LOCAL_BIN)/$(APP_NAME) $(CURDIR)/cmd
	@echo "Build complete: $(LOCAL_BIN)/$(APP_NAME)"

.PHONY: run
run: build ## Build and run the bot
	@$(LOCAL_BIN)/$(APP_NAME) -token=$(BOT_TOKEN) -seller-id=$(SELLER_ID)

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -race -count=1 -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -race -count=1 -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: lint
lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

.PHONY: lint-fix
lint-fix: ## Fix linter issues automatically
	@golangci-lint run --fix

.PHONY: fmt
fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w .

.PHONY: vet
vet: ## Check code with go vet
	@echo "Running go vet..."
	@go vet ./...

.PHONY: mod
mod: ## Update dependencies
	@echo "Updating dependencies..."
	@go mod download
	@go mod tidy
	@go mod verify

.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(LOCAL_BIN)
	@rm -f coverage.out coverage.html
	@go clean -cache -testcache

.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: ## Run Docker container
	@docker-compose up -d

.PHONY: docker-stop
docker-stop: ## Stop Docker container
	@docker-compose down

.PHONY: docker-logs
docker-logs: ## Show Docker container logs
	@docker-compose logs -f

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "Tools installed successfully"

.PHONY: all
all: clean mod fmt vet lint test build ## Run all checks and build