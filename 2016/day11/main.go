package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type chip struct {
	name  string
	floor int
}

type generator struct {
	name  string
	floor int
}

func main() {
	fmt.Println("--- Day 11: Radioisotope Thermoelectric Generators ---")
	fmt.Println()

	chips, generators := parse(aoc.GetInput(11))
	result := findMinSteps(chips, generators)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

// TODO: Start here.
func findMinSteps(chips []chip, generators []generator) int {
	panic("unimplemented")
}

func parse(input []string) ([]chip, []generator) {
	floors := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
		"fourth": 4,
	}

	ignoreValues := []string{"contains", "nothing", "relevant", "and", "a", "microchip", "generator"}

	chips := []chip{}
	generators := []generator{}

	for i := range input {
		elements := strings.Split(strings.ReplaceAll(strings.ReplaceAll(input[i], ",", ""), ".", ""), " ")
		floor := floors[elements[1]]

		for _, value := range elements[4:] {
			if !slices.Contains(ignoreValues, value) {
				valueElements := strings.Split(value, "-")

				// chip
				if len(valueElements) > 1 {
					chips = append(chips, chip{valueElements[0], floor})
				}

				// generator
				if len(valueElements) == 1 {
					generators = append(generators, generator{value, floor})
				}
			}
		}
	}

	return chips, generators
}
