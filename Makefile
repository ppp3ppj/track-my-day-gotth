.PHONY: dev
dev: ## run dev
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air -c .air.dev.toml

# Default environment
ENV ?= development

# Load the correct database path based on the environment
DB_PATH := $(shell yq e ".database.$(ENV)_db_path" appinfo.yaml)

# Migration Directory
MIGRATION_DIR := internal/infrastructure/database/migrations

.PHONY: migrate-create migrate-up migrate-down

# Target to create a new migration
migrate-create:
	@echo "Creating new migration in $(MIGRATION_DIR)"
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

# Target to apply migrations
migrate-up:
	@echo "Applying migrations to database at $(DB_PATH)"
	migrate -source file://$(MIGRATION_DIR) -database sqlite3://$(DB_PATH) -verbose up

# Target to roll back migrations
migrate-down:
	@echo "Rolling back migrations from database at $(DB_PATH)"
	migrate -source file://$(MIGRATION_DIR) -database sqlite3://$(DB_PATH) -verbose down

# Optional: Target to reset migrations
migrate-reset: migrate-down migrate-up
	@echo "Resetting database migrations"
