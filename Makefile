.PHONY: run build docker-compose

include .env

run:
	go run cmd/main.go

build:
	rm -f ./main
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-d -s -w" -o . ./cmd/main.go

docker-compose:
	docker-compose up -d --remove-orphans