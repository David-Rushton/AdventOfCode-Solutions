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

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 1: No Time for a Taxicab ---")
	fmt.Println()

	instructions := parse(aoc.GetInput(1)[0])
	distance, firstRevisit := solve(instructions)

	fmt.Println()
	fmt.Printf("Manhattan distance: %d\n", distance)
	fmt.Printf("First Revisit: %d\n", firstRevisit)
}

func solve(instructions []instruction) (distance, firstRevisit int) {
	firstRevisit = 0
	location := point{x: 0, y: 0}
	visited := map[point]int{}
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

		for ; instruction.steps > 0; instruction.steps-- {
			switch direction % 4 {
			// north
			case 0:
				location.y -= 1

			// east
			case 1:
				location.x += 1

			// south
			case 2:
				location.y += 1

			// west
			case 3:
				location.x -= 1

			default:
				panic(fmt.Sprintf("Unsupported direction: %d", direction))
			}

			visited[point{location.x, location.y}]++
			if visited[point{location.x, location.y}] == 2 && firstRevisit == 0 {
				firstRevisit = int(math.Abs(float64(location.x)) + math.Abs(float64(location.y)))
			}
		}

		fmt.Printf(" - Location: %d x %d\n", location.x, location.y)
	}

	return int(math.Abs(float64(location.x)) + math.Abs(float64(location.y))), firstRevisit
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
