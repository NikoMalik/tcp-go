build:
	@go build -o bin/cs

run: build
	@./bin/cs

test:
	@go test ./...


clean:
	@rm -rf bin/cs

.PHONY: run test clean
