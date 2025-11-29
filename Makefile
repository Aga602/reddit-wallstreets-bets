.PHONY: all build test lint clean run help

# Binary name
BINARY_NAME=greeter

# Build directory
BUILD_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOVET=$(GOCMD) vet
GOFMT=gofmt

# Default target
all: lint test build

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all      Run lint, test, and build"
	@echo "  build    Build the application"
	@echo "  test     Run tests"
	@echo "  lint     Run linters"
	@echo "  clean    Clean build artifacts"
	@echo "  run      Run the application"
	@echo "  help     Show this help message"

## build: Build the application
build:
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/greeter

## test: Run tests
test:
	$(GOTEST) -v -race ./...

## lint: Run linters
lint:
	$(GOVET) ./...
	@if [ -n "$$($(GOFMT) -l .)" ]; then \
		echo "The following files are not formatted correctly:"; \
		$(GOFMT) -l .; \
		exit 1; \
	fi

## clean: Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

## run: Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)
