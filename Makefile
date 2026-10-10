.DEFAULT_GOAL := build
.PHONY:clean fmt vet build
clean:
	go clean ./...
lint: clean
	golangci-lint run
lint-clean: lint
	golangci-lint run --fix
fmt: lint-clean
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size