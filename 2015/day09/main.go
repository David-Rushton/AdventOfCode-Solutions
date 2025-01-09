package main

import (
	"fmt"
	"maps"
	"math"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type locationName string
type destination struct {
	name     locationName
	distance int
}

func main() {
	fmt.Println("--- Day 9: All in a Single Night ---")
	fmt.Println()

	locationDistances := parse(aoc.GetInput(9))
	shortest := getShortestRoute(locationDistances)
	longest := getLongestRoute(locationDistances)

	fmt.Println()
	fmt.Printf("Shortest: %d\n", shortest)
	fmt.Printf("Shortest: %d\n", longest)
}

func getShortestRoute(locations map[locationName][]destination) int {
	var shortestRoute = math.MaxInt

	type state struct {
		at       locationName
		distance int
		route    string
		visited  map[locationName]bool
	}

	for start := range locations {
		queue := []state{{
			at:       start,
			distance: 0,
			route:    string(start),
			visited:  map[locationName]bool{start: true}}}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current.distance > shortestRoute {
				continue
			}

			if len(current.visited) == len(locations) {
				if current.distance < shortestRoute {
					fmt.Printf(" - Candidate: %v == %d\n", current.route, current.distance)
					shortestRoute = current.distance
				}

				continue
			}

			for _, candidate := range locations[current.at] {
				if !current.visited[candidate.name] {
					newVisited := map[locationName]bool{}
					maps.Copy(newVisited, current.visited)
					newVisited[candidate.name] = true

					queue = append(queue, state{
						at:       candidate.name,
						distance: current.distance + candidate.distance,
						route:    current.route + " > " + string(candidate.name),
						visited:  newVisited})
				}
			}
		}
	}

	return shortestRoute
}

func getLongestRoute(locations map[locationName][]destination) int {
	var longestRoute = 0

	type state struct {
		at       locationName
		distance int
		route    string
		visited  map[locationName]bool
	}

	for start := range locations {
		queue := []state{{
			at:       start,
			distance: 0,
			route:    string(start),
			visited:  map[locationName]bool{start: true}}}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if len(current.visited) == len(locations) {
				if current.distance > longestRoute {
					fmt.Printf(" - Candidate: %v == %d\n", current.route, current.distance)
					longestRoute = current.distance
				}

				continue
			}

			for _, candidate := range locations[current.at] {
				if !current.visited[candidate.name] {
					newVisited := map[locationName]bool{}
					maps.Copy(newVisited, current.visited)
					newVisited[candidate.name] = true

					queue = append(queue, state{
						at:       candidate.name,
						distance: current.distance + candidate.distance,
						route:    current.route + " > " + string(candidate.name),
						visited:  newVisited})
				}
			}
		}
	}

	return longestRoute
}

func parse(input []string) map[locationName][]destination {
	result := map[locationName][]destination{}

	add := func(from, to locationName, distance int) {
		var destinations []destination
		var exists bool
		if destinations, exists = result[from]; !exists {
			destinations = []destination{}
		}

		destinations = append(destinations, destination{
			name:     locationName(to),
			distance: distance})

		result[from] = destinations
	}

	for _, line := range input {
		if line != "" {
			elements := strings.Split(line, " ")

			add(locationName(elements[0]), locationName(elements[2]), aoc.ToInt(elements[4]))
			add(locationName(elements[2]), locationName(elements[0]), aoc.ToInt(elements[4]))
		}
	}

	return result
}
