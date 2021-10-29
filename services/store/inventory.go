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
	Inventory []Items `json:"inventory"`
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
	GetQuality() string
	SetName(string)
	SetPrice(cp, ss, gc int)
	SetWeight(float64)
	SetQuality(string) error
}

// Item contains  of items of multiude of different types
type Item struct {
	Name    string `json:"name"`
	Price   Money  `json:"price"`
	Weight  `json:"weight,omitempty"`
	Quality string `json:"quality"`
}

//GetName returns the Name of Item
func (i Item) GetName() string {
	return i.Name
}

//GetPrice returns the Money of Item
func (i Item) GetPrice() Money {
	return i.Price
}

//GetWeight returns the Weight of Item
func (i Item) GetWeight() Weight {
	return i.Weight
}

//GetQuality returns the Quality of Item
func (i Item) GetQuality() string {
	return i.Quality
}

//SetName sets the Name of Item
func (i Item) SetName(name string) {
	i.Name = name
}

func (i Item) SetPrice(cp, ss, gc int) {
	i.Price = NewMoney(cp, ss, gc)
}

func (i Item) SetWeight(value float64) {
	i.Weight = newWeight(value)
}

func (i Item) SetQuality(q string) error {
	i.Quality = Quality(1).String()

	return nil
}

// Weight contains the value and the units for measuring the weight
type Weight struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

func newWeight(value float64) Weight {
	return Weight{
		Value: value,
		Unit:  "lbs",
	}
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

var QualityName = []string{
	Broken:  "Broken",
	Poor:    "Poor",
	Average: "Average",
	Good:    "Good",
	Great:   "Great",
	Perfect: "Perfect",
}

func (q Quality) String() string {
	return QualityName[q]
}

func (q Quality) Len() int {
	return len(QualityName)
}

///////////////////////////////

type Items struct {
	Item
	Quantity int
}

// Coins /////////////////////////////

//12cp = 1ss
//20ss = 1gc
//240cp = 1gc
