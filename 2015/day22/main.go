package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 22: Wizard Simulator 20XX ---")
	fmt.Println()

	player := newPlayer(50, 500)
	boss := newBoss(71, 10)
	spells := getSpells()

	if aoc.TestMode {
		player.hitPoints = 10
		player.manaTotal = 250

		boss.hitPoints = 14
		boss.damage = 8
	}

	lowestManaWin := findLowestManaWin(player, boss, spells)

	fmt.Println()
	fmt.Printf("Lowest Mana Win: %d\n", lowestManaWin)
}

func findLowestManaWin(player Player, boss Boss, spells []Spell) int {
	result := math.MaxInt

	var solve func(Player, Boss, Effects)
	solve = func(player Player, boss Boss, effects Effects) {
		for _, spell := range spells {
			if effect, exists := effects[spell.name]; exists && effect.timer > 1 {
				continue
			}

			if spell.mana > player.manaTotal {
				continue
			}

			if player.manaSpent >= result {
				continue
			}

			nextPlayer, nextBoss, nextEffects := playTurn(player, boss, effects, spell)

			if nextPlayer.hitPoints <= 0 {
				continue
			}

			if nextBoss.hitPoints <= 0 {
				if nextPlayer.manaSpent < result {
					result = nextPlayer.manaSpent
					fmt.Printf(" - New best win found\n   Mana: %d\n   Cast: %v\n\n", nextPlayer.manaSpent, nextPlayer.castHistory)
				}
			}

			solve(nextPlayer, nextBoss, nextEffects)
		}
	}

	solve(player, boss, Effects{})

	return result
}

func playTurn(player Player, boss Boss, effects Effects, spellToCast Spell) (Player, Boss, Effects) {
	// player turn
	//	- apply effects
	//	- cast spell
	//	- check for win
	var damageModifier int
	nextEffects := effects.copy()
	nextEffects, damageModifier = nextEffects.apply(&player)

	player.manaTotal -= spellToCast.mana
	player.manaSpent += spellToCast.mana
	player.castHistory += spellToCast.name + " > "

	if !spellToCast.hasEffect {
		player.hitPoints += spellToCast.hitPoints
		damageModifier += spellToCast.damage
	}

	// We cast the spell during the players turn.
	// Effects are applied from the bosses turn onwards.
	if spellToCast.hasEffect {
		nextEffects[spellToCast.name] = spellToCast.effect
	}

	boss.hitPoints -= player.damage + damageModifier
	if boss.hitPoints <= 0 {
		return player, boss, nextEffects
	}

	// boss turn
	//	- apply effects
	// 	- check for win
	//	- attack
	// 	- check for loss
	nextEffects, damageModifier = nextEffects.apply(&player)
	boss.hitPoints -= damageModifier
	if boss.hitPoints <= 0 {
		return player, boss, nextEffects
	}

	bossAttack := boss.damage - player.armour
	if bossAttack < 1 {
		bossAttack = 1
	}

	player.hitPoints -= bossAttack

	return player, boss, nextEffects
}
