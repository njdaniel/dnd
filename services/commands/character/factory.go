package character

import (
	"log"
	"strings"

	"github.com/njdaniel/dnd/util/dice"
)

func createGender() string {
	return Gender(dice.Roll(2)).String()
}

func createRace() string {
	return Race(dice.Roll(3)).String()
}

func createHeritage(race string) string {
	switch race {
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

func createName(race, gender string) string {
	switch race {
	case "Human":
		if gender == "male" {
			log.Println("found human male")
			log.Println(BoxData.List())
			buf, err := BoxData.Find("names/human-male.txt")
			//open data/dnd/names/human-male.txt
			//slurp in memory, change to read by byte chunks if it gets to big
			//buf, err := ioutil.ReadFile("data/dnd/names/human-male.txt")
			if err != nil {
				log.Fatal(err)
			}
			s := string(buf)
			ss := strings.Split(s, "\n")
			result := Roll(len(ss))
			return ss[result-1]
		} else if gender == "female" {
			buf, err := BoxData.Find("names/human-female.txt")
			//slurp in memory, change to read by byte chunks if it gets to big
			//buf, err := ioutil.ReadFile("data/dnd/names/human-female.txt")
			if err != nil {
				log.Fatal(err)
			}
			s := string(buf)
			ss := strings.Split(s, "\n")
			result := Roll(len(ss))
			return ss[result-1]
		} else {
			return "Pierce the Dickish"
			log.Println("no gender assigned")
		}
	}
	return ""
}
