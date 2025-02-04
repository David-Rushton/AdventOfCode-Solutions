package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int
	y int
}

func (p *point) getOffset(x, y int) point {
	return point{
		p.x + x,
		p.y + y,
	}
}

func main() {
	fmt.Println("--- Day 2: Bathroom Security ---")
	fmt.Println()

	var result string

	keypad := getKeypad(aoc.Star == aoc.StarTwo)

	var current point
	for _, instruction := range aoc.GetInput(2) {
		for _, move := range instruction {
			var candidate point
			switch move {
			case 'U':
				candidate = current.getOffset(0, -1)

			case 'D':
				candidate = current.getOffset(0, 1)

			case 'L':
				candidate = current.getOffset(-1, 0)

			case 'R':
				candidate = current.getOffset(1, 0)

			default:
				panic(fmt.Sprintf("Unknown direction: %v", move))
			}

			if _, found := keypad[candidate]; found {
				current = candidate
			}
		}

		fmt.Printf(" - Key pressed : %v\n", keypad[current])
		result = fmt.Sprintf("%v%v", result, keypad[current])
	}

	fmt.Println()
	fmt.Printf("Result: %v\n", result)
}

func getKeypad(fancyMode bool) map[point]string {
	// In both modes the "5" key is 0, 0.  And therefore the starting point.

	if fancyMode {
		// 	     1
		// 	   2 3 4
		//   5 6 7 8 9
		// 	   A B C
		// 	     D
		return map[point]string{
			{x: +2, y: -2}: "1",
			{x: +1, y: -1}: "2",
			{x: +2, y: -1}: "3",
			{x: +3, y: -1}: "4",
			{x: +0, y: +0}: "5",
			{x: +1, y: +0}: "6",
			{x: +2, y: +0}: "7",
			{x: +3, y: +0}: "8",
			{x: +4, y: +0}: "9",
			{x: +1, y: +1}: "A",
			{x: +2, y: +1}: "B",
			{x: +3, y: +1}: "C",
			{x: +2, y: +2}: "D",
		}
	}

	// 1 2 3
	// 4 5 6
	// 7 8 9
	return map[point]string{
		{x: -1, y: -1}: "1",
		{x: +0, y: -1}: "2",
		{x: +1, y: -1}: "3",
		{x: -1, y: +0}: "4",
		{x: +0, y: +0}: "5",
		{x: +1, y: +0}: "6",
		{x: -1, y: +1}: "7",
		{x: +0, y: +1}: "8",
		{x: +1, y: +1}: "9",
	}
}
