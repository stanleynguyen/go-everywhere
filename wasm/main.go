package main

import "syscall/js"

func lol(this js.Value, args []js.Value) interface{} {
	for _, arg := range args {
		println(arg.String())
	}
	return nil
}

func registerCallbacks() {
	js.Global().Set("lol", js.FuncOf(lol))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("WASM Go initialized")
	<-c
}
