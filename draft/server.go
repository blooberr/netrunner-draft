package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
  "time"
)

type Server struct {
	Clients          []*Client
	AddClientChannel chan *Client
  RemoveClientChannel chan *Client
  AutoIncId int
}

func NewServer() *Server {
	clients := make([]*Client, 0)

	addClientConsumer := make(chan *Client, 10)
  removeClient := make(chan *Client, 10)

	return &Server{Clients: clients,
    AddClientChannel: addClientConsumer,
    RemoveClientChannel: removeClient,
  }

}

func (s *Server) RemoveClient(client *Client) {
  s.RemoveClientChannel <- client
}

func (s *Server) Launch() {
	fmt.Printf("Launch: running websocket server\n")

	onConnect := func(ws *websocket.Conn) {
		// numClients := len(s.Clients)
		playerName := fmt.Sprintf("player-%d", s.AutoIncId)
    s.AutoIncId = s.AutoIncId + 1

		client := NewClient(ws, s, playerName)
		s.AddClientChannel <- client
		client.Launch()

		defer ws.Close()
	}

	http.Handle("/ws", websocket.Handler(onConnect))

  heartBeat := time.NewTicker(time.Second * 5)

	for {
		select {
    case <-heartBeat.C:
      fmt.Printf("Send out heartbeat pulse to all clients\n")
      heartBeatMap := make(map[string]string)
      heartBeatMap["heart"] = "beat"
      cp := CommandPacket{Event: "HeartBeat", Data: heartBeatMap}

      s.Broadcast(cp)
    case delClient := <-s.RemoveClientChannel:
      fmt.Printf("Removing client: %#v \n", delClient)

      for i := range s.Clients {
        if s.Clients[i] == delClient {
          s.Clients = append(s.Clients[:i], s.Clients[i+1:]...)
          break
        }
      }

		case newClient := <-s.AddClientChannel:
			fmt.Printf("Adding new client: %#v \n", newClient)

			newNameMap := make(map[string]string)
			newNameMap["name"] = newClient.Player.Name
			cp := CommandPacket{Event: "New Player", Data: newNameMap}

			fmt.Printf("sending on writechan: %#v\n", cp)
			newClient.Write(cp)

			s.Clients = append(s.Clients, newClient)
			fmt.Printf("%d are in the lobby.\n", len(s.Clients))

			cp = CommandPacket{Event: "New Player Joined", Data: newNameMap}
			s.Broadcast(cp)

      lobbyRefreshMap := make(map[string][]string)

      for _, client := range s.Clients {
        lobbyRefreshMap["players"] = append(lobbyRefreshMap["players"], client.Player.Name)
      }

      cp = CommandPacket{Event: "Update Lobby", Data: lobbyRefreshMap}
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
