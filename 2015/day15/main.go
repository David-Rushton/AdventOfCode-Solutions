package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func (i *ingredient) getProperty(index int) int {
	return []int{i.capacity, i.durability, i.flavor, i.texture}[index]
}

func main() {
	fmt.Println("--- Day 15: Science for Hungry People ---")
	fmt.Println()

	ingredients := parse(aoc.GetInput(15))
	bestScore, bestDietScore := getMixtures(ingredients)

	fmt.Println()
	fmt.Printf("Best Score: %d\n", bestScore)
	fmt.Printf("Best Diet Score: %d\n", bestDietScore)
}

func getMixtures(ingredients []ingredient) (bestScore, bestDietScore int) {
	bestScore = 0
	bestDietScore = 0

	combinations := getCombinations(100, len(ingredients))
	for _, combination := range combinations {
		currentScore := 1

		for propertyIndex := 0; propertyIndex < 4; propertyIndex++ {
			var propertyScore int

			for i, ingredient := range ingredients {
				if combination[i] > 0 {
					propertyScore += combination[i] * ingredient.getProperty(propertyIndex)
				}
			}

			if propertyScore < 0 {
				propertyScore = 0
			}

			currentScore *= propertyScore
		}

		var calorieScore int
		for i, ingredient := range ingredients {
			if combination[i] > 0 {
				calorieScore += combination[i] * ingredient.calories
			}
		}

		if calorieScore == 500 && currentScore > bestDietScore {
			bestDietScore = currentScore
			fmt.Printf(" - New best diet score %d for combination %v\n", currentScore, combination)
		}

		if currentScore > bestScore {
			fmt.Printf(" - New best score %d for combination %v\n", currentScore, combination)
			bestScore = currentScore
		}
	}

	return bestScore, bestDietScore
}

func getCombinations(n, k int) [][]int {
	result := [][]int{}
	combination := make([]int, k)
	var backtrack func(int, int)

	backtrack = func(start, index int) {
		// For the final element, check if we have reached the target.
		if index == k {
			temp := make([]int, k)
			copy(temp, combination)
			if sliceSum(temp) == 100 {
				result = append(result, temp)
			}
			return
		}

		// Recurse over all values for current index.
		for i := start; i <= n; i++ {
			combination[index] = i
			backtrack(0, index+1)
		}
	}

	backtrack(0, 0)
	return result
}

func sliceSum(s []int) int {
	var result int
	for _, i := range s {
		result += i
	}
	return result
}

func parse(input []string) []ingredient {
	result := []ingredient{}

	for _, item := range input {
		elements := strings.Split(strings.ReplaceAll(item, ",", ""), " ")

		nextIngredient := ingredient{
			name:       elements[0],
			capacity:   aoc.ToInt(elements[2]),
			durability: aoc.ToInt(elements[4]),
			flavor:     aoc.ToInt(elements[6]),
			texture:    aoc.ToInt(elements[8]),
			calories:   aoc.ToInt(elements[10]),
		}

		result = append(result, nextIngredient)
	}

	return result
}
