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
		totalPriceByArea += area * perimeter
		totalPriceBySides += area * sides

		fmt.Printf(
			" A region of %s plants with price % 3d x % 3d = % 5d.\n",
			grid[region[0].y][region[0].x],
			area,
			sides,
			area*sides)
	}

	fmt.Println()
	fmt.Printf("Price by area: %d\n", totalPriceByArea)
	fmt.Printf("Price by sides: %d\n", totalPriceBySides)
}

func scanRegion(grid grid, region []point) (area, perimeter, sides int) {
	area = 0
	perimeter = 0
	sides = 0

	sameCrop := func(from point, xOffset, yOffset int) bool {
		offsetPoint := point{from.x + xOffset, from.y + yOffset}
		if inGrid(grid, offsetPoint) {
			return grid[from.y][from.x] == grid[offsetPoint.y][offsetPoint.x]
		}

		return false
	}

	perimeters := []point{}

	for _, point := range region {
		area++
		for _, offset := range getOffsets(point) {
			if !inGrid(grid, offset) || grid[point.y][point.x] != grid[offset.y][offset.x] {
				perimeter++

				if !slices.Contains(perimeters, point) {
					perimeters = append(perimeters, point)
				}
			}
		}

		// ext corner: top left.
		if !sameCrop(point, 0, -1) && !sameCrop(point, -1, 0) {
			sides++
		}

		// ext corner: top right.
		if !sameCrop(point, 0, -1) && !sameCrop(point, 1, 0) {
			sides++
		}

		// ext corner: bottom left.
		if !sameCrop(point, 0, 1) && !sameCrop(point, -1, 0) {
			sides++
		}

		// ext corner: bottom right.
		if !sameCrop(point, 0, 1) && !sameCrop(point, 1, 0) {
			sides++
		}

		// int corner: top left.
		if !sameCrop(point, -1, -1) && sameCrop(point, -1, 0) && sameCrop(point, 0, -1) {
			sides++
		}

		// int corner: top right.
		if !sameCrop(point, 1, -1) && sameCrop(point, 1, 0) && sameCrop(point, 0, -1) {
			sides++
		}

		// int corner: bottom left.
		if !sameCrop(point, -1, 1) && sameCrop(point, -1, 0) && sameCrop(point, 0, 1) {
			sides++
		}

		// int corner: bottom right.
		if !sameCrop(point, 1, 1) && sameCrop(point, 1, 0) && sameCrop(point, 0, 1) {
			sides++
		}
	}

	return area, perimeter, sides
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

					for _, offset := range getOffsets(current) {
						if inGrid(grid, offset) && !slices.Contains(visited, offset) && !slices.Contains(queue, offset) {
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

func getOffsets(from point) []point {
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
		result = append(result, point{from.x + offset.x, from.y + offset.y})
	}

	return result
}

func inGrid(grid grid, point point) bool {
	return point.y >= 0 && point.y < len(grid) && point.x >= 0 && point.x < len(grid[point.y])
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
