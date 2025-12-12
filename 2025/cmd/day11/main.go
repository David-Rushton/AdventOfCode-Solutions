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

	var paths uint64
	if aoc.Star == aoc.StarTwo {
		svrToFft := solvePaths("svr", "fft", 1, devices)
		fftToDac := solvePaths("fft", "dac", svrToFft, devices)
		paths = solvePaths("dac", "out", fftToDac, devices)
	} else {
		paths = solvePaths("you", "out", 1, devices)
	}

	fmt.Println()
	fmt.Printf("Paths: %d\n", paths)
}

func solvePaths(from, to string, count uint64, devices devices) uint64 {
	var pathsFound uint64

	q := map[string]uint64{from: count}
	for len(q) > 0 {
		nextQ := map[string]uint64{}

		for device, count := range q {
			for _, device := range devices[device] {
				if device == to {
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
