# justfile for greenhead

# Show this justfile's commands.
list:
	just --list

# Prepare assets etc.
prep:
	@echo prep is TBD, see if we have any assets first

# Run locally from source, with args passed. Args must not contain spaces.
run *ARGS='--version': prep
	go run ./cmd/orgn {{ARGS}}

# Build for current environment.
build: prep
	mkdir -p build
	go build -o build/orgn ./cmd/orgn

# Check for programmer errors.
vet: prep
	go vet ./...

# Run unit tests.
test: prep
	go test ./...

# Run unit tests with coverage, and open the coverage report.
cover: prep
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

# Run benchmarks, if any.
bench: prep
	go test ./... -bench=.

# Run pkgsite because godoc is deprecated. :-(
doc: prep
	pkgsite

# Remove the build and cover artifacts.
clean:
	/bin/rm -rf build cover.out


