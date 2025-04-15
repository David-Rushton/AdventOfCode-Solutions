package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

type nodeIndex map[int]*node

type nodes []*node

type node struct {
	id          int
	connectedTo []*node
}

func main() {
	fmt.Println("--- Day 12: Digital Plumber ---")
	fmt.Println()

	var nodes, nodeIndex = parse(aoc.GetInput(12))
	var connectedCount = len(getConnected(nodeIndex[0]))
	var groupCount = countGroups(nodes)

	fmt.Println()
	fmt.Printf("Nodes connected to 0: %d\n", connectedCount)
	fmt.Printf("Node groups: %d\n", groupCount)
}

func countGroups(nodes nodes) int {
	var connectionsFound = map[int]bool{}
	var groups int

	for _, node := range nodes {
		if !connectionsFound[node.id] {
			for connectionId, _ := range getConnected(node) {
				connectionsFound[connectionId] = true
			}

			connectionsFound[node.id] = true
			groups++
		}
	}

	return groups
}

func getConnected(root *node) map[int]bool {
	var connectionsFound = map[int]bool{}
	var queue = []*node{root}

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		for _, connectedTo := range current.connectedTo {
			if !connectionsFound[connectedTo.id] {
				queue = append(queue, connectedTo)
				connectionsFound[connectedTo.id] = true
			}
		}
	}

	return connectionsFound
}

func parse(input []string) (nodes, nodeIndex) {
	var nodes = nodes{}
	var index = nodeIndex{}

	var createIfNotExists = func(id int) *node {
		if _, exists := index[id]; !exists {
			var current = &node{id, []*node{}}
			nodes = append(nodes, current)
			index[id] = current
		}

		return index[id]
	}

	for i := range input {
		var idAndConnections = strings.Split(strings.ReplaceAll(input[i], " ", ""), "<->")
		var id = aoc.ToInt(idAndConnections[0])
		var connections = strings.Split(idAndConnections[1], ",")
		var current = createIfNotExists(id)

		for _, connection := range connections {
			var connectedToId = aoc.ToInt(connection)
			var connectedNode = createIfNotExists(connectedToId)
			current.connectedTo = append(current.connectedTo, connectedNode)
		}

	}

	return nodes, index
}
