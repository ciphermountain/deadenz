APP=deadenz
MUL=multiverse

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

generate: gen-grpc gen-grpc-multiverse

mockery:
	mockery

gen-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/proto/core/core.proto

gen-grpc-multiverse:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/proto/multiverse/multiverse.proto

build: gen-grpc
	go build -o $(GOBIN)/$(APP) ./cmd/$(APP)/*.go || exit

build-multiverse: gen-grpc-multiverse
	go build -o $(GOBIN)/$(MUL) ./cmd/$(MUL)/*.go || exit

build-all: fmt test build build-multiverse

run: build-all
	./bin/$(APP)

default: build

.PHONY: build project fmt
