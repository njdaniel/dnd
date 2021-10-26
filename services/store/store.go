package store

import (
	"errors"
	"fmt"

	"github.com/njdaniel/dnd/services/commands/character"
)

// Item is the generic physical object
//type Item interface {
//	Name() string
//	Description() string
//	Price() string
//}

type StoreType int

const (
	Tavern StoreType = iota + 1
	Fletcher
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

func (s StoreType) String() string {
	return [...]string{"Tavern", "Fletcher"}[s-1]
}

func (s StoreType) Len() int {
	return 0
}

// VerifyEnum verify that string is one of the enum values
func VerifyEnum(s string, enum character.Enum) bool {
	for i := 0; i < enum.Len(); i++ {
		if enum.String() == s {
			return true
		}
	}
	return false
}

///////////////////////
//Builder Pattern

type IStoreBuilder interface {
	setStoreType()
	setName()
	setOwner()
	setLocation()
	setInventory()
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

type Arrow struct {
	Weight
	Price Money
	Quality
	HeadMaterial Metal
	HeadType     ArrowHeadTypes
}

////////////////////

// Store represents the physical entity of a business
type Store struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Location  string `json:"location"`
	StoreType string `json:"store_type"`
	Inventory []Item `json:"inventory"`
	Money     Money  `json:"money"`
}

// NewStore simple factory of creating the object Store
func NewStore(storeType string) Store {
	ns := Store{}
	ns.Name = "Shop with no Sign"
	owner := character.NewCharacter("", "", "")
	ns.Owner = owner.Name
	ns.Location = "Heldheim"
	ns.StoreType = storeType
	ns.Inventory = generateInventoryForStore()
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
