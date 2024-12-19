package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 19: Linen Layout ---")
	fmt.Println()

	towels, patterns := parse(aoc.Input)
	var permutationsTotal int
	var possibleCount int

	for _, pattern := range patterns {
		permutations := getPermutations(towels, pattern)
		if permutations > 0 {
			permutationsTotal += permutations
			possibleCount++
		}
		fmt.Printf("  Permutations: % 15d | Pattern: %s\n", permutations, pattern)
	}

	fmt.Println()
	fmt.Printf("Possible: %d\n", possibleCount)
	fmt.Printf("Permutations: %d\n", permutationsTotal)
}

type state struct {
	pattern string
	score   int
}

func getPermutations(towels []string, pattern string) int {
	var result int

	queue := []state{{pattern: "", score: 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(queue) > 50 {
			fmt.Print("big")
		}

		if current.pattern == pattern {
			result += current.score
			continue
		}

		// Find possible next steps.
		for _, towel := range towels {
			candidate := current.pattern + towel

			if len(candidate) <= len(pattern) && candidate == pattern[0:len(candidate)] {
				queue = append(queue, state{candidate, current.score})
			}
		}

		// Compress queue.
		if len(queue) > 50 {
			patterns := make(map[string]int)
			for _, item := range queue {
				patterns[item.pattern] += item.score
			}

			queue = []state{}
			for k, v := range patterns {
				queue = append(queue, state{k, v})
			}
		}
	}

	return result
}

func parse(input []string) (towels []string, patterns []string) {
	towels = strings.Split(strings.Replace(input[0], ",", "", -1), " ")
	patterns = input[2:]

	return towels, patterns
}
