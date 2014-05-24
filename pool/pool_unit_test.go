package pool

import (
	"testing"
)

/*
func TestCreatePool(t *testing.T) {
  numPlayers := 4
  cardsPerPack := 10
  numPacksPerSide := 4

  SetSeed(12345)
  playerPools := CreateDraftPacks(numPlayers, numPacksPerSide, cardsPerPack, "../data/cards.json")

  for id, player := range playerPools {
    //t.Logf("player %d - %#v \n", id, player)
    for packId, corpPack := range player.Corp {
      t.Logf("Player [%d] Corp: %d - %#v \n", id, packId, corpPack)
    }

    for packId, runPack := range player.Runner {
      t.Logf("Player [%d] Runner: %d - %#v \n", id, packId, runPack)
    }

    t.Logf("*********************** \n")
  }
}
*/

func TestCreateCardPool(t *testing.T) {
	numPlayers := 4
	cardsPerPack := 10
	numPacksPerSide := 4

	p := InitPool(12345, "../data/cards.json")

	for i := 0; i < numPlayers; i++ {
		t.Logf("printing for player %d \n", i)
		for boosterNum := 0; boosterNum < numPacksPerSide; boosterNum++ {
			cb := p.GenerateCorpBooster(cardsPerPack)

			t.Logf("[[ corp ]] \n")
			for cardIndex, card := range cb {
				t.Logf("booster [%d] - (%d) %s \n", boosterNum, cardIndex, card.Title)
			}
			t.Logf("***** \n")
		}

		for boosterNum := 0; boosterNum < numPacksPerSide; boosterNum++ {
			cb := p.GenerateRunnerBooster(cardsPerPack)

			t.Logf("[[ runner ]]\n")
			for cardIndex, card := range cb {
				t.Logf("booster [%d] - (%d) %s \n", boosterNum, cardIndex, card.Title)
			}
			t.Logf("***** \n")
		}
	}

}
