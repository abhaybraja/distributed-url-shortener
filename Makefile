# Set project variables
APP_NAME = fiber-url-shortener
BUILD_DIR = build
CMD_DIR = cmd

# Docker variables
DOCKER_IMAGE = $(APP_NAME):latest
DOCKER_CONTAINER = $(APP_NAME)-container

# Go environment variables
GO_ENV = export GO111MODULE=on

# Default target
.PHONY: all
all: build

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Linting
.PHONY: lint
lint:
	@echo "Running static analysis..."
	@golangci-lint run

# Run application locally
.PHONY: run
run:
	@echo "Starting application..."
	@go run $(CMD_DIR)/main.go

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
.PHONY: coverage
coverage:
	@echo "Generating test coverage report..."
	@go test -cover ./...

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Dockerize application
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

# Run application in Docker
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	@docker run -d -p 3000:3000 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop and remove Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping Docker container..."
	@docker stop $(DOCKER_CONTAINER) || true
	@docker rm $(DOCKER_CONTAINER) || true

# Deploy using Docker Compose
.PHONY: deploy
deploy:
	@echo "Deploying with Docker Compose..."
	@docker-compose up --build -d

# Stop all services
.PHONY: stop
stop:
	@echo "Stopping all running services..."
	@docker-compose down

