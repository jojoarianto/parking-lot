.PHONY: build test run

build: 
	export GO111MODULE=on
	env GOOS=linux go build -o bin/main src/main.go

test: 
	go test ./... 

run:
	go run src/main.go
