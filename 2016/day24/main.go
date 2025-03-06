package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int
	y int
}

func (p *point) getNeighbours() []point {
	return []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

type routeKey struct {
	from point
	to   point
}

func main() {
	fmt.Println("--- Day 24: Air Duct Spelunking ---")
	fmt.Println()

	steps, locations := parse(aoc.GetInput(24))
	stepCount := findShortestRoute(steps, locations, aoc.Star == aoc.StarTwo)

	fmt.Println()
	fmt.Printf("Result: %d\n", stepCount)
}

func findShortestRoute(steps map[point]bool, locations map[int]point, returnHome bool) int {
	// find all locations to visit.
	var toVisit = []int{}
	for k := range locations {
		toVisit = append(toVisit, k)
	}
	slices.Sort(toVisit)

	// Find all possible routes.
	var routes = [][]int{}
	var findRoutes func(visited []int)
	findRoutes = func(visited []int) {
		if len(visited) == len(toVisit) {
			if returnHome {
				visited = append(visited, 0)
			}
			routes = append(routes, visited)
			return
		}

		for i := range toVisit {
			if !slices.Contains(visited, toVisit[i]) {
				nextVisited := make([]int, len(visited))
				copy(nextVisited, visited)
				nextVisited = append(nextVisited, toVisit[i])

				findRoutes(nextVisited)
			}
		}
	}
	findRoutes([]int{0})

	// Find, and cache, steps between locations.
	var stepCountCache = map[routeKey]int{}
	getStepCount := func(from, to point) int {
		if _, exists := stepCountCache[routeKey{from, to}]; !exists {
			// Dijkstra's_algorithm
			// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
			var unvisited = map[point]int{}
			for point := range steps {
				unvisited[point] = math.MaxInt
				if point == from {
					unvisited[point] = 0
				}
			}

			for {
				current, stepCount := getLowestValue(unvisited)
				stepCount++

				if current == to {
					stepCountCache[routeKey{from, to}] = unvisited[to]
					break
				}

				for _, neighbour := range current.getNeighbours() {
					if unvisitedSteps, exists := unvisited[neighbour]; exists {
						if stepCount < unvisitedSteps {
							unvisited[neighbour] = stepCount
						}
					}
				}

				delete(unvisited, current)
			}
		}

		return stepCountCache[routeKey{from, to}]
	}

	// find shortest route
	var shortestRoute = math.MaxInt
	for i := range routes {
		var currentRoute int
		for j := 0; j < len(routes[i])-1; j++ {
			currentRoute += getStepCount(locations[routes[i][j]], locations[routes[i][j+1]])
			if currentRoute > shortestRoute {
				break
			}
		}

		if currentRoute < shortestRoute {
			shortestRoute = currentRoute
			fmt.Printf(" - New shortest route found %v: %v\n", shortestRoute, routes[i])
		}
	}

	return shortestRoute
}

func getLowestValue(m map[point]int) (point, int) {
	var smallestK point
	var smallestV = math.MaxInt

	for k, v := range m {
		if v < smallestV {
			smallestV = v
			smallestK = k
		}
	}

	return smallestK, smallestV
}

func parse(input []string) (steps map[point]bool, locations map[int]point) {
	steps = map[point]bool{}
	locations = map[int]point{}

	for y := range input {
		for x := range input[y] {
			value := string(input[y][x])
			point := point{x, y}

			// Is location.
			if aoc.IsInt(value) {
				locations[aoc.ToInt(value)] = point
				steps[point] = true
			}

			// Is open step.
			if value == "." {
				steps[point] = true
			}
		}
	}

	return steps, locations
}
