package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type aunt struct {
	id         int
	properties map[string]int
}

func main() {
	fmt.Println("--- Day 16: Aunt Sue ---")
	fmt.Println()

	aunts := parse(aoc.GetInput(16))
	starOne := findAunt(aunts, false)
	starTwo := findAunt(aunts, true)

	fmt.Println()
	fmt.Printf("Result Star One: %d", starOne)
	fmt.Printf("Result Star Two: %d", starTwo)
}

func findAunt(aunts []aunt, fuzzyMode bool) int {
	predicates := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	var result int
	for _, aunt := range aunts {
		var mismatch bool
		for property, expected := range predicates {
			if actual, exists := aunt.properties[property]; exists {

				if fuzzyMode {
					if property == "cats" || property == "trees" {
						if expected > actual {
							mismatch = true
							break
						}

						continue
					}

					if property == "pomeranians" || property == "goldfish" {
						if expected < actual {
							mismatch = true
							break
						}

						continue
					}
				}

				if expected != actual {
					mismatch = true
					break
				}
			}
		}

		if !mismatch {
			fmt.Printf(" - Sue #%d is a match\n", aunt.id)
			result = aunt.id
		}
	}

	return result
}

func parse(input []string) []aunt {
	result := []aunt{}

	for _, current := range input {
		elements := strings.Split(current, " ")

		nextSue := aunt{
			id:         aoc.ToInt(strings.ReplaceAll(elements[1], ":", "")),
			properties: map[string]int{},
		}

		for i := 2; i < len(elements); i += 2 {
			key := strings.ReplaceAll(elements[i], ":", "")
			value := aoc.ToInt(strings.ReplaceAll(elements[i+1], ",", ""))
			nextSue.properties[key] = value
		}

		result = append(result, nextSue)
	}

	return result
}
