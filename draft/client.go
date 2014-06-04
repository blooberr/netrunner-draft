package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Client struct {
	Ws        *websocket.Conn
	Player    *Player
	WriteChan chan CommandPacket
}

func NewClient(ws *websocket.Conn, playerName string) *Client {
	fmt.Printf("NewClient: %#v \n", ws)
	writeChan := make(chan CommandPacket, 10)

	newPlayer := &Player{Name: playerName}

	return &Client{Ws: ws, Player: newPlayer, WriteChan: writeChan}
}

func (c *Client) Launch() {
  go c.listenRead()
	c.listenWrite()
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
      }
		}
	}
}

func (c *Client) listenRead() {
  fmt.Printf("listen on reads \n")
  for {
    select {

    // read data from websocket connection
    default:
      var cp CommandPacket
      err := websocket.JSON.Receive(c.Ws, &cp)
      if err != nil {
        fmt.Printf("error read: %s \n", err)
      } else {
        fmt.Printf("on read : %#v \n", cp)
      }
    }
  }
}
