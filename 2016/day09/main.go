package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 9: Explosives in Cyberspace ---")
	fmt.Println()

	for _, file := range aoc.GetInput(9) {
		label := file
		if len(label) > 20 {
			label = label[0:17] + "..."
		}

		fmt.Printf(
			" - %v has a decompressed len of %d\n",
			label,
			getDecompressedLen(file, aoc.Star == aoc.StarTwo))
	}

	fmt.Println()
}

func getDecompressedLen(compressed string, v2 bool) int {
	var solve func(s string) int
	solve = func(s string) int {
		var result int

		for i := 0; i < len(s); i++ {
			r := s[i]

			if r == '(' {
				markerEnd := indexOf(s, ")", i)
				take, repeat := parseMarker(s[i : markerEnd+1])

				if !v2 {
					result += take * repeat
				}

				if v2 {
					result += solve(s[markerEnd+1:markerEnd+take+1]) * repeat
				}

				i = markerEnd + take
				continue
			}

			result++
		}

		return result
	}

	return solve(compressed)
}

func parseMarker(marker string) (take, repeat int) {
	elements := strings.Split(marker[1:len(marker)-1], "x")
	return aoc.ToInt(elements[0]), aoc.ToInt(elements[1])
}

func indexOf(s string, substr string, from int) int {
	idx := strings.Index(s[from:], substr)
	if idx == -1 {
		return -1
	}

	return from + idx
}
