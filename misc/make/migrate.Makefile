# Makefile for Goose Migration
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Variables
GOOSE_CMD=goose
GOOSE_DIR=./database/migrations
DB_DRIVER=mysql
DB_STRING=${DATABASE_DSN}
MIGRATION_NAME?=migration
MIGRATION_PATH=$(GOOSE_DIR)/$(MIGRATION_NAME).sql

# Help
help:
	@echo "Usage:"
	@echo "    make install          Install goose tool"
	@echo "    make create           Create a new migration file"
	@echo "    make up               Run all pending migrations"
	@echo "    make down             Rollback the last migration"
	@echo "    make redo             Rollback the last migration and reapply it"
	@echo "    make status           Print current migration status"
	@echo "    make reset            Rollback all migrations"
	@echo "    make version          Print the current migration version"

# Targets
install:
	@echo "Installing goose..."
	go install github.com/pressly/goose/v3/cmd/goose@latest

create:
	@echo "Creating new migration: $(MIGRATION_NAME)"
	$(GOOSE_CMD) -dir $(GOOSE_DIR) create $(MIGRATION_NAME) sql

up:
	@echo "Applying all pending migrations..."
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" up

down:
	@echo "Rolling back the last migration..."
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" down

redo:
	@echo "Reapplying the last migration..."
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" redo

status:
	@echo "Checking migration status..."
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" status

reset:
	@echo "Rolling back all migrations..."
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" reset

version:
	@echo "Current migration version:"
	$(GOOSE_CMD) -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DB_STRING)" version
