package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 1: Inverse Captcha ---")
	fmt.Println()

	var result int
	for _, captcha := range parse(aoc.GetInput(1)) {
		result = solve(captcha)
		fmt.Printf(" - %v == %d\n", captcha, result)
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func solve(captcha []int) int {
	var result int

	for i := 1; i < len(captcha); i++ {
		if captcha[i] == captcha[i-1] {
			result += captcha[i]
		}
	}

	return result
}

func parse(input []string) [][]int {
	var result = [][]int{}

	for i := range input {
		var next = []int{}
		for j := range len(input[i]) {
			next = append(next, aoc.ToInt(string(input[i][j])))
		}
		next = append(next, next[0])
		result = append(result, next)
	}

	return result

}
