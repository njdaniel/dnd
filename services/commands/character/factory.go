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
				return v.HumanHeritage.String()
			}
		}
	case "Elf":
		weights := []int{45, 45, 10}
		wt := NewWeightedTable(High, weights)
		return ElvenHeritage(wt.Roll()).String()
	case "Dwarf":
		return "Mountain"
	default:
		log.Fatal("error: race not picked")
	}
	return ""
}