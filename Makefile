# Variables
BINARY_NAME=example
BUILD_DIR=build
MAIN_PATH=cmd/example/main.go

# Version information
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PKG_PATH=main
LDFLAGS=-ldflags "-X ${PKG_PATH}.Version=${VERSION} -X ${PKG_PATH}.Commit=${COMMIT} -X ${PKG_PATH}.BuildTime=${BUILD_TIME}"

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/$(BUILD_DIR)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## build: Build the binary
build:
	@echo "Building version ${VERSION}..."
	@echo "${LDFLAGS}"
	go build ${LDFLAGS} -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

## run: Run the application
run:
	go run $(MAIN_PATH)

## clean: Clean build files
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean

## test: Run tests
test:
	go test ./... -v -count=1

## test-coverage: Run tests with coverage
test-coverage:
	go test ./... -v -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out

## mock: Generate mocks for interfaces
mock:
	mockgen -source=internal/domain/user/application/service.go -destination=internal/domain/user/mocks/service_mock.go

## lint: Run linter
lint:
	golangci-lint run

## version: Show version info
version:
	@echo "Version:    ${VERSION}"
	@echo "Commit:     ${COMMIT}"
	@echo "Build Time: ${BUILD_TIME}"

## help: Display this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# Default to help if no target is specified
.DEFAULT_GOAL := help

.PHONY: build run clean test test-coverage mock lint help version
