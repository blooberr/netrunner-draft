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
  g.BeginRound(pool.Corp)
  for index, cards := range g.CurrentPacks {
    t.Logf("player [%d] starts with: \n", index)
    for _, card := range cards {
      t.Logf("[%d] card - %s \n", index, card.Title)
    }
  }

  // simulate players drafting a card.  (Using force random)
  for playerIndex, player := range g.Players {
    card := g.ForceRandom(playerIndex)
    t.Logf("player (%d) [%s] has been forced to randomly draft %s \n", playerIndex, player.Name, card.Title)
  }

/*
  // everyone randomly drafts a card
  for _, player := range g.Players {
    card := player.PickRandomCard(packNumber, corpRound)
    //t.Logf("card remaining: %#v \n", player.CorpPacks[0])
    t.Logf("\npicked -- %#v \n", card.Title)
    t.Logf("cards left: \n")
    for _, card := range player.CorpPacks[packNumber] {
      t.Logf("-- %s \n", card.Title)
    }
    player.AddCard(card)
  }
*/

/*
	packNumber := 0
  corpRound := true
	for _, player := range g.Players {
		card := player.PickRandomCard(packNumber, corpRound)
		//t.Logf("card remaining: %#v \n", player.CorpPacks[0])
		t.Logf("\npicked -- %#v \n", card.Title)
		t.Logf("cards left: \n")
		for _, card := range player.CorpPacks[packNumber] {
			t.Logf("-- %s \n", card.Title)
		}
		player.AddCard(card)
	}

	g.PrintDraftedCards()
*/

	// g.PassCards("r", packNumber, corpRound)
}
