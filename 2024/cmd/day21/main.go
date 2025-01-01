package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

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
		'^': {x: 11, y: 10},
		'A': {x: 12, y: 10},
		'<': {x: 10, y: 11},
		'v': {x: 11, y: 11},
		'>': {x: 12, y: 11},
	}
	directionPoints = points{
		{x: 11, y: 10}: true,
		{x: 12, y: 10}: true,
		{x: 10, y: 11}: true,
		{x: 11, y: 11}: true,
		{x: 12, y: 11}: true,
	}
)

func main() {
	fmt.Println("--- Day 21: Keypad Conundrum ---")
	fmt.Println()

	buildRouteCache()

	maxLevel := 2
	if aoc.Star == aoc.StarTwo {
		maxLevel = 25
	}

	var result int
	for _, code := range aoc.Input {
		codeValue, err := strconv.ParseInt(code[0:3], 10, 64)
		if err != nil {
			fmt.Printf("Cannot convert %v to a number.", codeValue)
		}

		bestLen := getCodeLen(code, maxLevel)
		fmt.Printf("  %v | % 4d * % 4d = %d\n", code, bestLen, codeValue, int(codeValue)*getCodeLen(code, maxLevel))
		result += int(codeValue) * getCodeLen(code, maxLevel)
	}

	fmt.Println()
	fmt.Printf("Results: %d\n", result)
}

func getNumericRoute(code string) sequence {
	route := ""
	from := 'A'
	for _, to := range code {
		route += getRoute(numericKeypad, numericPoints, from, to)
		from = to
	}

	result := make(sequence)
	for _, section := range strings.Split(route, "A")[0:4] {
		result[section+"A"]++
	}

	return result
}

func getDirectionalSequence(sequences sequence) sequence {
	result := make(sequence)

	for code, count := range sequences {
		sections := strings.Split(getDirectionalRoute(code), "A")
		sections = sections[0 : len(sections)-1]
		for _, section := range sections {
			result[section+"A"] += count
		}
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
			if current.score < bestScore {
				bestScore = current.score
			}

			if current.score == bestScore {
				result = current.path + "A"
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

// -----

func getCodeLen(code string, maxLevel int) int {
	result := math.MaxInt

	for _, route := range getNumericRoutes(code) {
		currentSequence := ToSequence(route)
		for i := 0; i < maxLevel; i++ {
			nextSequence := sequence{}

			for route, count := range currentSequence {
				from := 'A'
				for _, to := range route {
					nextSequence[getCachedRoute(from, to)] += count
					from = to
				}
			}

			currentSequence = nextSequence
		}

		if currentSequence.GetLen() < result {
			result = currentSequence.GetLen()
		}
	}

	return result
}

func getBestRoutes(routes []string) []string {
	result := []string{}
	bestLen := math.MaxInt

	for _, route := range routes {
		subRoutes := getDirectionalRoutes(route)
		routeLen := len(subRoutes[0])

		if routeLen < bestLen {
			bestLen = routeLen
			result = []string{}
		}

		if routeLen == bestLen {
			result = append(result, subRoutes...)
		}
	}

	return result
}

func getNumericRoutes(route string) []string {
	return getExtendedRoutes(numericKeypad, numericPoints, route)
}

func getDirectionalRoutes(route string) []string {
	return getExtendedRoutes(directionKeypad, directionPoints, route)
}

func getExtendedRoutes(keypad keypad, points points, route string) []string {
	type state struct {
		position int
		from     rune
		to       rune
		route    string
	}

	queue := []state{{
		position: 0,
		from:     'A',
		to:       rune(route[0]),
		route:    "",
	}}

	route = route + " "
	shortestRoute := math.MaxInt
	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current.route) > shortestRoute {
			continue
		}

		if current.position == len(route)-1 {
			if len(current.route) < shortestRoute {
				shortestRoute = len(current.route)
				result = []string{}
			}

			if len(current.route) == shortestRoute {
				result = append(result, current.route)
			}

			continue
		}

		for _, next := range getRoutes(keypad, points, current.from, current.to) {
			queue = append(queue, state{
				position: current.position + 1,
				from:     current.to,
				to:       rune(route[current.position+1]),
				route:    current.route + next})
		}
	}

	return result
}

func getRoutes(keypad keypad, points points, from, to rune) []string {
	toPoint := keypad[to]
	fromPoint := keypad[from]

	getX := func(start point) (string, bool) {
		result := ""
		current := start

		for current.x != toPoint.x {
			if current.x > toPoint.x {
				current.x--
				result += "<"
			} else {
				current.x++
				result += ">"
			}

			if !points[current] {
				return "", false
			}
		}

		return result, true
	}

	getY := func(start point) (string, bool) {
		result := ""
		current := start

		for current.y != toPoint.y {
			if current.y > toPoint.y {
				current.y--
				result += "^"
			} else {
				current.y++
				result += "v"
			}

			if !points[current] {
				return "", false
			}
		}

		return result, true
	}

	getXY := func() (string, bool) {
		if x, ok := getX(fromPoint); ok {
			if y, ok := getY(point{toPoint.x, fromPoint.y}); ok {
				return x + y + "A", true
			}
		}

		return "", false
	}

	getYX := func() (string, bool) {
		if y, ok := getY(fromPoint); ok {
			if x, ok := getX(point{fromPoint.x, toPoint.y}); ok {
				return y + x + "A", true
			}
		}

		return "", false
	}

	result := []string{}

	if xy, ok := getXY(); ok {
		result = append(result, xy)
	}

	if yx, ok := getYX(); ok {
		if !slices.Contains(result, yx) {
			result = append(result, yx)
		}
	}

	return result
}

// -----
