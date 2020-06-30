package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/stanleynguyen/go-everywhere/lighthttpcli"
	"github.com/stianeikeland/go-rpio/v4"
)

var serverURL = "http://localhost:8080" // Inject at build time with -ldflags "-X main.serverURL=http://something"
var pinNumberStr = "21"                 // Inject at build time with -ldflags "-X main.pinNumber=21"

var cli = lighthttpcli.NewCli(serverURL)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer rpio.Close()
	pinNumber, _ := strconv.Atoi(pinNumberStr)
	pin := rpio.Pin(pinNumber)
	pin.Output()
	stateChan := make(chan string)
	go pollLightState(stateChan)
	prevState := "OFF"
	pin.Low()
	for {
		state := <-stateChan
		if state != prevState {
			if state == "ON" {
				pin.High()
			} else {
				pin.Low()
			}
			prevState = state
		}
	}
}

func pollLightState(stateChan chan<- string) {
	for {
		state, _ := cli.GetState()
		stateChan <- state
		time.Sleep(500 * time.Millisecond)
	}
}
