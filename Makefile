# Project Variables
BINARY_NAME = express-delivery
OUTPUT_DIR = output
SRC_DIR = ./cmd
MAIN_SRC = ./main.go

# Get app version
GET_VERSION_CMD = grep -m1 '^\# 版本-' ./CHANGE.md | sed 's/\# 版本-//'
VERSION := $(shell $(GET_VERSION_CMD))
COMMIT := $(shell git rev-parse HEAD)
PACKAGE := github.com/xmh1011/express-delivery/pkg/variable

# Commands
.PHONY: all build run clean test help

# Default command: build the binary
all: build

# Build the binary
build:
	@echo "Building the project..."
	mkdir -p $(OUTPUT_DIR)
	go build -ldflags "-X $(PACKAGE).Version=$(VERSION) -X $(PACKAGE).GitCommit=$(COMMIT)" -o $(OUTPUT_DIR)/$(BINARY_NAME) $(MAIN_SRC)

# Run the application
run: build
	@echo "Running the application..."
	./$(OUTPUT_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean up binary and temporary files
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT_DIR)
	go clean

# Help message
help:
	@echo "Makefile commands:"
	@echo "  make           - Build the project (default)"
	@echo "  make build     - Build the binary in the output directory"
	@echo "  make run       - Run the application from the output directory"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Remove the output directory and other generated files"
	@echo "  make help      - Show this help message"
