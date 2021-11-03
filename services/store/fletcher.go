package store

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/njdaniel/dnd/util/dice"
)

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

type WoodType int

const (
	Ash WoodType = iota + 1
	Cedar
	DouglusFir
	Spruce
	Pine
)

var WoodTypeName = [...]string{
	Ash:        "Ash",
	Cedar:      "Cedar",
	DouglusFir: "Douglus Fir",
	Spruce:     "Spruce",
	Pine:       "Pine",
}

func (s WoodType) String() string {
	return WoodTypeName[s]
}

func (s WoodType) Len() int {
	return len(WoodTypeName)
}

func (s WoodType) IsValid() bool {
	switch s {
	case Ash, Cedar, DouglusFir, Spruce, Pine:
		return true
	}
	return false
}

func String2WoodType(s string) (WoodType, error) {
	for i := 1; i <= WoodType(i).Len(); i++ {
		if StoreType(i).String() == s {
			return WoodType(i), nil
		}
	}
	return 0, fmt.Errorf("error: Invalid WoodType")
}

type FletchMaterial string

const (
	GooseFeathers   FletchMaterial = "GooseFeathers"
	TurkeyFeathers                 = "TurkeyFeathers"
	ChickenFeathers                = "ChickenFeathers"
	PidgeonFeathers                = "PidgeonFeathers"
)

var Feathers = [...]string{
	"GooseFeathers",
	"TurkeyFeathers",
	"ChickenFeathers",
	"PidgeonFeathers",
}

func (f FletchMaterial) Len() int {
	return len(Feathers)
}

type Arrow struct {
	Item
	HeadMaterial Metal
	HeadType     ArrowHeadType
	WoodType
	FletchMaterial string
	Length         int
}

func NewArrow(hm Metal) Arrow {
	ht := ArrowHeadType(dice.Roll(ArrowHeadType(1).Len()))
	sm := WoodType(dice.Roll(WoodType(1).Len()))
	fm := Feathers[(dice.Roll(len(Feathers)) - 1)]
	arrowlengths := [2]int{24, 32}
	arrowweights := [2]float64{0.06, 0.08}
	r := dice.Roll(2) - 1
	length := arrowlengths[r]
	na := Arrow{
		Item: Item{
			Name:    hm.String() + " " + ht.String() + " " + strconv.Itoa(length) + "\" arrow",
			Price:   NewMoney(0, 0, 0),
			Weight:  newWeight(arrowweights[r]),
			Quality: Quality(dice.Roll(Quality(1).Len() - 1)).String(),
		},
		Length:         length,
		HeadMaterial:   hm,
		HeadType:       ht,
		WoodType:       sm,
		FletchMaterial: fm,
	}
	return na
}

func generateInventoryForFletcher() (Inventory, error) {
	fmt.Println("inside generateInventoryForFletcher")
	inv := Inventory{}
	items := make([]Items, 0)
	//how many times will be added to inventory
	//whats going to be in the inventory?
	//Add arrows
	for dice.Roll(20) > 1 {
		hm := Metal(dice.Roll(Metal(1).Len()))
		a := NewArrow(hm)
		fmt.Println(a)
		as := Items{
			Item: Item{
				Name:    a.GetName(),
				Price:   a.GetPrice(),
				Weight:  a.GetWeight(),
				Quality: a.GetQuality(),
			},
			Quantity: dice.Roll(100),
		}
		items = append(items, as)
	}
	inv.Inventory = items
	return inv, nil
}
