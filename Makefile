# Variables
CONFIG_FILE = ./config/local.yml
DATA_FILE = ./static/cities5000.txt
CMD_DIR = ./cmd
GOOSE_DRIVER = postgres
MIGRATION_DIR = ./migrations
DB_HOST = localhost
DB_PORT = 5443
DB_USER = postgres
DB_PASSWORD = 1234
DB_NAME = kiwi
DB_STRING = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

# Detect OS
ifeq ($(OS),Windows_NT)
    SET_ENV = set
else
    SET_ENV = export
endif

# Phony targets
.PHONY: start migrate-new migrate-up migrate-down jet-gen start-location docker-dev

# Start the application
start:
	go run $(CMD_DIR)/bot/main.go -config=$(CONFIG_FILE)

# Start the location service to fill database
start-location:
	go run $(CMD_DIR)/location/main.go -config=$(CONFIG_FILE)

# Create a new migration
migrate-new:
	$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER) && $(SET_ENV) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) && $(SET_ENV) GOOSE_DBSTRING=$(DB_STRING) && goose create new-migration sql

# Apply all migrations
migrate-up:
	$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER) && $(SET_ENV) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) && $(SET_ENV) GOOSE_DBSTRING=$(DB_STRING) && goose up

# Roll back the last migration
migrate-down:
	$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER) && $(SET_ENV) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) && $(SET_ENV) GOOSE_DBSTRING=$(DB_STRING) && goose down

# Generate Jet ORM code
jet-gen:
	jet -dsn=$(DB_STRING)?sslmode=disable -schema=public -path=./.gen

# Start the development environment
docker-dev:
	docker-compose -f ./docker/docker-compose.dev.yml up -d