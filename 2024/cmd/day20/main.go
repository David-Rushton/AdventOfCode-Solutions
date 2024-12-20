package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type shortcut struct {
	start point
	end   offset
}

type offset struct {
	point    point
	distance int
}

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

	radius := 2
	if aoc.Star == aoc.StarTwo {
		radius = 20
	}

	route := findRoute(start, end, walls)
	savings := findShortcuts(route, walls, radius)

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
			if _, found := visited[offset.point]; found {
				continue
			}

			if _, found := walls[offset.point]; found {
				continue
			}

			queue = append(queue, offset.point)
		}
	}

	panic("Cannot find route")
}

func findShortcuts(route, walls map[point]int, radius int) map[int]int {
	// Cheat!
	shortcuts := make(map[shortcut]int)
	for current, step := range route {

		offsets := getOffsets(current, radius)
		for _, offset := range offsets {
			routeStep, inRoute := route[offset.point]
			if inRoute {
				saving := routeStep - step - offset.distance
				if saving > 0 {
					shortcuts[shortcut{current, offset}] = saving
				}
			}
		}
	}

	// Transform to summary.
	result := make(map[int]int)
	for _, saving := range shortcuts {
		result[saving]++
	}

	return result
}

func getOffsets(from point, radius int) []offset {
	result := []offset{}

	for y := radius * -1; y <= radius; y++ {
		for x := radius * -1; x <= radius; x++ {
			candidatePoint := point{from.x + x, from.y + y}
			candidate := offset{
				candidatePoint,
				getDistance(from, candidatePoint),
			}

			if candidate.distance > 0 && candidate.distance <= radius {
				result = append(result, candidate)
			}
		}
	}

	return result
}

func getDistance(from, until point) int {
	xDistance := math.Abs(float64(from.x - until.x))
	yDistance := math.Abs(float64(from.y - until.y))
	return int(xDistance + yDistance)
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
