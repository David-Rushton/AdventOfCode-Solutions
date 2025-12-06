package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

func main() {
	fmt.Println("--- Day 6: Trash Compactor ---")
	fmt.Println()

	var total int

	switch aoc.Star {
	case aoc.StarOne:
		total = starOne()
	default:
		total = starTwo()
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", total)
}

func starOne() int {
	var problems [][]string

	for i, line := range aoc.Input {
		elements := split(line, " ")
		for j, element := range elements {
			if i == 0 {
				problems = append(problems, []string{})
			}

			problems[j] = append(problems[j], element)
		}
	}

	var total int

	for _, problem := range problems {
		var subTotal int
		op := problem[len(problem)-1]
		for _, element := range problem[:len(problem)-1] {
			switch op {
			case "+":
				subTotal += toInt(element)
			case "*":
				if subTotal == 0 {
					subTotal = 1
				}
				subTotal *= toInt(element)
			default:
				panic("unknown operation")
			}
		}

		fmt.Printf(" - Subtotal: %d\n", subTotal)
		total += subTotal
	}

	return total
}

func starTwo() int {
	var total int
	input := strings.Split(strings.ReplaceAll(aoc.InputRaw, "\r", ""), "\n")
	rows := len(input)
	columns := len(input[0])
	grid := input

	var numbers []int
	for c := columns - 1; c >= 0; c-- {
		var numberBuilder string
		var op rune

		for r := range rows {
			line := grid[r]
			if len(line) < columns {
				line += strings.Repeat(" ", columns-len(line))
			}

			switch line[c] {
			case ' ':
				// no-op.
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				numberBuilder += string(line[c])
			case '+', '*':
				op = rune(line[c])
			default:
				panic("unknown character")
			}
		}

		numberBuilder = strings.Trim(numberBuilder, " ")
		if len(numberBuilder) > 0 {
			numbers = append(numbers, toInt(numberBuilder))
		}
		numberBuilder = ""

		if op == '+' || op == '*' {
			var subTotal int
			for _, n := range numbers {
				switch op {
				case '+':
					subTotal += n
				case '*':
					if subTotal == 0 {
						subTotal = 1
					}
					subTotal *= n
				}
			}
			total += subTotal
			fmt.Printf(" - Subtotal: %d\n", subTotal)
			numbers = []int{}
			op = ' '
		}
	}

	return total
}

func split(s string, sep string) []string {
	var result []string
	for _, substring := range strings.Split(s, sep) {
		if substring != sep && substring != "" {
			result = append(result, substring)
		}
	}
	return result
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
