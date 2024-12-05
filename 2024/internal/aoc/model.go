package aoc

import "fmt"

type AocStar int

const (
	StarOne AocStar = iota + 1
	StarTwo
)

type AocState struct {
	Day         int
	Star        AocStar
	VerboseMode bool
	TestMode    bool
	InputPath   string
	Input       []string
	InputRaw    string
}

func (ae AocState) String() string {
	return fmt.Sprintf(
		"AoC = { Day = %d, Star = %d, Test Mode = %t, Debug Mode = %t, Input File = %v }",
		ae.Day,
		ae.Star,
		ae.TestMode,
		ae.VerboseMode,
		ae.InputPath)
}
