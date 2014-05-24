package pool

import (
  "testing"
)

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

