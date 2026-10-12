.DEFAULT_GOAL := build
.PHONY:clean fmt vet build
clean:
	go clean ./...
lint: clean
	golangci-lint run --fix
build: lint
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size