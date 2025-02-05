package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 6: Signals and Noise ---")
	fmt.Println()

	charCountByPos := []map[string]int{}
	maxCountByPos := []int{}
	maxFrequentByPos := []string{}
	for _, signal := range aoc.GetInput(6) {
		for len(charCountByPos) < len(signal) {
			charCountByPos = append(charCountByPos, map[string]int{})
			maxCountByPos = append(maxCountByPos, 0)
			maxFrequentByPos = append(maxFrequentByPos, "")
		}

		for pos, r := range signal {
			char := string(r)
			charCountByPos[pos][char]++

			if charCountByPos[pos][char] > maxCountByPos[pos] {
				maxCountByPos[pos] = charCountByPos[pos][char]
				maxFrequentByPos[pos] = char
				fmt.Printf(" - Updated max character at %d to %v\n", pos, char)
			}
		}
	}

	minCountByPos := make([]int, len(maxCountByPos))
	minFrequentByPos := make([]string, len(maxFrequentByPos))
	for pos := range charCountByPos {
		for char, count := range charCountByPos[pos] {
			if minCountByPos[pos] == 0 || minCountByPos[pos] > count {
				minCountByPos[pos] = count
				minFrequentByPos[pos] = char
				fmt.Printf(" - Updated min character at %d to %v\n", pos, char)
			}
		}
	}

	fmt.Println()
	fmt.Printf("Error corrected message (min): %v\n", strings.Join(minFrequentByPos, ""))
	fmt.Printf("Error corrected message (max): %v\n", strings.Join(maxFrequentByPos, ""))
}
