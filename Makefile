include .env

.PHONY: build
build:
	go build -v  -o exefile ./cmd/main.go

.PHONY: up
up:
	go run ./cmd/main.go

.PHONY: migrate-up
migrate-up:
	migrate -path ./schema -database $(DB_CONNECT) up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema -database $(DB_CONNECT) down

.DEFAULT_GOAL := up