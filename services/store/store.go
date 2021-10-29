package store

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/njdaniel/dnd/services/commands/character"
)

// Item is the generic physical object
//type Item interface {
//	Name() string
//	Description() string
//	Price() string
//}

type StoreType int

//type StoreType string
//
//const (
//	Fletcher StoreType = "Fletcher"
//)

const (
	Fletcher StoreType = iota + 1
	Tavern
	Blacksmith
	Apothecary
	Clothing
	Generalgoods
	Bowyer
	Woodworker
	Armorer
	Weaponsmith
	Grocer
	Leatherworker
	Stablemaster
)

var StoreTypeName = []string{
	Fletcher: "Fletcher",
}

func (s StoreType) String() string {
	return [...]string{"Fletcher"}[s-1]
}

func (s StoreType) Len() int {
	return 1
}

func (s StoreType) IsValid() bool {
	switch s {
	case Fletcher:
		return true
	}
	return false
}

func String2StoreType(s string) (StoreType, error) {
	for i := 1; i <= StoreType(i).Len(); i++ {
		if StoreType(i).String() == s {
			return StoreType(i), nil
		}
	}
	return 0, fmt.Errorf("error: Invalid StoreType")
}

// VerifyEnum verify that string is one of the enum values
//TODO: needs to be tested guessing this wont work
//func VerifyEnum(s string, enum character.Enum) bool {
//	for i := 0; i < enum.Len(); i++ {
//		if enum.String() == s {
//			return true
//		}
//	}
//	return false
//}

///////////////////////
//Builder Pattern

type IStoreBuilder interface {
	setStoreType()
	setName()
	setOwner()
	setLocation()
	setInventory()
	GetStore() Store
}

func GetStoreBuilder(storeBuilderType string) IStoreBuilder {
	fmt.Println(storeBuilderType)
	isb := FletcherBuilder{}
	return isb
}

type FletcherBuilder struct {
	Name string
	StoreType
	Owner    string
	Location string
	Inventory
}

func (f FletcherBuilder) setStoreType() {
	f.StoreType = 1
}

func (f FletcherBuilder) setName() {
	f.Name = "Fletcher Shop"
}

func (f FletcherBuilder) setOwner() {
	//TODO: call character.NewCharacter
	f.Owner = "Bob"
}

func (f FletcherBuilder) setLocation() {
	f.Location = "Heldheim"
}

func (f FletcherBuilder) setInventory() {
	f.Inventory, _ = generateInventoryForFletcher()
}

func (f FletcherBuilder) GetStore() Store {
	return Store{
		Name:     f.Name,
		Owner:    f.Owner,
		Location: f.Location,
	}
}

type Metal int

const (
	Brass Metal = iota + 1
	Iron
	Silver
	Steel
)

func (m Metal) String() string {
	return [...]string{"Brass", "Iron", "Silver", "Steel"}[m-1]
}

func (m Metal) Len() int {
	return 4
}

type ArrowHeadType int

const (
	Pointed ArrowHeadType = iota + 1
	Leaf
	Broadhead
	Barbed
	Bodkin
	Bludgeon
)

func (a ArrowHeadType) String() string {
	return [...]string{"Pointed", "Leaf", "Broadhead", "Barbed", "Bodkin", "Bludgeon"}[a-1]
}

func (a ArrowHeadType) Len() int {
	return 6
}

func (a ArrowHeadType) IsValid() error {
	switch a {
	case Pointed, Leaf, Broadhead, Barbed, Bodkin, Bludgeon:
		return nil
	}
	return errors.New("Invalid ArrowHeadType")
}

type ShaftMaterial int

const (
	Ash ShaftMaterial = iota + 1
	Cedar
	DouglusFir
	Spruce
	Pine
)

var ShaftMaterialName = []string{
	Ash:        "Ash",
	Cedar:      "Cedar",
	DouglusFir: "Douglus Fir",
	Spruce:     "Spruce",
	Pine:       "Pine",
}

func (s ShaftMaterial) String() string {
	return ShaftMaterialName[s]
}

func (s ShaftMaterial) Len() int {
	return len(ShaftMaterialName)
}

func (s ShaftMaterial) IsValid() bool {
	switch s {
	case Ash, Cedar, DouglusFir, Spruce, Pine:
		return true
	}
	return false
}

func String2ShaftMaterial(s string) (ShaftMaterial, error) {
	for i := 1; i <= ShaftMaterial(i).Len(); i++ {
		if StoreType(i).String() == s {
			return ShaftMaterial(i), nil
		}
	}
	return 0, fmt.Errorf("error: Invalid ShaftMaterial")
}

type FletchMaterial int

const (
	GooseFeathers FletchMaterial = iota + 1
	TurkeyFeathers
	ChickenFeathers
	PidgeonFeathers
)

func (f FletchMaterial) Len() int {
	return 4
}

type Arrow struct {
	Item
	HeadMaterial Metal
	HeadType     ArrowHeadType
	ShaftMaterial
	FletchMaterial
	Length int
}

func NewArrow() ItemInterface {
	hm := Metal(character.Roll(Metal(1).Len()))
	ht := ArrowHeadType(character.Roll(ArrowHeadType(1).Len()))
	sm := ShaftMaterial(character.Roll(ShaftMaterial(1).Len()))
	fm := FletchMaterial(character.Roll(FletchMaterial(1).Len()))
	arrowlengths := [2]int{24, 32}
	arrowweights := [2]float64{0.06, 0.08}
	r := character.Roll(2) - 1
	length := arrowlengths[r]
	na := Arrow{
		Item: Item{
			Name:    hm.String() + " " + ht.String() + " " + strconv.Itoa(length) + "\" arrow",
			Price:   NewMoney(0, 0, 0),
			Weight:  newWeight(arrowweights[r]),
			Quality: Quality(character.Roll(Quality(1).Len() - 1)).String(),
		},
		Length:         length,
		HeadMaterial:   hm,
		HeadType:       ht,
		ShaftMaterial:  sm,
		FletchMaterial: fm,
	}
	return na
}

////////////////////

// Store represents the physical entity of a business
type Store struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Location  string `json:"location"`
	StoreType string `json:"store_type"`
	Inventory `json:"inventory"`
	Money     `json:"money"`
}

// NewStore simple factory of creating the object Store
func NewStore(storeType string) Store {
	ns := Store{}
	ns.Name = "Shop with no Sign"
	owner := character.NewCharacter("", "", "")
	ns.Owner = owner.Name
	ns.Location = "Heldheim"
	ns.StoreType = storeType
	//ns.Inventory = generateInventoryForStore()
	cp := 0
	ss := 0
	gc := 0
	ns.Money = NewMoney(cp, ss, gc)

	return ns
}

//func generateStoreName() string {
//	buf, err := ioutil.ReadFile("data/dnd/names/stores.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	s := string(buf)
//	ss := strings.Split(s, "\n")
//	result := Roll(len(ss))
//	return ss[result-1]
//}

// NewMoney creates object of amount of money for each type
func NewMoney(cp, ss, gc int) Money {
	m := Money{
		CopperPennies:   cp,
		SilverShillings: ss,
		GoldCrowns:      gc,
	}
	return m
}
