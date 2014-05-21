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
		numClients := len(s.Clients)
		playerName := fmt.Sprintf("player-%d", numClients)

		client := NewClient(ws, playerName)
		s.AddClientChannel <- client
    client.Launch()

		defer ws.Close()
	}

	http.Handle("/ws", websocket.Handler(onConnect))

	for {
		select {
		case newClient := <-s.AddClientChannel:
			fmt.Printf("Adding new client: %#v \n", newClient)

			newNameMap := make(map[string]string)
			newNameMap["name"] = newClient.Player.Name
      cp := CommandPacket{Event: "New Player", Data: newNameMap}

			fmt.Printf("sending on writechan: %#v\n", cp)
			newClient.Write(cp)

			s.Clients = append(s.Clients, newClient)
			fmt.Printf("%d are in the lobby.\n", len(s.Clients))

      cp = CommandPacket{Event: "New Player Joined", Data: newNameMap }
      s.Broadcast(cp)
		}
	}
}

// Broadcast sends to all clients connected
func (s *Server) Broadcast(cp CommandPacket) {
  for id, client := range s.Clients {
    fmt.Printf("[%d] %#v \n", id, client)
    client.Write(cp)
  }
}
