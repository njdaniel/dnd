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

type Quality int

const (
	Poor Quality = iota + 1
	Average
	Good
	Great
)

///////////////////////////////

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quality     Quality `json:"quality"`
}

type Inventory struct {
	Items []Item `json:"items"`
}

type Store struct {
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	StoreType StoreType `json:"store_type"`
	Inventory Inventory `json:"inventory"`
}

func NewStore() Store {
	ns := Store{}

	return ns
}
