package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 19: An Elephant Named Joseph ---")
	fmt.Println()

	var elfCount int = aoc.ToInt(aoc.GetInput(19)[0])

	winnerStealNext := playStealFromNext(initialise(elfCount))
	fmt.Println()
	fmt.Printf("Steal from next won by Elf #%d\n\n", winnerStealNext.id)

	winnerStealOpposite := playStealFromOpposite(initialise(elfCount), elfCount)
	fmt.Println()
	fmt.Printf("Steal from opposite won by Elf #%d\n", winnerStealOpposite.id)
}

func playStealFromNext(current *elf) *elf {
	for !current.isWinner() {
		if aoc.VerboseMode {
			fmt.Printf(" - Elf %d steals from Elf %d\n", current.id, current.next.id)
		}

		current.next = current.next.next
		current = current.next
	}

	return current
}

func playStealFromOpposite(current *elf, elfCount int) *elf {

	for !current.isWinner() {
		stealFrom := current.skip(elfCount / 2)
		stealFrom.previous.next = stealFrom.next
		stealFrom.next.previous = stealFrom.previous

		if aoc.VerboseMode {
			fmt.Printf(" - Elf %d steals from Elf %d\n", current.id, stealFrom.id)
		}

		current = current.next

		elfCount--
	}

	return current
}

func initialise(elfCount int) *elf {
	var end = &elf{id: elfCount}
	var last = end
	var current *elf
	for i := elfCount - 1; i > 0; i-- {
		current = &elf{
			id:   i,
			next: last,
		}
		last.previous = current

		last = current
	}
	current.previous = end
	end.next = current

	return current
}
