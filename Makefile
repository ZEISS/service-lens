.DEFAULT_GOAL := build

BASE_DIR		?= $(CURDIR)
OUTPUT_DIR      ?= $(BASE_DIR)/dist

GO 				?= go
GO_RUN_TOOLS	?= $(GO) run -modfile ./tools/go.mod
GO_TEST 		?= $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname
GO_RELEASER 	?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_KO 			?= $(GO_RUN_TOOLS) github.com/google/ko
GO_MOD 			?= $(shell ${GO} list -m)

COMMANDS		:= $(notdir $(wildcard cmd/*))

IMAGE_TAG       ?= $(shell git rev-parse HEAD)
TAG_REGEX       := ^v([0-9]{1,}\.){2}[0-9]{1,}$
KOFLAGS         ?=

.PHONY: build
build: $(COMMANDS) ## Build the application.

$(filter-out $(CUSTOM_BUILD_BINARIES), $(COMMANDS)): ## Build artifact
	$(GO) build -ldflags "$(LDFLAGS_STATIC)" -o $(BIN_OUTPUT_DIR)/$@ ./cmd/$@

.PHONY: generate
generate: ## Generate code.
	$(GO) generate ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO_RUN_TOOLS) mvdan.cc/gofumpt -w .

.PHONY: vet
vet: ## Run go vet against code.
	$(GO) vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	mkdir -p .test/reports
	$(GO_TEST) --junitfile .test/reports/unit-test.xml -- -race ./... -count=1 -short -cover -coverprofile .test/reports/unit-test-coverage.out

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	rm -rf .test .dist
	find . -type f -name '*.gen.go' -exec rm {} +
	git checkout go.mod

.PHONY: release
release: ## Release the application.
	echo "Release $(IMAGE_TAG)"

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'