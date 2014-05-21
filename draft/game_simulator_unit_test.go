package draft

import (
	"testing"
)

func TestNewGame(t *testing.T) {

	// define 4 players
	players := []*Player{&Player{Name: "Jedi Bear"},
		&Player{Name: "Star Fox"},
		&Player{Name: "Captain Falcon"},
		&Player{Name: "Hiphop Rex"},
	}

	g := NewGame(players)
	t.Logf("new game: %#v \n", g)

	if len(g.Players) != 4 {
		t.Errorf("Incorrect number of players! \n")
	} else {
    t.Logf("Starting game with correct number of players. \n")
  }

}
