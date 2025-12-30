# Tideland Go Asserts - Makefile
#
# Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
#
# All rights reserved. Use of this source code is governed
# by the new BSD license.

# Variables
GO := go
GOLANGCI_LINT := golangci-lint
GOLANGCI_LINT_VERSION := v2.7.2
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Package selection
PACKAGE ?= all
PACKAGES := verify capture generators

# Validate package selection
ifeq ($(PACKAGE),all)
	TARGET_PACKAGES := $(PACKAGES)
else
	ifneq ($(filter $(PACKAGE),$(PACKAGES)),)
		TARGET_PACKAGES := $(PACKAGE)
	else
		$(error Invalid package '$(PACKAGE)'. Valid options: all, verify, capture, generators)
	endif
endif

# Colors for output
COLOR_RESET := \033[0m
COLOR_BOLD := \033[1m
COLOR_GREEN := \033[32m
COLOR_YELLOW := \033[33m
COLOR_BLUE := \033[34m
COLOR_CYAN := \033[36m

# Default target
.DEFAULT_GOAL := all

# Phony targets
.PHONY: all help fmt tidy lint build test bench coverage clean install-tools check-tools ci

## all: Run complete build process for selected package(s) (fmt, tidy, lint, build, test)
all:
	@echo "$(COLOR_CYAN)$(COLOR_BOLD)Running all tasks for: $(TARGET_PACKAGES)$(COLOR_RESET)"
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory tidy
	@$(MAKE) --no-print-directory lint
	@$(MAKE) --no-print-directory build
	@$(MAKE) --no-print-directory test
	@echo "$(COLOR_GREEN)$(COLOR_BOLD)✓ All tasks completed successfully for: $(TARGET_PACKAGES)$(COLOR_RESET)"

## help: Display this help message
help:
	@echo "$(COLOR_BOLD)Tideland Go Asserts - Available Targets:$(COLOR_RESET)"
	@echo ""
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
	@echo ""
	@echo "$(COLOR_BLUE)Usage:$(COLOR_RESET)"
	@echo "  make [target]                    # Run for all packages"
	@echo "  make [target] PACKAGE=verify     # Run for verify package only"
	@echo "  make [target] PACKAGE=capture    # Run for capture package only"
	@echo "  make [target] PACKAGE=generators # Run for generators package only"
	@echo ""
	@echo "$(COLOR_BLUE)Available packages:$(COLOR_RESET) $(PACKAGES)"
	@echo ""

## fmt: Format Go source files in selected package(s)
fmt:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Formatting $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && gofmt -s -w .) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Code formatting completed$(COLOR_RESET)"

## tidy: Update go.mod and go.sum files for selected package(s)
tidy:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Tidying $$pkg modules...$(COLOR_RESET)"; \
		(cd $$pkg && $(GO) mod tidy && $(GO) mod verify) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Module dependencies updated$(COLOR_RESET)"

## lint: Run golangci-lint on selected package(s)
lint:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Linting $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && $(GOLANGCI_LINT) run --timeout=5m) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Linting completed$(COLOR_RESET)"

## build: Build selected package(s) (verify compilation)
build:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Building $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && $(GO) build -v ./...) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Build successful$(COLOR_RESET)"

## test: Run tests for selected package(s)
test:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Testing $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && $(GO) test -v -race ./...) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Tests passed$(COLOR_RESET)"

## bench: Run benchmarks for selected package(s)
bench:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Benchmarking $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && $(GO) test -bench=. -benchmem -run=^$$ ./...) || exit 1; \
	done
	@echo "$(COLOR_GREEN)✓ Benchmarks completed$(COLOR_RESET)"

## coverage: Generate test coverage reports for selected package(s)
coverage:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Generating coverage for $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && \
			$(GO) test -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./... && \
			$(GO) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML) && \
			$(GO) tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print "Coverage: " $$3}') || exit 1; \
		echo "$(COLOR_GREEN)✓ Coverage report: $$pkg/$(COVERAGE_HTML)$(COLOR_RESET)"; \
	done
	@echo "$(COLOR_GREEN)✓ Coverage reports generated$(COLOR_RESET)"

## clean: Remove build artifacts and coverage files from selected package(s)
clean:
	@for pkg in $(TARGET_PACKAGES); do \
		echo "$(COLOR_YELLOW)→ Cleaning $$pkg...$(COLOR_RESET)"; \
		(cd $$pkg && rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)) || exit 1; \
	done
	@echo "$(COLOR_YELLOW)→ Cleaning Go caches...$(COLOR_RESET)"
	@$(GO) clean -cache -testcache -modcache
	@echo "$(COLOR_GREEN)✓ Clean completed$(COLOR_RESET)"

## install-tools: Install required development tools
install-tools:
	@echo "$(COLOR_YELLOW)→ Installing development tools...$(COLOR_RESET)"
	@which $(GOLANGCI_LINT) > /dev/null 2>&1 || \
		(echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..." && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION))
	@$(GOLANGCI_LINT) version | grep -q "has version $(shell echo $(GOLANGCI_LINT_VERSION) | sed 's/v//')" || \
		(echo "Updating golangci-lint to $(GOLANGCI_LINT_VERSION)..." && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION))
	@echo "$(COLOR_GREEN)✓ Tools installed$(COLOR_RESET)"

## check-tools: Check installed tool versions and compatibility
check-tools:
	@echo "$(COLOR_YELLOW)→ Checking tool versions...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)Go version:$(COLOR_RESET)"
	@$(GO) version
	@echo "$(COLOR_BLUE)golangci-lint version:$(COLOR_RESET)"
	@$(GOLANGCI_LINT) version || echo "$(COLOR_YELLOW)golangci-lint not found - run 'make install-tools'$(COLOR_RESET)"
	@echo "$(COLOR_GREEN)✓ Tool version check completed$(COLOR_RESET)"

## ci: Run CI pipeline for selected package(s) (used by GitHub Actions)
ci:
	@echo "$(COLOR_CYAN)$(COLOR_BOLD)Running CI pipeline for: $(TARGET_PACKAGES)$(COLOR_RESET)"
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory tidy
	@$(MAKE) --no-print-directory lint
	@$(MAKE) --no-print-directory build
	@$(MAKE) --no-print-directory test
	@echo "$(COLOR_GREEN)$(COLOR_BOLD)✓ CI pipeline completed for: $(TARGET_PACKAGES)$(COLOR_RESET)"

## list-packages: List all available packages
list-packages:
	@echo "$(COLOR_BLUE)Available packages:$(COLOR_RESET)"
	@for pkg in $(PACKAGES); do \
		echo "  - $$pkg"; \
	done
