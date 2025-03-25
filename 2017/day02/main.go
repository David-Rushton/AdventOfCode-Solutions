package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 2: Corruption Checksum ---")
	fmt.Println()

	data := parse(aoc.GetInput(2))
	checksum := getChecksum(data)
	evenlyDivisibleValues := getEvenlyDivisibleValues(data)

	fmt.Println()
	fmt.Printf("Checksum: %d\n", checksum)
	fmt.Printf("Evenly Divisible Values: %d\n", evenlyDivisibleValues)
}

func getChecksum(data [][]int) int {
	var result int

	for i := range data {
		var min int = math.MaxInt
		var max int

		for j := range data[i] {
			if data[i][j] < min {
				min = data[i][j]
			}

			if data[i][j] > max {
				max = data[i][j]
			}
		}

		result += max - min
	}

	return result
}

func getEvenlyDivisibleValues(data [][]int) int {
	var result int

nextRow:
	for i := range data {
		for j := range data[i] {
			for k := range data[i] {
				if j != k && data[i][j]%data[i][k] == 0 {
					fmt.Printf(" - Evenly divisible value found for: %v = %d & %d\n", data[i], data[i][j], data[i][k])
					result += data[i][j] / data[i][k]
					continue nextRow
				}
			}
		}
	}

	return result
}

func parse(input []string) [][]int {
	var result [][]int

	for i := range input {
		var next []int
		for number := range strings.SplitSeq(strings.ReplaceAll(input[i], "\t", " "), " ") {
			next = append(next, aoc.ToInt(number))
		}
		result = append(result, next)
	}

	return result
}
