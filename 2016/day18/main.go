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
	trapMap := getMap(firstRow, rows)
	fmt.Println(trapMap)
	safeTiles := countSafeTiles(trapMap)

	fmt.Println()
	fmt.Printf("Result: %d\n", safeTiles)
}

func getMap(firstRow string, rows int) string {
	result := firstRow

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
		nextRow := []rune{}
		for j := 0; j < len(lastRow); j++ {
			if itsATrap(getTraps(j, lastRow)) {
				nextRow = append(nextRow, '^')
			} else {
				nextRow = append(nextRow, '.')
			}
		}

		result += "\n" + string(nextRow)
		lastRow = nextRow
	}

	return result
}

func countSafeTiles(trapMap string) any {
	return len(trapMap) - len(strings.ReplaceAll(trapMap, ".", ""))
}
