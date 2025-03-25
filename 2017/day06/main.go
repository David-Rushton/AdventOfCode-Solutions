package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 6: Memory Reallocation ---")
	fmt.Println()

	var memoryBanks = parse(aoc.GetInput(6))
	var redistributionCycles = balanceMemoryBanks(memoryBanks)
	var cycleFrequency = balanceMemoryBanks(memoryBanks)

	fmt.Println()
	fmt.Printf("Redistribution Cycles: %d\n", redistributionCycles)
	fmt.Printf("Cycles Frequency: %d\n", cycleFrequency)
}

func balanceMemoryBanks(memoryBanks []int) int {
	var iterations int
	var visited = map[string]bool{ToKey(memoryBanks): true}

	for {
		iterations++

		// Find bank with most blocks.  Tie break on lowest index.
		var index = 0
		var maxValue = 0
		for i := range memoryBanks {
			if memoryBanks[i] > maxValue {
				maxValue = memoryBanks[i]
				index = i
			}
		}

		// Redistribute.
		var remaining = memoryBanks[index]
		memoryBanks[index] = 0
		for remaining > 0 {
			index++
			if index >= len(memoryBanks) {
				index = 0
			}

			memoryBanks[index]++
			remaining--
		}

		// Exit if we have already visited this configuration.
		key := ToKey(memoryBanks)
		if visited[key] {
			break
		}
		visited[key] = true
	}

	return iterations
}

func ToKey(memoryBanks []int) string {
	var result string

	for i := range memoryBanks {
		result += fmt.Sprint(memoryBanks[i]) + "-"
	}

	return result
}

func parse(input []string) []int {
	var result []int

	for i := range input {
		for _, strNumber := range strings.Split(input[i], " ") {
			result = append(result, aoc.ToInt(strNumber))
		}
	}

	return result
}
