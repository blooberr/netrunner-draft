package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Client struct {
	Ws          *websocket.Conn
	Player      *Player
	server      *Server
	WriteChan   chan CommandPacket
	DoneChannel chan bool
}

func NewClient(ws *websocket.Conn, s *Server, id int) *Client {
	fmt.Printf("NewClient: %#v \n", ws)
	writeChan := make(chan CommandPacket, 10)
	doneChannel := make(chan bool, 10)

	//newPlayer := &Player{Name: playerName}
  newPlayer := NewPlayer(id)

	return &Client{Ws: ws, server: s, Player: newPlayer, WriteChan: writeChan, DoneChannel: doneChannel}
}

func (c *Client) Launch() {
	go c.listenRead()
	go c.listenWrite()
	c.listenTerminate()
}

func (c *Client) listenTerminate() {
	for {
		select {
		case <-c.DoneChannel:
			fmt.Printf("listenTerminate(): terminating %#v \n", c)
			c.server.RemoveClient(c)
			return
		}
	}
}

// Write - lazy short cut way.  I don't think this actully saves any time, but
// will feel more intuitive
func (c *Client) Write(cp CommandPacket) {
	c.WriteChan <- cp
}

func (c *Client) listenWrite() {
	fmt.Printf("client.listenwrite \n")
	for {
		select {
		// send message to the client
		case msg := <-c.WriteChan:
			fmt.Printf("receive: %#v \n", msg)
			err := websocket.JSON.Send(c.Ws, msg)
			if err != nil {
				fmt.Printf("error: %#v \n", err)
				c.Terminate()
			}
		}
	}
}

func (c *Client) Terminate() {
	fmt.Printf("calling terminate ! \n")
	c.DoneChannel <- true
}

func (c *Client) listenRead() {
	fmt.Printf("listenRead \n")
	for {
		select {
		// read data from websocket connection
		default:
			var cp CommandPacket
			err := websocket.JSON.Receive(c.Ws, &cp)
			if err != nil {
				fmt.Printf("error read: %s \n", err)
				c.Terminate()
			} else {
				fmt.Printf("listnRead: %#v \n", cp)
        c.ProcessCommand(cp)
			}
		}
	}
}

func (c *Client) ProcessCommand(cp CommandPacket) {
  event := cp.Event
  fmt.Printf("event -- %s \n", event)
  fmt.Printf("data -- %#v \n", cp.Data)

  fmt.Printf("%s \n", cp.Data.(map[string]interface{}))

  //data := cp.Data

  if event == "NewName" {
    res := cp.Data.(map[string]interface{})
    fmt.Printf("%#v \n", res)
    fmt.Printf("%s \n", res["name"].(string))

    c.Player.SetName(res["name"].(string))

    newNameMap := make(map[string]string)
    newNameMap["name"] = c.Player.Name
    cp := CommandPacket{Event: "New Name", Data: newNameMap}
    c.Write(cp)
    c.server.LobbyRefresh()
  }

  if event == "Ready" {
    c.server.ReadyClient() <- c
  }

  //c.server.CommandChannel <- cp

}
