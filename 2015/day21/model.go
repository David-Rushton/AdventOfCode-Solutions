package main

type combatant struct {
	hitPoints int
	damage    int
	armour    int
}

type weapon item
type armour item
type ring item
type item struct {
	name   string
	cost   int
	damage int
	armour int
}

type loadout struct {
	weapon    weapon
	armour    armour
	leftRing  ring
	rightRing ring
}

func (lo *loadout) getDamage() int {
	return lo.weapon.damage + lo.leftRing.damage + lo.rightRing.damage
}

func (lo *loadout) getArmour() int {
	return lo.armour.armour + lo.leftRing.armour + lo.rightRing.armour
}

func (lo *loadout) getCost() int {
	return lo.weapon.cost + lo.armour.cost + lo.leftRing.cost + lo.rightRing.cost
}
