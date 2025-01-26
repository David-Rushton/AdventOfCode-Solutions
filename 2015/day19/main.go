package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type replacement struct {
	section string
	with    string
}

func main() {
	fmt.Println("--- Day 19: Medicine for Rudolph ---")
	fmt.Println()

	replacements, molecule := parse(aoc.GetInput(19))
	result := countMolecules(replacements, molecule)

	fmt.Println()
	fmt.Printf("Result: %d", result)
}

func countMolecules(replacements []replacement, molecule string) int {
	result := map[string]bool{}

	for _, replacement := range replacements {
		start := strings.Index(molecule, replacement.section)
		for start != -1 {
			candidate := molecule[0:start]
			candidate += replacement.with
			candidate += molecule[start+len(replacement.section):]

			if !result[candidate] {
				result[candidate] = true
				fmt.Printf(" - Molecule found: %v\n", candidate)
			}

			start++
			start = indexAt(molecule, replacement.section, start)
		}
	}

	return len(result)
}

func indexAt(s, substr string, n int) int {
	idx := strings.Index(s[n:], substr)
	if idx > -1 {
		idx += n
	}

	return idx
}

func parse(input []string) (replacements []replacement, molecule string) {
	replacements = []replacement{}

	for _, current := range input {
		if current != "" {
			elements := strings.Split(current, "=>")
			if len(elements) == 2 {
				replacements = append(replacements, replacement{
					section: strings.Trim(elements[0], " "),
					with:    strings.Trim(elements[1], " ")})
			} else {
				molecule = current
			}
		}
	}

	return replacements, molecule
}
