package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 18: Like a Rogue ---")
	fmt.Println()

	var rows int
	switch {
	case aoc.TestMode:
		rows = 10
	case aoc.Star == aoc.StarTwo:
		rows = 400000
	default:
		rows = 40
	}

	firstRow := aoc.GetInput(18)[0]
	trapMap, safeCount := getMap(firstRow, rows)
	printTrapMap(trapMap)

	fmt.Println()
	fmt.Printf("Result: %d\n", safeCount)
}

func getMap(firstRow string, rows int) (trapMap [][]rune, safeCount int) {
	safeCount = len(firstRow) - len(strings.ReplaceAll(firstRow, ".", ""))
	trapMap = make([][]rune, rows)
	trapMap[0] = []rune(firstRow)

	getTraps := func(position int, last []rune) (bool, bool, bool) {
		result := []bool{false, false, false}
		for i := -1; i < 2; i++ {
			if position+i >= 0 && position+i < len(last) {
				if last[position+i] == '^' {
					result[i+1] = true
				}
			}
		}

		return result[0], result[1], result[2]
	}

	itsATrap := func(left, centre, right bool) bool {
		if left && centre && !right {
			return true
		}

		if !left && centre && right {
			return true
		}

		if left && !centre && !right {
			return true
		}

		if !left && !centre && right {
			return true
		}

		return false
	}

	lastRow := []rune(firstRow)
	for i := 1; i < rows; i++ {
		nextRow := make([]rune, len(lastRow))
		for j := 0; j < len(lastRow); j++ {
			if itsATrap(getTraps(j, lastRow)) {
				nextRow[j] = '^'
			} else {
				nextRow[j] = '.'
				safeCount++
			}
		}

		trapMap[i] = nextRow
		lastRow = make([]rune, len(nextRow))
		copy(lastRow, nextRow)
	}

	return trapMap, safeCount
}

func printTrapMap(trapMap [][]rune) {
	for i := 0; i < len(trapMap); i++ {
		fmt.Printf("% 5d ", i+1)
		for j := 0; j < len(trapMap[i]); j++ {
			fmt.Print(string(trapMap[i][j]))
		}
		fmt.Println()
	}
}
