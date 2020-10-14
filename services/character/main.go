package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
)

var BoxData *packr.Box

func init() {
	BoxData = packr.Folder("data/dnd")

}

func main() {
	fmt.Println("starting...")
	//fmt.Println("Opening box")
	//fmt.Println(BoxData.List())
	//fmt.Println(ReadCSV("./background/eyecolor.csv"))
	//fmt.Println("Calculating..", CalcAttrBonus(20))
	fmt.Println(NewCharacter())
	fmt.Println("done.")
}


type Languages int

const (
	Common Languages = iota +1
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
	male = iota + 1
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

func (h HumanHeritage) Len() int {
	return 6
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

type WeightedRow struct {
	Enum int
	Weight int
	Min  int
	Max  int
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
		tmp := WeightedRow{i, w[i-1], ptr+1, totalWeight}
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
	Young AgeGroup = iota +1
	Adult
	MiddleAge
	Elderly
)
func (a AgeGroup) String() string {
	return [...]string{"Young", "Adult", "MiddleAge", "Elderly" }[a-1]
}
func (a AgeGroup) Len() int {
	return 4
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

type Bonuses struct {
	SB int
	DB int
	CB int
	PB int
	IB int
	WB int
	ChB int
	CombatBonus int
	Initiative int
	Encumburance int
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
	Agriculture ProfessionType = iota +1
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
		"health", "labor", "magic", "military", "outcast", "religion", "scholars", "transportation"}[p]
}
func (p ProfessionType) Len() int {
	return 16
}


type Character struct {
	Name     string
	Gender   string
	Race     string
	Ancestry string
	Age      string
	DamageThreshold int
	DamageConditionState string
	Injuries []string
	PerilsThreshold int
	PerilsConditionState string
	Bonuses
	Attributes
	DistinguishingMarks []string
	BodyType string
	Complexion string
	Season string
	Upbringing string
	SocialClass string
	EyeColor string
	HairColor string
	Languages []string
	Height string
	Weight int
	Professions []string
	Skills []Skill
	ProfessionSelection string
}

type Skill struct {
	Name string
	BonusUsed string
	Focuses []Focus
	Level int
}

type Focus struct {
	Name string
	Level int
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
	fmt.Println("creating character")
	fmt.Println(Roll(6))
	fmt.Println("done.")

	nc.Attributes = NewAttributes()

	// 2) determine gender
	nc.Gender = Gender(Roll(2)).String()

	// 3) Determine Race/Ancestry
	nc.Race = Race(Roll(3)).String()

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
	case "Elf":
		weights := []int{45, 45, 10}
		wt := NewWeightedTable(High, weights )
		nc.Ancestry = ElvenHeritage(wt.Roll()).String()
	case "Dwarf":
		nc.Ancestry = "Mountain"
	default:
		log.Fatal("error: race not picked")
	}

	// 4) Determine Age
	nc.Age=func() string {
		weights := []int{25, 35, 25, 10}
		wt := NewWeightedTable(Young, weights)
		return AgeGroup(wt.Roll()).String()
	}()

	// 4.1) Determine Profession(s)

	// 4.5) name
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
	nc.Encumburance = nc.SB +5

	// 12) Calc Base Combat Bonus = (DB + PB + WP)/3
	nc.CombatBonus = (nc.DB + nc.PB + nc.WB)/3

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
		n := Roll(2)-1
		fmt.Println(n)
		for i := 0; i < n; i ++ {
			result := Roll(len(ss))
			nc.DistinguishingMarks = append(nc.DistinguishingMarks, ss[result-1])
		}
	case "MiddleAge":
		n := Roll(3)-1
		fmt.Println(n)
		for i := 0; i < n; i ++ {
			result := Roll(len(ss))
			nc.DistinguishingMarks = append(nc.DistinguishingMarks, ss[result-1])
		}
	case "Elderly":
		n := Roll(4)
		fmt.Println(n)
		for i := 0; i < n; i ++ {
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

	//height and weight
	//based on race bring in height and weight chart
	//assign num for body type (frail, slender, normal, husky, corpulent)= iota
	nc.Height = func() string {
		racelc := strings.ToLower(nc.Race)
		filename := fmt.Sprintf("./background/%s-%s-height.txt", racelc, nc.Gender)
		bs, err := BoxData.Find(filename)
		if err != nil {
			log.Println(err)
		}
		s := string(bs)
		ss := strings.Split(s, "\n")
		result := Roll(len(ss))
		height := ss[result-1]
		return height
	}()

	nc.Weight = func() int {
		racelc := strings.ToLower(nc.Race)
		filename := fmt.Sprintf("./background/%s-%s-weight.csv", racelc, nc.Gender)
		ssw, err := ReadCSV(filename)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(ssw)
		weight := 0
		return weight
	}()

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
	// 17) Height and Weight
	// 18) Eye and hair color
	switch nc.Race {
	case "Human":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.EyeColor = sshc[r][2]
	case "Elf":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.EyeColor = sshc[r][1]
	case "Dwarf":
		sshc, err := ReadCSV("./background/eyecolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.EyeColor = sshc[r][0]
	default:
		log.Println("no race determined")
	}

	//Hair
	switch nc.Race {
	case "Human":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.HairColor = sshc[r][2]
	case "Elf":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.HairColor = sshc[r][1]
	case "Dwarf":
		sshc, err := ReadCSV("./background/haircolor.csv")
		if err != nil {
			log.Println(err)
		}
		r := Roll(len(sshc))-1
		nc.HairColor = sshc[r][0]
	default:
		log.Println("no race determined")
	}
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
		for i := 0; i < numLangs; i ++ {
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
		for i := 0; i < numLangs; i ++ {
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
		for i := 0; i < numLangs; i ++ {
			//wanted Languages.Len() but got error...
			newL := Languages(Roll(5))
			//check if lang is already know
			if !mapLanguages[newL.String()] {
				nc.Languages = append(nc.Languages, newL.String())
				mapLanguages[newL.String()] = true
			}
		}
	}

	//Religion

	//Hometown

	// 21) Drawbacks

	// skills
	nc.Skills = []Skill{{"Alchemy", "IB", []Focus{}, 0},{"Athletics", "SB", []Focus{}, 0},
		{"Awareness", "PB", []Focus{}, 0},{"Bargain", "ChB", []Focus{}, 0},
		{"Charm", "ChB", []Focus{}, 0},{"Coordination", "DB", []Focus{}, 0},
		{"Counterfeit", "IB", []Focus{}, 0},{"Disguise", "ChB", []Focus{}, 0},
		{"Drive", "SB", []Focus{}, 0},{"Eavesdrop", "PB", []Focus{}, 0},

		{"Education", "IB", []Focus{}, 0},{"Folklore", "IB", []Focus{}, 0},
		{"Gamble", "IB", []Focus{}, 0},{"Guile", "DB", []Focus{}, 0},
		{"HandleAnimal", "PB", []Focus{}, 0},{"Heal", "IB", []Focus{}, 0},

		{"Magic", "WB", []Focus{}, 0},{"Interrogation", "WB", []Focus{}, 0},
		{"Intimidate", "SB", []Focus{}, 0},{"Leadership", "ChB", []Focus{}, 0},
		{"Melee", "CombatBonus", []Focus{}, 0},{"Range", "PB", []Focus{}, 0},

		{"Throwing", "DB", []Focus{}, 0},{"Dodge", "DB", []Focus{}, 0},
		{"Navigation", "IB", []Focus{}, 0},{"Sailing", "DB", []Focus{}, 0},
		{"Resolve", "WB", []Focus{}, 0},{"Ride", "DB", []Focus{}, 0},

		{"Scruntinize", "PB", []Focus{}, 0},{"Pickpocketing", "DB", []Focus{}, 0},
		{"Lockpicking", "DB", []Focus{}, 0},{"Stealth", "DB", []Focus{}, 0},
		{"Survival", "WB", []Focus{}, 0},{"Tradecraft", "WB", []Focus{}, 0},
		{"Warfare", "WB", []Focus{}, 0},
	}

	//ProfessionSelection: normal, advance
	// Find all profession jsons
	boxlst := BoxData.List()
	pj := make([]string,0)
	for _, v := range boxlst {
		if strings.Contains(v, "professions/archtype") {
			ss := strings.Split(s, "/")
			if strings.Contains(ss[len(ss)-1], ".json") {
				pj = append(pj, v)
			}
		}
	}
	log.Println(pj)

	rp := Roll(100)
	if rp == 1 {
		//ArchType Professions
		nc.ProfessionSelection = "advance"
		//import professions from json file
		//ref to map[string]Profession
		//roll to add profession
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
			ss := strings.Split(s,"\n")
			result := Roll(len(ss))
			ps = append(ps, ss[result-1])
			return ps
		}()
	}




	return nc
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
	return result,nil
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
	tmp := a -10
	if a > 10 {
		return tmp / 2
	} else if a < 10 {
		return ((11-a)/2)*(-1)
	}
	return 0
}
