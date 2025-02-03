package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type group struct {
	packages            int
	weight              int
	quantumEntanglement uint64
}

func (g *group) add(item int) group {
	nextQuantumEntanglement := g.quantumEntanglement * uint64(item)
	if g.quantumEntanglement == 0 {
		nextQuantumEntanglement = uint64(item)
	}

	return group{
		packages:            g.packages + 1,
		weight:              g.weight + item,
		quantumEntanglement: nextQuantumEntanglement,
	}
}

func main() {
	fmt.Println("--- Day 24: It Hangs in the Balance ---")
	fmt.Println()

	packages := parse(aoc.GetInput(24))
	groups := 3
	if aoc.Star == aoc.StarTwo {
		groups = 4
	}

	result := findBestArrangement(packages, groups)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findBestArrangement(packages []int, groups int) uint64 {
	var totalWeight int
	for i := range packages {
		totalWeight += packages[i]
	}

	targetGroupA := totalWeight / groups
	targetGroupB := (totalWeight / groups) * (groups - 1)

	var fewestPackages int = math.MaxInt
	var smallestQuantumEntanglement uint64 = math.MaxUint64

	var solve func(index int, groupA, groupB group)
	solve = func(index int, groupA, groupB group) {
		index++

		if groupA.weight > targetGroupA {
			return
		}

		if groupA.packages > fewestPackages {
			return
		}

		if groupB.weight > targetGroupB {
			return
		}

		if index == len(packages) {
			if groupA.weight == targetGroupA {
				if groupB.weight == targetGroupB {
					if groupA.packages < fewestPackages {
						fewestPackages = groupA.packages
						smallestQuantumEntanglement = math.MaxUint64
					}

					if groupA.packages == fewestPackages {
						if groupA.quantumEntanglement < smallestQuantumEntanglement {
							smallestQuantumEntanglement = groupA.quantumEntanglement
							fmt.Printf(
								" - New best.\n   Packages: %d\n   Quantum Entanglement: %d\n\n",
								groupA.packages,
								groupA.quantumEntanglement)
						}
					}
				}
			}
		}

		if index < len(packages) {
			solve(index, groupA.add(packages[index]), groupB)
			solve(index, groupA, groupB.add(packages[index]))
		}
	}

	solve(-1, group{}, group{})

	return smallestQuantumEntanglement
}

func parse(input []string) []int {
	result := []int{}

	for i := range input {
		result = append(result, aoc.ToInt(input[i]))
	}

	return result
}
