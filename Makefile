.PHONY: install build

install:
	@go mod tidy
	@go mod download

build:
	@go build -o bin/server
