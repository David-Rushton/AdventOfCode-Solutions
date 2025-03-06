package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 23: Safe Cracking ---")
	fmt.Println()

	registers := map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	if aoc.Star == aoc.StarTwo {
		registers["a"] = 12
	}

	getX := func(s string) int {
		switch aoc.IsInt(s) {
		case true:
			return aoc.ToInt(s)
		default:
			return registers[s]
		}
	}

	skipper := newSkipper()
	instructions := aoc.GetInput(23)
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

			case "dec", "tgl":
				elements[cmd] = "inc"

			case "jnz":
				elements[cmd] = "cpy"

			case "cpy":
				elements[cmd] = "jnz"
			}
		}

		if i > 8 {
			fmt.Printf(
				" - Processing instruction (%v): % 3d | %v | %v\n",
				iteration,
				i,
				instruction,
				registers)
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

		case "jnz":
			if xValue := getX(elements[x]); xValue > 0 {
				var offset int
				switch aoc.IsInt(elements[y]) {
				case true:
					offset = aoc.ToInt(elements[y])
				case false:
					offset = registers[elements[y]]
				}

				shortcutAvailable, registerInterval := skipper.add(i, registers)
				if shortcutAvailable {
					if aoc.IsInt(elements[x]) {
						panic("Cannot jump")
					}

					// Find multiplying factor.
					factor := registers[elements[x]] / int(math.Abs(float64(registerInterval[elements[x]])))
					registers["a"] += registerInterval["a"] * factor
					registers["b"] += registerInterval["b"] * factor
					registers["c"] += registerInterval["c"] * factor
					registers["d"] += registerInterval["d"] * factor

					i++
					continue
				}

				i += offset
				continue
			}
		}

		i++
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", registers["a"])
}
