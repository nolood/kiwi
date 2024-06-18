# Variables
CONFIG_FILE = ./config/local.yml
CMD_DIR = ./cmd/bot
GOOSE_DRIVER = postgres
MIGRATION_DIR = ./migrations
DB_HOST = localhost
DB_PORT = 5443
DB_USER = postgres
DB_PASSWORD = 1234
DB_NAME = kiwi
DB_STRING = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

# Phony targets
.PHONY: start migrate-new migrate-up migrate-down jet-gen

# Start the application
start:
	go run $(CMD_DIR)/main.go -config=$(CONFIG_FILE)

# Create a new migration
migrate-new:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(DB_STRING) goose create new-migration sql

# Apply all migrations
migrate-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(DB_STRING) goose up

# Roll back the last migration
migrate-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(DB_STRING) goose down

# Generate Jet ORM code
jet-gen:
	jet -dsn=$(DB_STRING)?sslmode=disable -schema=public -path=./.gen
