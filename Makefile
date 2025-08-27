APP_NAME=orders
CMD_DIR=cmd

.PHONY: all build run test lint clean docker-build docker-run

# Default target
all: build

# Build the Go application
build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) $(CMD_DIR)/main.go

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	go run $(CMD_DIR)/main.go

# Run tests
test:
	@echo "Running tests..."
	go test ./... -v

# Lint using golangci-lint (make sure it's installed)
lint:
	@echo "Linting code..."
	golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 $(APP_NAME):latest




# With this Makefile, you can just run commands like:
# make build → compiles your app into bin/orders
# make run → runs directly with go run
# make test → runs all tests
# make lint → lints code
# make docker-build → builds docker image
# make docker-run → runs docker container