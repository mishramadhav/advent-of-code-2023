# Variables
BINARY_NAME=day1
GO=go
GOFLAGS=-v
MAIN_FILES=*.go

# Phony targets
.PHONY: all build run clean test

# Default target
all: build

# Build the program
build: clean
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build $(GOFLAGS) -o $(BINARY_NAME) $(MAIN_FILES)

# Run the program
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

# Clean up
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@$(GO) test $(GOFLAGS) ./...
