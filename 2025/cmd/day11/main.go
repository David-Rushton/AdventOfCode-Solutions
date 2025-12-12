package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type devices map[string][]string

func main() {
	fmt.Println("--- Day 11: Reactor ---")
	fmt.Println()

	devices := parse(aoc.Input)
	paths := solvePaths(devices)

	fmt.Println()
	fmt.Printf("Paths: %d\n", paths)
}

func solvePaths(devices devices) uint64 {
	var pathsFound uint64

	q := map[string]uint64{"you": 1}
	allTime := map[string]uint64{}

	for len(q) > 0 {
		nextQ := map[string]uint64{}

		for device, count := range q {
			allTime[device]++
			for _, device := range devices[device] {
				if device == "out" {
					pathsFound += count
					fmt.Printf(" - Paths found: %d                     \r", pathsFound)
					continue
				}

				nextQ[device] += count
			}
		}
		q = nextQ
	}
	fmt.Println()

	return pathsFound
}

func parse(input []string) devices {
	devices := devices{}

	for _, connections := range input {
		elements := strings.Split(connections, " ")

		device := elements[0][0 : len(elements[0])-1]
		connectedTo := elements[1:]
		if currentConnections, found := devices[device]; found {
			devices[device] = append(currentConnections, connectedTo...)
		} else {
			devices[device] = connectedTo
		}
	}

	return devices
}
