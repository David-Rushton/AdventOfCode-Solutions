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

	stoneQueue := make(map[int]int)
	for _, stone := range parse(aoc.Input[0]) {
		stoneQueue[stone]++
	}

	generations := 25
	if aoc.Star == aoc.StarTwo {
		generations = 75
	}

	blinkCache := make(map[int][]int)

	for generation := 0; generation < generations; generation++ {
		fmt.Printf("  Generation: %d\n", generation+1)

		nextStoneQueue := make(map[int]int)
		visitCache := make(map[int]int)

		for stone, visits := range stoneQueue {
			if _, found := visitCache[stone]; !found {
				blinkCache[stone] = blink(stone)
			}
			visitCache[stone] += visits

			for _, nextStone := range blinkCache[stone] {
				nextStoneQueue[nextStone] += visitCache[stone]
			}
		}

		stoneQueue = nextStoneQueue
	}

	var total int
	for _, visits := range stoneQueue {
		total += visits
	}

	fmt.Println()
	fmt.Printf("Result: %d", total)
}

// don't.  don't even...
func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	stoneLen := len(strconv.FormatInt(int64(stone), 10))
	if stoneLen%2 == 0 {
		splitAt := int(math.Pow10(stoneLen / 2.0))
		return []int{
			stone / splitAt,
			stone % splitAt,
		}
	}

	return []int{stone * 2024}
}

func parse(input string) []int {
	var result []int

	for _, valueTxt := range strings.Split(input, " ") {
		valueNum, err := strconv.ParseInt(string(valueTxt), 10, 64)
		if err != nil {
			log.Fatalf("Cannot convert %s to a number.", valueTxt)
		}

		result = append(result, int(valueNum))

	}

	return result
}
