.DEFAULT_GOAL := build
.PHONY: test

build:
	@go build -o tetris

clean:
	@rm tetris

test:
	@go test
