package store


type StoreType int

const (
	Tavern StoreType = iota + 1
	Blacksmith
	Alchemy
	Clothing
	General Goods
	Bowyer
)

type Store struct {
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	StoreType StoreType `json:"store_type"`
	Inventory Inventory `json:"inventory"`
}
