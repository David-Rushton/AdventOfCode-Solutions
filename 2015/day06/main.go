package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 6: Probably a Fire Hazard ---")
	fmt.Println()

	instructions := parse(aoc.GetInput(6))
	lit := countLit(instructions)
	brightness := getBrightness(instructions)

	fmt.Println()
	fmt.Printf("Lit: %d\n", lit)
	fmt.Printf("Brightness: %d\n", brightness)
}

func countLit(instructions []instruction) int {
	var result int

	grid := map[point]bool{}
	for _, instruction := range instructions {
		fmt.Printf(" - Action: %v. Range: %v -> %v\n", instruction.action, instruction.topLeft, instruction.bottomRight)

		for x := instruction.topLeft.x; x <= instruction.bottomRight.x; x++ {
			for y := instruction.topLeft.y; y <= instruction.bottomRight.y; y++ {
				current := point{x, y}

				switch instruction.action {
				case turnOn:
					if !grid[current] {
						grid[current] = true
						result++
					}
				case turnOff:
					if grid[current] {
						grid[current] = false
						result--
					}
				case toggle:
					grid[current] = !grid[current]
					if grid[current] {
						result++
					} else {
						result--
					}
				}
			}
		}
	}

	return result
}

func getBrightness(instructions []instruction) int {
	grid := map[point]int{}

	for _, instruction := range instructions {
		fmt.Printf(" - Action: %v. Range: %v -> %v\n", instruction.action, instruction.topLeft, instruction.bottomRight)

		for x := instruction.topLeft.x; x <= instruction.bottomRight.x; x++ {
			for y := instruction.topLeft.y; y <= instruction.bottomRight.y; y++ {
				current := point{x, y}

				switch instruction.action {
				case turnOn:
					grid[current]++
				case turnOff:
					grid[current]--
					if grid[current] < 0 {
						grid[current] = 0
					}
				case toggle:
					grid[current] += 2
				}
			}
		}
	}

	var result int
	for _, brightness := range grid {
		result += brightness
	}

	return result
}

func parse(input []string) []instruction {
	result := []instruction{}

	for _, line := range input {
		elements := strings.Split(line, " ")

		if strings.HasPrefix(line, "turn") {
			action := turnOn
			if elements[1] == "off" {
				action = turnOff
			}

			result = append(result, instruction{
				action:      action,
				topLeft:     toPoint(elements[2]),
				bottomRight: toPoint(elements[4])})

			continue
		}

		if strings.HasPrefix(line, "toggle") {
			result = append(result, instruction{
				action:      toggle,
				topLeft:     toPoint(elements[1]),
				bottomRight: toPoint(elements[3])})

		}
	}

	return result
}

func toPoint(xy string) point {
	elements := strings.Split(xy, ",")
	if len(elements) != 2 {
		panic("Cannot convert x,y to point")
	}

	return point{
		x: aoc.ToInt(elements[0]),
		y: aoc.ToInt(elements[1]),
	}
}
