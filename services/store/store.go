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
type Item interface {
	Name() string
	Description() string
	Price() string
}

// Inventory contains a list of items of multiude of different types
type Inventory struct {
	Items []Item `json:"items"`
}

// Store represents the physical entity of a business
type Store struct {
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	StoreType StoreType `json:"store_type"`
	Inventory Inventory `json:"inventory"`
}

// NewStore simple factory of creating the object Store
func NewStore(name, owner string, items []Item) Store {
	ns := Store{}
	ns.Name = name
	ns.Owner = owner

	inventory := Inventory{
		Items: items,
	}
	ns.Inventory = inventory

	return ns
}
