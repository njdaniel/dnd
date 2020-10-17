package main

//Profession is the object for the professions with advance description
type Profession struct {
	Name string
	Description string
	BonusAdvances []string
	AttributeBonusAdvances []string
	SkillAdvances []string
	Talents []Trait
	StartingMoney string
	StartingEquip []string
	UniqueTrait Trait
}

//Talent is a modifier to a character with a special ability
type Trait struct {
	Name string
	Description string
	Effect string
}

type Equipment interface {
	GetValue() string
}

type Weapon struct {
	Name string
}