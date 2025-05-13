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

test-reset:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose down
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose up
	go test ./...

# Clean the build binary
clean: 
	rm -rf bin tmp

# Run dev server with auto-reload
dev:
	air -c .air.toml

# export DATABASE_URL="postgres://postgres:2b4gp44g6wr607931@localhost:5432/snapcrumb_test?sslmode=disable"
migrate:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose force 1 || true
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose up

