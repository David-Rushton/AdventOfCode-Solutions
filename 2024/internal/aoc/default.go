/*
Provides the Advent of Code starting state.
Which star, mode and input?
*/
package aoc

import "fmt"

var (
	state     AocState
	Day       int
	Star      AocStar
	DebugMode bool
	TestMode  bool
	InputPath string
	Input     []string
)

func init() {
	state = getState()
	Day = state.Day
	Star = state.Star
	DebugMode = state.DebugMode
	TestMode = state.TestMode
	InputPath = state.InputPath
	Input = state.Input
}

// Prints Advent of Code state.
// Useful for debugging only.
func PrintState() {
	fmt.Println(state)
}
