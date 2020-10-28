package character

import (
	"github.com/njdaniel/dnd/util/dice"
	"log"
)

func createGender() string {
	return Gender(dice.Roll(3)).String()
}

func createRace() string {
	return Race(dice.Roll(3)).String()
}

func createHeritage(nc Character) string {
	switch nc.Race {
	case "Human":
		weights := []int{50, 10, 10, 10, 10, 10}
		wt := NewWeightedTable(Imperial, weights)
		return HumanHeritage(wt.Roll()).String()
	case "Elf":
		weights := []int{45, 45, 10}
		wt := NewWeightedTable(High, weights)
		return ElvenHeritage(wt.Roll()).String()
	case "Dwarf":
		weights := []int{45, 45, 10}
		wt := NewWeightedTable(DwarvenHeritage(1), weights)
		return DwarvenHeritage(wt.Roll()).String()
	default:
		log.Fatal("error: race not picked")
	}
	return ""
}