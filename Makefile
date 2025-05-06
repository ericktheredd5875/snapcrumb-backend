# Makefile for SnapCrumb

APP_NAME=snapcrumb

# Detect OS and set binary appropriately
ifeq ($(OS),Windows_NT)
	BINARY_EXT=.exe
else
	BINARY_EXT=
endif

BINARY=bin/$(APP_NAME)$(BINARY_EXT)

.PHONY: all build run test tidy clean dev

# Run go mod tidy
tidy:
	go mod tidy

# Build the Go binary
build:
	go build -o $(BINARY) ./cmd/server

# Run the server (Requires DB to be running)
run:
	go run ./cmd/server

# Run tests
test:
	go test ./...

# Clean the build binary
clean: 
	rm -rf bin tmp

# Run dev server with auto-reload
dev:
	air -c .air.toml
