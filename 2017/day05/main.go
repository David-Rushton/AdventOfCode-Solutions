package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 5: A Maze of Twisty Trampolines, All Alike ---")
	fmt.Println()

	var program = parse(aoc.GetInput(5))
	var iterations = runProgram(program, aoc.Star == aoc.StarTwo)

	fmt.Println()
	fmt.Printf("Result: %d", iterations)
}

func runProgram(program []int, complexMode bool) int {
	var iterations int

	var index int
	for {
		iterations++

		var nextIndex = index + program[index]
		if complexMode && program[index] >= 3 {
			program[index]--
		} else {
			program[index]++
		}

		index = nextIndex
		if nextIndex >= len(program) {
			break
		}
	}

	return iterations
}

func parse(input []string) []int {
	var result []int

	for i := range input {
		result = append(result, aoc.ToInt(input[i]))
	}

	return result
}
