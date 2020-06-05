package main

import (
	"machine"
	"time"
)

func main() {
	var outPin machine.Pin = 9
	var inPin machine.Pin = 7
	outPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	inPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	for {
		outPin.Set(inPin.Get())
		time.Sleep(time.Millisecond * 200)
	}
}
