package main

//Profession is the object for the professions with advance description
type Profession struct {
	Name string
	Description string
	BonusAdvances []string
	AttributeBonusAdvances []string
	SkillAdvances []string
	Talents []Talent
	StartingMoney []string
	StartingEquip []string
	Trait Talent
}

//Talent is a modifier to a character with a special ability
type Talent struct {
	Name string
	Description string
	Effect string
}

type Equipment interface {
	GetValue() map[string]int
}

type Weapon struct {
	Name string
}