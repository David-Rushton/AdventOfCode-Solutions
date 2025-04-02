package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 9: Stream Processing ---")
	fmt.Println()

	var score int
	var cancelled int
	for _, group := range aoc.GetInput(9) {
		score, cancelled = scoreGroups(group)
		fmt.Printf(" - `%v` scores %d points with %d cancelled\n", group, score, cancelled)
	}

	fmt.Println()
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Cancelled: %d\n", cancelled)
}

func scoreGroups(group string) (score, cancelled int) {
	var groups = []string{}

	var count func(r []rune, from int, level int) int
	count = func(r []rune, from int, level int) int {
		var isGarbage = r[from] == '<'

		for i := from + 1; i < len(r); i++ {
			if isGarbage {
				switch r[i] {
				case '>':
					return i

				case '!':
					i++
					continue
				default:
					cancelled++
				}
			}

			if !isGarbage {
				switch r[i] {
				case '<':
					i = count(r, i, level+1)

				case '{':
					i = count(r, i, level+1)

				case '}':
					score += level
					groups = append(groups, string(r[from:i+1]))
					return i
				}
			}
		}

		return -1
	}

	count([]rune(group), 0, 1)

	return score, cancelled
}
