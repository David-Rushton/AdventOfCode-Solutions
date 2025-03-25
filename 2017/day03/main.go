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
		if aoc.Star == aoc.StarOne {
			result = findDistance(target)
			fmt.Printf(" - Steps to %d == %d\n", target, result)
		} else {
			result = countStepsTo(target)
			fmt.Printf(" - Steps to %d == %d\n", target, result)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func countStepsTo(target int) int {
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

		// Check for exit
		grid[point{x, y}] = value
		if value > target {
			printGrid(top, bottom, left, right, grid)
			return value
		}
	}
}

func printGrid(top, bottom, left, right int, grid map[point]int) {
	for y := top; y >= bottom; y-- {
		for x := left; x <= right; x++ {
			if value, exists := grid[point{x, y}]; exists {
				fmt.Printf("% 8d", value)
			} else {
				fmt.Print("        ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findDistance(target int) int {
	var squareSide = 1
	var ceiling = 1
	var ySteps = 0
	for {

		if target <= ceiling {
			if target == ceiling {
				return ySteps * 2
			}

			var current = ceiling
			var xSteps = ySteps
			var descending = true
			var limit = ySteps
			for {
				if current == target {
					return ySteps + xSteps
				}

				if descending {
					xSteps--
					if xSteps <= 0 {
						descending = false
					}
				} else {
					xSteps++
					if xSteps >= limit {
						descending = true
					}
				}

				current--
			}
		}

		ySteps++
		squareSide += 2
		ceiling += (squareSide * 4) - 4
	}
}

func parse(input []string) []int {
	var result []int

	for i := range input {
		result = append(result, aoc.ToInt(input[i]))
	}

	return result
}
