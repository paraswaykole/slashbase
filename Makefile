# Try to get the semver from 1) git 2) the VERSION file 3) fallback.
VERSION := $(or $(shell git describe --tags --abbrev=0 2> /dev/null),"v0.0.0")

GOPATH ?= $(HOME)/go

BIN := slashbase
BINWIN := slashbase.exe
STATIC := web

.PHONY: build

build:
	env CGO_ENABLED=1 go build --o $(BIN) -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"

# THIS IS FOR BUILDING BIN FOR WINDOWS FROM MAC

# build-win: $(BINWIN)

# $(BINWIN):
# 	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build --o $(BINWIN) -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"
