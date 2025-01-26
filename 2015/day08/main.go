package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 8: Matchsticks ---")
	fmt.Println()

	var starOne int
	var starTwo int
	for _, current := range aoc.GetInput(8) {
		if current != "" {
			codeCharacters, literalCharacters := process(current)
			expendedLen := getExpandedLen(current)

			fmt.Printf("  - %v\n    %5d | %5d | %5d\n", current, codeCharacters, literalCharacters, expendedLen)

			starOne += codeCharacters - literalCharacters
			starTwo += expendedLen - codeCharacters
		}
	}

	fmt.Println()
	fmt.Printf("Star 1: %d\n", starOne)
	fmt.Printf("Star 2: %d\n", starTwo)
}

var hexadecimalRegexp = regexp.MustCompile(`\\x[0-9a-f]{2}`)

func process(value string) (codeCharacters, literalCharacters int) {
	literalValue := value[1 : len(value)-1]
	literalValue = strings.ReplaceAll(literalValue, `\\`, `\`)
	literalValue = strings.ReplaceAll(literalValue, `\"`, `"`)
	literalValue = hexadecimalRegexp.ReplaceAllLiteralString(literalValue, "?")

	return len(value), len(literalValue)
}

func getExpandedLen(value string) int {
	expandedValue := ""

	for _, r := range value {
		switch r {
		case '"':
			expandedValue += `\"`
		case '\\':
			expandedValue += `\\`
		default:
			expandedValue += string(r)
		}
	}

	return len(expandedValue) + 2
}
