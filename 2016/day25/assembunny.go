package main

import (
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func runProgram(seed int, instructions []string) []int {

	registers := map[string]int{"a": seed, "b": 0, "c": 0, "d": 0}
	result := []int{}

	getX := func(s string) int {
		switch aoc.IsInt(s) {
		case true:
			return aoc.ToInt(s)
		default:
			return registers[s]
		}
	}

	toggled := map[int]bool{}
	iteration := 0
	for i := 0; i < len(instructions); {
		const cmd int = 0
		const x int = 1
		const y int = 2

		iteration++
		instruction := instructions[i]
		elements := strings.Split(instruction, " ")

		if toggled[i] {
			switch elements[cmd] {
			case "inc":
				elements[cmd] = "dec"

			case "dec", "tgl", "out":
				elements[cmd] = "inc"

			case "jnz":
				elements[cmd] = "cpy"

			case "cpy":
				elements[cmd] = "jnz"
			}
		}

		switch elements[cmd] {
		case "tgl":
			offset := getX(elements[x])
			toggledValue := toggled[i+offset]
			toggled[i+offset] = !toggledValue

		case "cpy":
			registers[elements[y]] = getX(elements[x])

		case "inc":
			registers[elements[x]]++

		case "dec":
			registers[elements[x]]--

		case "out":
			xValue := getX(elements[x])
			result = append(result, xValue)
			if len(result) >= 10 {
				return result
			}

		case "jnz":
			if xValue := getX(elements[x]); xValue > 0 {
				var offset int
				switch aoc.IsInt(elements[y]) {
				case true:
					offset = aoc.ToInt(elements[y])
				case false:
					offset = registers[elements[y]]
				}

				i += offset
				continue
			}
		}

		i++
	}

	panic("Program halted unexpectedly")
}
