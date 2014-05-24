package pool

type Card struct {
	LastModified    string `json:"last-modified"`
	Code            string `json:"code"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	TypeCode        string `json:"type_code"`
	Subtype         string `json:"subtype"`
	SubtypeCode     string `json:"subtype_code"`
	Text            string `json:"text"`
	BaseLink        int    `json:"baselink,omitempty"`
	Faction         string `json:"faction"`
	FactionCode     string `json:"faction_code"`
	FactionLetter   string `json:"faction_letter"`
	Flavor          string `json:"flavor"`
	Illustrator     string `json:"illustrator"`
	InfluenceLimit  int    `json:"influencelimit,omitempty"`
	MinimumDeckSize int    `json:"minimumdecksize,omitempty"`
	Number          int    `json:"number"`
	Quantity        int    `json:'quantity"`
	SetName         string `json:"setname"`
	SetCode         string `json:"set_code"`
	Side            string `json:"side"`
	SideCode        string `json:"side_code"`
	Uniqueness      bool   `json:"uniqueness"`
	CycleNumber     int    `json:"cyclenumber"`
	Url             string `json:"url"`
	ImageSrc        string `json:"imagesrc"`
	LargeImageSrc   string `json:"largeimagesrc,omitempty"`
}

// no need to generate identities within the pool. totally optional though.
var ExcludeTypeCode = [...]string{"identity"}

// removing special /a lternative art cards
var ExcludeSetCode = [...]string{"special", "alt"}

// removing 6 since lunar cycle isn't out yet
var ExcludeCycleNumber = [...]int{6}

func ExcludeCard(card Card) (result bool) {

  for _, value := range ExcludeTypeCode {
    if card.TypeCode == value {
      return true
    }
  }

  for _, value := range ExcludeSetCode {
    if card.SetCode == value {
      return true
    }
  }

  for _, value := range ExcludeCycleNumber {
    if card.CycleNumber == value {
      return true
    }
  }
  return false
}

