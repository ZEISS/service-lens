.DEFAULT_GOAL := build

GO 				?= go
GO_RUN_TOOLS 	?= $(GO) run -modfile ./tools/go.mod
GO_TEST 		?= $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname
GO_RELEASER 	?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_MOD 			?= $(shell ${GO} list -m)

# Module name
MODULE_NAME ?= github.com/katallaxie/template-go

.PHONY: build
build: ## Build the binary file.
	$(GO_RELEASER) build --snapshot --clean

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

.PHONY: snapshot
snapshot: ## Create a snapshot release
	$(GO_RELEASER) release --clean --snapshot

.PHONY: release
release: ## Create a release
	$(GO_RELEASER) release --clean

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	rm -rf .test .dist
	find . -type f -name '*.gen.go' -exec rm {} +
	git checkout go.mod

.PHONY: up
up: ## Start the service.
	air

.PHONY: docs
docs: ## Generate documentation.
	npx @redocly/cli build-docs api/api.yml
	mv redoc-static.html public/index.html

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
