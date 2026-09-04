BIN_DIR := bin
BINARY  := $(BIN_DIR)/app

.PHONY: build run test clean

# El binario lee scripts/DDL.sql y templates/ con paths relativos al cwd,
# así que build siempre los deja al lado (cd bin && ./app para correrlo).
build:
	go build -o $(BINARY) ./cmd/bitacora
	cp -r ./scripts $(BIN_DIR)/
	cp -r ./templates $(BIN_DIR)/

run:
	go run ./cmd/bitacora

test:
	go test -v ./...

clean:
	rm -rf $(BIN_DIR)
