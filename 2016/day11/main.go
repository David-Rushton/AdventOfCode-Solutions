package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 11: Radioisotope Thermoelectric Generators ---")
	fmt.Println()

	facility := parse(aoc.GetInput(11), aoc.Star == aoc.StarTwo)
	result := findMinSteps(facility)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findMinSteps(f facility) int {
	var visited = map[uint64]bool{}
	var unvisited = map[facility]int{
		f: 0,
	}

	var maxCandidates int

	for {
		// fmt.Printf(" - visited %v\r", len(visited))

		if len(unvisited) == 0 {
			panic("Unable to find solution")
		}

		current, stepCount := getSmallest(unvisited)

		if current.isWin() {
			return stepCount
		}

		stepCount++

		// append any neighbours we haven't seen before.
		candidates := current.listMoves()
		if len(candidates) > maxCandidates {
			maxCandidates = len(candidates)
			fmt.Printf(" - New max candidates found %v\n", maxCandidates)
		}

		for _, candidate := range candidates {
			if _, exists := unvisited[candidate]; !exists {
				if !visited[candidate.state] {
					unvisited[candidate] = math.MaxInt
				}
			}

			if unvisitedStepCount, exists := unvisited[candidate]; exists {
				if stepCount < unvisitedStepCount {
					unvisited[candidate] = stepCount
				}
			}
		}

		delete(unvisited, current)
		visited[current.state] = true
	}
}

func getSmallest(m map[facility]int) (facility, int) {
	var smallestK facility
	var smallestV int = math.MaxInt

	for k, v := range m {
		if v < smallestV {
			smallestK = k
			smallestV = v
		}
	}

	return smallestK, smallestV
}

func parse(input []string, addMissingElements bool) facility {
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

	if addMissingElements {
		facilityFactory.addMicrochip("elerium", 0)
		facilityFactory.addGenerator("elerium", 0)
		facilityFactory.addMicrochip("dilithium", 0)
		facilityFactory.addGenerator("dilithium", 0)
	}

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
