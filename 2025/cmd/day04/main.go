package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type cell rune

const (
	cellEmpty  cell = '.'
	cellRoll   cell = '@'
	cellRemove cell = 'x'
)

type grid map[point]cell

func (g grid) getNeighbours(p point) []point {
	var result []point

	if _, found := g[p]; found {
		for _, candidate := range p.getNeighbours() {
			if _, found := g[candidate]; found {
				result = append(result, candidate)
			}
		}
	}

	return result
}

type point struct {
	x int
	y int
}

func (p point) getNeighbours() []point {
	return []point{
		{x: p.x - 1, y: p.y - 1},
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y - 1},
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},
		{x: p.x - 1, y: p.y + 1},
		{x: p.x, y: p.y + 1},
		{x: p.x + 1, y: p.y + 1},
	}
}

func main() {
	fmt.Println("--- Day 4: Printing Department ---")
	fmt.Println()

	// var rolls int
	var grid = make(grid)

	for y, line := range aoc.Input {
		for x, char := range line {
			p := point{x, y}
			grid[p] = cell(char)
		}
	}

	var rows = len(aoc.Input)
	var cols = len(aoc.Input[0])
	var rollsRemoved int

	for {
		var accessibleRolls = make(map[point]bool)

		for p, c := range grid {
			var rollsFound int
			for _, neighbour := range grid.getNeighbours(p) {
				if grid[neighbour] == cellRoll {
					rollsFound++
				}
			}

			if rollsFound < 4 && c == cellRoll {
				accessibleRolls[p] = true
			}
		}

		if len(accessibleRolls) == 0 {
			break
		}
		rollsRemoved += len(accessibleRolls)

		fmt.Printf(" - Removed %d rolls of paper\n", len(accessibleRolls))

		for y := range rows {
			for x := range cols {
				point := point{x, y}

				if accessibleRolls[point] {
					grid[point] = cellEmpty
					fmt.Print(string(cellRemove))
				} else {
					fmt.Print(string(grid[point]))
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("Rolls Removed: %d\n", rollsRemoved)
}
