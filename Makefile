SHELL := /bin/zsh

.PHONY: run tidy generate

run:
	go run ./cmd/server

tidy:
	go mod tidy

generate:
	go run github.com/99designs/gqlgen generate

.PHONY: migrate
migrate:
	go run ./cmd/migrate

