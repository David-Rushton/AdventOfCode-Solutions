package main

import (
	"fmt"
	"maps"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type guest struct {
	name       string
	neighbours map[string]int
}

func main() {
	fmt.Println("--- Day 13: Knights of the Dinner Table ---")
	fmt.Println()

	guests := parse(aoc.GetInput(13))

	if aoc.Star == aoc.StarTwo {
		const me string = "me"

		newGuests := map[string]guest{}
		maps.Copy(newGuests, guests)

		newGuests["me"] = guest{name: me, neighbours: map[string]int{}}

		for current := range guests {
			newGuests["me"].neighbours[current] = 0
			newGuests[current].neighbours[me] = 0
		}

		guests = newGuests
	}

	result := getBestRating(guests)

	fmt.Println()
	fmt.Printf("Result: %d", result)
}

func addMe(guests map[string]guest) map[string]guest {
	const me string = "me"

	result := map[string]guest{}
	maps.Copy(result, guests)

	result["me"] = guest{name: me, neighbours: map[string]int{}}
	for current := range guests {
		result["me"].neighbours[current] = 0
		result[current].neighbours[me] = 0
	}

	return result
}

func getBestRating(guests map[string]guest) int {
	type state struct {
		name    string
		start   string
		rating  int
		visited map[string]bool
	}

	queue := []state{}
	for guestName := range guests {
		queue = append(queue, state{
			name:    guestName,
			start:   guestName,
			rating:  0,
			visited: map[string]bool{guestName: true}})
	}

	bestRating := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current.visited) == len(guests) {
			// The table is a circle, we have to close the loop.
			current.rating += guests[current.name].neighbours[current.start]
			current.rating += guests[current.start].neighbours[current.name]

			if current.rating > bestRating {
				bestRating = current.rating
				fmt.Printf("  - New best: %d of %v\n", current.rating, toString(current.visited))
			}

			continue
		}

		for candidate, rating := range guests[current.name].neighbours {
			if !current.visited[candidate] {
				var newVisited = map[string]bool{}
				maps.Copy(newVisited, current.visited)
				newVisited[candidate] = true

				newRating := rating + guests[candidate].neighbours[current.name]

				queue = append(queue, state{
					name:    candidate,
					start:   current.start,
					rating:  current.rating + newRating,
					visited: newVisited})
			}
		}
	}

	return bestRating
}

func toString(visited map[string]bool) string {
	result := []string{}
	for key := range visited {
		result = append(result, key)
	}
	return strings.Join(result, " - ")
}

func parse(input []string) map[string]guest {
	result := map[string]guest{}

	for _, line := range input {
		elements := strings.Split(strings.ReplaceAll(line, ".", ""), " ")

		rating := aoc.ToInt(elements[3])
		if elements[2] == "lose" {
			rating *= -1
		}

		currentGuest, found := result[elements[0]]
		if !found {
			currentGuest = guest{
				name:       elements[0],
				neighbours: map[string]int{}}
		}

		currentGuest.neighbours[elements[10]] = rating
		result[currentGuest.name] = currentGuest
	}

	return result
}
