package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type triangle struct {
	side1 int
	side2 int
	side3 int
}

func (t *triangle) isPossible() bool {
	return t.side1+t.side2 > t.side3 &&
		t.side2+t.side3 > t.side1 &&
		t.side1+t.side3 > t.side2
}

func main() {
	fmt.Println("--- Day 3: Squares With Three Sides ---")
	fmt.Println()

	triangles := parse(aoc.GetInput(3), aoc.Star == aoc.StarTwo)

	var result int
	for i := range triangles {
		if triangles[i].isPossible() {
			result++
			fmt.Printf(
				" - Is valid triangle: %d x %d x %d\n",
				triangles[i].side1,
				triangles[i].side2,
				triangles[i].side3)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func parse(input []string, verticalMode bool) []triangle {
	result := []triangle{}

	for i := range input {
		current := strings.Trim(input[i], " ")
		current = strings.ReplaceAll(current, "    ", " ")
		current = strings.ReplaceAll(current, "   ", " ")
		current = strings.ReplaceAll(current, "  ", " ")
		elements := strings.Split(current, " ")

		result = append(result, triangle{
			side1: aoc.ToInt(elements[0]),
			side2: aoc.ToInt(elements[1]),
			side3: aoc.ToInt(elements[2])})
	}

	if verticalMode {
		temp := []triangle{}

		for i := 0; i < len(result); i += 3 {
			temp = append(temp, triangle{result[i].side1, result[i+1].side1, result[i+2].side1})
			temp = append(temp, triangle{result[i].side2, result[i+1].side2, result[i+2].side2})
			temp = append(temp, triangle{result[i].side3, result[i+1].side3, result[i+2].side3})
		}

		result = temp
	}

	return result
}
