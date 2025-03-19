package main

import (
	"fmt"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 25: Clock Signal ---")
	fmt.Println()

	instructions := aoc.GetInput(25)

	var target = []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	var seed int
	for {
		current := runProgram(seed, instructions)

		fmt.Printf(" - Seed: %d == %v\n", seed, current)

		if slices.Equal(current, target) {
			break
		}

		seed++
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", seed)
}
