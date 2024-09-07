# Use bash as the shell
SHELL := /bin/bash

# .ONESHELL ensures that all commands in a target are executed in a single shell instance.
# This allows for environment variables and working directory changes to persist across commands within a target.
# It's particularly useful for multi-line commands or when using shell-specific features.
# Ensure all targets are executed in a single shell
.ONESHELL:

# Silent mode by default, use VERBOSE=1 to enable verbose output
$(VERBOSE).SILENT:

.DEFAULT_GOAL := run

.PHONY: tidy run test help

## tidy: Add missing deps, remove unused ones, and download necessary modules
tidy:
	go mod tidy

## fmt: Format all Go source files
fmt:
	go fmt ./...

## run: Run the application
run: tidy fmt
	go run hello.go

## test: Run tests for the application
test: fmt
	go test

## help: Display help information about available commands
help:
	echo "Usage:"
	sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# catch-all target: Route all unknown targets to help
%:
	@make help
