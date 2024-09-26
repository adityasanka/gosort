# Use bash as the shell
SHELL := /bin/bash

APP=hello

# .ONESHELL ensures that all commands in a target are executed in a single shell instance.
# This allows for environment variables and working directory changes to persist across commands within a target.
# It's particularly useful for multi-line commands or when using shell-specific features.
# Ensure all targets are executed in a single shell
.ONESHELL:

# Silent mode by default, use VERBOSE=1 to enable verbose output
$(VERBOSE).SILENT:

.DEFAULT_GOAL :=run

.PHONY: install fmt lint run test bench help

## install: Install project dependencies
install:
	## Add missing deps, remove unused ones, and download necessary modules
	go mod tidy

## fmt: Format all Go source files
fmt:
	go fmt ./...

## lint: Lint all Go source files
lint:
	golangci-lint run

## test: Run tests for the application
test: install
	go test ./...

## bench: Run benchmarks for the application
bench: install
	go test -bench=. -benchmem ./...

## clean: Remove stale binaries
clean:
	rm -f ${APP}
	go clean

## build: Build the cmd line application
build: clean install
	go build -o ${APP} ./cmd/hello

## run: Run the cmd line application
run: build
	chmod +x ${APP}
	./${APP}
	@make clean

## help: Display help information about available commands
help:
	echo "Usage:"
	sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# catch-all target: Route all unknown targets to help
%:
	@make help
