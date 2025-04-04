package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 10: Knot Hash ---")
	fmt.Println()

	var rounds = 1
	if aoc.Star == aoc.StarTwo {
		rounds = 64
	}

	for _, input := range aoc.GetInput(10) {
		var lengths = getLengths(input, aoc.Star == aoc.StarTwo)
		var values = getValues(aoc.Star != aoc.StarTwo && aoc.TestMode)
		var sparseHash = getSparseHash(values, lengths, rounds)
		var denseHash = getDenseHash(sparseHash)
		var hex = getHex(denseHash)

		fmt.Println()
		fmt.Printf("Input: %v\n", input)
		fmt.Printf("Sparse Hash Checksum: %d\n", sparseHash[0]*sparseHash[1])
		fmt.Printf("Dense Hash: %v\n", denseHash)
		fmt.Printf("Hex: %v (%d)\n", hex, len(hex))
	}
}

func getSparseHash(values []int, lengths []int, rounds int) []int {
	var skip = 0
	var current = 0

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

			current = (current + skip + length) % len(values)
			skip++
		}
	}

	return values
}

func getDenseHash(values []int) []int {
	if len(values) < 16 || len(values)%16 != 0 {
		panic("Values must contain blocks of 16")
	}

	var result []int
	var subResult int

	for i := range values {
		subResult ^= values[i]
		if i%16 == 15 {
			result = append(result, subResult)
			subResult = 0
		}
	}

	return result
}

func getHex(denseHash []int) string {
	var result string

	for i := range denseHash {
		result += fmt.Sprintf("%02x", denseHash[i])
	}

	return result
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

func getLengths(input string, byteMode bool) []int {
	var result []int

	if byteMode {
		for _, r := range strings.Trim(input, "\n\t ") {
			result = append(result, int(r))
		}

		return append(result, 17, 31, 73, 47, 23)
	}

	for value := range strings.SplitSeq(strings.Trim(input, "\n\t "), ",") {
		result = append(result, aoc.ToInt(value))
	}

	return result
}
