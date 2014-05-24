package draft

import (
	"testing"
)

func TestNewGame(t *testing.T) {

	// define 4 players for now
	numPlayers := 4

	players := []*Player{&Player{Name: "Jedi Bear", Id: 0},
		&Player{Name: "Star Fox", Id: 2},
		&Player{Name: "Captain Falcon", Id: 8},
		&Player{Name: "Hiphop Rex", Id: 3},
	}

	g := NewGame(players)

	t.Logf("new game: %#v \n", g)

	if len(g.Players) != numPlayers {
		t.Errorf("Incorrect number of players! \n")
	} else {
		t.Logf("Starting game with correct number of players. \n")
	}

	// seating order check
	for i := g.SeatingOrder.Front(); i != nil; i = i.Next() {
		t.Logf("front of the list: %#v \n", i.Value.(*Player).Name)
	}

	// generate card packs
	pp := g.CreateDraftPacks(12345,
		len(players),
		4,
		10,
		"../data/cards.json")

  // load player struct with each individual booster pack.
	for id, player := range pp {
    players[id].CorpStartingPack = player.Corp
    players[id].RunnerStartingPack = player.Runner

		for packId, corpPack := range player.Corp {
			t.Logf("Player [%d] Corp: %d - %#v \n", id, packId, corpPack)
		}

		for packId, runPack := range player.Runner {
			t.Logf("Player [%d] Runner: %d - %#v \n", id, packId, runPack)
		}
	}

  t.Logf("Starting packs. \n")
  for _, player := range players {
    t.Logf("player %s - %#v \n", player.Name, player.CorpStartingPack[0])
    t.Logf("player %s - %#v \n", player.Name, player.RunnerStartingPack[0])
  }
}
