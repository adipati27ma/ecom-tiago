build:
	@go build -o bin/ecom-tiago cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-tiago

gorun:
	@go run cmd/main.go