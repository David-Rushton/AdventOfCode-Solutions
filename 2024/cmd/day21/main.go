package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

type directedPoint struct {
	direction rune
	point     point
}

var (
	numericKeypad = map[rune]point{
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
	numericPoints = map[point]rune{
		{x: 0, y: 0}: '7',
		{x: 1, y: 0}: '8',
		{x: 2, y: 0}: '9',
		{x: 0, y: 1}: '4',
		{x: 1, y: 1}: '5',
		{x: 2, y: 1}: '6',
		{x: 0, y: 2}: '1',
		{x: 1, y: 2}: '2',
		{x: 2, y: 2}: '3',
		{x: 1, y: 3}: '0',
		{x: 2, y: 3}: 'A',
	}
	directionalKeypad = map[rune]point{
		'^': {x: 1, y: 10},
		'A': {x: 2, y: 10},
		'<': {x: 0, y: 11},
		'v': {x: 1, y: 11},
		'>': {x: 2, y: 11},
	}
	directionalPoints = map[point]rune{
		{x: 1, y: 10}: '^',
		{x: 2, y: 10}: 'A',
		{x: 0, y: 11}: '<',
		{x: 1, y: 11}: 'v',
		{x: 2, y: 11}: '>',
	}
)

func main() {
	fmt.Println("--- Day 21: Keypad Conundrum ---")
	fmt.Println()

	robots := 2
	if aoc.Star == aoc.StarTwo {
		robots = 25
	}

	var complexity int

	codes := aoc.Input
	for _, code := range codes {
		shortestSequenceLen := math.MaxInt
		outerSequences := getNumericSequences(code)

		for i := 1; i <= robots; i++ {
			nextOuterSequences := make(map[string]int)

			for _, outerSequence := range outerSequences {
				for _, innerSequence := range getDirectionalSequences(outerSequence) {
					nextOuterSequences[innerSequence]++
				}
			}

			outerSequences = []string{}
			for sequence := range nextOuterSequences {
				outerSequences = append(outerSequences, sequence)
			}
		}

		for _, sequence := range outerSequences {
			if len(sequence) < shortestSequenceLen {
				shortestSequenceLen = len(sequence)
			}
		}

		codeValue, err := strconv.ParseInt(code[0:3], 10, 64)
		if err != nil {
			log.Fatalf("Cannot convert %v to a number.", code[0:3])
		}

		fmt.Printf("  - %v | % 7d * % 4d = %v\n", code, shortestSequenceLen, codeValue, shortestSequenceLen*int(codeValue))
		complexity += shortestSequenceLen * int(codeValue)
	}

	fmt.Println()
	fmt.Printf("Complexity : %d", complexity)
}

func getNumericSequences(code string) []string {
	return getSequences(code, numericPoints, numericKeypad)
}

func getDirectionalSequences(code string) []string {
	type state struct {
		position int
		path     string
	}

	sequences := strings.Split(code, "A")
	sequences = sequences[0 : len(sequences)-1]

	queue := []state{{position: 0, path: ""}}
	paths := make(map[string]int)
	shortestPath := math.MaxInt

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current.path) > shortestPath {
			continue
		}

		if current.position >= len(sequences) {
			if len(current.path) <= shortestPath {
				shortestPath = len(current.path)
				paths[current.path]++
			}

			continue
		}

		for _, sequence := range getSequences(sequences[current.position]+"A", directionalPoints, directionalKeypad) {
			queue = append(queue, state{
				position: current.position + 1,
				path:     current.path + sequence})
		}
	}

	result := []string{}
	for k := range paths {
		result = append(result, k)
	}

	return result
}

var sequenceCache = make(map[string][]string)

func getSequences(code string, points map[point]rune, keypad map[rune]point) []string {
	type state struct {
		position int
		from     rune
		path     string
	}

	// Use cached results if possible.
	if _, found := sequenceCache[code]; found {
		return sequenceCache[code]
	}

	// Find values.
	queue := []state{{
		position: 0,
		from:     'A',
		path:     "",
	}}

	shortestResult := math.MaxInt
	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current.path) > shortestResult {
			continue
		}

		if current.position == len(code) {
			if len(current.path) <= shortestResult {
				shortestResult = len(current.path)
				result = append(result, current.path)
			}
			continue
		}

		to := rune(code[current.position])
		for _, movement := range getMovements(points, keypad[current.from], keypad[to]) {
			queue = append(queue, state{
				position: current.position + 1,
				from:     to,
				path:     current.path + movement})
		}
	}

	// Cache result.
	sequenceCache[code] = result

	return result
}

type movementKey struct {
	from point
	to   point
}

var movementCache = make(map[movementKey][]string)

func getMovements(points map[point]rune, from, to point) []string {
	cacheKey := movementKey{from, to}
	if _, found := movementCache[cacheKey]; found {
		return movementCache[cacheKey]
	}

	type queueState struct {
		point point
		steps int
		route string
	}

	queue := []queueState{{from, 0, ""}}
	visited := make(map[point]bool)
	fewestSteps := math.MaxInt
	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.steps > fewestSteps {
			continue
		}

		if current.point == to && current.steps <= fewestSteps {
			result = append(result, current.route+"A")
			fewestSteps = current.steps
		}

		for _, offset := range getOffsets(current.point) {
			if _, found := points[offset.point]; found {
				if _, found := visited[offset.point]; !found {
					steps := 1
					if len(current.route) > 0 && current.route[len(current.route)-1] != byte(offset.direction) {
						steps = 10
					}

					queue = append(queue, queueState{
						offset.point,
						current.steps + steps,
						current.route + string(offset.direction)})
				}
			}
		}

		visited[current.point] = true
	}

	movementCache[cacheKey] = result

	return result
}

func getOffsets(from point) []directedPoint {
	return []directedPoint{
		{'^', point{from.x + 0, from.y - 1}},
		{'>', point{from.x + 1, from.y + 0}},
		{'v', point{from.x + 0, from.y + 1}},
		{'<', point{from.x - 1, from.y + 0}},
	}
}
