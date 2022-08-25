include tools/tools.mk

.PHONY: clean
clean:
	@ go clean

.PHONY: deps
deps:
	@ go mod tidy --compat=1.19

.PHONY: test
test: $(GOTESTSUM)
	@ $(GOTESTSUM) -- -tags=test $(COVERPROFILE) -cover ./...

.PHONY: lint
lint: $(GOLANGCI_LINT)
	@ $(GOLANGCI_LINT) run --config=.golangci.yaml --fix

.PHONY: package
package: $(GORELEASER)
	@ $(GORELEASER) release --config=.goreleaser.yml --snapshot --skip-publish --rm-dist

.PHONY: compile
compile:
	@ go build -o bin/tis100 cmd/main.go

.PHONY: build
build: clean-tools lint test package

.PHONY: clean-tools
clean-tools:
	@-rm -rf $(TOOLS_BIN_DIR)
