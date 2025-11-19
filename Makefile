# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go
	
	

build-arm:
	@echo "Cross compiling for arm64..."
	@GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o main-arm cmd/api/main.go


# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

sqlc:
	@echo "building sql queries..."
	@sqlc  generate -f internal/database/sqlc.yaml

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch

addMany:
	@echo "building addMany migrations"
	@go build -C cmd/migrations/addMany
	@mv cmd/migrations/addMany/addMany . 
	
