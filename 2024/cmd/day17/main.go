package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

const (
	registerA = 0
	registerB = 1
	registerC = 2
)

func main() {
	fmt.Println("Day 17: Chronospatial Computer")
	fmt.Println()

	stdOut := []int{}
	registers, program := parse(aoc.Input)

	if aoc.Star == aoc.StarOne {
		registers, stdOut = runProgram(registers, program)

		fmt.Printf("  Standard Out: %s\n", toString(stdOut))
		fmt.Println()
	}

	if aoc.Star == aoc.StarTwo {
		type state struct {
			result int
			level  int
		}

		queue := []state{{0, 0}}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			for j := 0; j < 8; j++ {
				candidate := current.result + j
				registers[registerA] = candidate
				registers[registerB] = 0
				registers[registerC] = 0
				registers, stdOut = runProgram(registers, program)

				if slices.Equal(stdOut, program) {
					fmt.Printf("\nResult: %d\n", candidate)
					os.Exit(0)
				}

				if stdOut[0] == program[len(program)-current.level-1] {
					if len(stdOut) < len(program) {
						queue = append(queue, state{candidate * 8, current.level + 1})
					}
					fmt.Printf("  -----------~> % 2d %d\n", current.level, candidate)
				}
			}
		}
	}

	fmt.Printf("Register A: %d\n", registers[registerA])
	fmt.Printf("Register B: %d\n", registers[registerB])
	fmt.Printf("Register C: %d\n", registers[registerC])
}

func runProgram(registersIn, program []int) (registers, stdOut []int) {
	registers = registersIn
	stdOut = []int{}

	getComboOperand := func(value int) int {
		if value >= 0 && value <= 3 {
			return value
		}

		if value >= 4 && value <= 6 {
			return registers[value-4]
		}

		panic(fmt.Sprintf("Combo operand %d not supported.", value))
	}

	var i int
	for {
		opCode := program[i]
		operand := program[i+1]

		switch opCode {
		// adv
		case 0:
			numerator := registers[registerA]
			denominator := math.Pow(2, float64(getComboOperand(operand)))
			result := numerator / int(denominator)
			registers[registerA] = result
		// bxl
		case 1:
			registers[registerB] = registers[registerB] ^ operand
		// bst
		case 2:
			registers[registerB] = getComboOperand(operand) % 8
		// jnz
		case 3:
			if registers[registerA] != 0 {
				i = operand
				continue
			}
		// bxc
		case 4:
			registers[registerB] = registers[registerB] ^ registers[registerC]
		// out
		case 5:
			stdOut = append(stdOut, getComboOperand(operand)%8)
		// bdv
		case 6:
			numerator := registers[registerA]
			denominator := math.Pow(2, float64(getComboOperand(operand)))
			result := numerator / int(denominator)
			registers[registerB] = result
		// cdv
		case 7:
			numerator := registers[registerA]
			denominator := math.Pow(2, float64(getComboOperand(operand)))
			result := numerator / int(denominator)
			registers[registerC] = result
		default:
			log.Fatalf("OpCode %d not supported.", opCode)
		}

		i += 2
		if i >= len(program) {
			break
		}
	}

	return registers, stdOut
}

func parse(input []string) (registers, program []int) {
	registers = []int{0, 0, 0}
	program = []int{}

	for _, line := range input {
		elements := strings.Split(line, ":")

		switch elements[0] {
		case "Register A":
			registers[registerA] = toInt(strings.Trim(elements[1], " "))
		case "Register B":
			registers[registerB] = toInt(strings.Trim(elements[1], " "))
		case "Register C":
			registers[registerC] = toInt(strings.Trim(elements[1], " "))
		case "Program":
			numbers := strings.Split(strings.Trim(elements[1], " "), ",")
			for _, number := range numbers {
				program = append(program, toInt(number))
			}
		}
	}

	return registers, program
}

func toInt(s string) int {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert %s to a number", s)
	}
	return int(result)
}

func toString(numbers []int) string {
	values := []string{}
	for _, number := range numbers {
		values = append(values, strconv.FormatInt(int64(number), 10))
	}
	return strings.Join(values, ",")
}
