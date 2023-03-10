NAME=rollback-buildkite-plugin

HAS_DOCKER=$(shell command -v docker;)
HAS_GORELEASER=$(shell command -v goreleaser;)

.PHONY: all
all: quality test

.PHONY: test-go
test-go:
	go test -race -coverprofile=coverage.out -covermode=atomic

.PHONY: build-docker-test
build-docker-test:


.PHONY: test-docker
test-docker: build-docker-test


.PHONY: test
test: test-go test-docker

.PHONY: quality
quality:
	go vet
	go fmt
	go mod tidy

.PHONY: build
build:
ifneq (${HAS_GORELEASER},)
	goreleaser build --rm-dist --skip-validate
else
	$(error goreleaser binary is missing, please install goreleaser)
endif
