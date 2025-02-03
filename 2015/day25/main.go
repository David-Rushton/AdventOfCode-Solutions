package main

import (
	"fmt"
	"os"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 25: Let It Snow ---")
	fmt.Println()

	row := aoc.ToInt(os.Args[1])
	column := aoc.ToInt(os.Args[2])
	result := findValue(row, column)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findValue(row, column int) int {
	maxRow := 1
	currentRow := 1
	currentColumn := 1
	value := 20151125

	for {
		// result!
		if currentRow == row && currentColumn == column {
			break
		}

		value = (value * 252533) % 33554393

		// iterate.
		// up and across until top row.
		// then restart from next highest row, 1st column
		currentRow--
		currentColumn++
		if currentRow == 0 {
			maxRow++
			currentRow = maxRow
			currentColumn = 1
		}
	}

	return value
}
