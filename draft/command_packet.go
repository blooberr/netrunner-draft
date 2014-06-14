package draft

type CommandPacket struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
  //Data map[string]interface{} `json:"data"`
}



