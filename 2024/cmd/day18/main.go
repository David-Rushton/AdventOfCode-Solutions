package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

var (
	width     = 71
	height    = 71
	timeLimit = 1024
)

func init() {
	if aoc.TestMode {
		width = 7
		height = 7
		timeLimit = 12
	}
}

func main() {
	fmt.Println("--- Day 18: RAM Run ---")
	fmt.Println()

	addresses := parse(aoc.Input)
	fewestSteps := plot(first(addresses, timeLimit))

	var firstFailingByte point
	corruption := timeLimit + 1
	for ; corruption < len(addresses); corruption++ {
		if steps := plot(first(addresses, corruption)); steps == math.MaxInt {
			firstFailingByte = addresses[corruption-1]
			break
		}
	}

	fmt.Println()
	fmt.Printf("Fewest Steps: %d\n", fewestSteps)
	fmt.Printf("Corrupted Bytes: %d\n", corruption)
	fmt.Printf("Last Byte: %d,%d\n", firstFailingByte.x, firstFailingByte.y)
}

type plotState struct {
	point point
	time  int
	steps int
	route []point
}

func plot(addresses []point) int {
	fewestSteps := math.MaxInt
	start := point{0, 0}
	exit := point{width - 1, height - 1}
	visited := make(map[point]int)
	queue := []plotState{{start, 0, 0, []point{start}}}
	iteration := 0

	if aoc.VerboseMode {
		fmt.Print("\x1b[2J")
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		iteration++
		if aoc.VerboseMode && iteration%500 == 0 {
			fmt.Println("\x1b[1;1H")
			fmt.Printf("Queue: %d\r", len(queue))
			printMemory(addresses, current)
		}

		if current.point == exit {
			if current.steps < fewestSteps {
				// printMemory(addresses, current)
				fewestSteps = current.steps
			}
			continue
		}

		if slices.Contains(addresses, current.point) {
			continue
		}

		if current.point != exit {
			if _, found := visited[current.point]; !found {
				visited[current.point] = math.MaxInt
			}

			if current.steps >= visited[current.point] {
				continue
			}
			visited[current.point] = current.steps
		}

		current.time++
		current.steps++

		for _, offset := range getOffsets(current.point) {
			if offset.x >= 0 && offset.x < width && offset.y >= 0 && offset.y < height {
				if steps, found := visited[offset]; !found || current.steps <= steps {
					if !slices.Contains(addresses, offset) {
						newRoute := make([]point, len(current.route))
						copy(newRoute, current.route)

						queue = append(queue, plotState{
							offset,
							current.time,
							current.steps,
							append(newRoute, offset)})
					}
				}
			}
		}

	}

	return fewestSteps
}

func printMemory(addresses []point, state plotState) {
	fmt.Printf("Memory @ %d | Steps = %d\n", state.time, state.steps)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			current := point{x, y}
			value := "."
			if slices.Contains(first(addresses, timeLimit), current) {
				value = "#"
			}
			if slices.Contains(state.route, current) {
				value = "0"
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
	fmt.Println()
}

func getOffsets(from point) []point {
	return []point{
		{from.x + 0, from.y - 1},
		{from.x + 1, from.y + 0},
		{from.x + 0, from.y + 1},
		{from.x - 1, from.y + 0},
	}
}

func parse(input []string) []point {
	result := []point{}

	for _, line := range input {
		coordinates := strings.Split(line, ",")
		result = append(result, point{toInt(coordinates[0]), toInt(coordinates[1])})
	}

	return result
}

func toInt(s string) int {
	number, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert %s to a number.\n", s)
	}
	return int(number)
}

func first[T any](s []T, n int) []T {
	if n < len(s) {
		return s[0:n]
	}
	return s
}
