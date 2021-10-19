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
	Inventory []Item `json:"inventory"`
	Money     Money  `json:"money"`
}

// Item contains  of items of multiude of different types
type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Weight string `json:"weight"`
}

// Money has the types of money
type Money struct {
	BrassPennies    int `json:"Brass Pennies"`
	SilverShillings int `json:"Silver Shillings"`
	GoldCrowns      int `json:"Gold Crowns"`
}

// NewStore simple factory of creating the object Store
func NewStore(name, owner string, items ...Item) Store {
	ns := Store{}
	ns.Name = name
	ns.Owner = owner

	inventory := items
	ns.Inventory = inventory

	return ns
}
