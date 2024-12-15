package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type dimensions point
type velocity point
type point struct {
	x int
	y int
}

type robot struct {
	id       int
	point    point
	velocity velocity
}

func main() {
	fmt.Println("---  Day 14: Restroom Redoubt ---")
	fmt.Println()

	grid := dimensions{101, 103}
	if aoc.TestMode {
		grid = dimensions{11, 7}
	}

	safetyFactor := 0
	robots := parse(aoc.Input)

	var seconds int
	for ; shouldContinue(seconds, grid, robots); seconds++ {
		robots, safetyFactor = moveRobots(grid, robots)
	}

	printGrid(grid, robots)

	fmt.Println()
	fmt.Printf("Result: %d at %d\n", safetyFactor, seconds)
}

func moveRobots(grid dimensions, robots []robot) ([]robot, int) {
	moveWithTeleport := func(from, velocity, max int) int {
		result := from + velocity

		if result < 0 {
			result = max + result
		}

		if result >= max {
			result = result % max
		}

		if !(result >= 0 && result < max) {
			panic("Need to extend wrapping")
		}

		return result
	}

	movedRobots := []robot{}
	quadrants := [4]int{}
	for _, robot := range robots {
		robot.point.x = moveWithTeleport(robot.point.x, robot.velocity.x, grid.x)
		robot.point.y = moveWithTeleport(robot.point.y, robot.velocity.y, grid.y)
		movedRobots = append(movedRobots, robot)
		if robot.point.x == grid.x/2 || robot.point.y == grid.y/2 {
			continue
		}
		if robot.point.x < grid.x/2 {
			if robot.point.y < grid.y/2 {
				quadrants[0]++
			} else {
				quadrants[1]++
			}
		} else {
			if robot.point.y < grid.y/2 {
				quadrants[2]++
			} else {
				quadrants[3]++
			}
		}
	}

	return movedRobots, quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func shouldContinue(seconds int, grid dimensions, robots []robot) bool {
	if aoc.Star == aoc.StarOne {
		return seconds < 100
	} else {
		return maybeChristmasTree(grid, robots)
	}
}

func maybeChristmasTree(grid dimensions, robots []robot) bool {
	// find robots
	locations := make(map[point]int)
	for _, robot := range robots {
		locations[robot.point]++
	}

	// check for opposite robots
	var matches int
	for location, _ := range locations {
		if location.x < grid.x/2 {
			oppositeLocation := point{
				x: grid.x - location.x,
				y: location.y,
			}

			if _, exists := locations[oppositeLocation]; exists {
				matches++
			}
		}
	}

	return matches < 70
}

func parse(input []string) []robot {
	result := []robot{}

	toPoint := func(cordinates string) point {
		result := point{}
		for i, element := range strings.Split(cordinates, ",") {
			num, err := strconv.ParseInt(element, 10, 64)
			if err != nil {
				log.Fatalf("Cannot convert %s to a number.\n", element)
			}

			switch i {
			case 0:
				result.x = int(num)
			case 1:
				result.y = int(num)
			}
		}
		return result
	}

	idSeed := 0
	for _, line := range input {
		cleanLine := strings.Replace(strings.Replace(line, "p=", "", 1), "v=", "", -1)
		elements := strings.Split(cleanLine, " ")
		result = append(result, robot{
			id:       idSeed,
			point:    toPoint(elements[0]),
			velocity: velocity(toPoint(elements[1]))})
		idSeed++
	}

	return result
}
