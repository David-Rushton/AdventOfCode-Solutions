package main

import (
	"fmt"
	"time"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

type cell struct {
	point       point
	visited     bool
	hasObstacle bool
}

type guard struct {
	point     point
	direction rune
}

var (
	offsets = map[rune]point{
		'^': {x: +0, y: -1},
		'>': {x: +1, y: +0},
		'v': {x: +0, y: +1},
		'<': {x: -1, y: +0},
	}
)

func main() {
	grid := make(map[point]cell)
	guard := guard{}

	// Build grid.
	for y, line := range aoc.Input {
		for x, r := range line {
			point := point{x, y}
			grid[point] = cell{
				point:       point,
				visited:     r == '^',
				hasObstacle: r == '#',
			}

			if grid[point].visited {
				guard.point = point
				guard.direction = r
			}
		}
	}

	// Simulate.
	_, visits := Simulate(guard, grid)
	loopsFound := findLoop(visits, guard, grid)

	fmt.Println()
	fmt.Printf("Loops found: %d\n", loopsFound)
}

type exitReason int

const (
	leftGrid exitReason = iota
	loop
)

func findLoop(visited map[point]int, guard guard, grid map[point]cell) int {
	var result int

	var skippedFirst bool
	for candidateObstacle := range visited {
		if !skippedFirst {
			// We can't obstruct the starting cell.
			skippedFirst = true
			continue
		}

		// copy grid
		candidateGrid := make(map[point]cell)
		for _, c := range grid {
			candidateGrid[c.point] = c

			if c.point == candidateObstacle {
				candidateGrid[c.point] = cell{candidateObstacle, false, true}
			}
		}

		if reason, _ := Simulate(guard, candidateGrid); reason == loop {
			result++
		}
	}

	return result
}

func Simulate(guard guard, grid map[point]cell) (reason exitReason, visits map[point]int) {
	visited := make(map[point]int)
	visited[guard.point]++

	fmt.Println("\x1b[2J")
	fmt.Print("\x1b[?25l")
	for {
		fmt.Println("\x1b[1;1H")
		fmt.Println("--- Day 6: Guard Gallivant ---")
		fmt.Printf("Visited: %d\n\n", len(visited))

		// Move Guard.
		for {
			candidatePoint := point{
				x: guard.point.x + offsets[guard.direction].x,
				y: guard.point.y + offsets[guard.direction].y,
			}

			if candidatePoint.x < 0 || candidatePoint.x >= len(aoc.Input[0]) {
				fmt.Print("\x1b[?25h")
				return leftGrid, visited
			}

			if candidatePoint.y < 0 || candidatePoint.y >= len(aoc.Input) {
				fmt.Print("\x1b[?25h")
				return leftGrid, visited
			}

			if grid[candidatePoint].hasObstacle {
				switch guard.direction {
				case '^':
					guard.direction = '>'
				case '>':
					guard.direction = 'v'
				case 'v':
					guard.direction = '<'
				case '<':
					guard.direction = '^'
				}

				continue
			}

			cell := grid[candidatePoint]
			cell.visited = true
			grid[candidatePoint] = cell
			guard.point = candidatePoint

			visited[guard.point]++
			if visited[guard.point] > 4 {
				return loop, visited
			}

			break
		}

		// Print Grid.
		if aoc.VerboseMode {
			for y := 0; y < len(aoc.Input); y++ {
				for x := 0; x < len(aoc.Input[0]); x++ {
					cell := grid[point{x, y}]

					var value string
					switch {
					case cell.point == guard.point:
						value = string(guard.direction)
					case cell.hasObstacle:
						value = "#"
					case cell.visited:
						value = "x"
					default:
						value = "."
					}

					fmt.Print(value)
				}
				fmt.Println()
			}

			time.Sleep(time.Millisecond * 100)
		}
	}
}
