color_red     := $(shell printf "\e[1;31m")
color_green   := $(shell printf "\e[1;32m")
color_yellow  := $(shell printf "\e[1;33m")
color_blue    := $(shell printf "\e[1;34m")
color_magenta := $(shell printf "\e[1;35m")
color_cyan    := $(shell printf "\e[1;36m")
color_reset   := $(shell printf "\e[0m")

EXAMPLES := \
	_examples/a-h-templ-single-toast \
	_examples/a-h-templ-multiple-toasts \
	_examples/custom-animations \
	_examples/custom-icons \
	_examples/theming \
	_examples/variants \
	_examples/go-html-template-single-toast \
	_examples/go-html-template-multiple-toasts

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
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$NF}' $(MAKEFILE_LIST) | sort
	@echo ""

# ==================================================================================== #
# PRIVATE BUILDERS
# ==================================================================================== #
_process_templ_files:
	@echo "$(color_magenta)Process TEMPL files$(color_reset)"

	@echo "$(color_cyan) * Formatting templ files...$(color_reset)"
	@templ fmt .

	@echo "$(color_cyan) * Generating templ funcs...$(color_reset)"
	@templ generate

	@echo "$(color_green)Done!$(color_reset)"

_go_tools:
	@echo ""
	@echo "$(color_magenta)Run go tools...$(color_reset)"

	@echo "$(color_cyan) * Downloading modules...$(color_reset)"
	@go mod download; go mod tidy
	@go build -v ./...
	@echo "$(color_cyan) * Running golangci-lint...$(color_reset)"
	@golangci-lint run ./...
	@echo "$(color_cyan) * Running tests...$(color_reset)"
	@$(MAKE) _test

	@echo "$(color_green)Done!$(color_reset)"

_test:
	@go test -race -covermode=atomic ./...

# ==================================================================================== #
# BUILDERS
# ==================================================================================== #

examples: ## Process templ files in the _examples folder
	@for example in $(EXAMPLES); do \
		pushd $$example && templ fmt . && templ generate && popd; \
	done

test: ## Run go tests
	@$(MAKE) _test

templ: ## Process TEMPL files
	@$(MAKE) _process_templ_files

clean:
	@echo ""
	@echo "$(color_magenta)Clean up$(color_reset)"
	@find . -type f -name '*.grc.bk' -exec rm -f {} +
	@find . -type f -name '*_templ.txt' -exec rm -f {} +
	@echo "$(color_green)Done!$(color_reset)"

build: ## The main build target
	@$(MAKE) templ _go_tools