package pool

import(
  "encoding/json"
  "log"
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "sort"
)

// GeneratePool pseudo-randomly generates a new pool of cards of size
// cardPoolSize
func GeneratePool(cardPoolSize int, cards []Card) (pool map[string]int) {
	originalPoolSize := len(cards)

	pool = make(map[string]int)

	for i := 0; i < cardPoolSize; i++ {
		index := rand.Intn(originalPoolSize)
		card := cards[index]
		cardTitle := card.Title
		numItems := pool[cardTitle]

		pool[cardTitle] = numItems + 1
	}

	return pool
}

func SortCards(cards map[string]int) (sortedCardNames []string) {
	cardNames := make([]string, len(cards))

	i := 0
	for cardName, _ := range cards {
		cardNames[i] = cardName
		i++
	}

	sort.Strings(cardNames)
	return cardNames
}

func ProcessFile(file []byte) (corp []Card, runner []Card) {
	// difference between make and new.
	raw := make([]json.RawMessage, 10)
	if err := json.Unmarshal(file, &raw); err != nil {
		log.Fatalf("error %v \n", err)
		os.Exit(1)
	}

	corp = make([]Card, 0)
	runner = make([]Card, 0)

	for i := 0; i < len(raw); i++ {
		card := Card{}
		if err := json.Unmarshal(raw[i], &card); err != nil {
			log.Fatalf("error %v \n", err)
			os.Exit(1)
		}

		// fmt.Printf("Card: %#v\n", card)

		// generate corp / runner lists
		if card.Side == "Corp" {
			if !ExcludeCard(card) {
				corp = append(corp, card)
			}
		}

		if card.Side == "Runner" {
			if !ExcludeCard(card) {
				runner = append(runner, card)
			}
		}
	}

	fmt.Printf("Number of corp cards: %d \n", len(corp))
	fmt.Printf("Number of runner cards: %d \n", len(runner))
	return corp, runner
}

func LoadPool(pathToCards string) (corp []Card, runner []Card) {
  file, err := ioutil.ReadFile(pathToCards)
  if err != nil {
    log.Fatalf("File error: %v \n", err)
    os.Exit(1)
  }

  corp, runner = ProcessFile(file)
  return corp, runner
}

func SetSeed(randSeed int64) {
  rand.Seed(randSeed)
}

type PlayerPacks struct {
  Corp []map[string]int
  Runner []map[string]int
}

func CreateDraftPacks(numPlayers int, numPacks int, cardsPerDeck int, pathsToCards string) (playerPools []PlayerPacks) {

  // load pool into memory
  corp, runner := LoadPool(pathsToCards)

  for i := 0; i < numPlayers; i++ {
    fmt.Printf("generating draft pack for player %d \n", i)

    corpPacks := []map[string]int{}
    runnerPacks := []map[string]int{}

    for n := 0; n < numPacks; n++ {
      newCorpPacks := GeneratePool(cardsPerDeck, corp)
      newRunnerPacks := GeneratePool(cardsPerDeck, runner)

      corpPacks = append(corpPacks, newCorpPacks)
      runnerPacks = append(runnerPacks, newRunnerPacks)
    }

    pp := PlayerPacks{Corp: corpPacks, Runner: runnerPacks}
    playerPools = append(playerPools, pp)
  }

  return playerPools
}

