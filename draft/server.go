package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

type Server struct {
	Clients          []*Client
	AddClientChannel chan *Client
}

func NewServer() *Server {
	clients := make([]*Client, 0)
	addClientConsumer := make(chan *Client, 10)

	return &Server{Clients: clients, AddClientChannel: addClientConsumer}
}

func (s *Server) Launch() {
	fmt.Printf("Launch: running websocket server\n")

	onConnect := func(ws *websocket.Conn) {
		client := NewClient(ws)
		s.AddClientChannel <- client
		defer ws.Close()
	}

	http.Handle("/ws", websocket.Handler(onConnect))

	for {
		select {
		case newClient := <-s.AddClientChannel:
			s.Clients = append(s.Clients, newClient)
			fmt.Printf("%d are in the lobby.\n", len(s.Clients))
		}
	}
}
