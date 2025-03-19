package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 22: Grid Computing ---")
	fmt.Println()

	nodes, nodeIndex := parse(aoc.GetInput(22))
	viablePairs := countViablePairs(nodes)
	printMap(nodeIndex)

	fmt.Println()
	fmt.Printf("Viable Pairs: %d\n", viablePairs)
}

// Manually count required steps from map.
// No really.
func printMap(nodeIndex map[point]node) {
	target := point{x: 29, y: 0}

	for y := range 35 {
		for x := range 35 {
			if node, exists := nodeIndex[point{x, y}]; exists {
				if node.address == target {
					fmt.Print("T")
					continue
				}

				if node.used > 100 {
					fmt.Print("#")
					continue
				}

				if node.used == 0 {
					fmt.Print("_")
					continue
				}

				if node.used > 0 {
					fmt.Print(".")
					continue
				}
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func countViablePairs(nodes []node) int {
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
