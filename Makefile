GOPATH ?= $(HOME)/go

BIN := slashbase
BIN-WIN := slashbase.exe
STATIC := web

.PHONY: build
build: $(BIN)

$(BIN):
	 env CGO_ENABLED=1 go build --o ${BIN} -trimpath -ldflags="-X 'main.Build=production'"

# THIS IS FOR BUILDING BIN FOR WINDOWS FROM MAC
.PHONY: build-win
build-win: $(BIN-WIN)

$(BIN-WIN):
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build --o ${BIN-WIN} -trimpath -ldflags="-X 'main.Build=production'"
