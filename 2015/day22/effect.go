package main

import "maps"

type Effects map[string]Effect

type Effect struct {
	name   string
	timer  int
	armour int
	damage int
	mana   int
}

func (e Effect) copy() Effect {
	return Effect{
		name:   e.name,
		timer:  e.timer,
		armour: e.armour,
		damage: e.damage,
		mana:   e.mana,
	}
}

func (e Effects) copy() Effects {
	result := Effects{}
	maps.Copy(result, e)
	return result
}

func (e Effects) apply(player *Player) (effects Effects, damageModify int) {
	nextEffects := Effects{}
	damageModify = 0

	player.armour = 0

	for _, effect := range e {
		if effect.timer > 0 {
			nextEffect := effect.copy()

			if nextEffect.armour > 0 {
				player.armour = nextEffect.armour
			}

			player.manaTotal += nextEffect.mana

			damageModify += nextEffect.damage

			nextEffect.timer--
			if nextEffect.timer > 0 {
				nextEffects[nextEffect.name] = nextEffect
			}
		}
	}

	return nextEffects, damageModify
}
