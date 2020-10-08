package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	fmt.Println("starting...")
	fmt.Println("Calculating..", CalcAttrBonus(20))
	fmt.Println(NewCharacter())
	fmt.Println("done.")
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
	HairColor string
	Complexion string
	Season string
	Upbringing string
	SocialClass string
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


	// 19) Upbringing

	// 20) Social Class add cash

	// Languages

	//Religion

	//Hometown

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
