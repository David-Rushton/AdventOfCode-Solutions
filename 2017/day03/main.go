package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 3: Spiral Memory ---")
	fmt.Println()

	var result int
	for _, target := range parse(aoc.GetInput(3)) {
		result = findDistance(target)
		fmt.Printf("   - % 6d == % 6d\n", target, result)
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
			var decending = true
			var limit = ySteps
			for {
				if current == target {
					return ySteps + xSteps
				}

				if decending {
					xSteps--
					if xSteps <= 0 {
						decending = false
					}
				} else {
					xSteps++
					if xSteps >= limit {
						decending = true
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
