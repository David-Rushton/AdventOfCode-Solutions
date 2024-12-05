/*
Provides the Advent of Code starting state.
Which star, mode and input?
*/
package aoc

import (
	"fmt"
	"os"
	"slices"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/iostr"
)

var (
	state       AocState
	Day         int
	Star        AocStar
	VerboseMode bool
	TestMode    bool
	InputPath   string
	Input       []string
	InputRaw    string
)

func init() {
	if slices.Contains(os.Args, "--help") || slices.Contains(os.Args, "-h") {
		ShowUsage()
	}

	state = getState()
	Day = state.Day
	Star = state.Star
	VerboseMode = state.VerboseMode
	TestMode = state.TestMode
	InputPath = state.InputPath
	Input = state.Input
	InputRaw = state.InputRaw

	if VerboseMode {
		iostr.Verboseln(state)
	}
}

// Prints Advent of Code state.
// Useful for debugging only.
func PrintState() {
	fmt.Println(state)
}

// Prints usage information and exits, normally.
func ShowUsage() {
	fmt.Println(`AoC usage:

  -t|--test            Reads the test input
  -2|--star-two        Runs the star 2 solver
  -v|--verbose         Enables additional logging`)

	os.Exit(0)
}
