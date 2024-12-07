package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type testCase struct {
	expected int64
	inputs   []int64
}

func main() {
	fmt.Println("--- Day 7: Bridge Repair ---")
	fmt.Println("")

	var total int64
	testCases := parse(aoc.InputRaw)

	for _, testCase := range testCases {

		fmt.Printf("Test Case: %d: %v == ", testCase.expected, testCase.inputs)

		if subtotal, _ := evaluate(testCase.expected, testCase.inputs[1:], testCase.inputs[0]); subtotal > 0 {
			fmt.Print(subtotal)
			total += testCase.expected
		}

		fmt.Println()
	}

	fmt.Println("")
	fmt.Printf("Result: %d.\n", total)
}

func evaluate(expected int64, testCases []int64, subTotal int64) (int64, bool) {
	for _, op := range []rune{'+', '*', '|'} {
		if op == '|' && aoc.Star == aoc.StarOne {
			continue
		}

		iterationTotal := subTotal
		switch op {
		case '+':
			iterationTotal += testCases[0]
		case '*':
			iterationTotal *= testCases[0]
		case '|':
			iterationTotal = numConcat(iterationTotal, testCases[0])
		}

		if iterationTotal == expected && len(testCases) == 1 {
			return iterationTotal, true
		}

		if len(testCases) > 1 {
			if result, found := evaluate(expected, testCases[1:], iterationTotal); found {
				return result, true
			}
		}
	}

	return 0, false
}

func numConcat(left, right int64) int64 {
	candidate, err := strconv.ParseInt(fmt.Sprintf("%d%d", left, right), 10, 64)
	if err != nil {
		log.Fatalf("Cannot concat %d and %d.  Because %v.", left, right, err)
	}
	return candidate
}

func parse(input string) []testCase {
	result := []testCase{}

	for _, line := range strings.Split(input, "\n") {
		elements := strings.Split(line, ":")

		numbers := []int64{}
		for _, num := range append([]string{elements[0]}, strings.Split(elements[1], " ")...) {
			num = strings.Trim(num, " ")
			if num != "" {
				i, err := strconv.ParseInt(num, 10, 64)
				if err != nil {
					log.Fatalf("Not a number: %v.\n", num)
				}
				numbers = append(numbers, i)
			}
		}

		result = append(result, testCase{
			expected: numbers[0],
			inputs:   numbers[1:],
		})
	}

	return result
}
