package draft

import (
	"container/list"
	"fmt"
)

type Game struct {
	Players      []*Player
	SeatingOrder list.List
}

func NewGame(players []*Player) (game *Game) {
	fmt.Printf("Starting new game with %d players.\n", len(players))

	g := &Game{Players: players}
	g.SeatPlayers()
	return g
}

func (g *Game) SeatPlayers() {

	for _, player := range g.Players {
		fmt.Printf("player: %#v \n", player)
		g.SeatingOrder.PushBack(player)
	}
}
