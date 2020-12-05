.PHONY: dep run build test test-ci

run:
	go run main.go

dep:
	go mod download
	go mod verify

build:
	go build .

lint:
	go fmt ./...

docs-update:
	- rm -rf docs
	swag init

test:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

test-ci:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
