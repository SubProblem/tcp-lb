# Variables
BINARY_NAME=tcplb
MAIN_FILE=main.go

build:
	go build -o ${BINARY_NAME} .

run:
	go run .

test:
	go test ./...

test-cov:
	go test -v -cover ./...