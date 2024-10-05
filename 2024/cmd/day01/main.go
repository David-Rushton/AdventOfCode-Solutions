package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func init() {
	if aoc.DebugMode {
		aoc.PrintState()
	}
}

func main() {
	for _, line := range aoc.Input {
		fmt.Println(line)
	}
}
