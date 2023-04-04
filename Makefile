all: build install

.PHONY: build install

build:
	@go build ./...

install:
	@go install ./...
