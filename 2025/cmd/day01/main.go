package main

import (
	"fmt"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

// too low 2396
func main() {
	fmt.Println("--- Day 1: Secret Entrance ---")
	fmt.Println()

	var zeros int
	var wraps int
	current := 50

	for _, line := range aoc.Input {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		var rotateWraps int
		current, rotateWraps = rotate(current, distance, rune(direction))
		wraps += rotateWraps

		if current == 0 {
			zeros++
		}

		fmt.Printf(" - Rotate %s %d to %d\t%d\n", string(direction), distance, current, wraps)
	}

	fmt.Println()
	fmt.Printf("Zeros %d\n", zeros)
	fmt.Printf("Wraps %d\n", wraps)
}

func rotate(from, distance int, direction rune) (int, int) {
	increment := 1
	if direction == 'L' {
		increment = -1
	}

	wraps := 0
	result := from

	for range distance {
		result += increment

		switch {
		case result < 0:
			result = 99
		case result > 99:
			result = 0
		}

		if result == 0 {
			wraps++
		}
	}

	return result, wraps
}
