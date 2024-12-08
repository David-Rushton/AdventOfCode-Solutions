package main

import (
	"fmt"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

type antenna struct {
	id        int
	point     point
	frequency rune
}

func main() {
	fmt.Println("--- Day 8: Resonant Collinearity ---")
	fmt.Println()

	antennas, antennasByPoint, bottom, right := parse(aoc.Input)
	antennaPairs := [][]antenna{}
	antinodes := []point{}

	// Find pairs.
	foundPairIds := []string{}
	for _, outer := range antennas {
		for _, inner := range antennas {
			if outer.frequency == inner.frequency {
				pairId := getPairId(outer, inner)
				if outer.id != inner.id && !slices.Contains(foundPairIds, pairId) {
					antennaPairs = append(antennaPairs, []antenna{outer, inner})
					foundPairIds = append(foundPairIds, pairId)
				}
			}
		}
	}

	// Find antinodes
	for _, antennaPair := range antennaPairs {
		xOffset := antennaPair[0].point.x - antennaPair[1].point.x
		yOffset := antennaPair[0].point.y - antennaPair[1].point.y

		limit := 1
		if aoc.Star == aoc.StarTwo {
			limit = -1
		}

		candidates := slices.Concat(
			getAntinodes(antennaPair[0].point, +1, xOffset, yOffset, bottom, right, limit),
			getAntinodes(antennaPair[1].point, -1, xOffset, yOffset, bottom, right, limit))
		for _, candidate := range candidates {
			if !slices.Contains(antinodes, candidate) {
				antinodes = append(antinodes, candidate)
			}
		}
	}

	// print result
	for y := 0; y <= bottom; y++ {
		for x := 0; x <= right; x++ {
			point := point{x, y}
			value := "."

			if antenna, found := antennasByPoint[point]; found {
				value = string(antenna.frequency)
			}

			if slices.Contains(antinodes, point) {
				value = "#"
			}

			fmt.Print(value)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", len(antinodes))
}

func getAntinodes(from point, direction, xOffset, yOffset, bottom, right, limit int) []point {
	result := []point{}
	offset := func(p point) point {
		if direction > 0 {
			return point{p.x + xOffset, p.y + yOffset}
		} else {
			return point{p.x - xOffset, p.y - yOffset}
		}
	}

	if limit < 0 {
		result = append(result, from)
	}

	candidate := from
	for {
		candidate = offset(candidate)
		if candidate.x < 0 || candidate.x > right || candidate.y < 0 || candidate.y > bottom {
			break
		}
		result = append(result, candidate)

		if limit > 0 && limit >= len(result) {
			break
		}
	}

	return result
}

func getPairId(left, right antenna) string {
	if left.id < right.id {
		return fmt.Sprintf("%d.%d", left.id, right.id)
	} else {
		return fmt.Sprintf("%d.%d", right.id, left.id)
	}
}

func parse(input []string) (antennas []antenna, antennasByPoint map[point]antenna, bottom, right int) {
	antennas = []antenna{}
	antennasByPoint = make(map[point]antenna)
	idSeed := 0

	for y, line := range input {
		for x, r := range line {
			if r != '.' {
				point := point{x, y}
				antenna := antenna{
					id:        idSeed,
					point:     point,
					frequency: r,
				}
				antennas = append(antennas, antenna)
				antennasByPoint[point] = antenna
				idSeed++
			}

			if x > right {
				right = x
			}
		}

		if y > bottom {
			bottom = y
		}
	}

	return antennas, antennasByPoint, bottom, right
}
