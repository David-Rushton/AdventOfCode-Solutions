package main

import (
	"fmt"
)

func print(start, end point, width, height int, route, walls map[point]int) {
	for y := 0; y < height; y++ {
		fmt.Print("  ")
		for x := 0; x < width; x++ {
			current := point{x, y}
			value := "."

			if _, found := walls[current]; found {
				value = "#"
			}

			if current == start {
				value = "S"
			}

			if current == end {
				value = "E"
			}

			if _, found := route[current]; found {
				value = "*"
			}

			fmt.Print(value)
		}
		fmt.Println()
	}
	fmt.Println()
}
