package draft

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Client struct {
	Ws         *websocket.Conn
	PlayerName string
	WriteChan  chan CommandPacket
}

func NewClient(ws *websocket.Conn, playerName string) *Client {
	fmt.Printf("NewClient: %#v \n", ws)
	writeChan := make(chan CommandPacket, 10)
	return &Client{Ws: ws, PlayerName: playerName, WriteChan: writeChan}
}

func (c *Client) Launch() {
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
			websocket.JSON.Send(c.Ws, msg)
		}
	}
}
