package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 21: RPG Simulator 20XX ---")
	fmt.Println()

	player := getPlayer(aoc.TestMode)
	boss := getBoss(aoc.TestMode)
	loadouts := getLoadouts()

	bestWin, worstLost := getResults(player, boss, loadouts)

	fmt.Println()
	fmt.Printf("Best Win: %d\n", bestWin)
	fmt.Printf("Worst Lost: %d\n", worstLost)
}

func getResults(player, boss combatant, loadouts []loadout) (bestWin, worstLost int) {
	bestWin = math.MaxInt
	worstLost = 0

	for _, loadout := range loadouts {
		goldSpent := loadout.getCost()
		if playerWins(player, boss, loadout) {
			if goldSpent < bestWin {
				bestWin = goldSpent
				fmt.Printf(
					" - New best loadout found.  Cost: % 4d.  Equipped: %v, %v, %v & %v.\n",
					goldSpent,
					loadout.weapon.name,
					loadout.armour.name,
					loadout.leftRing.name,
					loadout.rightRing.name)
			}
		} else {
			if goldSpent > worstLost {
				worstLost = goldSpent
				fmt.Printf(
					" - New worst loadout found.  Cost: % 4d.  Equipped: %v, %v, %v & %v.\n",
					goldSpent,
					loadout.weapon.name,
					loadout.armour.name,
					loadout.leftRing.name,
					loadout.rightRing.name)
			}
		}
	}

	return bestWin, worstLost
}

func playerWins(player, boss combatant, loadout loadout) bool {
	player.damage = loadout.getDamage()
	player.armour = loadout.getArmour()

	playTurn := func(attacker, defender *combatant) {
		attack := attacker.damage - defender.armour
		if attack <= 0 {
			attack = 1
		}
		defender.hitPoints -= attack
	}

	for {
		playTurn(&player, &boss)
		if boss.hitPoints <= 0 {
			return true
		}

		playTurn(&boss, &player)
		if player.hitPoints <= 0 {
			return false
		}
	}
}

func getPlayer(testMode bool) combatant {
	if testMode {
		return combatant{hitPoints: 8}
	}

	return combatant{hitPoints: 100}
}

func getBoss(testMode bool) combatant {
	if testMode {
		return combatant{hitPoints: 12, damage: 7, armour: 2}
	}

	return combatant{hitPoints: 109, damage: 8, armour: 2}
}
