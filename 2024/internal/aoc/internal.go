package aoc

import (
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

// Returns the Advent of Code starting state.
func getState() AocState {
	// Day
	day := day()

	// Star one, unless stated.
	star := StarOne
	if slices.Contains(os.Args, "--star-two") || slices.Contains(os.Args, "-2") {
		star = StarTwo
	}

	// Test Mode.
	testMode := slices.Contains(os.Args, "--test") || slices.Contains(os.Args, "-t")

	// Input path.
	inputPath := path.Join(".", "cmd", fmt.Sprintf("day%02d", day))
	if testMode {
		inputPath = path.Join(inputPath, "input.test.txt")
	} else {
		inputPath = path.Join(inputPath, "input.txt")
	}

	// Everyting we need to get started.
	return AocState{
		Day:         day,
		Star:        star,
		VerboseMode: slices.Contains(os.Args, "--verbose") || slices.Contains(os.Args, "-v"),
		TestMode:    testMode,
		InputPath:   inputPath,
		Input:       input(inputPath),
	}
}

// Returns the input file.
// Content is split by line endings, and returned row by row.
func input(path string) []string {
	// Get the content
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read file %v.  Because %v.", path, err)
	}

	content := strings.Trim(string(data), " \r\n\t")

	return strings.Split(strings.Replace(content, "\r", "", -1), "\n")
}

func day() int {
	dayText := os.Getenv("aoc_day")

	if dayText == "" {
		// HACK: We read the current day from the executable.
		// Where we assume the file will be called dayXX.exe.
		exePath, err := os.Executable()
		if err != nil {
			log.Fatalf("Cannot read current day from exe path %v.  Because %v.", exePath, err)
		}

		exeInfo, err := os.Stat(exePath)
		if err != nil {
			log.Fatalf("Cannot read current day from exe path %v.  Because %v.", exePath, err)
		}

		dayText = strings.Replace(strings.Replace(exeInfo.Name(), "day", "", -1), ".exe", "", -1)
	}

	day, err := strconv.ParseInt(dayText, 10, 64)
	if err != nil {
		log.Fatalf("Cannot read current day from %v.  Because %v.", dayText, err)
	}

	if day < 0 || day > 25 {
		log.Fatalf("Current day %v should be between 1 and 25, inclusive.", day)

	}

	return int(day)
}
