package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 2: Bathroom Security ---")
	fmt.Println()

	var result string

	// TODO: Start here.
	keypad := map[point]int{
		{x: -1, y: -1}: 1,
		{x: +0, y: -1}: 2,
		{x: +1, y: -1}: 3,
		{x: -1, y: +0}: 4,
		{x: +0, y: +0}: 5,
		{x: +1, y: +0}: 6,
		{x: -1, y: +1}: 7,
		{x: +0, y: +1}: 8,
		{x: +1, y: +1}: 9,
	}

	var current point
	for _, instruction := range aoc.GetInput(2) {
		for _, move := range instruction {
			switch move {
			case 'U':
				current.y--
				if current.y < -1 {
					current.y = -1
				}
			case 'D':
				current.y++
				if current.y > 1 {
					current.y = 1
				}
			case 'L':
				current.x--
				if current.x < -1 {
					current.x = -1
				}
			case 'R':
				current.x++
				if current.x > 1 {
					current.x = 1
				}
			default:
				panic(fmt.Sprintf("Unknown direction: %v", move))
			}
		}

		if keypad[current] > 0 {
			fmt.Printf(" - Key pressed : %v\n", keypad[current])
			result = fmt.Sprintf("%v%v", result, keypad[current])
		}
	}

	fmt.Println()
	fmt.Printf("Result: %v\n", result)
}
