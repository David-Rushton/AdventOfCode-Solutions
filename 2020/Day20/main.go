package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	testMode bool
)

func init() {
	if slices.Contains(os.Args, "-t") {
		testMode = true
	}
}

func main() {
	fmt.Println("--- Day 20: Jurassic Jigsaw ---")
	fmt.Println()

	tiles := parse(getInput())
	tiles = linkTiles(tiles)

	result := 1
	for _, tile := range tiles {
		if len(tile.connectedTo) == 2 {
			result *= tile.id
			fmt.Printf("  - Corner found: %d\n", tile.id)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func linkTiles(tiles []tile) []tile {
	result := make([]tile, len(tiles))
	copy(result, tiles)

	for _, outer := range result {
		for _, inner := range result {
			if outer.id == inner.id {
				continue
			}

			for _, outerEdge := range outer.getEdgeIds() {
				for _, innerEdge := range inner.getEdgeIds() {
					if outerEdge == innerEdge {
						outer.addConnection(inner.id)
						inner.addConnection(outer.id)
					}
				}
			}
		}
	}

	return result
}

func getInput() []string {
	path := "./input.txt"
	if testMode {
		path = "./input.test.txt"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read file: %v\n.", err)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r", "")

	return strings.Split(content, "\n")
}

func parse(input []string) []tile {
	result := []tile{}

	var next tile
	for _, line := range input {
		if line == "" {
			if next.id != 0 {
				result = append(result, next)
			}
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			// Format `Tile nnnn:`.
			numId, err := strconv.ParseInt(line[5:9], 10, 64)
			if err != nil {
				log.Fatalf("Cannot convert %v to a number.", line[5:9])
			}

			next = tile{
				id:          int(numId),
				cells:       [][]string{},
				connectedTo: map[int]bool{},
			}

			continue
		}

		nextRow := []string{}
		for _, r := range line {
			nextRow = append(nextRow, string(r))
		}
		next.cells = append(next.cells, nextRow)
	}

	return result
}
