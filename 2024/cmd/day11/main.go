package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 11: Plutonian Pebbles ---")
	fmt.Println()

	stones := parse(aoc.Input[0])
	fmt.Printf("Initial arrangement:\n%v\n\n", stones)

	for i := 1; i < 76; i++ {
		stones = blink(stones)
		fmt.Printf("After %d blinks: %v\n", i, len(stones))
	}

	fmt.Println()
	fmt.Printf("Result: %d", len(stones))
}

// don't.  don't even...
func blink(stones []int64) []int64 {
	var result []int64

	for _, stone := range stones {
		if stone == 0 {
			result = append(result, 1)
			continue
		}

		stoneLen := len(strconv.FormatInt(stone, 10))
		if stoneLen%2 == 0 {
			splitAt := int64(math.Pow10(stoneLen / 2.0))
			result = append(
				result,
				stone/splitAt,
				stone%splitAt)
			continue
		}

		result = append(result, stone*2024)
	}

	return result
}

func parse(input string) []int64 {
	var result []int64

	for _, valueTxt := range strings.Split(input, " ") {
		valueNum, err := strconv.ParseInt(string(valueTxt), 10, 64)
		if err != nil {
			log.Fatalf("Cannot convert %s to a number.", valueTxt)
		}

		result = append(result, valueNum)

	}

	return result
}
