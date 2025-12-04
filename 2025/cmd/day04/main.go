package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type cell rune

const (
	cellEmpty cell = '.'
	cellRoll  cell = '@'
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

	var rolls int
	var grid = make(grid)

	for y, line := range aoc.Input {
		for x, char := range line {
			p := point{x, y}
			grid[p] = cell(char)
		}
	}

	for p, c := range grid {
		var rollsFound int
		for _, neighbours := range grid.getNeighbours(p) {
			if grid[neighbours] == cellRoll {
				rollsFound++
			}
		}

		if rollsFound < 4 && c == cellRoll {
			rolls++
		}
	}

	fmt.Println()
	fmt.Printf("Accessible Rolls: %d\n", rolls)
}
