package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 20: Firewall Rules ---")
	fmt.Println()

	var blockedRanges = parse(aoc.GetInput(20))
	blockedRanges = compactIpRanges(blockedRanges)

	// find lowest allowed ip.
	var minAllowed int
	for _, blocked := range blockedRanges {
		if minAllowed >= blocked.from {
			if minAllowed < blocked.until {
				minAllowed = blocked.until + 1
			}
		}
	}

	// find all open addresses.
	var openAddresses = 10 //math.MaxInt32
	for _, blocked := range blockedRanges {
		openAddresses -= blocked.count()
	}

	fmt.Println()
	fmt.Printf("Lowest allowed IP: %d\n", minAllowed)
	fmt.Printf("Count of allowed IPs: %d\n", openAddresses)
}

func parse(input []string) []ipRange {
	result := []ipRange{}

	for i := range input {
		elements := strings.Split(input[i], "-")
		result = append(result, ipRange{
			from:  aoc.ToInt(elements[0]),
			until: aoc.ToInt(elements[1])})
	}

	return result
}

func compactIpRanges(ranges []ipRange) []ipRange {
	// Sort.
	slices.SortFunc(ranges, func(a, b ipRange) int {
		return a.from - b.from
	})

	// Compact
	var result = []ipRange{}
	var from, until int
	for i, blocked := range ranges {
		extendsCurrent := i == 0 || (blocked.from >= ranges[i-1].from && blocked.from <= ranges[i-1].until)
		isLast := i == len(ranges)-1

		// extend current.
		if extendsCurrent {
			if blocked.until > until {
				until = blocked.until
			}

			if isLast {
				result = append(result, ipRange{from, until})
			}

			continue
		}

		// extend existing.
		result = append(result, ipRange{from, until})

		// start new range.
		from = blocked.from
		until = blocked.until

		// commit what we have.
		if isLast {
			result = append(result, ipRange{from, until})
		}
	}

	return result
}
