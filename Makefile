MODULE_NAME=ligaturizer

BUILD_DIR ?= out
VENDOR_DIR = vendor

APP ?= ligaturizer

PYTHON_VERSION ?= 3.12
ALPINE_VERSION ?= 3.20
GOLANGCI_LINT_VERSION ?= v2.2.1

GO ?= go
GOLANGCI_LINT ?= $(shell go env GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION)

.PHONY: $(VENDOR_DIR)
$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)
	@$(GO) mod vendor
	@$(GO) mod tidy

.PHONY: lint
lint:
	@$(GOLANGCI_LINT) run

.PHONY: test
test: test-unit

## Run unit tests
.PHONY: test-unit
test-unit:
	@echo ">> unit test"
	@$(GO) test -gcflags=-l -coverprofile=unit.coverprofile -covermode=atomic -race ./...

#.PHONY: test-integration
#test-integration:
#	@echo ">> integration test"
#	@$(GO) test ./features/... -gcflags=-l -coverprofile=features.coverprofile -coverpkg ./... -race --godog

.PHONY: build
build:
	@env CGO_ENABLED=1 \
	$(GO) build -ldflags "$(shell ./resources/scripts/build_args)" -o $(BUILD_DIR)/$(APP) cmd/$(APP)/main.go && \
		chmod +x $(BUILD_DIR)/$(APP)

.PHONY: $(GITHUB_OUTPUT)
$(GITHUB_OUTPUT):
	@echo "MODULE_NAME=$(MODULE_NAME)" >>"$@"
	@echo "PYTHON_VERSION=$(PYTHON_VERSION)" >>"$@"
	@echo "ALPINE_VERSION=$(ALPINE_VERSION)" >>"$@"
	@echo "GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION)" >>"$@"

.PHONY: $(GITHUB_ENV)
$(GITHUB_ENV):
	@echo "MODULE_NAME=$(MODULE_NAME)" >>"$@"
	@echo "PYTHON_VERSION=$(PYTHON_VERSION)" >>"$@"
	@echo "ALPINE_VERSION=$(ALPINE_VERSION)" >>"$@"
	@echo "GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION)" >>"$@"

$(GOLANGCI_LINT):
	@echo "$(OK_COLOR)==> Installing golangci-lint $(GOLANGCI_LINT_VERSION)$(NO_COLOR)"; \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin "$(GOLANGCI_LINT_VERSION)"
	@mv ./bin/golangci-lint $(GOLANGCI_LINT)
