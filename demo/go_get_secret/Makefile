GOCMD=go
PROJECT_NAME=devoxx
SRC_DIR=.
BUILD_DIR=bin

all: build

build:
	@echo "Building the project..."
	@$(GOCMD) build -o $(BUILD_DIR)/$(PROJECT_NAME) $(SRC_DIR)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)/$(PROJECT_NAME)

test:
	@echo "Running tests..."
	@$(GOCMD) test $(SRC_DIR)/...

run:
	@echo "Running the project..."
	@$(BUILD_DIR)/$(PROJECT_NAME)

.PHONY: all build clean test run
