package draft

import (
	"testing"

	"github.com/blooberr/netrunner-draft/pool"
)

func TestNewGame(t *testing.T) {

	players := []*Player{&Player{Name: "Jedi Bear", Id: 0},
		&Player{Name: "Star Fox", Id: 2},
		&Player{Name: "Captain Falcon", Id: 8},
		&Player{Name: "Hiphop Rex", Id: 3},
	}

	numPlayers := len(players)
	randomSeed := int64(12345)
	numberOfPacks := 4
	cardsPerPack := 10

	g := NewGame(randomSeed, numberOfPacks, cardsPerPack, players)
	t.Logf("New game: %#v \n", g)

	if len(g.Players) != numPlayers {
		t.Errorf("Incorrect number of players! \n")
	} else {
		t.Logf("Starting game with correct number of players. \n")
	}

	// Start with corp first

	direction := Left
	for packs := 0; packs < numberOfPacks; packs++ {
		g.BeginRound(pool.Corp)

		// simulate players drafting a card.  (Using force random)
		for turns := 0; turns < cardsPerPack; turns++ {
			for playerIndex, player := range g.Players {
				card := g.ForceRandom(playerIndex)
				t.Logf("player (%d) [%s] has been forced to randomly draft %s \n", playerIndex, player.Name, card.Title)
			}

			g.PassCards(direction)
			g.PrintCurrentPacks()
		}

		g.PrintDraftedCards()

		direction = !direction
	}

	// now with runner
	direction = Left
	for packs := 0; packs < numberOfPacks; packs++ {
		g.BeginRound(pool.Runner)

		// simulate players drafting a card.  (Using force random)
		for turns := 0; turns < cardsPerPack; turns++ {
			for playerIndex, player := range g.Players {
				card := g.ForceRandom(playerIndex)
				t.Logf("player (%d) [%s] has been forced to randomly draft %s \n", playerIndex, player.Name, card.Title)
			}

			g.PassCards(direction)
			g.PrintCurrentPacks()
		}

		g.PrintDraftedCards()

		direction = !direction
	}

}
