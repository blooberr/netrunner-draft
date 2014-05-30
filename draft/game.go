package draft

import (
	"fmt"
  "math/rand"

	"github.com/blooberr/netrunner-draft/pool"
)

type Game struct {
	Players []*Player
	Pool    *pool.Pool

	Round          int
	CurrentFaction pool.Faction

	NumberOfPacks int
	CardsPerPack  int

	CurrentPacks [][]pool.Card
}

func NewGame(seed int64, numberOfPacks int, cardsPerPack int, players []*Player) (game *Game) {
	fmt.Printf("Starting new game with %d players.\n", len(players))

	p := pool.InitPool(seed, "../data/cards.json")

	g := &Game{Players: players, Pool: p, NumberOfPacks: numberOfPacks, CardsPerPack: cardsPerPack}
	g.InitPlayers(numberOfPacks)

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

func (g *Game) BeginRound(faction pool.Faction) {
	fmt.Printf("Faction -- %s \n", faction)

	g.CurrentFaction = faction

	// create draft packs for everyone.
	for i := 0; i < len(g.Players); i++ {
		if faction == pool.Corp {
			g.CurrentPacks = append(g.CurrentPacks, g.Pool.GenerateCorpBooster(g.CardsPerPack))
		} else {
			g.CurrentPacks = append(g.CurrentPacks, g.Pool.GenerateRunnerBooster(g.CardsPerPack))
		}
	}

}

// ForceRandom forces a player to choose a card from the available booster
func (g *Game) ForceRandom(playerIndex int) (selectedCard pool.Card) {
	player := g.Players[playerIndex]

	packLength := len(g.CurrentPacks[playerIndex])
	cardPosition := rand.Intn(packLength)
	selectedCard = g.CurrentPacks[playerIndex][cardPosition]

	// remove card from position
	g.CurrentPacks[playerIndex] = g.CurrentPacks[playerIndex][:cardPosition+copy(g.CurrentPacks[playerIndex][cardPosition:], g.CurrentPacks[playerIndex][cardPosition+1:])]

	player.AddCard(selectedCard, g.CurrentFaction)

	return selectedCard
}

func (g *Game) PrintDraftedCards() {
	for _, player := range g.Players {
		player.PrintDraftedCards()
	}
}

// PassCards when this is called, everyone rotates hands.
// passing right is defined as shifting +1 (player order)
// passing left is defined as shifting -1

func (g *Game) PassCards(direction string, packNumber int, isCorp bool) {
	if direction == "l" {
		fmt.Printf("Passing left.\n")

		if isCorp == true {
		}

	} else {
		fmt.Printf("Passing right.\n")
	}

}

// RunDraft steps:
// players pick a card in their starting pack and pass in a direction.
func (g *Game) RunDraft() {
}
