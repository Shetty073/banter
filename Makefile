# Go environment variables
GO      := go
GOINST  := $(GO) install
GOTIDY  := $(GO) mod tidy
GOBUILD := $(GO) build
GORUN   := $(GO) run
GOTEST  := $(GO) test

# Dev tools
SWAGGO  := github.com/swaggo/swag/cmd/swag@latest
AIR     := github.com/air-verse/air@latest

# Application settings
APP_NAME := banter
MAIN_FILE := main.go
OUTPUT_DIR := bin
OUTPUT_BIN := $(OUTPUT_DIR)/$(APP_NAME)

# Install dependencies
install:
	$(GO) mod download

# Install development tools
install-dev:
	$(GOINST) $(SWAGGO)
	$(GOINST) $(AIR)

# Generate Swagger documentation
swagger:
	swag init --generalInfo $(MAIN_FILE) --output docs

# Run the application
run:
	$(GORUN) $(MAIN_FILE)

# Run the application with Air live reloading
watch:
	air

# Build the application
build:
	mkdir -p $(OUTPUT_DIR)
	$(GOBUILD) -o $(OUTPUT_BIN) $(MAIN_FILE)

# Run tests
test:
	$(GOTEST) ./...

# Clean up generated files
clean:
	rm -rf $(OUTPUT_DIR)
	rm -rf docs/swagger*

# Tidy Go modules (remove unused dependencies)
tidy:
	$(GOTIDY)

# Default command
.DEFAULT_GOAL := run
