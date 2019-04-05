package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func echo(w http.ResponseWriter, r *http.Request) {
	con, _ := upgrader.Upgrade(w, r, nil)
	for {
		// read message from browser
		msgType, msg, err := con.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("%s sent: %s\n", con.RemoteAddr(), string(msg))

		// write the message back to the browser
		if err = con.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
