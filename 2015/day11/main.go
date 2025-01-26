package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 11: Corporate Policy ---")
	fmt.Println("")

	for _, current := range aoc.GetInput(11) {
		if current == "" {
			continue
		}

		next := getNextPassword(current)
		nextNext := getNextPassword(next)
		fmt.Printf(" - %v -> %v -> %v\n", current, next, nextNext)
	}

	fmt.Println("")
	fmt.Println("")
}

func getNextPassword(password string) string {
	const a rune = 'a'
	const z rune = 'z'

	// Convert to array of runes.
	candidate := []rune{}
	for _, r := range password {
		candidate = append(candidate, r)
	}

	// Find next.
	for {
		// Increment.
		position := len(password) - 1
		for {
			candidate[position]++
			if candidate[position] > z {
				candidate[position] = a
				position--
				continue
			}

			break
		}

		if isValidPassword(candidate) {
			return string(candidate)
		}
	}
}

func isValidPassword(candidate []rune) bool {
	var hasIncreasingSection bool
	pairs := map[rune]int{}

	increasingCount := 0
	var last = '!'
	for _, current := range append(candidate, '!') {

		// fail fast.
		if current == 'i' || current == 'o' || current == 'l' {
			return false
		}

		if current == last {
			pairs[current]++
		}

		if getNext(last) == current {
			increasingCount++
		} else {
			if increasingCount >= 3 {
				hasIncreasingSection = true
			}

			increasingCount = 1
		}

		last = current
	}

	return hasIncreasingSection && len(pairs) >= 2
}

func getNext(current rune) rune {
	next := current + 1
	if next == 'i' || next == 'o' || next == 'l' {
		next++
	}
	return next
}
