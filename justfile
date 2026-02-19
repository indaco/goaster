# List available recipes
default:
    @just --list

# ==================================================================================== #
# Private recipes
# ==================================================================================== #

[private]
modernize:
    modernize -test ./...

[private]
lint:
    golangci-lint run ./...

[private]
view-total-coverage:
    @echo ""
    @echo "Total Coverage:"
    @go tool cover -func=coverage.txt | grep total | awk -F '[[:space:]]+' '{print $NF}'

[private]
view-coverage:
    go tool cover -html=coverage.txt
    @echo "Coverage report displayed in your default browser."

[private]
live-templ:
    templ generate --watch --proxy="http://localhost:3300" --open-browser=true

[private]
live-server:
    go run github.com/air-verse/air@v1.61.7 \
        --build.cmd "go build -o tmp/bin/main ./examples" \
        --build.bin "tmp/bin/main" \
        --build.delay "100" \
        --build.exclude_dir "assets,docs" \
        --build.include_ext "go" \
        --build.stop_on_error "false" \
        --misc.clean_on_exit true

[private]
live-sync:
    go run github.com/air-verse/air@v1.61.7 \
        --build.cmd "tempo sync && templ generate -notify-proxy" \
        --build.bin "true" \
        --build.delay "100" \
        --build.exclude_dir "" \
        --build.include_dir "assets" \
        --build.include_ext "js,css"

# ==================================================================================== #
# Public recipes
# ==================================================================================== #

# Run all tests and generate coverage report
test:
    go test -count=1 -timeout 30s $(go list ./... | grep -Ev 'examples|components') -covermode=atomic -coverprofile=coverage.txt
    @just view-total-coverage

# Clean go tests cache and run all tests
test-force:
    go clean -testcache
    @just test

# Run go tests and use go tool cover
test-coverage:
    @just test-force
    @just view-coverage

# Run templ fmt and templ generate commands
templ:
    templ fmt . && templ generate

# Run the dev server with live reload
dev:
    just live-templ & just live-server & just live-sync & wait

# Run pre-build tasks
pre-build:
    @just modernize
    @just lint
    @just test-force

# Build for production with minified asset files
build:
    @just pre-build
    tempo sync --prod
    @just templ
