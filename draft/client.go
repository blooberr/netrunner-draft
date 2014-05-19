package draft

import(
  "code.google.com/p/go.net/websocket"
  "fmt"
)

type Client struct {
  Ws *websocket.Conn
}

func NewClient(ws *websocket.Conn) *Client {
  fmt.Printf("NewClient: %#v \n", ws)
  return &Client{Ws: ws}
}

