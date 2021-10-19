## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the cmd/web application
.PHONY: run
run:
	@go run ./...

## audit: tidy deps and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

## build: builds daypaper
current_time = $(shell date -u --iso-8601=seconds)
current_tag = $(shell git describe --tags --dirty)
commit_hash = $(shell git describe --always)
linker_flags = '-s -w -X main.buildTime=${current_time} -X main.version=${current_tag}+${commit_hash}' -o=./dist/daypaper

.PHONY: build
build: audit
	rm -rf ./dist
	mkdir ./dist
	go build -ldflags=${linker_flags} -o=./dist ./...

## build/snapshot: build a snapshot using goreleaser
.PHONY: build/snapshot
build/snapshot:
	goreleaser release --snapshot --rm-dist

## build/release: publish a release using goreleaser
.PHONY: build/release
build/release: audit
	goreleaser release --rm-dist