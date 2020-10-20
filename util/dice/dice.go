package dice

import (
	"fmt"
	"log"
	"regexp"
)

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

//ParseRollString parses a string of roll info into a DiceInfo's fields
func ParseRollString(s string) DiceInfo {
	//result := 0
	//numberOfDice = 1
	//typeOfDice = 0
	di := newDiceInfo()

	rm := regexp.MustCompile(`\dd`)
	rh := regexp.MustCompile(`\dkh\dd\d`)
	rl := regexp.MustCompile(`\dkl\dd\d`)


	switch  {
	case rm.MatchString(s):
		fmt.Println("has more than one dice")
		if _, err := fmt.Sscanf(s, "%dd%d", di.NumberOfDice, di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
	case rh.MatchString(s):
		fmt.Println("keep the highest dice")
		if _, err := fmt.Sscanf(s, "%kh%dd%d", di.NumberOfDice, di.KeepDice, di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
		di.HighRoll=true
	case rl.MatchString(s):
		fmt.Println("keep the lowest dice")
		if _, err := fmt.Sscanf(s, "%dkl%dd%d", di.NumberOfDice, di.KeepDice, di.TypeOfDice); err != nil {
			log.Fatal(err)
		}
		di.LowRoll=true
	default:
		log.Println("just one dice")
		//ex d6
		if _, err := fmt.Sscanf(s, "d%d", di.TypeOfDice); err != nil {
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
