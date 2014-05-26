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

/*
  // draft pass 1
  // set beginnig direction
  round := 0

  for playerId, player := range players {
    cardInBoosterNum := 0

    pickedCard := player.CorpPacks[round][cardInBoosterNum]

    t.Logf("[%d] %s picked Card: %#v \n", playerId, player.Name, pickedCard)
  }
*/

}

