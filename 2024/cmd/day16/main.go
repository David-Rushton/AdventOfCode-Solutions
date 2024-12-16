package main

import (
	"fmt"
	"math"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type directedPoint struct {
	point     point
	direction direction
}

type direction rune

const (
	north direction = '^'
	east  direction = '>'
	south direction = 'v'
	west  direction = '<'
)

type point struct {
	x int
	y int
}

const (
	wall = '#'
	path = '.'
)

func main() {
	fmt.Println("--- Day 16: Reindeer Maze ---")
	fmt.Println()

	reindeer, end, maze := parse(aoc.Input)
	score, steps := findBestRoute(reindeer, end, maze)

	fmt.Println()
	fmt.Printf("Score: %d | Steps: %d", score, steps)
}

type routeState struct {
	reindeer directedPoint
	steps    []directedPoint
	score    int
}

func findBestRoute(reindeer directedPoint, end point, maze [][]rune) (bestScore int, totalSteps int) {
	bestScore = math.MaxInt
	queue := []routeState{{reindeer, []directedPoint{reindeer}, 0}}
	visited := make(map[directedPoint]int)
	bestRoutes := make(map[int][]directedPoint)

	// Best route.
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.steps) > 10000 {
			continue
		}

		if state.score > bestScore {
			continue
		}

		if state.reindeer.point == end {
			printMaze(maze, state)
			if state.score <= bestScore {
				bestRouteTmp := bestRoutes[state.score]
				bestRouteTmp = append(bestRouteTmp, state.steps...)
				bestRoutes[state.score] = bestRouteTmp
				bestScore = state.score
				continue
			}
		}

		for _, candidate := range getOffsets(state.reindeer) {
			if maze[candidate.point.y][candidate.point.x] == path {
				score := state.score + 1
				if state.reindeer.direction != candidate.direction {
					score += 1000
				}

				if previous, found := visited[candidate]; found {
					if previous < score {
						continue
					}
				}
				visited[candidate] = score

				newSteps := make([]directedPoint, len(state.steps))
				copy(newSteps, state.steps)

				queue = append(queue, routeState{
					reindeer: candidate,
					steps:    append(newSteps, candidate),
					score:    score})
			}
		}
	}

	// Best steps.
	buffer := make(map[point]int)
	for _, v := range bestRoutes[bestScore] {
		buffer[v.point]++
	}

	return bestScore, len(buffer)
}

func getOffsets(reindeer directedPoint) []directedPoint {
	n := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y - 1}, north}
	e := directedPoint{point{x: reindeer.point.x + 1, y: reindeer.point.y + 0}, east}
	s := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y + 1}, south}
	w := directedPoint{point{x: reindeer.point.x - 1, y: reindeer.point.y + 0}, west}

	switch reindeer.direction {
	case north:
		return []directedPoint{n, e, w}
	case east:
		return []directedPoint{e, s, n}
	case south:
		return []directedPoint{s, w, e}
	case west:
		return []directedPoint{w, n, s}
	}

	panic("Reindeer facing unknown direction.")
}

func printMaze(maze [][]rune, route routeState) {
	fmt.Printf("Route Found: %d\n", route.score)

	getStep := func(p point) (rune, bool) {
		for i, step := range route.steps {
			if step.point == p {
				switch i {
				case 0:
					return 'S', true
				case len(route.steps) - 1:
					return 'E', true
				default:
					return rune(step.direction), true
				}
			}
		}

		return 0, false
	}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			current := point{x, y}
			value := maze[current.y][current.x]

			if r, found := getStep(current); found {
				value = r
			}

			fmt.Printf("%s", string(value))

		}
		fmt.Println()
	}

	fmt.Println()
}

func parse(input []string) (reindeer directedPoint, end point, maze [][]rune) {
	maze = [][]rune{}

	for y, line := range input {
		maze = append(maze, []rune{})
		for x, r := range line {
			switch r {
			case wall:
				maze[y] = append(maze[y], r)
			case path:
				maze[y] = append(maze[y], r)
			case 'S':
				reindeer = directedPoint{point{x, y}, east}
				maze[y] = append(maze[y], path)
			case 'E':
				end = point{x, y}
				maze[y] = append(maze[y], path)
			}
		}
	}

	return reindeer, end, maze
}
