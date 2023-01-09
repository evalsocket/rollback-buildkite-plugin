all: test

test:
	go test -timeout=3s -v ./...

build:
	go build main.go

.PHONY: all build test
