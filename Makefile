SHELL := /bin/bash

export PROJECT = go-bookclub

# ==============================================================================
# Development

run: up dev

up:
	docker-compose up -d db

dev:
	go run ./cmd/app --db-host=localhost:5678

dcup:
	docker-compose up

dcdown:
	docker-compose down -v --remove-orphans

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Administration

migrate:
	go run ./cmd/admin --db-host=localhost:5678 migrate

dcmigrate:
	docker-compose exec app go run ./cmd/admin --db-host=db:5678 migrate

seed: migrate
	go run ./cmd/admin --db-host=localhost:5678 seed

dcseed: dcmigrate
	docker-compose exec app go run ./cmd/admin --db-host=db:5678 seed

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1
	staticcheck ./...
