package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 4: Ceres Search ---")
	fmt.Println()

	// build grid
	grid := [][]rune{}
	for y, line := range aoc.Input {
		grid = append(grid, []rune{})

		for _, r := range line {
			grid[y] = append(grid[y], r)
		}
	}

	// scan grid
	xmasTotal := 0
	masTotal := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			xmasTotal += findXMas(grid, x, y)
			masTotal += findMas(grid, x, y)
		}
	}

	fmt.Printf("XMAS Hits: %d\n", xmasTotal)
	fmt.Printf("MAS Hits: %d\n", masTotal)
}

func findXMas(grid [][]rune, x, y int) int {
	offsets := []struct {
		x int
		y int
	}{
		{y: 1, x: 0},
		{y: 0, x: 1},
		{y: 1, x: 1},
		{y: 1, x: -1},
	}

	if grid[y][x] != 'X' {
		return 0
	}

	var found int
	for _, multiplier := range []int{1, -1} {
		for _, offset := range offsets {
			for i, expected := range []rune{'X', 'M', 'A', 'S'} {
				testY := y + (offset.y * i * multiplier)
				testX := x + (offset.x * i * multiplier)

				if testY < 0 || testX < 0 {
					break
				}

				if !(testY < len(grid) && testX < len(grid[testY])) {
					break
				}

				if grid[testY][testX] != expected {
					break
				}

				if grid[testY][testX] == 'S' {
					found++
				}
			}
		}
	}

	return found
}

func findMas(grid [][]rune, x, y int) int {
	if grid[y][x] != 'A' {
		return 0
	}

	if y < 1 || y > len(grid)-2 {
		return 0
	}

	if x < 1 || x > len(grid[y])-2 {
		return 0
	}

	isMS := func(left, right rune) bool {
		if left == 'M' && right == 'S' {
			return true
		}

		if left == 'S' && right == 'M' {
			return true
		}

		return false
	}

	found := 0
	topLeft := grid[y-1][x-1]
	topRight := grid[y-1][x+1]
	bottomLeft := grid[y+1][x-1]
	bottomRight := grid[y+1][x+1]

	if isMS(topLeft, bottomRight) && isMS(topRight, bottomLeft) {
		found++
	}

	return found
}
