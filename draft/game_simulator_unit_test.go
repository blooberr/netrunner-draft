package draft

import (
	"testing"
)

func TestNewGame(t *testing.T) {

	// define 4 players
	players := []*Player{&Player{Name: "Jedi Bear", Id: 0},
		&Player{Name: "Star Fox", Id: 2},
		&Player{Name: "Captain Falcon", Id: 8},
		&Player{Name: "Hiphop Rex", Id: 3},
	}

	g := NewGame(players)

	t.Logf("new game: %#v \n", g)

	if len(g.Players) != 4 {
		t.Errorf("Incorrect number of players! \n")
	} else {
		t.Logf("Starting game with correct number of players. \n")
	}

	// seating order check
	for i := g.SeatingOrder.Front(); i != nil; i = i.Next() {
		//front := g.SeatingOrder.Front()
		t.Logf("front of the list: %#v \n", i.Value.(*Player).Name)
	}
}
