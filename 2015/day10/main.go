package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 10: Elves Look, Elves Say ---")
	fmt.Println()

	seed := toChunks("1321131112")
	iterations := 40

	if aoc.TestMode {
		seed = toChunks("1")
		iterations = 5
	}

	if aoc.Star == aoc.StarTwo {
		iterations = 50
	}

	current := seed
	for i := 0; i < iterations; i++ {
		fmt.Printf("  - % 3d >> %v\n", i, current.getLen())
		current = getNextV2(current)
	}
	fmt.Printf("  - % 3d >> %v\n", iterations, current.getLen())

	fmt.Println()
	fmt.Printf("Result: %d\n", current.getLen())
}

func toChunks(seed string) chunks {
	result := newChunks()

	for _, r := range seed {
		result.add(aoc.ToInt(string(r)))
	}

	result.flush()

	return result
}

func getNextV2(values chunks) chunks {
	result := newChunks()

	for _, item := range values.items {
		result.add(item.count)
		result.add(item.number)
	}

	result.flush()

	return result
}

func getNext(value string) string {
	buffer := ""
	counter := 0
	current := 'x'

	// HACK: We apply a trailing character to force the final flush to the buffer.
	for _, r := range value + "-" {
		if r != current {
			if counter > 0 {
				buffer += fmt.Sprintf("%v", counter) + string(current)
			}

			counter = 1
			current = r
		} else {
			counter++
		}
	}

	return buffer
}
