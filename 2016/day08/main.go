package main

import (
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type instruction struct {
	name   string
	value1 int
	value2 int
}

type point struct {
	x int
	y int
}

func main() {
	aoc.ClearScreen()
	aoc.HideCursor()

	fmt.Println("--- Day 8: Two-Factor Authentication ---")
	fmt.Println()

	width := 50
	height := 6
	if aoc.TestMode {
		width = 7
		height = 3
	}

	instructions := parse(aoc.GetInput(8))
	result := displayImage(instructions, width, height)

	aoc.ShowCursor()
	fmt.Println()
	fmt.Printf("Lit: %d", result)
}

func displayImage(instructions []instruction, width, height int) int {
	var result int

	screen := map[point]bool{}

	for _, instruction := range instructions {
		nextScreen := map[point]bool{}
		maps.Copy(nextScreen, screen)

		// Update display.
		if instruction.name == "rect" {
			for y := 0; y < instruction.value2; y++ {
				for x := 0; x < instruction.value1; x++ {
					nextScreen[point{x, y}] = true
				}
			}
		}

		if instruction.name == "rotate row" {
			// Clear row.
			for x := 0; x < width; x++ {
				nextScreen[point{x, instruction.value1}] = false
			}

			// Shift existing.
			for pixel, state := range screen {
				if state && pixel.y == instruction.value1 {
					nextPoint := point{(pixel.x + instruction.value2) % width, pixel.y}
					nextScreen[nextPoint] = true
				}
			}
		}

		if instruction.name == "rotate column" {
			// Clear column.
			for y := 0; y < height; y++ {
				nextScreen[point{instruction.value1, y}] = false
			}

			// Shift existing.
			for pixel, state := range screen {
				if state && pixel.x == instruction.value1 {
					nextPoint := point{pixel.x, (pixel.y + instruction.value2) % height}
					nextScreen[nextPoint] = true
				}
			}
		}

		screen = nextScreen

		// Print to console.
		aoc.MoveCursor(3, 1)

		result = 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if screen[point{x, y}] {
					fmt.Print("#")
					result++
				} else {
					fmt.Print(".")
				}
			}

			fmt.Println()
		}

		time.Sleep(time.Millisecond * 50)
	}

	return result
}

func parse(input []string) []instruction {
	result := []instruction{}

	for _, current := range input {
		elements := strings.Split(current, " ")

		if strings.HasPrefix(current, "rect") {
			xy := strings.Split(elements[1], "x")

			result = append(result, instruction{
				name:   "rect",
				value1: aoc.ToInt(xy[0]),
				value2: aoc.ToInt(xy[1])})
			continue
		}

		value1 := aoc.ToInt(strings.Split(elements[2], "=")[1])

		if strings.HasPrefix(current, "rotate row") {
			result = append(result, instruction{
				name:   "rotate row",
				value1: value1,
				value2: aoc.ToInt(elements[4])})
			continue
		}

		if strings.HasPrefix(current, "rotate column") {
			result = append(result, instruction{
				name:   "rotate column",
				value1: value1,
				value2: aoc.ToInt(elements[4])})
			continue
		}
	}

	return result
}
