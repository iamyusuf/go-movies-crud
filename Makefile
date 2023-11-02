build:
	@go build -o bin/go-movies-crud

run: build
	@./bin/go-movies-crud

test:
	@go test -v ./...-v ./...