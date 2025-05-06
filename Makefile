# Makefile for SnapCrumb

APP_NAME=snapcrumb

.PHONY: all build run test tidy clean dev

# Run go mod tidy
tidy:
	go mod tidy

# Build the Go binary
build:
	go build -o bin/$(APP_NAME) ./cmd/server

# Run the server (Requires DB to be running)
run:
	go run ./cmd/server

# Run tests
test:
	go test ./...

# Clean the build binary
clean: 
	rm -rf bin/

# Run dev server with auto-reload
dev:
	air -c .air.toml
	# air -c .air.toml
