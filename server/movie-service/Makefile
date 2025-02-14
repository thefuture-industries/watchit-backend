build:
	@echo "Running build..."
	@go build -o bin/movie-service main.go

test:
	@go test -v ./...

run: build
	@./bin/movie-service

run-to-test:
	@cmd /c "$(CURDIR)/$(BATCH_FILE_TEST)"

migrate-up:
	@echo "Running migrations..."
	@go run cmd/migrate/main.go

migrate-down:
	@go run cmd/migrate/main-down.go
