.PHONY: all lint lint-fix fmt vet tidy build test clean

all: tidy clean lint fmt vet build test

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 10m

lint-fix:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix --timeout 10m

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build:
	mkdir -p bin
	go build -o ./bin/celery .

test:
	go test ./...

clean:
	rm -rf bin
