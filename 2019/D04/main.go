package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Day 4: Secure Container ---")
	fmt.Println()

	// Rules
	// - six-digit number.
	// - value is within the range given in your puzzle input.
	// - two adjacent digits are the same (like 22 in 122345).
	// - left to right, the digits never decrease (like 111123 or 135679).

	from := 284639
	until := 748759
	starOneMatches := 0
	starTwoMatches := 0

	for current := from; current < until; current++ {
		digits := [...]int{
			current / 100000,
			(current / 10000) % 10,
			(current / 1000) % 10,
			(current / 100) % 10,
			(current / 10) % 10,
			(current / 1) % 10,
		}

		if isMonotonic(digits[:]) && hasAdjacentDigits(digits[:]) {
			// fmt.Printf("\tCandidate: %d\n", current)
			starOneMatches++
		}

		if isMonotonic(digits[:]) && hasAdjacentDigitsExclusive(digits[:]) {
			// fmt.Printf("\tCandidate: %d\n", current)
			starTwoMatches++
		}
	}

	fmt.Println()
	fmt.Printf("Star 1 Matches Found: %d\n", starOneMatches)
	fmt.Printf("Star 2 Matches Found: %d\n", starTwoMatches)
}

func isMonotonic(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}

	return true
}

func hasAdjacentDigits(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func hasAdjacentDigitsExclusive(s []int) bool {
	matches := make(map[int]int)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			matches[s[i]]++
		}
	}

	for _, v := range matches {
		if v == 1 {
			return true
		}
	}

	return false
}
