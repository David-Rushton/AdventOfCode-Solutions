package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type point struct {
	id int
	x  int
	y  int
	z  int
}

func (p point) String() string {
	return fmt.Sprintf("%d: %d,%d,%d", p.id, p.x, p.y, p.z)
}

type connection struct {
	from     point
	to       point
	distance int
}

type circuit struct {
	from int
	to   map[int]bool
}

func (p point) distance(other point) int {
	xDist := math.Abs(float64(p.x - other.x))
	yDist := math.Abs(float64(p.y - other.y))
	zDist := math.Abs(float64(p.z - other.z))

	xSq := math.Pow(xDist, 2)
	ySq := math.Pow(yDist, 2)
	zSq := math.Pow(zDist, 2)

	return int(math.Sqrt(xSq + ySq + zSq))
}

func main() {
	fmt.Println("--- Day 8: Playground ---")
	fmt.Println()

	junctionBoxes := parseJunctionBoxes(aoc.Input)
	distances := []connection{}

	// Get distances between all junction boxes
	processed := map[string]bool{}
	for i := range len(junctionBoxes) {
		for j := range len(junctionBoxes) {
			if i == j {
				continue
			}

			key := fmt.Sprintf("%d-%d", min(junctionBoxes[i].id, junctionBoxes[j].id), max(junctionBoxes[i].id, junctionBoxes[j].id))
			if processed[key] {
				continue
			}
			processed[key] = true

			distance := junctionBoxes[i].distance(junctionBoxes[j])
			distances = append(distances, connection{
				from:     junctionBoxes[i],
				to:       junctionBoxes[j],
				distance: distance})
		}
	}

	// Sort distances.
	slices.SortFunc(distances, func(a, b connection) int {
		return a.distance - b.distance
	})

	// Make connections.
	connectionsRequired := 1000
	if aoc.TestMode {
		connectionsRequired = 10
	}
	if aoc.Star == aoc.StarTwo {
		connectionsRequired = len(distances)
	}

	locator := map[int]*circuit{}
	for i := range connectionsRequired {
		current := distances[i]

		fromCircuit, fromExists := locator[current.from.id]
		toCircuit, toExists := locator[current.to.id]

		if !fromExists {
			fromCircuit = &circuit{
				from: current.from.id,
				to:   map[int]bool{},
			}
		}

		if !toExists {
			toCircuit = &circuit{
				from: current.to.id,
				to:   map[int]bool{},
			}
		}

		fromCircuit.to[current.to.id] = true
		toCircuit.to[current.from.id] = true

		locator[current.from.id] = fromCircuit
		locator[current.to.id] = toCircuit

		if isOneCircuit(current.from.id, len(junctionBoxes), locator) {
			fmt.Println()
			fmt.Printf("Distance: %s x %s == %d\n", current.from, current.to, current.from.x*current.to.x)
			os.Exit(0)
		}
	}

	// Count circuits and sizes.
	visited := map[int]bool{}
	circuits := []int{}

	for _, circuit := range locator {
		if visited[circuit.from] {
			continue
		}

		nodes := 0
		q := append([]int{}, toSlice(circuit.to)...)

		for len(q) > 0 {
			current := q[0]
			q = q[1:]

			if visited[current] {
				continue
			}

			visited[current] = true
			nodes++

			q = append(q, toSlice(locator[current].to)...)
		}

		circuits = append(circuits, nodes)
		fmt.Printf(" - Circuit added %d\n", nodes)
	}

	// Star 1: Sum of top 3 largest circuits.
	total := 1
	slices.SortFunc(circuits, func(a, b int) int {
		return b - a
	})
	for i := range 3 {
		total *= circuits[i]
	}

	fmt.Println()
	fmt.Printf("Circuits: %d\n", len(circuits))
	fmt.Printf("Connections: %v\n", total)
}

func isOneCircuit(from, target int, locator map[int]*circuit) bool {
	q := []int{from}
	visited := map[int]bool{}

	if len(locator) < target {
		return false
	}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		q = append(q, toSlice(locator[current].to)...)
	}

	return len(visited) == target
}

func toSlice(m map[int]bool) []int {
	result := []int{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func parseJunctionBoxes(input []string) []point {
	var result []point
	var idSeed int

	for _, current := range input {
		elements := strings.Split(current, ",")
		result = append(result, point{
			id: idSeed,
			x:  aoc.ToInt(elements[0]),
			y:  aoc.ToInt(elements[1]),
			z:  aoc.ToInt(elements[2])})
		idSeed++
	}

	return result
}
