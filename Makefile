
# Makefile

build-wasm:
	cd ./cmd/wasm && \
	GOOS=js GOARCH=wasm go build -o ../../web/json.wasm

run-server:
	cd ./cmd/server/ && go run main.go
