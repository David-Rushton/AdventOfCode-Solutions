package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type container struct {
	id       int
	capacity int
}

func (c *container) getPath() string {
	return fmt.Sprintf("%v", c.capacity)
}

func main() {
	fmt.Println("--- Day 17: No Such Thing as Too Much ---")
	fmt.Println()

	containers := parse(aoc.GetInput(17))

	targetCapacity := 150
	if aoc.TestMode {
		targetCapacity = 25
	}

	combinations, countOfMinContainers := getCombinations(targetCapacity, containers)

	fmt.Println()
	fmt.Printf("Combinations: %d\n", combinations)
	fmt.Printf("Count of minimum containers: %d\n", countOfMinContainers)
}

func getCombinations(targetCapacity int, containers []container) (combinations, countOfMinContainers int) {
	var combinationsCount int
	var minCombinationsCount int
	minCombinations := math.MaxInt

	var solver func(int, int, int)
	solver = func(index, count, total int) {
		if total == targetCapacity {
			combinationsCount++

			if count < minCombinations {
				minCombinations = count
				minCombinationsCount = 0
			}

			if count == minCombinations {
				minCombinationsCount++
			}

			return
		}

		if total > targetCapacity {
			return
		}

		if total < targetCapacity {
			if index < len(containers) {
				solver(index+1, count+1, total+containers[index].capacity)
				solver(index+1, count, total)
			}
		}
	}

	solver(0, 0, 0)

	return combinationsCount, minCombinationsCount
}

func parse(input []string) []container {
	result := []container{}

	nextId := 2
	for _, capacity := range input {
		result = append(result, container{
			id:       nextId,
			capacity: aoc.ToInt(capacity)})

		nextId *= 2
	}

	return result
}
