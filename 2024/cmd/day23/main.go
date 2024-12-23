package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 23: LAN Party ---")
	fmt.Println()

	network := parse(aoc.Input)
	maxClique := []string{}

	subnets := make(map[string]int)
	for node, computers := range network {
		for i := 0; i < len(computers); i++ {
			for j := 0; j < len(computers); j++ {
				if i != j {
					if slices.Contains(network[computers[i]], computers[j]) {
						candidate := []string{node, computers[i], computers[j]}

						if maybeHistorian(candidate...) {
							subnets[newSubnet(candidate...)]++
						}
					}
				}
			}
		}

		clique := getMaximumClique(node, network)
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	fmt.Println()
	fmt.Printf("Candidate Parties: %d\n", len(subnets))
	fmt.Printf("Password: %v\n", strings.Join(maxClique, ","))
}

func newSubnet(computers ...string) string {
	slices.Sort(computers)
	return strings.Join(computers, ",")
}

func maybeHistorian(computers ...string) bool {
	for _, machine := range computers {
		if strings.HasPrefix(machine, "t") {
			return true
		}
	}

	return false
}

func getMaximumClique(from string, network map[string][]string) []string {
	clique := []string{from}
	discarded := []string{}

	for _, machine := range network[from] {
		for _, candidate := range clique {
			if slices.Contains(network[candidate], machine) {
				if !slices.Contains(clique, machine) {
					clique = append(clique, machine)
				}
			} else {
				if !slices.Contains(discarded, machine) {
					discarded = append(discarded, machine)
				}
			}
		}
	}

	result := []string{}
	for _, c := range clique {
		if !slices.Contains(discarded, c) {
			result = append(result, c)
		}
	}

	slices.Sort(result)
	return result
}

func parse(input []string) map[string][]string {
	network := make(map[string][]string)

	for _, line := range input {
		computers := strings.Split(line, "-")
		left := computers[0]
		right := computers[1]

		for _, connection := range [][2]string{{left, right}, {right, left}} {
			if connections, found := network[connection[0]]; found {
				network[connection[0]] = append(connections, connection[1])
				continue
			}
			network[connection[0]] = []string{connection[1]}
		}
	}

	return network
}
