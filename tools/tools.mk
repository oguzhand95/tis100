XDG_CACHE_HOME ?= $(HOME)/.cache
TOOLS_BIN_DIR := $(abspath $(XDG_CACHE_HOME)/oguzhand95/tis100/bin)
TOOLS_MOD := tools/go.mod

GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint
GORELEASER := $(TOOLS_BIN_DIR)/goreleaser
GOTESTSUM := $(TOOLS_BIN_DIR)/gotestsum

$(TOOLS_BIN_DIR):
	@ mkdir -p $(TOOLS_BIN_DIR)

$(GOLANGCI_LINT): $(TOOLS_BIN_DIR)
	@ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLS_BIN_DIR)

$(GORELEASER): $(TOOLS_BIN_DIR)
	@ GOBIN=$(TOOLS_BIN_DIR) go install -modfile=$(TOOLS_MOD) github.com/goreleaser/goreleaser

$(GOTESTSUM): $(TOOLS_BIN_DIR)
	@ GOBIN=$(TOOLS_BIN_DIR) go install -modfile=$(TOOLS_MOD) gotest.tools/gotestsum
