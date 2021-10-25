package store

import "github.com/njdaniel/dnd/services/commands/character"

// Item is the generic physical object
//type Item interface {
//	Name() string
//	Description() string
//	Price() string
//}

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
func NewStore(name, storeType string) Store {
	ns := Store{}
	ns.Name = name
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
