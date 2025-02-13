package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 11: Radioisotope Thermoelectric Generators ---")
	fmt.Println()

	facility := parse(aoc.GetInput(11))
	result := findMinSteps(facility)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findMinSteps(f facility) int {
	minSteps := math.MaxInt

	var iterations int
	var solve func(steps int, f facility, visited map[uint64]bool)
	solve = func(steps int, f facility, visited map[uint64]bool) {
		iterations++

		if !f.isValid() {
			return
		}

		if steps > minSteps {
			return
		}

		if f.isWin() {
			if steps < minSteps {
				minSteps = steps
				fmt.Printf(" - Found new best of %d steps\n", steps)
				return
			}
		}

		moves := f.listMoves()
		for _, move := range moves {
			if !visited[move.state] {
				nextVisited := map[uint64]bool{}
				maps.Copy(nextVisited, visited)
				nextVisited[move.state] = true
				solve(steps+1, move, nextVisited)
			}
		}
	}

	solve(0, f, map[uint64]bool{})

	fmt.Printf(" - Total iterations %d\n", iterations)

	return minSteps
}

func parse(input []string) facility {
	floorMap := map[string]int{
		"first":  0,
		"second": 1,
		"third":  2,
		"fourth": 3,
	}
	ignoreValues := []string{
		"contains",
		"nothing",
		"relevant",
		"and",
		"a",
		"microchip",
		"generator",
	}

	facilityFactory := newFacilityFactory()

	for i := range input {
		elements := strings.Split(strings.ReplaceAll(strings.ReplaceAll(input[i], ",", ""), ".", ""), " ")
		floor := floorMap[elements[1]]

		for _, value := range elements[4:] {
			if !slices.Contains(ignoreValues, value) {
				valueElements := strings.Split(value, "-")

				// chip
				if len(valueElements) > 1 {
					facilityFactory.addMicrochip(valueElements[0], floor)
				}

				// generator
				if len(valueElements) == 1 {
					facilityFactory.addGenerator(valueElements[0], floor)
				}
			}
		}
	}

	return facilityFactory.build()
}
