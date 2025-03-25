package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 4: High-Entropy Passphrases ---")
	fmt.Println()

	var validPassphrases int
	for _, passphrase := range aoc.GetInput(4) {
		if isValid(passphrase, aoc.Star == aoc.StarTwo) {
			validPassphrases++
			fmt.Printf(" - Valid passphrase: %v\n", passphrase)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d", validPassphrases)
}

func isValid(passphrase string, strictMode bool) bool {
	var phrases = map[string]bool{}
	var sortPhrase = func(phrase string) string {
		var runes = []rune(phrase)
		slices.Sort(runes)
		return string(runes)
	}

	for current := range strings.SplitSeq(passphrase, " ") {
		if strictMode {
			current = sortPhrase(current)
		}

		if phrases[current] {
			return false
		}

		phrases[current] = true
	}

	return true
}
