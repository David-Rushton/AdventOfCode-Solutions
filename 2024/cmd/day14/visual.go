package main

import (
	"fmt"
	"strconv"
)

func printGrid(grid dimensions, robots []robot) {
	// find robots
	locations := make(map[point]int)
	for _, robot := range robots {
		locations[robot.point]++
	}

	// print robots
	for y := 0; y < grid.y; y++ {
		for x := 0; x < grid.x; x++ {
			value := "."
			if count, found := locations[point{x, y}]; found {
				value = strconv.FormatInt(int64(count), 10)
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
}
