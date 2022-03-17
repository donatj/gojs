BIN=wasmtest.wasm

.PHONY: build
build: $(BIN)

$(BIN): go.mod go.sum $(shell find . -name "*.go")
	GOOS=js GOARCH=wasm go build -o wasmtest.wasm