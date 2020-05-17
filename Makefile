start_server:
	go run main.go
compile_wasm:
	GOARCH=wasm GOOS=js go build -o static/main.wasm wasm/main.go