package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type instruction struct {
	turn  string
	steps int
}

func main() {
	fmt.Println("--- Day 1: No Time for a Taxicab ---")
	fmt.Println()

	instructions := parse(aoc.GetInput(1)[0])
	result := getManhattanDistance(instructions)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func getManhattanDistance(instructions []instruction) int {
	var x int
	var y int
	direction := 0

	for _, instruction := range instructions {
		if instruction.turn == "L" {
			direction--
			if direction < 0 {
				direction = 3
			}
		} else {
			direction++
			if direction > 3 {
				direction = 0
			}
		}

		switch direction % 4 {
		// north
		case 0:
			y -= instruction.steps

		// east
		case 1:
			x += instruction.steps

		// south
		case 2:
			y += instruction.steps

		// west
		case 3:
			x -= instruction.steps

		default:
			panic(fmt.Sprintf("Unsupported direction: %d", direction))
		}

		fmt.Printf(" - Location: %d x %d\n", x, y)
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func parse(input string) []instruction {
	result := []instruction{}

	elements := strings.Split(input, ",")
	for i := range elements {
		value := strings.Trim(elements[i], " ")
		result = append(result, instruction{
			turn:  value[0:1],
			steps: aoc.ToInt(value[1:])})
	}

	return result
}
