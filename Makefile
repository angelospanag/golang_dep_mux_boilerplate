.PHONY: build lint run

build:
	go build

# Lint all Go project files except vendor folder
lint:
	golint -set_exit_status $$(go list ./... | grep -v /vendor/)

run:
	go run main.go
