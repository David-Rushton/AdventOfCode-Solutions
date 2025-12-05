package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type idRange struct {
	from  int64
	until int64
}

func (r idRange) includes(i int64) bool {
	return i >= r.from && i <= r.until
}

func main() {
	fmt.Println("--- Day 5: Cafeteria ---")
	fmt.Println()

	// Parse.
	idRanges := []idRange{}
	ingredients := []int64{}

	ingredientMode := false
	for _, line := range aoc.Input {
		if line == "" {
			ingredientMode = true
			continue
		}

		switch ingredientMode {
		case true:
			ingredients = append(ingredients, toInt(line))

		default:
			elements := strings.Split(line, "-")
			idRanges = append(idRanges, idRange{
				from:  toInt(elements[0]),
				until: toInt(elements[1])})
		}
	}

	// Find fresh ingredients.
	var freshIngredients int

	for _, ingredient := range ingredients {
		for _, idRange := range idRanges {
			if idRange.includes(ingredient) {
				fmt.Printf(" - Fresh ingredient found %d\n", ingredient)
				freshIngredients++
				break
			}
		}
	}

	// Find total possible fresh ingredients.

	fmt.Println()
	fmt.Printf("Result %d\n", freshIngredients)
}

func toInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
