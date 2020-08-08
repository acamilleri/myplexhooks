.PHONY: lint build binary test docker

VERSION = $(shell git rev-parse --verify HEAD)

build: binary docker

binary:
	go build -ldflags "-X main.version=$(VERSION)" cmd/myplexhooks.go

docker:
	docker build --build-arg version=$(VERSION) .

lint:
	@go get -u golang.org/x/lint/golint
	golint ./...

test:
	go test -v -cover ./...
