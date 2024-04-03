include .env

.PHONY: all db_create db_up db_drop docker_volume migrate_app migrate_create migrate_up migrate_down install-golangci-lint lint

LOCAL_BIN=$(CURDIR)/bin
SWAG_PATH=$(CURDIR)/report
MIGRATION_PATH=$(CURDIR)/migrations
LOCAL_MIGRATION_DSN="host=$(DB_DSN) port=$(DB_PORT) dbname=$(DB_NAME) user=$(DB_USER) password=$(DB_PASSWORD) sslmode=disable"
CONTAINER_PORT=5432
PG_V=14-alpine3.17

all:
	@go run ./... -path .env

db_create:
	docker pull postgres:$(PG_V)

db_up: docker_volume
	docker run --name $(DB_NAME) -p $(DB_PORT):$(CONTAINER_PORT) -e POSTGRES_DB=$(DB_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -v taskDB:$(DB_VOLUME) -d postgres:$(PG_V)

db_drop:
	docker kill $(DB_NAME)
	docker rm $(DB_NAME)

docker_volume:
	docker volume create tasks_data

migrate_app:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.19.1

migrate_create:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_PATH) create main_table sql

migrate_up:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_PATH) postgres $(LOCAL_MIGRATION_DSN) up -v

migrate_down:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_PATH) postgres $(LOCAL_MIGRATION_DSN) down -v

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml
	
# goimports -w <filename>