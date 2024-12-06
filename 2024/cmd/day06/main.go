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
	visited := make(map[point]int)
	grid := make(map[point]cell)
	guard := guard{}
	gameOver := false

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
				visited[point]++
				guard.point = point
				guard.direction = r
			}
		}
	}

	// Simulate.
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
				gameOver = true
				break
			}

			if candidatePoint.y < 0 || candidatePoint.y >= len(aoc.Input) {
				gameOver = true
				break
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

			time.Sleep(time.Millisecond * 300)
		}

		if gameOver {
			break
		}
	}

	fmt.Print("\x1b[?25h")
}
