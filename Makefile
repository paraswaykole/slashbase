VERSION := $(or $(shell git describe --tags --abbrev=0 2> /dev/null),"v0.0.0")

GOPATH ?= $(HOME)/go
WAILS ?= $(GOPATH)/bin/wails

.PHONY: build

build:
	env CGO_ENABLED=1 $(WAILS) build -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"

.PHONY: build-win

build-win:
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" $(WAILS) build -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'" -skipbindings