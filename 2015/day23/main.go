package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type instruction struct {
	code     string
	register string
	value    int
}

func main() {
	fmt.Println("--- Day 23: Opening the Turing Lock ---")
	fmt.Println()

	var initialValue int
	if aoc.Star == aoc.StarTwo {
		initialValue = 1
	}

	instructions := parse(aoc.GetInput(23))
	result := run(instructions, initialValue)

	fmt.Println()
	fmt.Printf("Result: %d", result)
}

func run(instructions []instruction, initialValue int) int {
	registers := map[string]int{"a": initialValue, "b": 0}

	var index int
	for {
		if index >= len(instructions) {
			break
		}

		instruction := instructions[index]

		switch instruction.code {
		// half
		case "hlf":
			registers[instruction.register] /= 2

		// triple
		case "tpl":
			registers[instruction.register] *= 3

		// increment
		case "inc":
			registers[instruction.register]++

		// jump
		case "jmp":
			index += instruction.value
			continue

		// jump if even
		case "jie":
			if registers[instruction.register]%2 == 0 {
				index += instruction.value
				continue
			}

		// jump if one
		case "jio":
			if registers[instruction.register] == 1 {
				index += instruction.value
				continue
			}
		}

		index++
	}

	for k, v := range registers {
		fmt.Printf(" - Register: %v == %d\n", k, v)
	}

	return registers["b"]
}

func parse(input []string) []instruction {
	result := []instruction{}

	for _, current := range input {
		current = strings.ReplaceAll(current, ",", "")
		elements := strings.Split(current, " ")

		if len(elements) == 3 {
			result = append(result, instruction{
				code:     elements[0],
				register: elements[1],
				value:    aoc.ToInt(elements[2])})
		}

		if len(elements) == 2 {
			if elements[1] == "a" || elements[1] == "b" {
				result = append(result, instruction{
					code:     elements[0],
					register: elements[1]})
			} else {
				result = append(result, instruction{
					code:  elements[0],
					value: aoc.ToInt(elements[1])})
			}
		}
	}

	return result
}
