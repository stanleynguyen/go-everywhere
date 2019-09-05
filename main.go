package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ON indicate that light is on, can be sent to client
const ON = "ON"

// OFF indicate that light is off, can be sent to client
const OFF = "OFF"

type connPresent struct{}

func main() {
	light := OFF
	clients := map[*websocket.Conn]connPresent{}

	http.HandleFunc("/led", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		defer func() {
			delete(clients, conn)
		}()

		// send initial state of led until client gets it
		retry := 5
		for {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(light)); err == nil {
				clients[conn] = connPresent{}
				break
			}

			if retry > 0 {
				retry--
			} else {
				conn.Close()
				return
			}
		}

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				return
			}

			if light != OFF {
				light = OFF
			} else {
				light = ON
			}

			for c := range clients {
				if err = c.WriteMessage(websocket.TextMessage, []byte(light)); err != nil {
					log.Println(err.Error())
				}
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
