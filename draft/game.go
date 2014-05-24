package draft

import (
	"container/list"
	"fmt"
  "github.com/blooberr/netrunner-draft/pool"
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

func (g *Game) CreateDraftPacks(seed int64,
  numPlayers int,
  numPacksPerSide int,
  cardsPerPack int,
  dataPath string) (playerPools []pool.PlayerPacks) {

  pool.SetSeed(seed)
  playerPools = pool.CreateDraftPacks(numPlayers, numPacksPerSide, cardsPerPack, dataPath)

  // fmt.Printf("%#v \n", playerPools)
  return playerPools
}

