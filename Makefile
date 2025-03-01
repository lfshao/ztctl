# Go parameters
GOBIN = go
GOFMT = gofmt
GOLINT = golangci-lint
BINARY_NAME = ztctl

.PHONY: all build clean test lint fmt

all: lint test build

build:
	$(GOBIN) build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)

test:
	$(GOBIN) test -v ./...

lint:
	$(GOLINT) run

fmt:
	$(GOFMT) -w .

install:
	$(GOBIN) install

.DEFAULT_GOAL := build