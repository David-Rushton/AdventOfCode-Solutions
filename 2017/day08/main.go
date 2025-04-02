package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

type operations int
type operators int

const (
	increment operations = iota
	decrement
)

const (
	greaterThan operators = iota
	greaterThanOrEqualTo
	lessThan
	lessThanOrEqualTo
	equalTo
	notEqualTo
)

type instruction struct {
	registerName string
	operation    operations
	value        int
	expression   expression
}

type expression struct {
	registerName string
	operator     operators
	value        int
}

func main() {
	fmt.Println("--- Day 8: I Heard You Like Registers ---")
	fmt.Println()

	var instructions = parse(aoc.GetInput(8))
	var result, memoryRequired = runProgram(instructions)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Memory Required: %d\n", memoryRequired)
}

func runProgram(instructions []instruction) (result, memoryRequired int) {
	var registers = map[string]int{}

	var evaluateExpression = func(e expression) bool {
		value := registers[e.registerName]

		switch e.operator {
		case greaterThan:
			return value > e.value
		case greaterThanOrEqualTo:
			return value >= e.value
		case lessThan:
			return value < e.value
		case lessThanOrEqualTo:
			return value <= e.value
		case equalTo:
			return value == e.value
		case notEqualTo:
			return value != e.value
		}

		panic("Cannot evaluate expression")
	}

	// Run program.
	for _, instruction := range instructions {
		if evaluateExpression(instruction.expression) {
			switch instruction.operation {
			case increment:
				registers[instruction.registerName] += instruction.value
			case decrement:
				registers[instruction.registerName] -= instruction.value
			}

			if registers[instruction.registerName] > memoryRequired {
				memoryRequired = registers[instruction.registerName]
			}
		}
	}

	// Get result.
	for _, value := range registers {
		if value > result {
			result = value
		}
	}

	return result, memoryRequired
}

func parse(input []string) []instruction {
	var result []instruction

	var toOperation = func(s string) operations {
		switch s {
		case "inc":
			return increment
		case "dec":
			return decrement
		default:
			panic(fmt.Sprintf("Operation not supported: %v", s))
		}
	}

	var toOperator = func(s string) operators {
		switch s {
		case ">":
			return greaterThan
		case "<":
			return lessThan
		case ">=":
			return greaterThanOrEqualTo
		case "<=":
			return lessThanOrEqualTo
		case "==":
			return equalTo
		case "!=":
			return notEqualTo
		default:
			panic(fmt.Sprintf("Operator not supported: %v", s))
		}
	}

	for i := range input {
		elements := strings.Split(input[i], " ")

		result = append(result, instruction{
			registerName: elements[0],
			operation:    toOperation(elements[1]),
			value:        aoc.ToInt(elements[2]),
			expression: expression{
				registerName: elements[4],
				operator:     toOperator(elements[5]),
				value:        aoc.ToInt(elements[6])}})
	}

	return result
}
