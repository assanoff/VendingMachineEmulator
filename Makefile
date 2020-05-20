build:
	go build -o bin/vendingmachine -v ./cmd/vendingmachine

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	./bin/vendingmachine	

.DEFAULT_GOAL := build./