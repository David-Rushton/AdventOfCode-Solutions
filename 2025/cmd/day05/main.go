package main

import (
	"fmt"
	"slices"
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

func (r idRange) overlaps(other idRange) bool {
	if r.from >= other.from && r.from <= other.until {
		return true
	}

	if r.until >= other.from && r.until <= other.until {
		return true
	}

	if r.from <= other.from && r.until >= other.until {
		return true
	}

	return false
}

func (r idRange) merge(other idRange) idRange {
	return idRange{
		from:  min(r.from, other.from),
		until: max(r.until, other.until),
	}
}

func (r idRange) String() string {
	return fmt.Sprintf("% 15d -> % 15d", r.from, r.until)
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

	// Sort ranges.
	slices.SortFunc(idRanges, func(a, b idRange) int {
		fromCmp := int(a.from - b.from)
		if fromCmp != 0 {
			return fromCmp
		}
		return int(a.until - b.until)
	})

	// Merge overlapping ranges.
	mergedIdRanges := []idRange{idRanges[0]}
	for _, idRange := range idRanges {
		lastIdx := len(mergedIdRanges) - 1
		if idRange.overlaps(mergedIdRanges[lastIdx]) {
			mergedIdRanges[lastIdx] = mergedIdRanges[lastIdx].merge(idRange)
		} else {
			mergedIdRanges = append(mergedIdRanges, idRange)
		}
	}

	// Find fresh ingredients.
	var freshIngredients int

	for _, ingredient := range ingredients {
		for _, idRange := range mergedIdRanges {
			if idRange.includes(ingredient) {
				fmt.Printf(" - Fresh ingredient found %d\n", ingredient)
				freshIngredients++
				break
			}
		}
	}
	fmt.Println()

	// Find total possible fresh ingredients.
	var availableIngredients int64
	for _, idRange := range mergedIdRanges {
		availableIngredients += idRange.until - idRange.from + 1
		fmt.Println(" - Available range extended:", availableIngredients)
	}
	fmt.Println()

	fmt.Println()
	fmt.Printf("Fresh %d\n", freshIngredients)
	fmt.Printf("Available %d\n", availableIngredients)
}

func toInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
