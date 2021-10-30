package store

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/njdaniel/dnd/services/commands/character"
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

type ShaftMaterial int

const (
	Ash ShaftMaterial = iota + 1
	Cedar
	DouglusFir
	Spruce
	Pine
)

var ShaftMaterialName = []string{
	Ash:        "Ash",
	Cedar:      "Cedar",
	DouglusFir: "Douglus Fir",
	Spruce:     "Spruce",
	Pine:       "Pine",
}

func (s ShaftMaterial) String() string {
	return ShaftMaterialName[s]
}

func (s ShaftMaterial) Len() int {
	return len(ShaftMaterialName)
}

func (s ShaftMaterial) IsValid() bool {
	switch s {
	case Ash, Cedar, DouglusFir, Spruce, Pine:
		return true
	}
	return false
}

func String2ShaftMaterial(s string) (ShaftMaterial, error) {
	for i := 1; i <= ShaftMaterial(i).Len(); i++ {
		if StoreType(i).String() == s {
			return ShaftMaterial(i), nil
		}
	}
	return 0, fmt.Errorf("error: Invalid ShaftMaterial")
}

type FletchMaterial int

const (
	GooseFeathers FletchMaterial = iota + 1
	TurkeyFeathers
	ChickenFeathers
	PidgeonFeathers
)

func (f FletchMaterial) Len() int {
	return 4
}

type Arrow struct {
	Item
	HeadMaterial Metal
	HeadType     ArrowHeadType
	ShaftMaterial
	FletchMaterial
	Length int
}

func NewArrow() ItemInterface {
	hm := Metal(character.Roll(Metal(1).Len()))
	ht := ArrowHeadType(character.Roll(ArrowHeadType(1).Len()))
	sm := ShaftMaterial(character.Roll(ShaftMaterial(1).Len()))
	fm := FletchMaterial(character.Roll(FletchMaterial(1).Len()))
	arrowlengths := [2]int{24, 32}
	arrowweights := [2]float64{0.06, 0.08}
	r := character.Roll(2) - 1
	length := arrowlengths[r]
	na := Arrow{
		Item: Item{
			Name:    hm.String() + " " + ht.String() + " " + strconv.Itoa(length) + "\" arrow",
			Price:   NewMoney(0, 0, 0),
			Weight:  newWeight(arrowweights[r]),
			Quality: Quality(character.Roll(Quality(1).Len() - 1)).String(),
		},
		Length:         length,
		HeadMaterial:   hm,
		HeadType:       ht,
		ShaftMaterial:  sm,
		FletchMaterial: fm,
	}
	return na
}

func generateInventoryForFletcher() (Inventory, error) {
	inv := new(Inventory)
	items := make([]Items, 0)
	//how many times will be added to inventory
	//whats going to be in the inventory?
	//Add arrows
	for dice.Roll(20) > 0 {
		a := NewArrow()
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
	return *inv, nil
}
