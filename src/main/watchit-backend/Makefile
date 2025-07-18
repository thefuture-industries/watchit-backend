DB_URL ?= postgres://postgres:password@localhost:5432/watchit?sslmode=disable
MIGRATIONS_DIR ?= cmd/migrate/migrations

GOOSE ?= goose

.PHONY: fmt lint test

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run

check: fmt lint test
	@echo "All checks passed!"
build:
	@echo "Running build..."
	@go build -o bin/watchit_http ./src/main/watchit-backend/cmd/server

test:
	@go test -v ./...

run: build
	@./bin/watchit_http

run-to-test:
	@cmd /c "$(CURDIR)/$(BATCH_FILE_TEST)"

## make create NAME=create_users_table
migrate-create:
	@echo "Creating new migration: $(NAME)"
	$(GOOSE) -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-up:
	@echo "Running up migrations"
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate-down:
	@echo "Running down migrations..."
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

database-drop:
	@echo "Dropping all tables in database..."
	@psql "$(DB_URL)" -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
