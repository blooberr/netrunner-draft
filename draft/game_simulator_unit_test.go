package draft

import (
	"testing"
)

func TestNewGame(t *testing.T) {

	// define 4 players for now
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

  t.Logf("Starting packs for all players: \n")
  for _, player := range g.Players {
    player.PrintCorpPacks()
    player.PrintRunnerPacks()
  }

  packNumber := 0
  for _, player := range g.Players {
    card := player.PickRandomCard(packNumber, true)
    //t.Logf("card remaining: %#v \n", player.CorpPacks[0])
    t.Logf("\npicked -- %#v \n", card.Title)
    t.Logf("cards left: \n")
    for _, card := range player.CorpPacks[packNumber] {
      t.Logf("-- %s \n", card.Title)
    }

  }
}

