VERSION := $(or $(shell git describe --tags --abbrev=0 2> /dev/null),"v0.0.0")

GOPATH ?= $(HOME)/go
WAILS ?= $(GOPATH)/bin/wails

# ALWAYS USE BUILD PHONY RECIPE TO BUILD FROM SOURCE

.PHONY: build
build:
	env CGO_ENABLED=1 $(WAILS) build -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'"

# DO NOT USE THE FOLLOWING PHONY RECIPIES, THEY ARE ONLY FOR DISTRIBUTION

.PHONY: build-win
build-win:
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" $(WAILS) build -trimpath -ldflags="-s -w -X 'main.build=production' -X 'main.version=$(VERSION)'" -skipbindings

.PHONY: sign
sign:
	codesign --timestamp --options=runtime -s "Developer ID Application: Paras Waykole (EGSVK8P42D)" -v --entitlements ./build/darwin/entitlements.plist ./build/bin/Slashbase.app

.PHONY: dmg
dmg:
	create-dmg ./build/bin/Slashbase.app