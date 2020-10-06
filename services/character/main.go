package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("starting...")
	fmt.Println(NewCharacter())
	fmt.Println("done.")
}

type Gender int

const (
	male = iota + 1
	female
)

func (g Gender) String() string {
	return [...]string{"male", "female"}[g]
}

type Race int

const (
	Human = iota + 1
	Elf
	Dwarf
)

func (r Race) String() string {
	return [...]string{"Human", "Elf", "Dwarf"}[r]
}

type DwarvenHeritage int

const (
	IronHills = iota + 1
	Mountain
	Deep
)

func (d DwarvenHeritage) String() string {
	return [...]string{"IronHills", "Mountain", "Deep"}[d]
}

type ElvenHeritage int

const (
	High = iota + 1
	Drow
	Wood
)

func (e ElvenHeritage) String() string {
	return [...]string{"High", "Drow", "Wood"}[e]
}

type HumanHeritage int

const (
	Taldan = iota + 1
	Ulfen
	Varisan
	Vudrani
	Nidalese
	Keleshite
)

func (h HumanHeritage) String() string {
	return [...]string{"Taldan", "Ulfen", "Varisan", "Vudrani", "Nidalese", "Keleshite"}[h]
}

type HumanHeritageWeighted struct {
	HumanHeritage
	Weight int
}

type HumanHeritageRange struct {
	HumanHeritage
	Min int
	Max int
}

type WeightedTable struct {
	enum int
	Min  int
	Max  int
}

type WeightedTables []WeightedTable

func (w *WeightedTables) Roll() int {
	return 0
}

type Attributes struct {
	Strength     int
	Dexterity    int
	Constitution int
	Perception   int
	Intelligence int
	Willpower    int
	Charisma     int
	Movement     int
}

type Character struct {
	Name     string
	Gender   string
	Race     string
	Ancestry string
	Attributes
}

func NewAttributes() Attributes {
	return Attributes{
		Strength:     SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Dexterity:    SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Constitution: SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Perception:   SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Intelligence: SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Willpower:    SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Charisma:     SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
		Movement:     SumRolls(KeepHighestRolls(3, MultiRolls(4, 6))),
	}
}

func NewCharacter() Character {
	nc := Character{}
	// 1) Add attributes
	//roll 4d6 take sum of highest 3
	fmt.Println("starting...")
	fmt.Println(Roll(6))
	fmt.Println("done.")

	nc.Attributes = NewAttributes()

	// 2) determine gender
	nc.Gender = Gender(Roll(2)).String()

	// 3) Determine Race/Ancestry
	nc.Race = Race(Roll(2)).String()

	switch nc.Race {
	case "Human":
		whhs := []HumanHeritageWeighted{
			{0, 50},
			{1, 10},
			{2, 10},
			{3, 10},
			{4, 10},
			{5, 10},
		}
		rhhs := make([]HumanHeritageRange, 0)
		totalWeight := 0
		ptr := 0
		for i, v := range whhs {
			totalWeight += v.Weight
			tmp := HumanHeritageRange{HumanHeritage(i), ptr + 1, totalWeight}
			rhhs = append(rhhs, tmp)
			ptr += v.Weight
		}
		result := Roll(totalWeight)
		for _, v := range rhhs {
			if result >= v.Min && result <= v.Max {
				nc.Ancestry = v.HumanHeritage.String()
			}
		}
	}

	// 4) Determine Profession(s)

	// 4.5) name

	// 5) add to inventory based on profession(s)

	// 6) Determine Damage Threshold

	// 7) Determine Damage Condition Track

	// 8) Determine Perils Threshold

	// 9) Determine Perils Condition Track

	// 10) Calculate Bonuses

	// 11) Determine Encumberance limit = SB + 5

	// 12) Calc Base Combat Bonus = (DB + PB + WP)/3

	// 13) Calc Initiative/Reflex = (DB + PB)

	// 14) Determine Age

	// 15) Distinguishing Marks

	// 16) Build Type

	// 17) Height and Weight

	// 18) Eye and hair color

	// 19) Upbringing

	// 20) Social Class add cash

	// 21) Drawbacks

	return nc
}

func Roll(d int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := d
	return rand.Intn(max-min+1) + min
}

func MultiRolls(n, d int) []int {
	rs := []int{}
	for i := 0; i < n; i++ {
		rs = append(rs, Roll(d))
	}
	return rs
}

func MultiRollSum(n, d int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += Roll(d)
	}
	return sum
}

func SumRolls(r []int) int {
	sum := 0
	for _, v := range r {
		sum += v
	}
	return sum
}

func KeepHighestRolls(h int, rs []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(rs)))
	rs = rs[:h]
	return rs
}

func KeepLowestRolls(l int, rs []int) []int {
	sort.Ints(rs)
	rs = rs[:l]
	return rs
}