package main

import (
	"fmt"
	"math"
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
	calibration := getCalibration(replacements, molecule)
	fabrication := getFabrication(replacements, molecule)

	fmt.Println()
	fmt.Printf("Calibration: %d\n", calibration)
	fmt.Printf("Fabrication: %d\n", fabrication)
}

func getCalibration(replacements []replacement, molecule string) int {
	result := map[string]bool{}

	for _, replacement := range replacements {
		start := strings.Index(molecule, replacement.section)
		for start != -1 {
			candidate := molecule[0:start]
			candidate += replacement.with
			candidate += molecule[start+len(replacement.section):]

			if !result[candidate] {
				result[candidate] = true
				fmt.Printf(" - Molecule found %v\n", candidate)
			}

			start++
			start = indexAt(molecule, replacement.section, start)
		}
	}

	return len(result)
}

// TODO: Returns correct result, but then hangs.
// Suspect we are stuck in some endless solve recursion.
func getFabrication(replacements []replacement, molecule string) int {
	const target = "e"
	bestSteps := math.MaxInt
	var solve func(string, int)

	solve = func(current string, steps int) {
		if current == target {
			if steps < bestSteps {
				bestSteps = steps
				fmt.Printf(" - Molecule fabricated in %d steps\n", steps)
			}

			return
		}

		if steps >= bestSteps {
			return
		}

		for _, replacement := range replacements {
			if idx := strings.Index(current, replacement.with); idx > -1 {
				candidate := current[0:idx]
				candidate += replacement.section
				candidate += current[idx+len(replacement.with):]

				solve(candidate, steps+1)
			}
		}
	}

	solve(molecule, 0)

	return bestSteps
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
