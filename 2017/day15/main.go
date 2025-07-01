package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 15: Duelling Generators ---")
	fmt.Println()

	var strictMode = aoc.Star == aoc.StarTwo
	var seedA, seedB = getSeeds(aoc.GetInput(15))
	var generatorA = newGenerator(seedA, 16807, 4, strictMode)
	var generatorB = newGenerator(seedB, 48271, 8, strictMode)
	var matches = 0

	var sampleSize = 40_000_000
	if strictMode {
		sampleSize = 5_000_000
	}

	for i := range sampleSize {
		var valueA = generatorA()
		var binaryA = truncate(toBinary(valueA), 16)
		var valueB = generatorB()
		var binaryB = truncate(toBinary(valueB), 16)

		if binaryA == binaryB {
			matches++
		}

		if i%10_000_000 == 0 {
			fmt.Printf("  - Processed %d\n", i)
		}
	}

	fmt.Println()
	fmt.Printf("Matches: %d\n", matches)
}

func toBinary(i int64) string {
	return fmt.Sprintf("%b", i)
}

func truncate(s string, length int) string {
	if len(s) > length {
		return s[len(s)-16:]
	}

	return s
}

func newGenerator(seed, factor, multiple int64, strictMode bool) func() int64 {
	const product int64 = 2147483647
	var last = seed

	return func() int64 {
		for {
			var result = (last * factor) % product
			last = result

			if !strictMode || result%multiple == 0 {
				return result
			}
		}
	}
}

func getSeeds(input []string) (seedA, seedB int64) {
	var rawA = aoc.ToInt(strings.Split(input[0], " ")[4])
	var rawB = aoc.ToInt(strings.Split(input[1], " ")[4])

	return int64(rawA), int64(rawB)
}
