package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 19: An Elephant Named Joseph ---")
	fmt.Println()

	var elfCount int = aoc.ToInt(aoc.GetInput(19)[0])

	// Set up game.
	var end = &elf{id: elfCount, value: 1}
	var last = end
	var current *elf
	for i := elfCount - 1; i > 0; i-- {
		current = &elf{
			id:    i,
			next:  last,
			value: 1,
		}

		last = current
	}
	end.next = current

	// Play the game.
	for !current.isWinner() {
		if aoc.VerboseMode {
			fmt.Printf(" - Elf %d steals from Elf %d\n", current.id, current.next.id)
		}

		current.next = current.next.next
		current = current.next
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", current.id)
}
