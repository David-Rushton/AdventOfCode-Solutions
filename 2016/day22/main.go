package main

import (
	"fmt"
	"maps"
	"math"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 22: Grid Computing ---")
	fmt.Println()

	nodes, nodeIndex := parse(aoc.GetInput(22))
	viablePairs := countViablePairs(nodes, nodeIndex)
	minSteps := getMinSteps(nodeIndex)

	fmt.Println()
	fmt.Printf("Viable Pairs: %d\n", viablePairs)
	fmt.Printf("Minimum Steps: %d\n", minSteps)
}

func countViablePairs(nodes []node, nodeIndex map[point]node) int {
	var result int

	for i := range nodes {
		for j := range nodes {
			if i != j {
				if nodes[i].used > 0 && nodes[i].used <= nodes[j].available {
					result++
				}
			}
		}
	}

	return result
}

func getMinSteps(nodeIndex map[point]node) int {
	// Target is cell with y of 0 and highest x.
	var target point
	for i := 0; ; i++ {
		if node, exists := nodeIndex[point{x: i, y: 0}]; exists {
			target = node.address
			continue
		}

		break
	}

	// Solve.
	type state struct {
		cluster map[point]node
		payload point
		steps   int
	}

	var queue = []state{{nodeIndex, target, 0}}
	var destination = point{x: 0, y: 0}
	var minSteps = math.MaxInt
	var visited = map[int64]bool{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf(" - Queue: %v\r", len(queue))

		// Win.
		if current.payload == destination {
			if current.steps < minSteps {
				minSteps = current.steps
				fmt.Printf(" - New best found: %v\n", minSteps)
				continue
			}
		}

		// Lose.
		if current.steps > minSteps {
			continue
		}

		// Play on.
		connections := getViablePairs(current.cluster, current.payload)
		for _, connection := range connections {
			var newCluster = map[point]node{}
			from, to := connection.move()
			maps.Copy(newCluster, current.cluster)
			newCluster[from.address] = from
			newCluster[to.address] = to

			newPayload := current.payload
			if newPayload == from.address {
				newPayload = to.address
			}

			if !visited[connection.hash] {
				// visited[connection.hash] = true

				queue = append(queue, state{
					cluster: newCluster,
					payload: newPayload,
					steps:   current.steps + 1})
			}

		}
	}

	return minSteps
}

func getViablePairs(nodes map[point]node, payload point) []connectedNodes {
	var result []connectedNodes

	for _, left := range nodes {
		offsets := []point{
			{x: left.address.x + 1, y: left.address.y},
			{x: left.address.x - 1, y: left.address.y},
			{x: left.address.x, y: left.address.y + 1},
			{x: left.address.x, y: left.address.y - 1},
		}
		for _, offset := range offsets {
			if right, exists := nodes[offset]; exists {
				if left.used > 0 && left.used <= right.available {
					result = append(result, connectedNodes{left, right, getConnectedHash(left, right)})
				}
			}
		}
	}

	return result
}

func getConnectedHash(left, right node) int64 {
	return int64(left.address.x*1e11) +
		int64(left.address.y*1e9) +
		int64(right.address.x*1e7) +
		int64(right.address.y*1e5) +
		int64(left.used)
}

func parse(input []string) (nodes []node, nodeIndex map[point]node) {
	cleanString := func(s string) string {
		for strings.Contains(s, "  ") {
			s = strings.ReplaceAll(s, "  ", " ")
		}

		s = strings.ReplaceAll(s, "/dev/grid/node-x", "dev ")
		s = strings.ReplaceAll(s, "-y", " ")
		s = strings.ReplaceAll(s, "T", "")

		return s
	}

	nodes = []node{}
	nodeIndex = map[point]node{}

	for i := range len(input) {
		elements := strings.Split(
			cleanString(input[i]),
			" ")

		if elements[0] == "dev" {
			address := point{aoc.ToInt(elements[1]), aoc.ToInt(elements[2])}
			next := node{
				address:   address,
				size:      aoc.ToInt(elements[3]),
				used:      aoc.ToInt(elements[4]),
				available: aoc.ToInt(elements[5])}

			nodes = append(nodes, next)
			nodeIndex[address] = next
		}
	}

	return nodes, nodeIndex
}
