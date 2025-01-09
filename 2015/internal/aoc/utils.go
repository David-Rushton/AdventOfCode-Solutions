package aoc

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type star int

const (
	StarOne star = iota
	StarTwo
)

var (
	Star        star = StarOne
	TestMode    bool
	VerboseMode bool
)

func init() {
	if slices.Contains(os.Args, "-2") {
		Star = StarTwo
	}

	if slices.Contains(os.Args, "-t") {
		TestMode = true
	}

	if slices.Contains(os.Args, "-v") {
		VerboseMode = true
	}
}

func GetInput(day int) []string {
	path := fmt.Sprintf("./day%02d/input.txt", day)
	if TestMode {
		path = fmt.Sprintf("./day%02d/input.test.txt", day)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read file %v", path)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r", "")

	return strings.Split(content, "\n")
}

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func ToInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert %s to a number.", s)
	}

	return int(num)
}
