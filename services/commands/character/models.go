package character

//go:generate easytags $GOFILE

//Profession is the object for the professions with advance description
type Profession struct {
	Name                   string `yaml:"name"`
	Description            string
	BonusAdvances          []string
	AttributeBonusAdvances []string
	SkillAdvances          []string
	Talents                []string
	StartingMoney          map[string]string
	StartingEquip          [][]string
	UniqueTrait            Trait
}

//Talent is a modifier to a character with a special ability
type Trait struct {
	Name        string
	Description string
	Effect      string
}

type Equipment interface {
	GetValue() string
}

type Weapon struct {
	Name string
}

type Money struct {
	GoldCrowns
	SilverShillings
	CopperPennies
}
type GoldCrowns struct {
	Name      string
	ShortHand string
	Count     int
}
type SilverShillings struct {
	Name      string
	ShortHand string
	Count     int
}
type CopperPennies struct {
	Name      string
	ShortHand string
	Count     int
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
