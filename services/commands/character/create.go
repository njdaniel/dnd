package character

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/njdaniel/dnd/util/dice"
)

var BoxData *packr.Box

func init() {
	//log.Println("character init() called")
	//dir, _ := os.Getwd()
	//log.Println("Current dir: " + dir)
	BoxData = packr.New("boxdata", "../../../data/dnd")
	//BoxData = packr.Folder("data/dnd")
	//log.Println(BoxData.List())
	//log.Println("character init() done")

}

type Languages int

const (
	Common Languages = iota + 1
	Undercommon
	Elvish
	Dwarvish
	Orcish
)

func (l Languages) String() string {
	return [...]string{"Common", "Undercommon", "Elvish", "Dwarvish", "Orcish"}[l-1]
}

func (l Languages) Len() int {
	return 5
}

type Gender int

const (
	male Gender = iota + 1
	female
)

func (g Gender) String() string {
	return [...]string{"male", "female"}[g-1]
}

type Race int

const (
	Human = iota + 1
	Elf
	Dwarf
)

func (r Race) String() string {
	return [...]string{"Human", "Elf", "Dwarf"}[r-1]
}

//List return slice of strings that contain all the elements of Race
func (r Race) List() []string {
	return []string{"Human", "Elf", "Dwarf"}
}

type DwarvenHeritage int

const (
	IronHills DwarvenHeritage = iota + 1
	Mountain
	Deep
)

func (d DwarvenHeritage) String() string {
	return [...]string{"IronHills", "Mountain", "Deep"}[d-1]
}

func (d DwarvenHeritage) Len() int {
	return 3
}

type ElvenHeritage int

const (
	High ElvenHeritage = iota + 1
	Drow
	Wood
)

func (e ElvenHeritage) String() string {
	return [...]string{"High", "Drow", "Wood"}[e-1]
}

func (e ElvenHeritage) Len() int {
	return 3
}

type HumanHeritage int

const (
	Imperial HumanHeritage = iota + 1
	Nord
	Vardisan
	Lumdrani
	Nimalese
	Minskite
)

func (h HumanHeritage) String() string {
	return [...]string{"Imperial", "Nord", "Vardisan", "Lumdrani", "Nimalese", "Minskite"}[h-1]
}

func (h HumanHeritage) Len() int {
	return 6
}

type HumanHeritageWeighted struct {
	HumanHeritage
	Weight int `json:"weight"`
}

type HumanHeritageRange struct {
	HumanHeritage
	Min int `json:"min"`
	Max int `json:"max"`
}

type WeightedRow struct {
	Enum   int `json:"enum"`
	Weight int `json:"weight"`
	Min    int `json:"min"`
	Max    int `json:"max"`
}

type WeightedTable []WeightedRow

type Enum interface {
	String() string
	Len() int
}

func NewWeightedTable(enum Enum, w []int) WeightedTable {
	wt := make(WeightedTable, 0)
	totalWeight := 0
	ptr := 0
	for i := 1; i <= enum.Len(); i++ {
		totalWeight += w[i-1]
		tmp := WeightedRow{i, w[i-1], ptr + 1, totalWeight}
		wt = append(wt, tmp)
		ptr += w[i-1]
	}
	return wt
}

func (wt *WeightedTable) Roll() int {
	totalWeight := 0
	for _, v := range *wt {
		totalWeight += v.Weight
	}
	result := Roll(totalWeight)
	for _, v := range *wt {
		if result >= v.Min && result <= v.Max {
			return v.Enum
		}
	}
	return 0
}

type AgeGroup int

const (
	Young AgeGroup = iota + 1
	Adult
	MiddleAge
	Elderly
)

func (a AgeGroup) String() string {
	return [...]string{"Young", "Adult", "MiddleAge", "Elderly"}[a-1]
}
func (a AgeGroup) Len() int {
	return 4
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Perception   int `json:"perception"`
	Intelligence int `json:"intelligence"`
	Willpower    int `json:"willpower"`
	Charisma     int `json:"charisma"`
	Movement     int `json:"movement"`
}

type Bonuses struct {
	SB           int `json:"sb"`
	DB           int `json:"db"`
	CB           int `json:"cb"`
	PB           int `json:"pb"`
	IB           int `json:"ib"`
	WB           int `json:"wb"`
	ChB          int `json:"ch_b"`
	CombatBonus  int `json:"combat_bonus"`
	Initiative   int `json:"initiative"`
	Encumburance int `json:"encumburance"`
}

type DamageConditionState int

const (
	Unharmed DamageConditionState = iota
	LightlyWounded
	ModeratelyWounded
	SeriouslyWounded
	GreivouslyWounded
	Slain
)

func (d DamageConditionState) String() string {
	return [...]string{"Unharmed", "LightlyWounded", "ModeratelyWounded", "SeriouslyWounded", "GreivouslyWounded", "Slain"}[d]
}
func (d DamageConditionState) Len() int {
	return 6
}

type PerilsConditonState int

const (
	Unhindered PerilsConditonState = iota
	Imperiled
	IgnoreOneSkill
	IgnoreTwoSkills
	IgnoreThreeSkills
	Incapacitated
)

func (p PerilsConditonState) String() string {
	return [...]string{"Unhindered", "Imperiled", "IgnoreOneSkill", "IgnoreTwoSkills", "IgnoreThreeSkills", "Incapacitated"}[p]
}

func (p PerilsConditonState) Len() int {
	return 6
}

type ProfessionType int

const (
	Agriculture ProfessionType = iota + 1
	Arts
	BusinessTrade
	Communications
	Construction
	Craftman
	Crime
	Government
	Health
	Labor
	Magic
	Military
	Outcast
	Religion
	Scholars
	Transportation
)

func (p ProfessionType) String() string {
	return [...]string{"agriculture", "arts", "business-trade", "communications", "construction", "craftman", "crime", "governament",
		"health", "labor", "magic", "military", "outcast", "religion", "scholars", "transportation"}[p-1]
}
func (p ProfessionType) Len() int {
	return 16
}

//go:generate easytags $GOFILE json

//Character is the fields of information
type Character struct {
	Name                 string   `json:"name"`
	Gender               string   `json:"gender"`
	Race                 string   `json:"race"`
	Ancestry             string   `json:"ancestry"`
	Age                  string   `json:"age"`
	DamageThreshold      int      `json:"damage_threshold"`
	DamageConditionState string   `json:"damage_condition_state"`
	Injuries             []string `json:"injuries"`
	PerilsThreshold      int      `json:"perils_threshold"`
	PerilsConditionState string   `json:"perils_condition_state"`
	Bonuses
	Attributes
	DistinguishingMarks []string `json:"distinguishing_marks"`
	BodyType            string   `json:"body_type"`
	Complexion          string   `json:"complexion"`
	Season              string   `json:"season"`
	Upbringing          string   `json:"upbringing"`
	SocialClass         string   `json:"social_class"`
	EyeColor            string   `json:"eye_color"`
	HairColor           string   `json:"hair_color"`
	Languages           []string `json:"languages"`
	Height              string   `json:"height"`
	Weight              int      `json:"weight"`
	Professions         []string `json:"professions"`
	Skills              []Skill  `json:"skills"`
	ProfessionSelection string   `json:"profession_selection"`
	Money
	Items []InventoryItem `json:"items"`
}

type Skill struct {
	Name      string  `json:"name"`
	BonusUsed string  `json:"bonus_used"`
	Focuses   []Focus `json:"focuses"`
	Level     int     `json:"level"`
}

type Focus struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
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

func NewCharacter(name, gender, race string) Character {
	nc := Character{}
	// 1) Add attributes
	//roll 4d6 take sum of highest 3
	fmt.Println("creating character..")
	//fmt.Println(Roll(6))
	//fmt.Println("done.")

	nc.Attributes = NewAttributes()

	// 2) determine gender
	//nc.Gender = Gender(Roll(2)).String()
	if gender == "" {
		nc.Gender = createGender()
	} else {
		nc.Gender = gender
	}

	// 3) Determine Race/Ancestry
	//nc.Race = Race(Roll(3)).String()
	if race == "" {
		nc.Race = createRace()
	} else {
		//capitalize first char of name
		race = strings.Title(race)
		nc.Race = race
	}

	nc.Ancestry = createHeritage(nc.Race)
	//switch nc.Race {
	//case "Human":
	//	whhs := []HumanHeritageWeighted{
	//		{0, 50},
	//		{1, 10},
	//		{2, 10},
	//		{3, 10},
	//		{4, 10},
	//		{5, 10},
	//	}
	//	rhhs := make([]HumanHeritageRange, 0)
	//	totalWeight := 0
	//	ptr := 0
	//	for i, v := range whhs {
	//		totalWeight += v.Weight
	//		tmp := HumanHeritageRange{HumanHeritage(i), ptr + 1, totalWeight}
	//		rhhs = append(rhhs, tmp)
	//		ptr += v.Weight
	//	}
	//	result := Roll(totalWeight)
	//	for _, v := range rhhs {
	//		if result >= v.Min && result <= v.Max {
	//			nc.Ancestry = v.HumanHeritage.String()
	//		}
	//	}
	//case "Elf":
	//	weights := []int{45, 45, 10}
	//	wt := NewWeightedTable(High, weights)
	//	nc.Ancestry = ElvenHeritage(wt.Roll()).String()
	//case "Dwarf":
	//	nc.Ancestry = "Mountain"
	//default:
	//	log.Fatal("error: race not picked")
	//}

	// 4) Determine Age
	nc.Age = func() string {
		weights := []int{25, 35, 25, 10}
		wt := NewWeightedTable(Young, weights)
		return AgeGroup(wt.Roll()).String()
	}()
	log.Printf("Age Created %s\n", nc.Age)

	// 4.1) Determine Profession(s)

	// 4.5) name
	if name == "" {
		switch nc.Race {
		case "Human":
			if nc.Gender == "male" {
				//open data/dnd/names/human-male.txt
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/human-male.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else if nc.Gender == "female" {
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/human-female.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else {
				nc.Name = "Pierce the Dickish"
				log.Println("no gender assigned")
			}
		case "Dwarf":
			if nc.Gender == "male" {
				//open data/dnd/names/human-male.txt
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/dwarf-male.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else if nc.Gender == "female" {
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/dwarf-female.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else {
				nc.Name = "Carlos"
			}
		case "Elf":
			if nc.Gender == "male" {
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/elf-male.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else if nc.Gender == "female" {
				//slurp in memory, change to read by byte chunks if it gets to big
				buf, err := ioutil.ReadFile("data/dnd/names/elf-female.txt")
				if err != nil {
					log.Fatal(err)
				}
				s := string(buf)
				ss := strings.Split(s, "\n")
				result := Roll(len(ss))
				nc.Name = ss[result-1]
			} else {
				nc.Name = "Brutalitops"
			}
		}
	} else {
		nc.Name = name
	}

	// 5) add to inventory based on profession(s)

	// 10) Calculate Bonuses
	nc.SB = CalcAttrBonus(nc.Strength)
	nc.CB = CalcAttrBonus(nc.Constitution)
	nc.PB = CalcAttrBonus(nc.Perception)
	nc.IB = CalcAttrBonus(nc.Intelligence)
	nc.ChB = CalcAttrBonus(nc.Charisma)
	nc.WB = CalcAttrBonus(nc.Willpower)
	nc.DB = CalcAttrBonus(nc.Dexterity)
	// 6) Determine Damage Threshold
	// 7) Determine Damage Condition Track
	// 8) Determine Perils Threshold
	// 9) Determine Perils Condition Track
	nc.DamageThreshold = nc.CB
	nc.DamageConditionState = DamageConditionState(0).String()
	nc.PerilsThreshold = nc.WB + 3
	nc.PerilsConditionState = PerilsConditonState(0).String()

	// 11) Determine Encumberance limit = SB + 5
	nc.Encumburance = nc.SB + 5

	// 12) Calc Base Combat Bonus = (DB + PB + WP)/3
	nc.CombatBonus = (nc.DB + nc.PB + nc.WB) / 3

	// 13) Calc Initiative/Reflex = (DB + PB)
	nc.Initiative = nc.DB + nc.PB

	// 15) Distinguishing Marks

	//slurp in memory, change to read by byte chunks if it gets to big
	buf, err := ioutil.ReadFile("data/dnd/background/marks.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(buf)
	ss := strings.Split(s, "\n")
	switch nc.Age {
	case "Young":
		log.Println("young and unblemished")
	case "Adult":
		n := Roll(2) - 1
		fmt.Println(n)
		for i := 0; i < n; i++ {
			result := Roll(len(ss))
			nc.DistinguishingMarks = append(nc.DistinguishingMarks, ss[result-1])
		}
	case "MiddleAge":
		n := Roll(3) - 1
		fmt.Println(n)
		for i := 0; i < n; i++ {
			result := Roll(len(ss))
			nc.DistinguishingMarks = append(nc.DistinguishingMarks, ss[result-1])
		}
	case "Elderly":
		n := Roll(4)
		fmt.Println(n)
		for i := 0; i < n; i++ {
			result := Roll(len(ss))
			nc.DistinguishingMarks = append(nc.DistinguishingMarks, ss[result-1])
		}
	default:
		log.Println("undetermined age and no distinguishing features")
	}
	// 16) Build Type
	//slurp in memory, change to read by byte chunks if it gets to big
	nc.BodyType = func() string {
		buf, err := ioutil.ReadFile("data/dnd/background/body_type.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		return ss[result-1]
	}()
	log.Println("added build type")

	//height and weight
	//based on race bring in height and weight chart
	//assign num for body type (frail, slender, normal, husky, corpulent)= iota
	nc.Height = func() string {
		racelc := strings.ToLower(nc.Race)
		filename := fmt.Sprintf("background/%s-%s-height.txt", racelc, nc.Gender)
		log.Println(filename)
		bs, err := BoxData.Find(filename)
		if err != nil {
			log.Fatalf("error finding file %s: %v", filename, err)
		}
		s := string(bs)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		height := ss[result-1]
		return height
	}()
	log.Println("add height")

	nc.Weight = func() int {
		racelc := strings.ToLower(nc.Race)
		filename := fmt.Sprintf("./background/%s-%s-weight.csv", racelc, nc.Gender)
		ssw, err := ReadCSV(filename)
		if err != nil {
			log.Fatalf("error reading csv: %v", err)
		}
		fmt.Println(ssw)
		weight := 0
		return weight
	}()
	log.Println("add weight")

	// Complexion
	nc.Complexion = func() string {
		buf, err := ioutil.ReadFile("data/dnd/background/complexion.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		return ss[result-1]
	}()
	log.Println("weight")

	//Season of Birth
	nc.Season = func() string {
		buf, err := ioutil.ReadFile("data/dnd/background/season.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		return ss[result-1]
	}()
	log.Println("add season of birth")

	// 19) Upbringing

	// Upbringing
	nc.Upbringing = func() string {
		buf, err := ioutil.ReadFile("data/dnd/background/upbringing.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		return ss[result-1]
	}()
	log.Println("add upbringing")

	// Social Class
	nc.SocialClass = func() string {
		buf, err := ioutil.ReadFile("data/dnd/background/socialclass.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		return ss[result-1]
	}()
	log.Println("add social class")
	// 17) Height and Weight
	// 18) Eye and hair color
	switch nc.Race {
	case "Human":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.EyeColor = sshc[r][2]
	case "Elf":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.EyeColor = sshc[r][1]
	case "Dwarf":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.EyeColor = sshc[r][0]
	default:
		log.Println("no race determined")
	}
	log.Println("add eye color")

	//Hair
	switch nc.Race {
	case "Human":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.HairColor = sshc[r][2]
	case "Elf":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.HairColor = sshc[r][1]
	case "Dwarf":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc)) - 1
		nc.HairColor = sshc[r][0]
	default:
		log.Println("no race determined")
	}
	log.Println("add hair color")
	//nc.HairColor = func() string {
	//	buf, err := ioutil.ReadFile("data/dnd/background/.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	s := string(buf)
	//	ss := strings.Split(s, "\n")
	//	result := Roll(len(ss))
	//	return ss[result-1]
	//}()

	// 20) Social Class add cash

	// Languages
	nc.Languages = append(nc.Languages, "Common")
	switch nc.Race {
	case "Human":
		mapLanguages := map[string]bool{"Common": true, "Undercommon": false, "Elvish": false, "Dwarvish": false, "Orcish": false}
		numLangs := Roll(nc.IB)
		for i := 0; i < numLangs; i++ {
			//wanted Languages.Len() but got error...
			newL := Languages(Roll(5))
			//check if lang is already know
			if !mapLanguages[newL.String()] {
				nc.Languages = append(nc.Languages, newL.String())
				mapLanguages[newL.String()] = true
			}
		}
	case "Elf":
		nc.Languages = append(nc.Languages, "Elvish")
		mapLanguages := map[string]bool{"Common": true, "Undercommon": false, "Elvish": true, "Dwarvish": false, "Orcish": false}
		numLangs := Roll(nc.IB)
		for i := 0; i < numLangs; i++ {
			//wanted Languages.Len() but got error...
			newL := Languages(Roll(5))
			//check if lang is already know
			if !mapLanguages[newL.String()] {
				nc.Languages = append(nc.Languages, newL.String())
				mapLanguages[newL.String()] = true
			}
		}
	case "Dwarf":
		nc.Languages = append(nc.Languages, "Dwarvish")
		mapLanguages := map[string]bool{"Common": true, "Undercommon": false, "Elvish": false, "Dwarvish": true, "Orcish": false}
		numLangs := Roll(nc.IB)
		for i := 0; i < numLangs; i++ {
			//wanted Languages.Len() but got error...
			newL := Languages(Roll(5))
			//check if lang is already know
			if !mapLanguages[newL.String()] {
				nc.Languages = append(nc.Languages, newL.String())
				mapLanguages[newL.String()] = true
			}
		}
	}
	log.Println("add languages")

	//Religion

	//Hometown

	// 21) Drawbacks

	// skills
	nc.Skills = []Skill{{"Alchemy", "IB", []Focus{}, 0}, {"Athletics", "SB", []Focus{}, 0},
		{"Awareness", "PB", []Focus{}, 0}, {"Bargain", "ChB", []Focus{}, 0},
		{"Charm", "ChB", []Focus{}, 0}, {"Coordination", "DB", []Focus{}, 0},
		{"Counterfeit", "IB", []Focus{}, 0}, {"Disguise", "ChB", []Focus{}, 0},
		{"Drive", "SB", []Focus{}, 0}, {"Eavesdrop", "PB", []Focus{}, 0},

		{"Education", "IB", []Focus{}, 0}, {"Folklore", "IB", []Focus{}, 0},
		{"Gamble", "IB", []Focus{}, 0}, {"Guile", "DB", []Focus{}, 0},
		{"HandleAnimal", "PB", []Focus{}, 0}, {"Heal", "IB", []Focus{}, 0},

		{"Magic", "WB", []Focus{}, 0}, {"Interrogation", "WB", []Focus{}, 0},
		{"Intimidate", "SB", []Focus{}, 0}, {"Leadership", "ChB", []Focus{}, 0},
		{"Melee", "CombatBonus", []Focus{}, 0}, {"Range", "PB", []Focus{}, 0},

		{"Throwing", "DB", []Focus{}, 0}, {"Dodge", "DB", []Focus{}, 0},
		{"Navigation", "IB", []Focus{}, 0}, {"Sailing", "DB", []Focus{}, 0},
		{"Resolve", "WB", []Focus{}, 0}, {"Ride", "DB", []Focus{}, 0},

		{"Scruntinize", "PB", []Focus{}, 0}, {"Pickpocketing", "DB", []Focus{}, 0},
		{"Lockpicking", "DB", []Focus{}, 0}, {"Stealth", "DB", []Focus{}, 0},
		{"Survival", "WB", []Focus{}, 0}, {"Tradecraft", "WB", []Focus{}, 0},
		{"Warfare", "WB", []Focus{}, 0},
	}
	log.Println("add skills")
	//ProfessionSelection: normal, advance, detailed, special
	// Find all profession jsons
	boxlst := BoxData.List()
	advanceProfessions := make([]string, 0)
	for _, v := range boxlst {
		if strings.Contains(v, "professions/archtype") {
			ss := strings.Split(s, "/")
			if strings.Contains(ss[len(ss)-1], ".json") {
				advanceProfessions = append(advanceProfessions, v)
			}
		}
	}
	log.Println(advanceProfessions)

	rp := Roll(100)
	if rp == 1 {
		//ArchType Professions
		nc.ProfessionSelection = "advance"
		//import professions from json file
		//ref to map[string]Profession
		//roll to add profession
		newProfession := advanceProfessions[Roll(len(advanceProfessions))]
		nc.Professions = append(nc.Professions, newProfession)
		//read in profession
		bs, err := BoxData.Find(newProfession)
		if err != nil {
			log.Fatal(err)
		}
		p := Profession{}
		json.Unmarshal(bs, &p)
		//add money from profession
		nc.Money = NewMoney()
		func(nm map[string]string, m Money) {
			//for each preciousMetal(GoldCrowns, SilverShillings, CopperPennies)
			coinTypes := []string{"gc", "ss", "cp"}
			for i, ct := range coinTypes {
				if _, ok := nm[ct]; ok {
					//numberOfDice := ""
					fmt.Println("debug")
					di := dice.ParseRollString(nm[ct])
					if coinTypes[i] == m.GoldCrowns.ShortHand {
						m.GoldCrowns.Count += dice.SumRolls(di.RollDice())
					} else if coinTypes[i] == m.SilverShillings.ShortHand {
						m.SilverShillings.Count += dice.SumRolls(di.RollDice())
					} else if coinTypes[i] == m.CopperPennies.ShortHand {
						m.CopperPennies.Count += dice.SumRolls(di.RollDice())
					}

				}
			}
			//find number of dice
			//find die type eg d6, d20
			//add to preciousMetal.count
		}(p.StartingMoney, nc.Money)
	} else {
		nc.ProfessionSelection = "normal"
		//Professions
		nc.Professions = func() []string {
			pt := Roll(Agriculture.Len())
			ps := make([]string, 0)
			filename := fmt.Sprintf("./professions/types/%s.txt", ProfessionType(pt).String())
			bs, err := BoxData.Find(filename)
			if err != nil {
				log.Println(err)
			}
			s := string(bs)
			ss := strings.Split(s, "\n")
			result := Roll(len(ss))
			ps = append(ps, ss[result-1])
			return ps
		}()
	}

	return nc
}

func RollString(s string) (numberOfDice, typeOfDice int) {
	//result := 0
	numberOfDice = 1
	typeOfDice = 0

	rm := regexp.MustCompile(`\dd`)

	switch {
	case rm.MatchString(s):
		fmt.Println("has more than one dice")
		if _, err := fmt.Sscanf(s, "%dd%d", &numberOfDice, &typeOfDice); err != nil {
			log.Fatal(err)
		}
	default:
		log.Println("just one dice")
		//ex d6
		if _, err := fmt.Sscanf(s, "d%d", &typeOfDice); err != nil {
			log.Fatal(err)
		}
	}

	return

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

//ReadCSV reads in csv file to be used
func ReadCSV(filename string) ([][]string, error) {
	//verify csv
	//sss := make([][]string, 0)
	bs, err := BoxData.Find(filename)
	if err != nil {
		log.Println(err)
	}
	r := csv.NewReader(bytes.NewReader(bs))
	result, err := r.ReadAll()
	if err != nil {
		log.Println()
	}
	return result, nil
}

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

// Calculates Attribute Bonus
func CalcAttrBonus(a int) int {
	tmp := a - 10
	if a > 10 {
		return tmp / 2
	} else if a < 10 {
		return ((11 - a) / 2) * (-1)
	}
	return 0
}
