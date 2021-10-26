package store

func generateInventoryForStore() []Item {
	si := make([]Item, 0)
	//Tavern
	//Blacksmith
	//Bowyer
	//Apothecary
	//Clothing
	return si
}

type Inventory struct {
	Items []Items
}

type ItemsInterface interface {
	ItemInterface
	GetQuantity() int
	SetQuantity(int)
}

// ItemInterface is the abstract object of all physical objects
// Price vs value: the value would be the "real" average value overall
//   where price would be the asking price for object.
type ItemInterface interface {
	GetName() string
	GetPrice() Money
	GetWeight() Weight
	GetQuality() Quality
	SetName(string)
	SetPrice(cp, ss, gc int)
	SetWeight()
	SetQuality()
}

// Item contains  of items of multiude of different types
type Item struct {
	Name   string `json:"name"`
	Price  Money  `json:"price"`
	Weight `json:"weight,omitempty"`
	Quality
}

// Weight contains the value and the units for measuring the weight
type Weight struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Money has the types of money
type Money struct {
	CopperPennies   int `json:"Copper_Pennies"`
	SilverShillings int `json:"Silver_Shillings"`
	GoldCrowns      int `json:"Gold_Crowns"`
}

////////Item Quality ////////////

// Quality states the condition of physical objects
type Quality int

const (
	Broken Quality = iota
	Poor
	Average
	Good
	Great
	Perfect
)

///////////////////////////////

type Items struct {
	Item
	Quantity int
}

// Coins /////////////////////////////

//12cp = 1ss
//20ss = 1gc
//240cp = 1gc
