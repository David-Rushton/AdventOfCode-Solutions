package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 14: Reindeer Olympics ---")
	fmt.Println()

	reindeers := parse(aoc.GetInput(14))
	bestScore, bestDistance := runRace(reindeers)

	fmt.Println()
	fmt.Printf("Best Distance: %d\n", bestDistance)
	fmt.Printf("Best Score: %d\n", bestScore)
}

func runRace(reindeers []*reindeer) (bestScore, bestDistance int) {
	bestScore = 0
	bestDistance = 0

	// Run race.
	for i := 0; i < 2503; i++ {
		roundDistances := map[int][]*reindeer{}
		roundBestDistance := 0

		for _, current := range reindeers {
			current.Advance()

			currentTravelled := current.GetDistanceTravelled()
			currentList := roundDistances[currentTravelled]
			if currentList == nil {
				currentList = []*reindeer{}
			}
			currentList = append(currentList, current)
			roundDistances[currentTravelled] = currentList

			if currentTravelled > roundBestDistance {
				roundBestDistance = currentTravelled
			}
		}

		for _, current := range roundDistances[roundBestDistance] {
			current.AddPoint()
		}
	}

	// Find winners.
	for _, current := range reindeers {
		fmt.Printf(
			" - %d travelled %d and scored %d\n",
			current.Name,
			current.GetDistanceTravelled(),
			current.GetPoints())

		if current.GetPoints() > bestScore {
			bestScore = current.GetPoints()
		}

		if current.GetDistanceTravelled() > bestDistance {
			bestDistance = current.GetDistanceTravelled()
		}
	}

	return bestScore, bestDistance
}

func parse(input []string) []*reindeer {
	result := []*reindeer{}

	for _, item := range input {
		elements := strings.Split(item, " ")
		result = append(result, &reindeer{
			Name:           elements[0],
			FlightSpeed:    aoc.ToInt(elements[3]),
			FlightDuration: aoc.ToInt(elements[6]),
			RestDuration:   aoc.ToInt(elements[13])})
	}

	return result
}
