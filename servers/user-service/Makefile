build:
	@go build -o bin/user-service main.go

test:
	@go test -v ./...

run: build
	@./bin/user-service

run-to-test:
	@cmd /c "$(CURDIR)/$(BATCH_FILE_TEST)"

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
