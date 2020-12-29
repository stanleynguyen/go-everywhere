start_server:
	go run main.go
compile_wasm:
	GOARCH=wasm GOOS=js go build -o static/main.wasm wasm/main.go
bundle_static:
	statik -src static/ -include=*.html,*.css,*.js,*.wasm