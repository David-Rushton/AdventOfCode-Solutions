package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

/*
      0   1   2
	+---+---+---+
	| 7 | 8 | 9 | 0
	+---+---+---+
	| 4 | 5 | 6 | 1
	+---+---+---+
	| 1 | 2 | 3 | 2
	+---+---+---+
		| 0 | A | 3
		+---+---+

      0   1   2
		+---+---+
		| ^ | A | 10
	+---+---+---+
	| < | v | > | 11
	+---+---+---+

*/

type directedPoint struct {
	direction rune
	point     point
}

type point struct {
	x int
	y int
}

type keypad map[rune]point
type points map[point]bool

var (
	numericKeypad = keypad{
		'7': {x: 0, y: 0},
		'8': {x: 1, y: 0},
		'9': {x: 2, y: 0},
		'4': {x: 0, y: 1},
		'5': {x: 1, y: 1},
		'6': {x: 2, y: 1},
		'1': {x: 0, y: 2},
		'2': {x: 1, y: 2},
		'3': {x: 2, y: 2},
		'0': {x: 1, y: 3},
		'A': {x: 2, y: 3},
	}
	numericPoints = points{
		{x: 0, y: 0}: true,
		{x: 1, y: 0}: true,
		{x: 2, y: 0}: true,
		{x: 0, y: 1}: true,
		{x: 1, y: 1}: true,
		{x: 2, y: 1}: true,
		{x: 0, y: 2}: true,
		{x: 1, y: 2}: true,
		{x: 2, y: 2}: true,
		{x: 1, y: 3}: true,
		{x: 2, y: 3}: true,
	}
	directionKeypad = keypad{
		'^': {x: 1, y: 10},
		'A': {x: 2, y: 10},
		'<': {x: 0, y: 11},
		'v': {x: 1, y: 11},
		'>': {x: 2, y: 11},
	}
	directionPoints = points{
		{x: 1, y: 10}: true,
		{x: 2, y: 10}: true,
		{x: 0, y: 11}: true,
		{x: 1, y: 11}: true,
		{x: 2, y: 11}: true,
	}
)

func main() {
	fmt.Println("--- Day 21: Keypad Conundrum ---")
	fmt.Println()

	levels := 2
	if aoc.Star == aoc.StarTwo {
		levels = 25
	}

	var result int
	for _, code := range aoc.Input {
		fmt.Printf("%v\n", code)

		sequence := getNumericRoute(code)

		for i := 0; i < levels; i++ {
			sequence = getDirectionalRoute(sequence)
			fmt.Printf(" %v(%v) %v\n", i, len(sequence), "sequence")
		}

		codeValue, err := strconv.ParseInt(code[0:3], 10, 64)
		if err != nil {
			log.Fatalf("Cannot convert %s to number.", code[0:3])
		}

		fmt.Printf("  %v * %v = %v\n", len(sequence), codeValue, int(codeValue)*len(sequence))
		fmt.Println()

		result += int(codeValue) * len(sequence)
	}

	fmt.Println()
	fmt.Printf("Results: %d\n", result)
}

func getNumericRoute(code string) string {
	result := ""

	from := 'A'
	for _, to := range code {
		result += getRoute(numericKeypad, numericPoints, from, to)
		from = to
	}

	return result
}

func getDirectionalRoute(code string) string {
	result := ""

	from := 'A'
	for _, to := range code {
		result += getRoute(directionKeypad, directionPoints, from, to)
		from = to
	}

	return result
}

func getRoute(keypad keypad, points points, from, to rune) string {
	type state struct {
		point     point
		score     int
		direction rune
		path      string
	}

	queue := []state{{
		point:     keypad[from],
		score:     0,
		direction: 'X',
		path:      ""}}
	visited := make(map[point]bool)
	bestScore := math.MaxInt
	result := ""

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > bestScore {
			continue
		}

		if current.point == keypad[to] {
			if current.score <= bestScore {
				result = current.path + "A"
				bestScore = current.score
			}
			continue
		}

		for _, neighbour := range getNeighbours(current.point) {
			if points[neighbour.point] && !visited[neighbour.point] {
				score := 1
				if current.direction != neighbour.direction {
					score = 100
				}

				queue = append(queue, state{
					neighbour.point,
					current.score + score,
					neighbour.direction,
					current.path + string(neighbour.direction)})
			}
		}

		visited[current.point] = true
	}

	return result
}

func getNeighbours(from point) []directedPoint {
	return []directedPoint{
		{'^', point{from.x + 0, from.y - 1}},
		{'>', point{from.x + 1, from.y + 0}},
		{'v', point{from.x + 0, from.y + 1}},
		{'<', point{from.x - 1, from.y + 0}},
	}
}
