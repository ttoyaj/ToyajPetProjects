# Default binary name
BINARY_NAME=ping-pong

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Linter (using golangci-lint)
LINT_CMD=golangci-lint

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) main.go

# Run the application
run:
	$(GORUN) main.go

# Run tests
test:
	$(GOTEST) ./...

# Run linter
lint:
	$(LINT_CMD) run

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install dependencies (including linter)
setup:
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: run build test lint clean setup
