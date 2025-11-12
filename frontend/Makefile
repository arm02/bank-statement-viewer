# Variables
NODE_VERSION := 18
PACKAGE_MANAGER := npm

# Install dependencies
install:
	$(PACKAGE_MANAGER) install

# Development server
dev: install
	$(PACKAGE_MANAGER) run dev

# Build for production
build: install
	$(PACKAGE_MANAGER) run build

# Run tests
test:
	$(PACKAGE_MANAGER) run test

# Run tests with coverage
test-coverage:
	$(PACKAGE_MANAGER) run test:coverage

# Clean build artifacts and dependencies
clean:
	rm -rf dist
	rm -rf node_modules
	

# Check Node.js version
check-node:
	@node --version | grep -E "v$(NODE_VERSION)" > /dev/null || (echo "Node.js $(NODE_VERSION) is required" && exit 1)

# Setup development environment
setup: check-node install
	cp .env.example .env || touch .env
	@echo "Vue.js environment setup complete!"
	@echo "Run 'make dev' to start the development server"

# Health check - test if server is running
health:
	@curl -s http://localhost:3000 || echo "Vue frontend server not running on port 3000"