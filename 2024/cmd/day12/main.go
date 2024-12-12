package main

import (
	"fmt"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type plot string
type grid [][]plot
type regions [][]point

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("---  Day 12: Garden Groups ---")
	fmt.Println()

	grid := parse(aoc.Input)
	regions := toRegions(grid)

	var totalPriceByArea int
	var totalPriceBySides int
	for _, region := range regions {
		area, perimeter, sides := scanRegion(grid, region)
		price := area * perimeter
		totalPriceByArea += area * perimeter
		totalPriceBySides += area * sides

		fmt.Printf(
			" A region of %s plants with price %d x %d = %d.\n",
			grid[region[0].y][region[0].x],
			area,
			perimeter,
			price)
	}

	fmt.Println()
	fmt.Printf("Price by area: %d\n", totalPriceByArea)
	fmt.Printf("Price by sides: %d\n", totalPriceBySides)
}

func scanRegion(grid grid, region []point) (area, perimeter, sides int) {
	area = 0
	perimeter = 0

	crop := grid[region[0].y][region[0].x]
	perimeters := []point{}

	for _, point := range region {
		area++
		offsets := getOffset(grid, point)
		perimeter += 4 - len(offsets)
		for _, offset := range offsets {
			if grid[offset.y][offset.x] != crop {
				perimeter++

				if !slices.Contains(perimeters, offset) {
					perimeters = append(perimeters, point)
				}
			}
		}
	}

	return area, perimeter, len(perimeters)
}

func toRegions(grid grid) regions {
	visited := []point{}
	regions := regions{}

	for y, row := range grid {
		for x, _ := range row {
			p := point{x, y}
			if !slices.Contains(visited, p) {
				queue := []point{p}
				crop := grid[p.y][p.x]
				region := []point{}

				for len(queue) > 0 {
					current := queue[0]

					if grid[current.y][current.x] == crop {
						region = append(region, current)
						visited = append(visited, current)
					}

					for _, offset := range getOffset(grid, current) {
						if !slices.Contains(visited, offset) && !slices.Contains(queue, offset) {
							if grid[offset.y][offset.x] == crop {
								queue = append(queue, offset)
							}
						}
					}

					queue = queue[1:]
				}

				regions = append(regions, region)
			}
		}
	}

	return regions
}

func getOffset(grid grid, from point) []point {
	offsets := []struct {
		x int
		y int
	}{
		{x: +1, y: +0},
		{x: -1, y: +0},
		{x: +0, y: +1},
		{x: +0, y: -1},
	}

	result := []point{}
	for _, offset := range offsets {
		candidate := point{from.x + offset.x, from.y + offset.y}
		if candidate.y >= 0 && candidate.y < len(grid) && candidate.x >= 0 && candidate.x < len(grid[candidate.y]) {
			result = append(result, candidate)
		}
	}

	return result
}

func parse(input []string) grid {
	result := grid{}
	ignoreRunes := []rune{'\r', '\n', '\t', ' '}

	for y, line := range input {
		result = append(result, []plot{})
		for _, r := range line {
			if !slices.Contains(ignoreRunes, r) {
				result[y] = append(result[y], plot(r))
			}
		}
	}

	return result
}
