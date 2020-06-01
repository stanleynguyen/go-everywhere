package main

import (
	"time"

	"github.com/stanleynguyen/go-everywhere/lighthttpcli"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

var lightHTTPCli = lighthttpcli.NewCli("http://localhost:8080")

func main() {
	stateChan := make(chan string)
	go checkState(stateChan)
	app.Main(func(a app.App) {
		var glctx gl.Context
		state := "OFF"
		for {
			select {
			case state = <-stateChan:
				a.Send(paint.Event{})
			case e := <-a.Events():
				switch e := a.Filter(e).(type) {
				case lifecycle.Event:
					glctx, _ = e.DrawContext.(gl.Context)
				case paint.Event:
					if glctx == nil {
						continue
					}
					if state == "ON" {
						glctx.ClearColor(1, 1, 0, 1)
					} else {
						glctx.ClearColor(0, 0, 0, 1)
					}
					glctx.Clear(gl.COLOR_BUFFER_BIT)
					a.Publish()
				case touch.Event:
					if state == "ON" {
						lightHTTPCli.SetState("OFF")
					} else {
						lightHTTPCli.SetState("ON")
					}
				}
			}
		}
	})
}

func checkState(stateChan chan<- string) {
	t := time.NewTicker(500 * time.Millisecond)
	for {
		<-t.C // wait for 500 Millisecond
		state, _ := lightHTTPCli.GetState()
		stateChan <- state
	}
}
