// This is Go WASM equivalent of main.js file
package main

import (
	"syscall/js"

	"github.com/stanleynguyen/go-everywhere/lighthttpcli"
)

func getStateBtnHandlerFunc(state string, cli lighthttpcli.LightHttpCli) js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			go func() {
				err := cli.SetState(state)
				if err != nil {
					println(err.Error())
				}
			}()
			return nil
		},
	)
}

func getRefreshStateFunc(bulbElem js.Value, cli lighthttpcli.LightHttpCli) js.Func {
	var prevState string
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			go func() {
				state, err := cli.GetState()
				if err != nil {
					println(err.Error())
				}

				if state != prevState {
					if state == lighthttpcli.StateOn {
						bulbElem.Get("classList").Call("add", "on")
					} else {
						bulbElem.Get("classList").Call("remove", "on")
					}
				}
			}()
			return nil
		},
	)
}

func setup() {
	cli := lighthttpcli.NewCli(js.Global().Get("location").Get("origin").String())
	bulbElem := js.Global().Get("document").Call("getElementById", "bulb")

	js.Global().Set("turnOn", getStateBtnHandlerFunc("on", cli))
	js.Global().Set("turnOff", getStateBtnHandlerFunc("off", cli))
	js.Global().Call("setInterval", getRefreshStateFunc(bulbElem, cli), 500)
}

func main() {
	c := make(chan struct{}, 0)
	setup()
	println("WASM Go initialized")
	<-c
}
