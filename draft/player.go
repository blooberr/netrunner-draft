package draft

import (
  "fmt"
  "math/rand"

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

func (p *Player) PrintCorpPacks() {
  fmt.Printf("[ Player %s's corp cards ] \n", p.Name)

  for corpPackNumber, cardsInPack := range p.CorpPacks {
    fmt.Printf("Corp pack #%d: \n", corpPackNumber)
      for _, card := range cardsInPack {
        fmt.Printf("-- %s\n", card.Title)
      }
  }
  fmt.Printf("\n")
}

func (p *Player) PrintRunnerPacks() {
  fmt.Printf("[ Player %s's runner cards ] \n", p.Name)

  for runnerPackNumber, cardsInPack := range p.RunnerPacks {
    fmt.Printf("Runner pack #%d: \n", runnerPackNumber)
      for _, card := range cardsInPack {
        fmt.Printf("-- %s\n", card.Title)
      }
  }
  fmt.Printf("\n")
}

// PickRandomCard forces a player to select a card randomly. Called if the
// player takes forever or does it for the luls.
func (p *Player) PickRandomCard(packNumber int, isCorp bool) (selectedCard pool.Card){
  if isCorp {
    packLength := len(p.CorpPacks[packNumber])
    cardPosition := rand.Intn(packLength)
    selectedCard = p.CorpPacks[packNumber][cardPosition]

    p.CorpPacks[packNumber] = p.CorpPacks[packNumber][:cardPosition + copy(p.CorpPacks[packNumber][cardPosition:], p.CorpPacks[packNumber][cardPosition+1:])]
  } else {
    packLength := len(p.RunnerPacks[packNumber])
    cardPosition := rand.Intn(packLength)
    selectedCard = p.RunnerPacks[packNumber][cardPosition]

    p.RunnerPacks[packNumber] = p.RunnerPacks[packNumber][:cardPosition + copy(p.RunnerPacks[packNumber][cardPosition:], p.RunnerPacks[packNumber][cardPosition+1:])]
  }

  return selectedCard
}

