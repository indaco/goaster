# Colors
color_green   := $(shell printf "\e[32m")  # Green color
color_reset   := $(shell printf "\e[0m")

# ==================================================================================== #
# HELPERS
# ==================================================================================== #
.PHONY: help
help: ## Print this help message
	@echo ""
	@echo "Usage: make [action]"
	@echo ""
	@echo "Available Actions:"
	@echo ""
	@awk -v green="$(color_green)" -v reset="$(color_reset)" -F ':|##' \
		'/^[^\t].+?:.*?##/ {printf "%s* %-15s%s %s\n", green, $$1, reset, $$NF}' $(MAKEFILE_LIST) | sort
	@echo ""

# ==================================================================================== #
# PRIVATE TARGETS
# ==================================================================================== #
.PHONY: modernize
modernize:
	@modernize -test ./...

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test/view-total-coverage
test/view-total-coverage:
	@echo ""
	@echo "Total Coverage:"
	@go tool cover -func=profile.cov | grep total | awk -F '[[:space:]]+' '{print $$NF}'

.PHONY: test/view-coverage
test/view-coverage:
	@go tool cover -html=profile.cov
	@echo "Coverage report displayed in your default browser."

.PHONY: live/templ
live/templ:
	@templ generate --watch --proxy="http://localhost:3300" --open-browser=true

.PHONY: live/server
live/server:
	@go run github.com/air-verse/air@v1.61.7 \
		--build.cmd "go build -o tmp/bin/main ./examples" \
		--build.bin "tmp/bin/main" \
		--build.delay "100" \
		--build.exclude_dir "assets,docs" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

.PHONY: live/sync
live/sync:
	@go run github.com/air-verse/air@v1.61.7 \
		--build.cmd "tempo sync && templ generate -notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "assets" \
		--build.include_ext "js,css"

# ==================================================================================== #
# PUBLIC TARGETS
# ==================================================================================== #
.PHONY: test
test: ## Run all tests and generate coverage report
	@go test -count=1 -timeout 30s $(shell go list ./... | grep -Ev 'examples|components') -covermode=atomic -coverprofile=profile.cov
	@$(MAKE) test/view-total-coverage

.PHONY: test/coverage
test/coverage: ## Run go tests and use go tool cover
	@$(MAKE) test/force
	@$(MAKE) test/view-coverage

.PHONY: test/force
test/force: ## Clean go tests cache and run all tests
	@go clean -testcache
	@$(MAKE) test

.PHONY: templ
templ: ## Run all tests and generate coverage report
	@templ fmt . && templ generate

.PHONY: dev
dev: ## Run the dev server with live reload
	make -j3 live/templ live/server live/sync

.PHONY: pre-build
pre-build: ## Run pre-build tasks
	@$(MAKE) modernize
	@$(MAKE) lint
	@$(MAKE) test/force

.PHONY: build
build: ## Build for production with minified asset files
	@$(MAKE) pre-build
	@tempo sync --prod
	@$(MAKE) templ