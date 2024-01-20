help:
	@cat Makefile

build:
	@go build .

test:
	@go test 

coverage:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out
