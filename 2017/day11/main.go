package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

type point struct {
	x int
	y int
}

type direction point

var (
	north     = direction{0, -2}
	northEast = direction{1, -1}
	southEast = direction{1, 1}
	south     = direction{0, 2}
	southWest = direction{-1, 1}
	northWest = direction{-1, -1}
)

func main() {
	fmt.Println("--- Day 11: Hex Ed ---")
	fmt.Println()

	for _, input := range aoc.GetInput(11) {
		directions := getDirections(input)
		childProcess, maxSteps := findChildProcess(directions)
		currentSteps := getDistance(childProcess)
		fmt.Printf(" - Point %v is %d step(s) away now - %d step(s) at peak\n", childProcess, currentSteps, maxSteps)
	}
}

func getDistance(to point) int {
	var steps int
	var current = to

	for current.x != 0 {
		switch current.x > 0 {
		case true:
			current.x--
		default:
			current.x++
		}

		switch current.y > 0 {
		case true:
			current.y--
		default:
			current.y++
		}

		steps++
	}

	if current.y != 0 {
		steps += current.y / 2
	}

	return steps
}

func findChildProcess(directions []direction) (point, int) {
	var result = point{0, 0}
	var maxSteps int

	for i := range directions {
		result.x += directions[i].x
		result.y += directions[i].y

		currentSteps := getDistance(result)
		if currentSteps > maxSteps {
			maxSteps = currentSteps
		}
	}

	return result, maxSteps
}

func getDirections(input string) []direction {
	var result []direction

	for current := range strings.SplitSeq(input, ",") {
		switch current {
		case "n":
			result = append(result, north)
		case "ne":
			result = append(result, northEast)
		case "se":
			result = append(result, southEast)
		case "s":
			result = append(result, south)
		case "sw":
			result = append(result, southWest)
		case "nw":
			result = append(result, northWest)
		}
	}

	return result
}
