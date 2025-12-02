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

func main() {
	fmt.Println("--- Day 2: Gift Shop ---")
	fmt.Println()

	result := 0
	idRanges := parse(aoc.InputRaw)

	for _, currentRange := range idRanges {
		current := currentRange.from

		for current.num <= currentRange.until.num {
			if !current.isValid() {
				fmt.Printf(
					" - invalid id %s in range %s - %s\n",
					current.txt,
					currentRange.from.txt,
					currentRange.until.txt)

				result += current.num
			}

			current = current.next()
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
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
