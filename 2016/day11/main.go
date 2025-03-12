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
	var stepIndex = map[int]int{}
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

		current, stepCount := getSmallest(unvisited, stepIndex)

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
			if visited[candidate.state] {
				continue
			}

			if unvisitedStepCount, exists := unvisited[candidate]; exists {
				if stepCount < unvisitedStepCount {
					unvisited[candidate] = stepCount
					stepIndex[stepCount]++
				}
			} else {
				unvisited[candidate] = stepCount
				stepIndex[stepCount]++
			}

		}

		stepIndex[stepCount-1]--
		delete(unvisited, current)
		visited[current.state] = true
	}
}

func getSmallest(m map[facility]int, stepIndex map[int]int) (facility, int) {
	var smallestK facility
	var smallestV int = math.MaxInt

	var i int
	for i = 1; i < 1_000; i++ {
		if stepIndex[i] > 0 {
			break
		}
	}

	for k, v := range m {
		if i < 1_000 && v == i {
			return k, v
		}

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
