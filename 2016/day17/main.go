package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Day 17: Two Steps Forward ---")
	fmt.Println()

	grid := map[point]bool{
		{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true,
		{0, 1}: true, {1, 1}: true, {2, 1}: true, {3, 1}: true,
		{0, 2}: true, {1, 2}: true, {2, 2}: true, {3, 2}: true,
		{0, 3}: true, {1, 3}: true, {2, 3}: true, {3, 3}: true,
	}

	for _, passcode := range aoc.GetInput(17) {
		fmt.Printf(" - Shorted route: %v\n", findShortestRoute(grid, passcode))
		fmt.Printf(" - Longest route: %d\n", len(findLongestRoute(grid, passcode)))
		fmt.Println()
	}
}

func findShortestRoute(grid map[point]bool, passcode string) string {
	var shortestRoute string
	var destination = point{3, 3}
	var unlocked = []rune{'b', 'c', 'd', 'e', 'f'}

	var solve func(from point, route string)
	solve = func(from point, route string) {
		// fmt.Printf(" - Visiting %vx%v | Route %v\n", from.x, from.y, route)

		// We've arrived at our destination.
		if from == destination {
			if len(route) < len(shortestRoute) || len(shortestRoute) == 0 {
				shortestRoute = route
				fmt.Printf(
					" - Found new shortest route for passcode %v: %v (%d)\n",
					passcode,
					shortestRoute,
					len(shortestRoute))
			}

			return
		}

		// We've already found a better route.
		if len(route) > len(shortestRoute) && len(shortestRoute) > 0 {
			return
		}

		// Visit points behind unlocked doors.
		hash := getMd5(passcode + route)
		for i := 0; i < 4 && i < len(hash); i++ {
			if slices.Contains(unlocked, rune(hash[i])) {
				var candidate point = point{from.x, from.y}
				var direction string
				switch i {
				// up
				case 0:
					candidate.y--
					direction = "U"
					// down
				case 1:
					candidate.y++
					direction = "D"
					// left
				case 2:
					candidate.x--
					direction = "L"
					// right
				case 3:
					candidate.x++
					direction = "R"
				}

				if grid[candidate] {
					solve(candidate, route+direction)
				}
			}
		}
	}

	solve(point{0, 0}, "")

	return shortestRoute
}

func findLongestRoute(grid map[point]bool, passcode string) string {
	var longestRoute string
	var destination = point{3, 3}
	var unlocked = []rune{'b', 'c', 'd', 'e', 'f'}

	var solve func(from point, route string)
	solve = func(from point, route string) {
		// fmt.Printf(" - Visiting %vx%v | Route %v\n", from.x, from.y, route)

		// We've arrived at our destination.
		if from == destination {
			if len(route) > len(longestRoute) {
				longestRoute = route
				// fmt.Printf(
				// 	" - Found new longest route for passcode %v: %d)\n",
				// 	passcode,
				// 	len(longestRoute))
			}

			return
		}

		// Visit points behind unlocked doors.
		hash := getMd5(passcode + route)
		for i := 0; i < 4 && i < len(hash); i++ {
			if slices.Contains(unlocked, rune(hash[i])) {
				var candidate point = point{from.x, from.y}
				var direction string
				switch i {
				// up
				case 0:
					candidate.y--
					direction = "U"
					// down
				case 1:
					candidate.y++
					direction = "D"
					// left
				case 2:
					candidate.x--
					direction = "L"
					// right
				case 3:
					candidate.x++
					direction = "R"
				}

				if grid[candidate] {
					solve(candidate, route+direction)
				}
			}
		}
	}

	solve(point{0, 0}, "")

	return longestRoute
}

func getMd5(s string) string {
	data := md5.Sum([]byte(s))
	return hex.EncodeToString(data[:])
}
