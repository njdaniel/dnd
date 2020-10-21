package dice

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"time"
)

//DiceInfo contains the information about a roll
type DiceInfo struct {
	NumberOfDice int
	TypeOfDice int
	HighRoll bool
	LowRoll bool
	KeepDice int
	Explodes bool
	ExplodesOn []int
}

func newDiceInfo() DiceInfo {
	di := DiceInfo{
		NumberOfDice: 1,
		TypeOfDice: 0,
		HighRoll:false,
		LowRoll:false,
		KeepDice:0,
		Explodes:false,
		ExplodesOn:[]int{6,},
	}
	return di
}

//Roll base on the information of DiceInfo
func (d *DiceInfo) RollDice() []int {
	log.Println("rolling dice...")
	dice := make([]int, 0)
	switch  {
	case d.HighRoll:
		fmt.Println("use high roll")
		return KeepHighestRolls(d.KeepDice,MultiRolls(d.NumberOfDice, d.TypeOfDice))
	case d.LowRoll:
		fmt.Println("use low roll")
		return KeepLowestRolls(d.KeepDice,MultiRolls(d.NumberOfDice, d.TypeOfDice))
	case d.Explodes:
		fmt.Println("explodes")
		return RollExplodes(d.TypeOfDice, d.ExplodesOn)
	case d.NumberOfDice > 1:
		fmt.Println("multiple rolls")
		return MultiRolls(d.NumberOfDice, d.TypeOfDice)
	default:
		fmt.Println("default roll")
		dice = append(dice, Roll(d.TypeOfDice))
	}
	return dice
}

//RollExplodes keeps rerolling until not hitting an explode die number
func RollExplodes(d int, e []int) []int {
	rolls := make([]int, 0)
	r := Roll(d)
	rolls = append(rolls, r)
	me := make(map[int]struct{}, 0)
	for _, v := range e {
		if _, ok := me[v]; !ok {
			me[v] = struct{}{}
		}
	}
	//explodes := func(){
	//
	//}
	exploded := false
	if _, ok := me[r]; ok {
		exploded = true
	}
	for exploded {
		r := Roll(d)
		rolls = append(rolls, r)
		if _, ok := me[r]; !ok {
			exploded = false
		}
	}
	return rolls
}
//func explodes(d int, m map[int]struct{}) []int {
//
//	r := Roll(d)
//	rolls := make([]int, 0)
//	rolls = append(rolls, r)
//	if _, ok := m[r]; ok {
//
//	}
//	return rolls
//}
func Roll(d int) int {
	if d <= 0 {
		return -1
	}
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
func SumRolls(r []int) int {
	sum := 0
	for _, v := range r {
		sum += v
	}
	return sum
}

//ParseRollString parses a string of roll info into a DiceInfo's fields
func ParseRollString(s string) DiceInfo {
	//result := 0
	//numberOfDice = 1
	//typeOfDice = 0
	di := newDiceInfo()

	rm := regexp.MustCompile(`\dd`)
	rh := regexp.MustCompile(`\dkh\dd\d`)
	rl := regexp.MustCompile(`\dkl\dd\d`)
	rme := regexp.MustCompile(`\dd\d!`)
	re := regexp.MustCompile(`d\d!`)

	switch  {
	case rh.MatchString(s):
		fmt.Println("keep the highest dice")
		if _, err := fmt.Sscanf(s, "%dkh%dd%d", &di.NumberOfDice, &di.KeepDice, &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
		di.HighRoll=true
	case rl.MatchString(s):
		fmt.Println("keep the lowest dice")
		if _, err := fmt.Sscanf(s, "%dkl%dd%d", &di.NumberOfDice, &di.KeepDice, &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
		di.LowRoll=true
	case rme.MatchString(s):
		fmt.Println("multiple explodes")
		if _, err := fmt.Sscanf(s, "%dd%d!", &di.NumberOfDice, &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
	case re.MatchString(s):
		fmt.Println("explodes")
		if _, err := fmt.Sscanf(s, "d%d!", &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
		di.Explodes=true
	case rm.MatchString(s):
		fmt.Println("has more than one dice")
		if _, err := fmt.Sscanf(s, "%dd%d", &di.NumberOfDice, &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
	default:
		log.Println("just one dice")
		//ex d6
		if _, err := fmt.Sscanf(s, "d%d", &di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
	}

	return di

	//ex 2d6
	//roll 2d6 return [2]int

	//ex 4kh3d6
	//roll 4d6 keep highest 3

	//ex 4kl3d6
	//roll 4d6 keep lowest 3

	//explode means to add to total and reroll and add to total, keep rolling until not hitting an explode number

	//ex d6! or d6!6
	//explode on 6

	//ex d6!1,6
	//explode on 6 and 1

	//ex d6!1
	//explode on 1

	//return result
}
