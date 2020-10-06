package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("starting...")
	rs := MultiRolls(6, 6)
	fmt.Println("rs: ", rs)
	hrs := KeepHighestRolls(3, rs)
	fmt.Println("hrs: ", hrs)
	fmt.Println("done.")
}

type Gender int

const (
	male = iota
	female
	)

func (g Gender) String() string {
	return[...]string{"male", "female"}[g]
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
	Name string
	Gender string
	Attributes
}

func NewAttributes() Attributes {
	return Attributes{
		Strength:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Dexterity:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Constitution:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Perception:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Intelligence:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Willpower:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Charisma:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
		Movement:SumRolls(KeepHighestRolls(3,MultiRolls(4,6))),
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
