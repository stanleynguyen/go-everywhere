package main

import (
	"machine"
	"strconv"
	"time"
)

var outPinStr = "9" // Inject at build time with -ldflags "-X main.outPinStr=9"
var inPinStr = "7"  // Inject at build time with -ldflags "-X main.outPinStr=7"

func main() {
	outPinNumber, _ := strconv.Atoi(outPinStr)
	inPinNumber, _ := strconv.Atoi(inPinStr)
	var outPin machine.Pin = outPinNumber
	var inPin machine.Pin = inPinNumber
	outPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	inPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	for {
		outPin.Set(inPin.Get())
		time.Sleep(time.Millisecond * 200)
	}
}
