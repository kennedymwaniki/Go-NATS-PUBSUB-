.PHONY: help build run-publisher run-subscriber start-nats stop-nats clean test

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build publisher and subscriber binaries
	@echo "Building publisher..."
	@go build -o bin/publisher cmd/publisher/main.go
	@echo "Building subscriber..."
	@go build -o bin/subscriber cmd/subscriber/main.go
	@echo "Build complete!"

run-publisher: ## Run the publisher
	@go run cmd/publisher/main.go

run-subscriber: ## Run the subscriber
	@go run cmd/subscriber/main.go

test: ## Run tests
	@go test -v ./...

fmt: ## Format code
start-nats: ## Start NATS server using Docker Compose
	@docker-compose up -d
	@echo "NATS server started. Monitoring available at http://localhost:8222"

stop-nats: ## Stop NATS server
	@docker-compose down

clean: ## Clean build artifacts
	@rm -rf bin/
	@echo "Cleaned build artifacts"

test: ## Run tests
	@go test -v ./...

fmt: ## Format Go code
	@go fmt ./...

vet: ## Run go vet
	@go vet ./...

clean: ## Clean build artifacts
	@rm -rf bin/
	@echo "✓ Cleaned build artifacts"

docker-up: ## Start NATS server using Docker Compose
	@docker-compose up -d
	@echo "✓ NATS server started"

docker-down: ## Stop NATS server
	@docker-compose down
	@echo "✓ NATS server stopped"

deps: ## Download dependencies
	@go mod download
	@echo "✓ Dependencies downloaded"

all: fmt vet build ## Format, vet, and build
deps: ## Download dependencies
	@go mod download
	@go mod tidy

all: deps build ## Download dependencies and build all binaries
