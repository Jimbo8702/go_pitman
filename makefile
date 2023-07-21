build: 
	@go build -o bin/gocrawl

run: build
	@./bin/gocrawl

test:
	@go test -v ./...


