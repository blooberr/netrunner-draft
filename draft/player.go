package draft

type Player struct {
	Name string
  Id int
  CardsDrafted map[string]int // strCode -> number of items
  CardsInHand map[string]int

}
