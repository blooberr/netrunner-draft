package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

type Command struct {
	Event string
}

func ReceiveCommand(ws *websocket.Conn) {
	var command Command
	err := websocket.JSON.Receive(ws, &command)
	if err != nil {
		fmt.Println("Error receiving command.")
	} else {
		fmt.Println("Event: " + command.Event)
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.Handle("/ws", websocket.Handler(ReceiveCommand))

	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatal("listenandserve:", err)
	}
}
