APP=console

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

GOFILES = $(shell find . -name '*.go')
GOPACKAGES = $(shell go list ./...)

all: dependencies build

dependencies:
	go mod download

test: dependencies
	@go test -v $(GOPACKAGES)

benchmark: dependencies fmt
	@go test $(GOPACKAGES) -bench=.

fmt:
	gofmt -w .

build: 
	go build -o $(GOBIN)/$(APP) ./cmd/$(APP)/*.go || exit

build-all: fmt test build

run: build-all
	./bin/$(APP)

default: build

.PHONY: build project fmt
