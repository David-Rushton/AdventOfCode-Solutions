package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 12: Leonardo's Monorail ---")
	fmt.Println()

	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	if aoc.Star == aoc.StarTwo {
		registers["c"] = 1
	}

	getX := func(s string) int {
		switch aoc.IsInt(s) {
		case true:
			return aoc.ToInt(s)
		default:
			return registers[s]
		}
	}

	instructions := aoc.GetInput(12)
	for i := 0; i < len(instructions); {
		const cmd int = 0
		const x int = 1
		const y int = 2
		instruction := instructions[i]
		elements := strings.Split(instruction, " ")

		fmt.Printf(" - Processing instruction: % 3d | %v\n", i, instruction)

		switch elements[cmd] {
		case "cpy":
			registers[elements[y]] = getX(elements[x])

		case "inc":
			registers[elements[x]]++

		case "dec":
			registers[elements[x]]--

		case "jnz":
			if xValue := getX(elements[x]); xValue > 0 {
				if yValue := aoc.ToInt(elements[y]); yValue == -2 {
					incReg := strings.Split(instructions[i-2], " ")[1]
					decReg := strings.Split(instructions[i-1], " ")[1]
					registers[incReg] += registers[decReg]
					registers[decReg] = 0
					i++
					continue
				}

				i += aoc.ToInt(elements[y])
				continue
			}
		}

		i++
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", registers["a"])
}
