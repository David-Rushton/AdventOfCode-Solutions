package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 25: Code Chronicle ---")
	fmt.Println()

	keys, locks := parse(aoc.Input)
	fmt.Printf("  Keys:  %d\n", len(keys))
	fmt.Printf("  Locks: %d\n", len(locks))

	fmt.Println()
	fmt.Printf("Result: %v", countCandidates(keys, locks))
}

func countCandidates(keys []key, locks []lock) int {
	var result int

	for _, key := range keys {
		for _, lock := range locks {
			var overlaps int

			for column := 0; column < 5; column++ {
				if key.columns[column]+lock.columns[column] > 5 {
					overlaps++
				}
			}

			if overlaps == 0 {
				result++
			}
		}
	}

	return result
}

func parse(input []string) (keys []key, locks []lock) {
	keys = []key{}
	locks = []lock{}

	next := []int{-1, -1, -1, -1, -1}
	isKey := false
	y := -1

	input = append(input, "")
	for _, line := range input {
		if line == "" {
			// Append current.
			if isKey {
				newKey := key{}
				newKey.columns = make([]int, 5)
				copy(newKey.columns, next)
				keys = append(keys, newKey)
			} else {
				newLock := lock{}
				newLock.columns = make([]int, 5)
				copy(newLock.columns, next)
				locks = append(locks, newLock)
			}

			// Reset for next key/lock.
			next = []int{-1, -1, -1, -1, -1}
			isKey = false
			y = -1
			continue
		}

		y++

		for x, r := range line {
			if y == 0 && x == 0 {
				isKey = r == '#'
			}

			if r == '#' {
				next[x]++
			}
		}
	}

	return keys, locks
}
