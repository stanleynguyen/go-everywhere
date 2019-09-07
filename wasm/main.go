// This is Go WASM equivalent of main.js file
package main

import "syscall/js"

func newWS() js.Value {
	wsURI := wsURI()

	return js.Global().Get("WebSocket").New(wsURI)
}

func wsURI() string {
	loc := js.Global().Get("window").Get("location")
	wsURI := ""
	if loc.Get("protocol").String() == "https:" {
		wsURI = "wss:"
	} else {
		wsURI = "ws:"
	}
	wsURI += "//" + loc.Get("host").String()
	wsURI += loc.Get("pathname").String() + "led"

	return wsURI
}

func newOpenHandler() js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			println("Connected")
			return nil
		})
}

func newMessageHandler(lightElem js.Value) js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			event := args[0]
			lightElem.Set("innerHTML", event.Get("data").String())
			return nil
		})

}

func newSendHandler(ws js.Value) js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			ws.Call("send", "switch")
			return nil
		})
}

func registerCallbacks() {
	ws := newWS()
	lightElem := js.Global().Get("document").Call("getElementById", "light")

	ws.Call("addEventListener", "open", newOpenHandler())
	ws.Call("addEventListener", "message", newMessageHandler(lightElem))
	js.Global().Set("send", newSendHandler(ws))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("WASM Go initialized")
	<-c
}
