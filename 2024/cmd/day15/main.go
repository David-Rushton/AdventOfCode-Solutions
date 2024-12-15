package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("---  Day 15: Warehouse Woes ---")
	fmt.Println()

	var robot point
	var grid map[point]rune
	var directions []rune

	width := len(aoc.Input[0])
	height := width

	if aoc.Star == aoc.StarOne {
		robot, grid, directions = parse(aoc.Input, false)
	} else {
		width *= 2
		robot, grid, directions = parse(aoc.Input, true)
	}

	printGrid(' ', robot, grid, width, height)

	for i, direction := range directions {
		offset := getOffset(direction)
		robot, grid = moveRobot(robot, offset, grid)

		if aoc.VerboseMode || i == height-1 {
			printGrid(direction, robot, grid, width, height)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", getGPS(grid))
}

func getGPS(grid map[point]rune) int {
	var result int

	for k, v := range grid {
		if v == 'O' || v == '[' {
			result += 100*k.y + k.x
		}
	}

	return result
}

func moveRobot(robot, offset point, grid map[point]rune) (point, map[point]rune) {
	candidateRobot := point{
		x: robot.x + offset.x,
		y: robot.y + offset.y,
	}

	// Robot can move into cell.
	if _, found := grid[candidateRobot]; !found {
		return candidateRobot, grid
	}

	// We can't move.
	// Exit early.
	if cell := grid[candidateRobot]; cell == '#' {
		return robot, grid
	}

	// Can we move left or right?
	if offset.x != 0 {
		xMoves := []point{}
		for x := 1; ; x++ {
			nextPoint := point{robot.x + (offset.x * x), robot.y}

			if grid[nextPoint] == '#' {
				return robot, grid
			}

			xMoves = append(xMoves, nextPoint)

			if _, found := grid[nextPoint]; !found {
				for i := len(xMoves) - 1; i > 0; i-- {
					grid[xMoves[i]] = grid[xMoves[i-1]]
				}
				delete(grid, xMoves[0])

				return candidateRobot, grid
			}
		}
	}

	// Can we move up or down?
	if offset.y != 0 {
		fromX := robot.x
		toX := robot.x
		nextY := robot.y + offset.y
		movesY := []point{}
		for y := 1; ; y++ {
			if grid[point{fromX, nextY}] == ']' {
				fromX--
			}

			if grid[point{toX, nextY}] == '[' {
				toX++
			}

			for ; grid[point{fromX, nextY}] == 0; fromX++ {
			}

			for ; grid[point{toX, nextY}] == 0; toX-- {
			}

			var boxesFound bool
			for x := fromX; x <= toX; x++ {
				nextPoint := point{x, nextY}
				cell := grid[nextPoint]

				if cell == '#' {
					return robot, grid
				}

				if cell == 'O' || cell == '[' || cell == ']' {
					boxesFound = true
					movesY = append(movesY, nextPoint)
				}
			}

			if !boxesFound {
				for i := len(movesY) - 1; i >= 0; i-- {
					box := movesY[i]
					grid[point{box.x, box.y + offset.y}] = grid[box]
					delete(grid, box)
				}

				return candidateRobot, grid
			}

			nextY += offset.y
		}
	}

	panic("Cannot move robot")
}

func getOffset(direction rune) point {
	switch direction {
	case '^':
		return point{x: +0, y: -1}
	case '>':
		return point{x: +1, y: +0}
	case 'v':
		return point{x: +0, y: +1}
	case '<':
		return point{x: -1, y: +0}
	}

	panic("Unknown direction")
}

func parse(input []string, doubleWidthMode bool) (robot point, grid map[point]rune, directions []rune) {
	grid = make(map[point]rune)
	directions = []rune{}

	// Grid.
	var directionMode bool
	validDirections := []rune{'^', '>', 'v', '<'}
	for y, line := range input {
		if line == "" {
			directionMode = true
			continue
		}

		processedLine := ""
		if !directionMode && doubleWidthMode {
			for _, r := range line {
				switch r {
				case '#':
					processedLine += "##"
				case '@':
					processedLine += "@."
				case 'O':
					processedLine += "[]"
				case '.':
					processedLine += ".."
				}
			}
		} else {
			processedLine = line
		}

		for x, r := range processedLine {
			// Directions.
			if directionMode {
				if !slices.Contains(validDirections, r) {
					log.Fatalf("Cannot parse direction %v", string(r))
				}

				directions = append(directions, r)
				continue
			}

			// Grid.
			if r == '#' || r == 'O' || r == '[' || r == ']' {
				grid[point{x, y}] = r
			}

			if r == '@' {
				robot = point{x, y}
			}
		}
	}

	return robot, grid, directions
}
