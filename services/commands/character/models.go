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

type InventoryItem interface {
	GetValue() string
	GetCount() int
}

type Item struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	WeightLbs   int    `json:"weight_lbs"`
}

type MeleeWeapon struct {
	Item
	ReachInch     int           `json:"reach_inch"`
	InitBonusNear int           `json:"init_bonus_near"`
	InitBonusFar  int           `json:"init_bonus_far"`
	HandsUsed     []int         `json:"hands_used"`
	MeleeType     string        `json:"melee_type"`
	WeaponTraits  []WeaponTrait `json:"weapon_traits"`
}

type RangeWeapon struct {
	Item
	OptimalRangeYds int    `json:"optimal_range_yds"`
	FallOffYds      int    `json:"fall_off_yds"`
	MaxRangeYds     int    `json:"max_range_yds"`
	AmmunitionType  string `json:"ammunition_type"`
	ReloadTimeMS    int    `json:"reload_time_ms"`
	DrawTimeMS      int    `json:"draw_time_ms"`
}

type WeaponTrait struct {
	Name   string `json:"name"`
	Effect string `json:"effect"`
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
