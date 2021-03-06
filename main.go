package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rakyll/statik/fs"
	_ "github.com/stanleynguyen/go-everywhere/statik"
)

type lightState int

func (s lightState) string() string {
	switch s {
	case on:
		return "ON"
	case off:
		fallthrough
	default:
		return "OFF"
	}
}

const (
	on = lightState(iota)
	off
)

type light struct {
	state lightState
	mux   sync.Mutex
}

func newLight() light {
	return light{
		state: off,
	}
}

func (l *light) setState(s lightState) {
	l.mux.Lock()
	l.state = s
	l.mux.Unlock()
}

func (l *light) getState() lightState {
	return l.state
}

func main() {
	light := newLight()
	http.HandleFunc("/led", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Method not allowed")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, light.getState().string())
	})
	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Method not allowed")
			return
		}

		w.WriteHeader(http.StatusOK)
		light.setState(on)
		fmt.Fprint(w, "")
	})
	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Method not allowed")
			return
		}

		w.WriteHeader(http.StatusOK)
		light.setState(off)
		fmt.Fprint(w, "")
	})

	// build wasm file
	err := os.Setenv("GOARCH", "wasm")
	if err != nil {
		log.Fatal("error setting env", err.Error())
	}
	err = os.Setenv("GOOS", "js")
	if err != nil {
		log.Fatal("error setting env", err.Error())
	}

	// serve static site
	statikFS, err := fs.New()
	fileSys := http.FileServer(statikFS)
	http.Handle("/", fileSys)

	http.ListenAndServe(":8080", nil)
}
