VERSION := $(or $(shell git describe --tags --abbrev=0 2> /dev/null),"v0.0.0")

GOPATH ?= $(HOME)/go

BIN := slashbase
BINWIN := slashbase.exe
STATIC := web

.PHONY: build

build:
	env go build --o $(BIN) -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"

.PHONY: build-win

build-win:
	env GOOS=windows GOARCH=amd64 go build --o $(BINWIN) -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"
