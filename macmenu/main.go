package main

import (
	"fmt"
	"time"

	"github.com/caseymrm/menuet"
	"github.com/stanleynguyen/go-everywhere/lighthttpcli"
)

var serverURL = "http://localhost:8080" // Inject at build time with -ldflags "-X main.serverURL=http://something"

var cli = lighthttpcli.NewCli(serverURL)

func intervalStateRefresh() {
	ticker := time.NewTicker(500 * time.Millisecond)
	var prevState string
	for {
		<-ticker.C
		state, _ := cli.GetState()
		if state != prevState {
			menuet.App().SetMenuState(&menuet.MenuState{
				Title: fmt.Sprintf("Light is: %s", state),
			})
			prevState = state
		}
	}
}

func menuItems() []menuet.MenuItem {
	onBtn := menuet.MenuItem{
		Text: "Turn On",
		Clicked: func() {
			cli.SetState(lighthttpcli.StateOn)
		},
	}
	offBtn := menuet.MenuItem{
		Text: "Turn Off",
		Clicked: func() {
			cli.SetState(lighthttpcli.StateOff)
		},
	}
	return []menuet.MenuItem{onBtn, offBtn}
}

func main() {
	go intervalStateRefresh()
	menuet.App().Label = "com.github.stanleynguyen.goeverywhere"
	menuet.App().Children = menuItems
	menuet.App().RunApplication()
}
