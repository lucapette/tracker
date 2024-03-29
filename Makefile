SOURCE_FILES?=$$(go list ./... | grep -v '/tracker/vendor/')
TEST_PATTERN?=.
TEST_OPTIONS?=

test: ## Run all the tests
	go test $(TEST_OPTIONS) -cover $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=30s

lint: ## Run all the linters
	golangci-lint run

ci: lint test ## Run all the tests and code checks

build: ## Build a dev version of tracker
	go build cmd/tracker.go

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
