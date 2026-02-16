.PHONY: help build run-publisher run-subscriber test fmt vet clean docker-up docker-down

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

build: ## Build publisher and subscriber binaries
	@echo "Building binaries..."
	@go build -o bin/publisher cmd/publisher/main.go
	@go build -o bin/subscriber cmd/subscriber/main.go
	@echo "✓ Build complete"

run-publisher: ## Run the publisher
	@go run cmd/publisher/main.go

run-subscriber: ## Run the subscriber
	@go run cmd/subscriber/main.go

test: ## Run tests
	@go test -v ./...

fmt: ## Format code
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
