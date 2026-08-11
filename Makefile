SERVICE=auth-service
BIN=bin/auth-service
DOCKER_IMAGE?=marketing-digest-auth-service
ROOT:=$(shell git -C ../.. rev-parse --show-toplevel 2>/dev/null || realpath ../../..)
# Workspace root containing backend/ and protos/
WS_ROOT:=$(abspath ../../..)

.PHONY: build test lint run docker-build atlas-diff atlas-lint atlas-validate atlas-hash tidy

build:
	mkdir -p bin
	go build -o $(BIN) ./cmd/server

test:
	go test ./...

lint:
	go vet ./...

tidy:
	go mod tidy

run: build
	set -a && [ -f .env ] && . ./.env; set +a; ./$(BIN)

docker-build:
	docker build -f $(abspath Dockerfile) -t $(DOCKER_IMAGE) $(WS_ROOT)

atlas-hash:
	atlas migrate hash --dir file://migrations

atlas-diff:
	@echo "Set DATABASE_URL and DEV_DATABASE_URL, then:"
	@echo "  atlas migrate diff --env local --to \"file://migrations\""
	@echo "Prefer: atlas schema diff against GORM-desired schema after review."

atlas-lint:
	atlas migrate lint --env local --latest 1

atlas-validate:
	atlas migrate validate --dir file://migrations
