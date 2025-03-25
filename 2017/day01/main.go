package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

func main() {
	fmt.Println("--- Day 1: Inverse Captcha ---")
	fmt.Println()

	var result1 int
	var result2 int
	for _, captcha := range parse(aoc.GetInput(1)) {
		result1 = solve1(captcha)
		fmt.Printf(" - 1: %v == %d \n", captcha, result1)

		result2 = solve2(captcha[0 : len(captcha)-1])
		fmt.Printf(" - 2: %v == %d \n", captcha, result2)

		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}

func solve1(captcha []int) int {
	var result int

	for i := 1; i < len(captcha); i++ {
		if captcha[i] == captcha[i-1] {
			result += captcha[i]
		}
	}

	return result
}

func solve2(captcha []int) int {
	var result int

	for i := range captcha {
		j := (i + len(captcha)/2) % len(captcha)
		if captcha[i] == captcha[j] {
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
