package main

import (
	"fmt"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

func main() {
	switch aoc.Star {
	case aoc.StarOne:
		v1()
	default:
		v2()
	}
}

func v1() {
	fmt.Println("--- Day 3: Lobby --")
	fmt.Println()

	var totalJoltage int

	for _, line := range aoc.Input {
		var maxJoltage int

		for i := range len(line) {
			for j := i + 1; j < len(line); j++ {

				candidate := fmt.Sprintf("%c%c", line[i], line[j])
				num, _ := strconv.Atoi(candidate)
				if num > maxJoltage {
					maxJoltage = num
				}
			}
		}

		totalJoltage += maxJoltage
		fmt.Printf(" - %s -> %d\n", line, maxJoltage)
	}

	fmt.Println()
	fmt.Printf("Joltage: %d\n", totalJoltage)
}

func v2() {
	fmt.Println("--- Day 3: Lobby --")
	fmt.Println()

	var totalJoltage int

	for _, line := range aoc.Input {
		var joltage string
		var from int

		for i := range 12 {
			var mostSignificantDigit byte
			var mostSignificantPosition int

			for j := from; j < len(line)-11+i; j++ {
				if line[j] > mostSignificantDigit {
					mostSignificantDigit = line[j]
					mostSignificantPosition = j
				}
			}

			from = mostSignificantPosition + 1
			joltage += string(line[mostSignificantPosition])
		}

		joltageNum, _ := strconv.Atoi(joltage)
		totalJoltage += joltageNum
		fmt.Printf(" - %s -> %s\n", line, joltage)
	}

	fmt.Println()
	fmt.Printf("Joltage: %d\n", totalJoltage)
}
