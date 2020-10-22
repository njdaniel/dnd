package character

//go:generate easytags $GOFILE json

//Profession is the object for the professions with advance description
type Profession struct {
	Name                   string            `json:"name"`
	Description            string            `json:"description"`
	BonusAdvances          []string          `json:"bonus_advances"`
	AttributeBonusAdvances []string          `json:"attribute_bonus_advances"`
	SkillAdvances          []string          `json:"skill_advances"`
	Talents                []string          `json:"talents"`
	StartingMoney          map[string]string `json:"starting_money"`
	StartingEquip          [][]string        `json:"starting_equip"`
	UniqueTrait            Trait             `json:"unique_trait"`
}

//Talent is a modifier to a character with a special ability
type Trait struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
}

type Equipment interface {
	GetValue() string
}

type Weapon struct {
	Name string `json:"name"`
}

type Money struct {
	GoldCrowns
	SilverShillings
	CopperPennies
}
type GoldCrowns struct {
	Name      string `json:"name"`
	ShortHand string `json:"short_hand"`
	Count     int    `json:"count"`
}
type SilverShillings struct {
	Name      string `json:"name"`
	ShortHand string `json:"short_hand"`
	Count     int    `json:"count"`
}
type CopperPennies struct {
	Name      string `json:"name"`
	ShortHand string `json:"short_hand"`
	Count     int    `json:"count"`
}

func NewMoney() Money {

	return Money{
		GoldCrowns{
			Name:      "GoldCrowns",
			ShortHand: "gc",
			Count:     0,
		}, SilverShillings{
			Name:      "SilverShillings",
			ShortHand: "ss",
			Count:     0,
		}, CopperPennies{
			Name:      "CopperPennies",
			ShortHand: "cp",
			Count:     0,
		},
	}
}

type CoinTypes int

const (
	GoldCrown CoinTypes = iota
	SilverShilling
	CopperPenny
)

func (c CoinTypes) String() string {
	return [...]string{"GoldCrowns", "SilverShillings", "CopperPennies"}[c]
}
func (c CoinTypes) Shorthand() string {
	return [...]string{"gc", "ss", "cp"}[c]
}
