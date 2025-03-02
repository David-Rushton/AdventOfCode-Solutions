package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type point struct {
	x int
	y int
}

type node struct {
	address   point
	size      int
	used      int
	available int
}

func main() {
	fmt.Println("--- Day 22: Grid Computing ---")
	fmt.Println()

	nodes, nodeIndex := parse(aoc.GetInput(22))
	viablePairs := countViablePairs(nodes, nodeIndex)

	fmt.Println()
	fmt.Printf("Viable Pairs: %d\n", viablePairs)
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
