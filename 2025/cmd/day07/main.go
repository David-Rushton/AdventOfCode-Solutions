package main

import (
	"fmt"
	"maps"
	"time"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 7: Laboratories ---")
	fmt.Println()

	grid := map[point]rune{}
	start := point{}

	// Parse input.
	for y, line := range aoc.Input {
		for x, char := range line {
			p := point{x, y}
			grid[p] = char

			if char == 'S' {
				start = p
			}
		}
	}

	splits := getSplits(grid, start)
	timelines := getTimelines(grid, start)

	fmt.Println()
	fmt.Printf("Splits: %d\n", splits)
	fmt.Printf("Timelines: %d\n", timelines)
}

func getSplits(grid map[point]rune, start point) int {
	var splits int

	tracking := map[point]bool{}
	tracking[start] = true

	// Parse input.
	for y, line := range aoc.Input {
		for x, char := range line {
			p := point{x, y}
			grid[p] = char

			if char == 'S' {
				tracking[p] = true
			}
		}
	}

	// Simulate.
	for len(tracking) > 0 {
		newTracking := map[point]bool{}

		for p := range tracking {
			if char, ok := grid[point{p.x, p.y + 1}]; ok {
				grid[p] = '|'

				if char == '^' {
					splits++
					newTracking[point{p.x + 1, p.y + 1}] = true
					newTracking[point{p.x - 1, p.y + 1}] = true
				} else {
					newTracking[point{p.x, p.y + 1}] = true
				}
			}
		}

		// printGrid(grid)
		tracking = maps.Clone(newTracking)
	}

	return splits
}

func getTimelines(grid map[point]rune, start point) int {
	var timelines int

	tracking := map[point]int{start: 1}

	for len(tracking) > 0 {
		newTracking := map[point]int{}

		for p, count := range tracking {
			if char, ok := grid[p]; ok {
				if char == '^' {
					newTracking[point{p.x + 1, p.y + 1}] += count
					newTracking[point{p.x - 1, p.y + 1}] += count
				} else {
					newTracking[point{p.x, p.y + 1}] += count
				}
			} else {
				timelines += count
			}
		}

		tracking = maps.Clone(newTracking)
	}

	return timelines
}

func printGrid(grid map[point]rune) {
	rows := len(aoc.Input)
	cols := len(aoc.Input[0])

	// Ansi escape code for clear screen.
	fmt.Print("\033[H\033[2J")

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			fmt.Print(string(grid[point{x, y}]))
		}
		fmt.Println()
	}

	time.Sleep(time.Millisecond * 50)
}
