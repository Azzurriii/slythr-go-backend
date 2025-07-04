.PHONY: build run test docker-build docker-run dev clean fmt lint vet staticcheck install-lint depend-services

APP_NAME=slythr

build:
	mkdir -p bin
	go build -o bin/app cmd/api/main.go

depend-services:
	docker compose up -d slither
	docker compose up -d postgres
	docker compose up -d redis

run:
	if command -v air >/dev/null 2>&1; then air; else go run cmd/api/main.go; fi

test:
	go test -v ./...

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker compose up -d

dev:
	air

clean:
	rm -rf bin/ tmp/

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
