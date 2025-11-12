# Development commands
install: 
	@echo "$(GREEN)Installing backend dependencies...$(RESET)"
	cd backend && go mod download && go mod tidy
	@echo "$(GREEN)Installing frontend dependencies...$(RESET)"
	cd frontend && npm install

dev: ## Start both backend and frontend in development mode
	@echo "$(GREEN)Starting development servers...$(RESET)"
	@echo "$(YELLOW)Backend will start on http://localhost:8080$(RESET)"
	@echo "$(YELLOW)Frontend will start on http://localhost:3000$(RESET)"
	@echo "$(YELLOW)Press Ctrl+C to stop all servers$(RESET)"
	@echo ""
	@trap 'kill %1 %2 2>/dev/null; exit' INT; \
	cd backend && make run & \
	cd frontend && make dev & \
	wait

# Environment setup
setup:
	@echo "$(GREEN)Setting up development environment...$(RESET)"
	@echo "Checking prerequisites..."
	@command -v go >/dev/null 2>&1 || { echo "$(RED)Go is required but not installed$(RESET)"; exit 1; }
	@command -v node >/dev/null 2>&1 || { echo "$(RED)Node.js is required but not installed$(RESET)"; exit 1; }
	@command -v npm >/dev/null 2>&1 || { echo "$(RED)npm is required but not installed$(RESET)"; exit 1; }
	@echo "$(GREEN)Prerequisites check passed!$(RESET)"
	@echo "Installing dependencies..."
	$(MAKE) install
	@echo "$(GREEN)Environment setup complete!$(RESET)"
	@echo ""
	@echo "$(YELLOW)Next steps:$(RESET)"
	@echo "  1. Run 'make dev' to start development servers"
	@echo "  2. Open http://localhost:3000 in your browser"

# Individual service commands
backend: 
	@echo "$(GREEN)Starting backend server...$(RESET)"
	cd backend && make run

frontend:
	@echo "$(GREEN)Starting frontend server...$(RESET)"
	cd frontend && make dev

# Build commands
build: 
	@echo "$(GREEN)Building backend...$(RESET)"
	cd backend && make build
	@echo "$(GREEN)Building frontend...$(RESET)"
	cd frontend && make build

build-backend:
	cd backend && make build

build-frontend:
	cd frontend && make build


# Test commands
test: 
	@echo "$(GREEN)Running backend tests...$(RESET)"
	cd backend && make test
	@echo "$(GREEN)Running frontend tests...$(RESET)"
	cd frontend && make test

test-backend:
	cd backend && make test

test-frontend:
	cd frontend && make test

test-coverage:
	@echo "$(GREEN)Generating backend coverage...$(RESET)"
	cd backend && make coverage
	@echo "$(GREEN)Generating frontend coverage...$(RESET)"
	cd frontend && make test-coverage

# Health and status checks
health:
	@echo "$(GREEN)Checking service health...$(RESET)"
	@echo -n "Backend: "
	@curl -s http://localhost:8080/health > /dev/null && echo "$(GREEN)✓ Running$(RESET)" || echo "$(RED)✗ Not running$(RESET)"
	@echo -n "Frontend: "
	@curl -s http://localhost:3000 > /dev/null && echo "$(GREEN)✓ Running$(RESET)" || echo "$(RED)✗ Not running$(RESET)"

status: health
