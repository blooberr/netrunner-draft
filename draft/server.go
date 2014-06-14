package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Server struct {
	Clients             []*Client
	AddClientChannel    chan *Client
	RemoveClientChannel chan *Client
	CommandChannel      chan CommandPacket
	ReadyClientChannel  chan *Client
	AutoIncId           int
	l                   sync.Mutex
}

func NewServer() *Server {
	clients := make([]*Client, 0)

	addClientConsumer := make(chan *Client, 10)
	removeClient := make(chan *Client, 10)
	cmdChannel := make(chan CommandPacket, 10)
	rdyClientChan := make(chan *Client, 10)

	return &Server{Clients: clients,
		AddClientChannel:    addClientConsumer,
		RemoveClientChannel: removeClient,
		CommandChannel:      cmdChannel,
		ReadyClientChannel:  rdyClientChan,
	}
}

func (s *Server) RemoveClient(client *Client) {
	s.RemoveClientChannel <- client
}

func (s *Server) Launch() {
	fmt.Printf("Launch: running websocket server\n")

	// onConnect handles a new client connecting to the server.
	onConnect := func(ws *websocket.Conn) {

		//playerName := fmt.Sprintf("player-%d", s.AutoIncId)
		s.l.Lock()
		s.AutoIncId = s.AutoIncId + 1
		client := NewClient(ws, s, s.AutoIncId)
    s.l.Unlock()

		s.AddClientChannel <- client
		client.Launch()

		defer ws.Close()
	}

	http.Handle("/ws", websocket.Handler(onConnect))

	heartBeat := time.NewTicker(time.Second * 30)

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

      s.LobbyRefresh()
		case cmd := <-s.CommandChannel:
			fmt.Printf("command channel - %#v \n", cmd)
			s.ProcessCommandPacket(cmd)

		case c := <-s.ReadyClientChannel:
			fmt.Printf("player %s is ready. \n", c.Player.Name)
      c.Player.SetReady(true)

      for _, client := range s.Clients {
        fmt.Printf("player %s ready: %s \n", client.Player.Name, client.Player.Ready)
      }
			// once we have all n players ready, we can begin the draft.
		}
	}
}

// ClientReady handles if the client is ready to begin the draft.
func (s *Server) ReadyClient() chan<- *Client {
	return (chan<- *Client)(s.ReadyClientChannel)
}

func (s *Server) ProcessCommandPacket(cp CommandPacket) {
	switch cp.Event {
	case "Ready":
		fmt.Printf("are you ready ready ready\n")
	}
}

// Broadcast sends to all clients connected
func (s *Server) Broadcast(cp CommandPacket) {
	for id, client := range s.Clients {
		fmt.Printf("[%d] %#v \n", id, client)
		client.Write(cp)
	}
}

func (s *Server) LobbyRefresh() {
	lobbyRefreshMap := make(map[string][]string)
	names := []string{}

	for _, client := range s.Clients {
		names = append(names, client.Player.Name)
	}

	sort.Strings(names)

	for _, name := range names {
		lobbyRefreshMap["players"] = append(lobbyRefreshMap["players"], name)
	}

	cp := CommandPacket{Event: "Update Lobby", Data: lobbyRefreshMap}
	s.Broadcast(cp)
}
