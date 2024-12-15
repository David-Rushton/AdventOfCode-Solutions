package main

import "fmt"

func printGrid(direction rune, robot point, grid map[point]rune, width, height int) {
	fmt.Printf("Move: %s\n", string(direction))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			current := point{x, y}

			value, found := grid[current]
			if !found {
				value = '.'
			}

			if current == robot {
				value = '@'
			}

			fmt.Printf("%s", string(value))
		}
		fmt.Println()
	}
	fmt.Println()
}
