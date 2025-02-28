build:
	@echo "Running build..."
	@go build -o bin/gateway-service main.go

test:
	@go test -v ./...

run: build
	@./bin/gateway-service

run-to-test:
	@cmd /c "$(CURDIR)/$(BATCH_FILE_TEST)"
