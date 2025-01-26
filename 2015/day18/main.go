package main

import (
	"fmt"
	"time"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("--- Day 18: Like a GIF For Your Yard ---")
	fmt.Println()

	frames := 100
	if aoc.TestMode {
		frames = 4
	}

	cornersOn := false
	if aoc.Star == aoc.StarTwo {
		cornersOn = true
	}

	lights := parse(aoc.GetInput(18))
	lights = animate(frames, cornersOn, lights)
	result := countIlluminated(lights)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func animate(frames int, cornersOn bool, lights [][]bool) [][]bool {
	printFrame(-1, lights)

	// Assumes rectangle.
	firstX := 0
	firstY := 0
	lastX := len(lights[firstY]) - 1
	lastY := len(lights) - 1

	// When requested, override corners.
	// Must occur before we start comparing lights to each other.
	current := lights
	if cornersOn {
		current[firstY][firstX] = true
		current[firstY][lastX] = true
		current[lastY][firstX] = true
		current[lastY][lastX] = true
	}

	for frame := 0; frame < frames; frame++ {
		next := deepCopy(current)

		for y, row := range current {
			for x, illuminated := range row {
				if y == firstY || y == lastY {
					if x == firstX || x == lastX {
						next[x][y] = true
						continue
					}
				}

				illuminatedNeighbours := countIlluminatedNeighbours(x, y, current)

				if illuminated && (illuminatedNeighbours < 2 || illuminatedNeighbours > 3) {
					next[y][x] = false
				}

				if !illuminated && illuminatedNeighbours == 3 {
					next[y][x] = true
				}
			}
		}

		current = next
		printFrame(frame, current)
	}

	return current
}

func printFrame(frame int, lights [][]bool) {
	fmt.Print("\033[3;1H")
	fmt.Printf("Frame %d:     \n", frame)

	for _, row := range lights {
		for _, illuminated := range row {
			if illuminated {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println()
	time.Sleep(time.Millisecond * 350)
}

func countIlluminatedNeighbours(x, y int, lights [][]bool) int {
	var result int

	// Count illuminated neighbours.
	// Where illuminated is `lights[y][x] == true`.
	// Where neighbour is any adjacent cell, including diagonals.
	for yOffset := -1; yOffset < 2; yOffset++ {
		for xOffset := -1; xOffset < 2; xOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}
			yCandidate := y + yOffset
			if yCandidate < 0 || yCandidate >= len(lights) {
				continue
			}

			xCandidate := x + xOffset
			if xCandidate < 0 || xCandidate >= len(lights[yCandidate]) {
				continue
			}

			if lights[yCandidate][xCandidate] {
				result++
			}
		}
	}

	return result
}

func countIlluminated(lights [][]bool) int {
	var result int

	for _, row := range lights {
		for _, illuminated := range row {
			if illuminated {
				result++
			}
		}
	}

	return result
}

func deepCopy(source [][]bool) [][]bool {
	result := make([][]bool, len(source))

	for i := range source {
		result[i] = make([]bool, len(source[i]))
		copy(result[i], source[i])
	}

	return result
}

func parse(input []string) [][]bool {
	result := [][]bool{}

	for y, row := range input {
		result = append(result, make([]bool, len(row)))

		for x, cell := range row {
			if cell == '#' {
				result[y][x] = true
			}
		}
	}

	return result
}
