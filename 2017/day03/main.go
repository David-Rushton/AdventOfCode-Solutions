package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 3: Spiral Memory ---")
	fmt.Println()

	targets := parse(aoc.GetInput(3))

	var result int
	for _, target := range targets {
		result = countStepsTo(target)
		fmt.Printf(" - Steps to %d == %d\n", target, result)
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func countStepsTo(target int) int {
	var steps = 0
	var grid = map[point]int{{x: 0, y: 0}: 1}

	const east int = 0
	const north int = 1
	const west int = 2
	const south int = 3
	var direction = south

	var x int
	var y int
	var top, bottom, left, right int
	for {
		// Move
		switch direction {
		case north:
			if y == top {
				direction = west
				left--
				x--
			} else {
				y++
			}

		case east:
			if x == right {
				direction = north
				top++
				y++
			} else {
				x++
			}

		case south:
			if y == bottom {
				direction = east
				right++
				x++
			} else {
				y--
			}

		case west:
			if x == left {
				direction = south
				bottom--
				y--
			} else {
				x--
			}
		}

		// Add next cell.
		var value int
		for _, yOffset := range []int{-1, 0, 1} {
			for _, xOffset := range []int{-1, 0, 1} {
				if yOffset == 0 && xOffset == 0 {
					continue
				}

				if neighbourValue, exists := grid[point{x + xOffset, y + yOffset}]; exists {
					value += neighbourValue
				}
			}
		}

		// Increment steps
		steps++

		grid[point{x, y}] = value
		if value > target {
			return steps
		}
	}
}

func parse(input []string) []int {
	var result []int

	for i := range input {
		result = append(result, aoc.ToInt(input[i]))
	}

	return result
}
