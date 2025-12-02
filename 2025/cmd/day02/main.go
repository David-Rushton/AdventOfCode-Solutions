package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type idRange struct {
	from  value
	until value
}

type value struct {
	num int
	txt string
}

func (v value) next() value {
	return value{
		num: v.num + 1,
		txt: strconv.FormatInt(int64(v.num+1), 10),
	}
}

func (v value) isValid() bool {
	if len(v.txt)%2 != 0 {
		return true
	}

	half := len(v.txt) / 2
	firstHalf := v.txt[:half]
	secondHalf := v.txt[half:]

	return firstHalf != secondHalf
}

func (v value) isValidV2() bool {
	repeats := func(s string) bool {
		if len(v.txt)%len(s) == 0 && len(v.txt)/len(s) > 1 {
			repeated := strings.Repeat(s, len(v.txt)/len(s))
			return v.txt == repeated
		}

		return false
	}

	for i := range len(v.txt) {
		if repeats(v.txt[0 : i+1]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("--- Day 2: Gift Shop ---")
	fmt.Println()

	result1 := 0
	result2 := 0
	idRanges := parse(aoc.InputRaw)

	for _, currentRange := range idRanges {
		current := currentRange.from

		for current.num <= currentRange.until.num {
			if !current.isValid() {
				result1 += current.num
			}

			if !current.isValidV2() {
				fmt.Printf(" - invalid id %s\n", current.txt)

				result2 += current.num
			}

			current = current.next()
		}
	}

	fmt.Println()
	fmt.Printf("Star 1: %d\n", result1)
	fmt.Printf("Star 2: %d\n", result2)
}

func parse(input string) []idRange {
	var ranges []idRange

	for _, current := range strings.Split(input, ",") {
		elements := strings.Split(current, "-")
		ranges = append(ranges, idRange{
			from: value{
				num: toInt(elements[0]),
				txt: elements[0],
			},
			until: value{
				num: toInt(elements[1]),
				txt: elements[1],
			},
		})
	}

	return ranges
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
