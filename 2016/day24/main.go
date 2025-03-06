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
	stepCount := findShortestRoute(steps, locations)

	fmt.Println()
	fmt.Printf("Result: %d\n", stepCount)
}

func findShortestRoute(steps map[point]bool, locations map[int]point) int {
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
		var minSteps = math.MaxInt
		if _, exists := stepCountCache[routeKey{from, to}]; !exists {
			type state struct {
				from    point
				visited []point
			}

			var queue = []state{{
				from:    from,
				visited: []point{},
			}}
			for len(queue) > 0 {
				fmt.Printf(" - Queue size: % 5d\r", len(queue))

				current := queue[0]
				queue = queue[1:]

				if len(current.visited) > minSteps {
					continue
				}

				if current.from == to {
					if len(current.visited) < minSteps {
						minSteps = len(current.visited)
					}

					continue
				}

				for _, neighbour := range current.from.getNeighbours() {
					if steps[neighbour] && !slices.Contains(current.visited, neighbour) {
						var newVisited = make([]point, len(current.visited))
						copy(newVisited, current.visited)
						newVisited = append(newVisited, current.from)
						queue = append(queue, state{from: neighbour, visited: newVisited})
					}
				}
			}

			stepCountCache[routeKey{from, to}] = minSteps
		}

		return stepCountCache[routeKey{from, to}]
	}

	// find shortest route
	var shortestRoute = math.MaxInt
	for i := range routes {
		fmt.Printf(" - Testing route % 4d of % 4d: %v \r", i+1, len(routes), routes[i])

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
