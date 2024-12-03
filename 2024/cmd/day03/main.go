package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/iostr"
)

const (
	starOnePattern = `mul\(\d{1,3},\d{1,3}\)`
	starTwoPattern = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
)

func main() {
	iostr.Outln("--- Day 3: Mull It Over ---")
	iostr.Outln("---------------------------")

	pattern := starOnePattern
	if aoc.Star == aoc.StarTwo {
		pattern = starTwoPattern
	}
	var re = regexp.MustCompile(pattern)

	var total int64
	enabled := true

	for _, memory := range aoc.Input {
		for _, match := range re.FindAllString(memory, -1) {
			if match == "do()" {
				iostr.Verbose("enabling instructions\n")
				enabled = true
			}

			if match == "don't()" {
				iostr.Verbose("disabling instructions\n")
				enabled = false
			}

			if enabled && match[:3] == "mul" {
				iostr.Verbosef("found match: %s.", match)
				left, right := parse(match)
				total += left * right
			}
		}
	}

	iostr.Outln("---------------------------")
	iostr.Outf("--- total: %d ---", total)
}

func parse(instruction string) (left, right int64) {
	numbers := strings.Split(instruction[4:len(instruction)-1], ",")
	for i, number := range numbers {
		result, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.Fatalf("Cannot convert %s to an int.", number)
		}

		if i == 0 {
			left = result
		} else {
			right = result
		}
	}

	return
}
