# Makefile

ENV ?= local
BINARY_NAME = iam-system-$(ENV)
MAIN_PATH = cmd/main/main.go

# Command to load the appropriate .env file and run the application with air
run:
	@echo "Running in $(ENV) environment..."
	@if [ -f .env.$(ENV) ]; then \
		export GIN_ENV=$(ENV); \
		echo "Loaded .env.$(ENV)"; \
		cp .env.$(ENV) .env; \
	else \
		echo "Environment file .env.$(ENV) not found, using default .env"; \
	fi
	air

# Command to build the application with the appropriate environment
build:
	@echo "Building in $(ENV) environment..."
	@if [ -f .env.$(ENV) ]; then \
		export GIN_ENV=$(ENV); \
		echo "Loaded .env.$(ENV)"; \
		cp .env.$(ENV) .env; \
	else \
		echo "Environment file .env.$(ENV) not found, using default .env"; \
	fi
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Specific commands for each environment
run-dev:
	@make run ENV=development

run-local:
	@make run ENV=local

run-prod:
	@make run ENV=production

build-dev:
	@make build ENV=development

build-prod:
	@make build ENV=production

