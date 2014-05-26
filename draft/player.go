package draft

import (
  "fmt"

	"github.com/blooberr/netrunner-draft/pool"
)

type Player struct {
	Name         string
	Id           int
	CardsDrafted map[string]int // strCode -> number of items
	CardsInHand  map[string]int

	CorpPacks   [][]pool.Card
	RunnerPacks [][]pool.Card
}

func (p *Player) InitPlayer(numPacks int) {
  fmt.Printf("num packs:  %d \n", numPacks)
  //p.CorpPacks = make([][]pool.Card, numPacks)
  //fmt.Printf("corp packs: %#v \n", p.CorpPacks)
}
