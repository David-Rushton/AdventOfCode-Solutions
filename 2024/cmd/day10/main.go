package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 10: Hoof It ---")
	fmt.Println()

	grid, trailHeads := parse(aoc.Input)

	var starOne int
	var starTwo int
	for _, trailHead := range trailHeads {
		destinations, rating := countTrails(trailHead, grid)
		starOne += destinations
		starTwo += rating

	}

	fmt.Println()
	fmt.Printf("Result: %d & %d\n", starOne, starTwo)
}

func countTrails(from point, grid [][]int) (destinations, rating int) {
	offsets := []point{
		{x: +1, y: +0},
		{x: +0, y: +1},
		{x: -1, y: +0},
		{x: +0, y: -1},
	}
	queue := []point{from}
	found := []point{}
	routes := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, offset := range offsets {
			offsetPoint := point{current.x + offset.x, current.y + offset.y}

			if offsetPoint.y >= 0 && offsetPoint.y < len(grid) {
				if offsetPoint.x >= 0 && offsetPoint.x < len(grid[offsetPoint.y]) {
					diff := grid[offsetPoint.y][offsetPoint.x] - grid[current.y][current.x]
					if diff == 1 {
						if grid[offsetPoint.y][offsetPoint.x] == 9 {
							routes++
							if !slices.Contains(found, offsetPoint) {
								found = append(found, offsetPoint)
							}
						} else {
							queue = append(queue, offsetPoint)
						}
					}
				}
			}
		}
	}

	return len(found), routes
}

func parse(input []string) (grid [][]int, trailHeads []point) {
	grid = [][]int{}
	trailHeads = []point{}

	for y, line := range input {
		grid = append(grid, []int{})

		for x, heightTxt := range line {
			heightNum, err := strconv.ParseInt(string(heightTxt), 10, 64)
			if err != nil {
				if heightNum == '.' {
					heightNum = -99
				} else {
					log.Fatalf("Cannot covert %v to a number.\n", heightTxt)
				}
			}

			grid[y] = append(grid[y], int(heightNum))

			if heightNum == 0 {
				trailHeads = append(trailHeads, point{x, y})
			}
		}
	}

	return grid, trailHeads
}
