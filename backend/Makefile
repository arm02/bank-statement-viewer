BINARY_NAME=flip-test-assignment
MAIN_PATH=./cmd/main.go
BUILD_DIR=./bin

# Default target
all: clean fmt lint test build

# Install dependencies
install:
	go mod download
	go mod tidy

# Run the application
run:
	go run $(MAIN_PATH)

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Run tests
test:
	TMPDIR=/tmp go test -v -race -coverprofile=coverage.out ./...

# Show test coverage
coverage: test
	TMPDIR=/tmp go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Health check
health:
	@curl -s http://localhost:8080/health || echo "Server not running"

# Stop background server
stop:
	pkill -f "go run $(MAIN_PATH)" || true