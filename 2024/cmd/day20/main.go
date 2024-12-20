package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 20: Race Condition ---")
	fmt.Println()

	start, end, walls := parse((aoc.Input))
	width := len(aoc.Input[0])
	height := len(aoc.Input)
	route := findRoute(start, end, walls)
	savings := findShortcuts(route, walls)

	print(start, end, width, height, make(map[point]int), walls)

	var found int
	var greatShortcuts int
	for k := 0; found < len(savings); k++ {
		if v, ok := savings[k]; ok {
			found++
			fmt.Printf("  - There are % 3d cheats that save % 3d picoseconds\n", v, k)

			if k >= 100 {
				greatShortcuts += v
			}
		}
	}

	fmt.Println()
	fmt.Printf("Picoseconds: %d\n", len(route)+1)
	fmt.Printf("Shortcuts >= 100ps: %d\n", greatShortcuts)
}

func findRoute(start, end point, walls map[point]int) map[point]int {
	queue := []point{start}
	visited := make(map[point]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current] = len(visited)

		if current == end {
			return visited
		}

		for _, offset := range getOffsets(current, 1) {
			if _, found := visited[offset]; found {
				continue
			}

			if _, found := walls[offset]; found {
				continue
			}

			queue = append(queue, offset)
		}
	}

	panic("Cannot find route")
}

func findShortcuts(route, walls map[point]int) map[int]int {
	result := make(map[int]int)

	// Cheat!
	for current, step := range route {
		fmt.Printf("Checking %d,%d\r", current.x, current.y)

		candidates := map[point]int{current: 1}
		for i := 0; i < 2; i++ {
			nextCandidates := make(map[point]int)
			for candidate := range candidates {
				for _, offset := range getOffsets(candidate, 1) {
					if offset == current {
						continue
					}

					// 1st move *must* deviate from route (or we are not cheating).
					_, inWall := walls[offset]
					if i == 0 && inWall {
						nextCandidates[offset]++
					}

					// 2nd move *must* rejoin route.
					routeStep, inRoute := route[offset]
					if i == 1 && inRoute {
						saving := routeStep - step - i - 1
						if saving > 0 {
							result[saving]++

							if aoc.VerboseMode {
								shortcut := make(map[point]int)
								for k, v := range route {
									if v <= step || v >= routeStep {
										shortcut[k] = v
									}
								}

								fmt.Printf("Saving: %d\n", saving)
								print(candidate, offset, 15, 15, shortcut, walls)
							}
						}
					}
				}
			}
			candidates = nextCandidates
		}
	}

	return result
}

func getOffsets(from point, radius int) []point {
	result := []point{}

	for y := radius * -1; y <= radius; y++ {
		for x := radius * -1; x <= radius; x++ {
			move := math.Abs(float64(y)) + math.Abs(float64(x))
			if move <= float64(radius) {
				result = append(result, point{from.x + x, from.y + y})
			}
		}
	}

	return result
}

func parse(input []string) (start, end point, walls map[point]int) {
	walls = make(map[point]int)

	for y, line := range input {
		for x, r := range line {
			switch r {
			case 'S':
				start = point{x, y}
			case 'E':
				end = point{x, y}
			case '#':
				walls[point{x, y}]++
			}
		}
	}

	return start, end, walls
}
