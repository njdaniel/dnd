package store

import (
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

type iStoreBuilder interface {
	setStoreType()
	setName()
	setOwner()
	setLocation()
	setInventory()
}

func getStoreBuilder(storeBuilderType string) iStoreBuilder {
	fmt.Println(storeBuilderType)
	isb := fletcherBuilder{}
	return isb
}

type fletcherBuilder struct {
	Name string
	StoreType
	Owner    string
	Location string
}

func (f fletcherBuilder) setStoreType() {

}

func (f fletcherBuilder) setName() {

}

func (f fletcherBuilder) setOwner() {

}

func (f fletcherBuilder) setLocation() {

}

func (f fletcherBuilder) setInventory() {

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
