build:
	go build -o bin/vendingmachine -v ./cmd/vendingmachine

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build./