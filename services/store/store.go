package store

type StoreType int

const (
	Tavern StoreType = iota + 1
	Blacksmith
	Apothecary
	Clothing
	GeneralGoods
	Bowyer
)

////////Item Quality ////////////

// Quality states the condition of physical objects
type Quality int

const (
	Poor Quality = iota + 1
	Average
	Good
	Great
)

///////////////////////////////

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

// Item contains  of items of multiude of different types
type Item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Weight   string `json:"weight,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

// Money has the types of money
type Money struct {
	CopperPennies   int `json:"Copper_Pennies"`
	SilverShillings int `json:"Silver_Shillings"`
	GoldCrowns      int `json:"Gold_Crowns"`
}

// NewStore simple factory of creating the object Store
func NewStore(name, owner, location, storeType string, money Money, items ...Item) Store {
	ns := Store{}
	ns.Name = name
	ns.Owner = owner
	ns.Location = location
	ns.StoreType = storeType
	ns.Inventory = items
	ns.Money = money

	return ns
}

// NewMoney creates object of amount of money for each type
func NewMoney(cp, ss, gc int) Money {
	m := Money{
		CopperPennies:   cp,
		SilverShillings: ss,
		GoldCrowns:      gc,
	}
	return m
}
