.DEFAULT_GOAL := build

.PHONY: fmt vet test build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

build: vet
	mkdir -p build/bin
	go build -o build/bin/txtpix ./cmd/txtpix

run: build
	go run ./cmd/txtpix