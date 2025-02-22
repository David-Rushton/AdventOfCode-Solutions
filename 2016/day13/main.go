package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int64
	y int64
}

func main() {
	fmt.Println("--- Day 13: A Maze of Twisty Little Cubicles ---")
	fmt.Println()

	grid := getGrid(int64(aoc.ToInt(aoc.GetInput(13)[0])))
	destination := point{31, 39}
	if aoc.TestMode {
		destination = point{7, 4}
	}

	steps := findBestPath(point{1, 1}, destination, grid)
	visitable := visitableWithinRange(point{1, 1}, grid)

	printGrid(grid, map[point]bool{})

	fmt.Println()
	fmt.Printf("Steps: %d\n", steps)
	fmt.Printf("Visitable: %d\n", visitable)
}

func findBestPath(from, to point, grid map[point]bool) int {
	var result int = math.MaxInt

	queue := []point{from}
	visited := map[point]bool{}
	best := map[point]int{from: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true

		if current == to && best[current] < result {
			result = best[current]
		}

		for _, candidate := range getNeighbours(current) {
			if visited[candidate] {
				continue
			}

			if !grid[candidate] {
				continue
			}

			if _, exists := best[candidate]; !exists {
				best[candidate] = math.MaxInt
			}

			if best[current]+1 < best[candidate] {
				best[candidate] = best[current] + 1
				if !slices.Contains(queue, candidate) {
					queue = append(queue, candidate)
				}
			}
		}
	}

	return result
}

func visitableWithinRange(from point, grid map[point]bool) int {
	visited := map[point]bool{}

	var solve func(from point, steps int)
	solve = func(from point, steps int) {
		if steps > 52 {
			return
		}

		visited[from] = true

		for _, candidate := range getNeighbours(from) {
			if visited[candidate] {
				continue
			}

			if !grid[candidate] {
				continue
			}

			solve(candidate, steps+1)
		}
	}

	solve(from, 0)

	printGrid(grid, visited)

	return len(visited)
}

func getNeighbours(p point) []point {
	return []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

func getGrid(seed int64) map[point]bool {
	result := map[point]bool{}

	var y int64
	var x int64
	for y = 0; y < 100; y++ {
		for x = 0; x < 100; x++ {
			var gridValue int64 = (x * x) + (3 * x) + (2 * x * y) + y + (y * y) + seed
			binaryValue := strconv.FormatInt(int64(gridValue), 2)
			bitsSet := len(strings.ReplaceAll(binaryValue, "0", ""))
			if bitsSet%2 == 0 {
				result[point{x, y}] = true
			}
		}
	}

	return result
}

func printGrid(grid, visited map[point]bool) {
	var y int64
	var x int64
	for y = 0; y < 40; y++ {
		for x = 0; x < 100; x++ {
			p := point{x, y}
			if grid[p] {
				if visited[p] {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			} else {
				if visited[p] {
					panic("Cannot visit a wall")
				}
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
