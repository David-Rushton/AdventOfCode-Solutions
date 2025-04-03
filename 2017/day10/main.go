package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 10: Knot Hash ---")
	fmt.Println()

	var lengths = getLengths(aoc.GetInput(10)[0])
	var values = getValues(aoc.TestMode)
	var rounds = 1

	var result = getSparseHash(values, lengths, rounds)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func getSparseHash(values []int, lengths []int, rounds int) int {
	var skip = 0
	var current = 0

	fmt.Printf(" - %v\n", values)

	for range rounds {
		for _, length := range lengths {
			// Reverse section of length from current.
			// Remember this span wraps.
			for i := range length / 2 {
				var firstIdx = (current + i) % len(values)
				var secondIdx = (current + length - i - 1) % len(values)

				// Swap
				var temp = values[firstIdx]
				values[firstIdx] = values[secondIdx]
				values[secondIdx] = temp
			}

			fmt.Printf(" - %v\n", values)

			current = (current + skip + length) % len(values)
			skip++
		}
	}

	return values[0] * values[1]
}

func getValues(testMode bool) []int {
	var result []int

	until := 256
	if testMode {
		until = 5
	}

	for i := range until {
		result = append(result, i)
	}

	return result
}

func getLengths(input string) []int {
	var result []int

	for value := range strings.SplitSeq(strings.Trim(input, "\n\t "), ",") {
		result = append(result, aoc.ToInt(value))
	}

	return result
}
