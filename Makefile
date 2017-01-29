.PHONY: all test build

all: build


test:
	go test

build: test
	go build
