package draft

import (
	"fmt"

  "github.com/blooberr/netrunner-draft/pool"
)

type Game struct {
	Players      []*Player
  Pool         *pool.Pool
}

func NewGame(seed int64, numberOfPacks int, cardsPerPack int, players []*Player) (game *Game) {
	fmt.Printf("Starting new game with %d players.\n", len(players))

  p := pool.InitPool(seed, "../data/cards.json")

	g := &Game{Players: players, Pool: p}
	g.InitPlayers(numberOfPacks)
  g.CreateDraftPacks(numberOfPacks, cardsPerPack)

	return g
}

// InitPlayers is a wrapper to call player.InitPlayer() on all players
func (g *Game) InitPlayers(numPacks int) {
	for _, player := range g.Players {
    player.InitPlayer(numPacks)
	}
}

func (g *Game) CreateDraftPacks(numPacksPerSide int,
  cardsPerPack int) {

  for _, player := range g.Players {
    for pack := 0; pack < numPacksPerSide; pack++ {

      //fmt.Printf("CreateDraftPacks: %#v \n", player)
      player.CorpPacks = append(player.CorpPacks, g.Pool.GenerateCorpBooster(cardsPerPack))
      player.RunnerPacks = append(player.RunnerPacks, g.Pool.GenerateRunnerBooster(cardsPerPack))
    }
  }

}

// RunDraft steps:
// players pick a card in their starting pack and pass in a direction.
func (g *Game) RunDraft() {
}
