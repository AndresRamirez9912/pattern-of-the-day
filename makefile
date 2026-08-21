# Build configuration for FlowCheck-backend
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
BUILD_DATE := $(shell date -u +'%y-%m-%dT%H:%M:%SZ')

# Set VERSION onlyt if not already defined
ifndef VERSION
	VERSION := $(shell git describe --tags 2>/dev/null)

# Use BRANCH-COMMIT if no tags found
	ifeq ($(VERSION),)
		VERSION := $(BRANCH)-$(COMMIT)
	endif 
endif

# Binary name and build target
BINARY_NAME := pattern${EXT}
BUILD_DIR := $(CURDIR)/build
FULL_BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)

###############################################################################
###                              Build                                      ###
###############################################################################
.PHONY: build clean

# Define our LDFlags
LD_FLAGS := -X github.com/AndresRamirez9912/pattern-of-the-day/types.VERSION=$(VERSION) \
			-X github.com/AndresRamirez9912/pattern-of-the-day/types.Commit=$(COMMIT) \
			-X github.com/AndresRamirez9912/pattern-of-the-day/types.BuildDate=$(BUILD_DIR)

# Build the binary
build: 
	@mkdir -p $(BUILD_DIR)
	@echo "building The patern of the day with version:$(VERSION)"
	@go build -a -ldflags "$(LD_FLAGS)" -o $(FULL_BINARY_PATH) ./cmd/pattern-of-the-day

# Clean up build folder
clean:
	rm -rf $(BUILD_DIR)/


###############################################################################
###                                Database                                 ###
###############################################################################

# Define the tools and their versions
SQLC             := $(BUILD_DIR)/sqlc${EXT}
SQLC_VERSION     := latest

# Install all the database tools
db-tools:
	@echo "Checking for required tools..."
	@mkdir -p $(BUILD_DIR)
	@if [ ! -x "$(SQLC)" ]; then \
	  echo "Installing sqlc..."; \
	  GOBIN=$(BUILD_DIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION); \
	fi

# Generate SQLC code
db-gen: db-tools
	@echo "Generating SQLC code..."
	# Run SQLC
	@$(SQLC) generate --file sqlc.yaml
