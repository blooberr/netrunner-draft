package draft

import(
  "fmt"
)

type Game struct {
	Players []*Player
}

func NewGame(players []*Player) (game *Game) {
  fmt.Printf("Starting new game with %d players.\n", len(players))

  return &Game{Players: players}
}


