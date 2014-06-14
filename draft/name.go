package draft

import(
  "fmt"
  "math/rand"
)

// Generates an awesome random name

var NamePrefix = [...]string{
  "Legendary",
  "Captivating",
  "Consummate",
  "Liberated",
  "Dispossessed",
  "Honey",
  "Jedi",
  "Punitive",
  "Project",
  "Biotic",
  "National",
  "Professional",
  "Corporate",
  "Private",
  "Quality",
  "Stout",
  "Star",
  "Inside",
  "Pop-up",
  "Future",
  "Fetal",
  "Same Old",
  "Acidic",
  "Chaotic",
  "Magnum",
  "Replicating",
  "Self-modifying",
  "Light",
  "Ghost",
  "Cheesy",
  "Dark",
  "Lucky",
  "Archived",
  "Prepaid",
  "Enormous",
  "Subliminal",
  "Clone",
  "Sea",
  "Scorched",
  "Cyber",
  "Special",
  "Happy",
  "Dirrty",
  "Lightning",
  "Hammer",
  "Grumpy",
  "Mighty",
  "Compromised",
  "Giant",
  "Robot",
  "Stealthy",
  "Million Cred",
  "Mr.",
  "Ms.",
  "Uncle",
}

var NameSuffix = [...]string{
  "Bear",
  "Fox",
  "Wolf",
  "Duck",
  "Horse",
  "Jeffrey",
  "Tiger",
  "Cheetah",
  "Snowflake",
  "Russian",
  "Cheater",
  "Ichi",
  "Viktor",
  "Scrub",
  "Window",
  "Kati",
  "Hero",
  "Tollbooth",
  "Sensei",
  "Badger",
  "Peacock",
  "Sharpshooter",
  "Stirling",
  "Faerie",
  "Ristie",
  "Barn",
  "Ninja",
  "Eli",
  "Mustang",
  "Li",
  "Viper",
  "Andy",
  "Gabe",
  "Professor",
  "Kit",
  "Kate",
}

func CreateName() string {
  prefix := NamePrefix[rand.Intn(len(NamePrefix))]
  suffix := NameSuffix[rand.Intn(len(NameSuffix))]

  return fmt.Sprintf("%s %s", prefix, suffix)
}
