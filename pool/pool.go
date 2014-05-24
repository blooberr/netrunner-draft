package pool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
)

type Pool struct {
	DataPath string
	Corp     []Card
	Runner   []Card
}

func InitPool(randSeed int64, pathToData string) *Pool {
  rand.Seed(randSeed)
	p := Pool{DataPath: pathToData}
	p.LoadCardPool(pathToData)
	return &p
}

func (p *Pool) LoadCardPool(pathToData string) {
  file, err := ioutil.ReadFile(pathToData)
  if err != nil {
    log.Fatalf("File error: %v \n", err)
    os.Exit(1)
  }

  corp, runner := ProcessFile(file)
  p.Corp = corp
  p.Runner = runner
}

func (p *Pool) GenerateCorpBooster(numCards int) (booster []Card){
  poolSize := len(p.Corp)
  booster = []Card{}
  for i := 0; i < numCards; i++ {
    index := rand.Intn(poolSize)
    card := p.Corp[index]
     booster = append(booster, card)
  }

  return booster
}

func (p *Pool) GenerateRunnerBooster(numCards int) (booster []Card){
  poolSize := len(p.Runner)
  booster = []Card{}
  for i := 0; i < numCards; i++ {
    index := rand.Intn(poolSize)
    card := p.Runner[index]
     booster = append(booster, card)
  }

  return booster
}

/*
func (p *Pool) GenerateRunnerBooster(numCards int) (booster []Card) {

}
*/
func GenerateCardPool(cardPoolSize int, cards []Card) (pool []Card) {
	originalPoolSize := len(cards)
	pool = []Card{}

	for i := 0; i < cardPoolSize; i++ {
		index := rand.Intn(originalPoolSize)
		card := cards[index]
		pool = append(pool, card)
	}

	return pool
}

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

