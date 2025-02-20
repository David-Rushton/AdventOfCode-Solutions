package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type disc struct {
	id             int
	startPosition  int
	totalPositions int
}

func main() {
	fmt.Println("--- Day 15: Timing is Everything ---")
	fmt.Println()

	discs := parse(aoc.GetInput(15))
	if aoc.Star == aoc.StarTwo {
		discs = append(discs, disc{
			id:             discs[len(discs)-1].id + 1,
			startPosition:  0,
			totalPositions: 11})
	}

	result := findDropTime(discs)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findDropTime(discs []disc) int {
	for i := 0; ; i++ {
		winning := true

		for _, disc := range discs {
			time := i + disc.id
			position := (disc.startPosition + time) % disc.totalPositions
			if position == 0 {
				indent := strings.Repeat(" ", disc.id)
				fmt.Printf(" %v- Passed slot %d at time %d\n", indent, disc.id, i)
			} else {
				winning = false
				break
			}
		}

		if winning {
			return i
		}
	}
}

func parse(input []string) []disc {
	result := []disc{}

	for i := range input {
		elements := strings.Split(strings.ReplaceAll(input[i], ".", ""), " ")
		result = append(result, disc{
			id:             i + 1,
			startPosition:  aoc.ToInt(elements[11]),
			totalPositions: aoc.ToInt(elements[3])})
	}

	return result
}
