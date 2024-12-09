package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	testMode bool
	star     int
)

func init() {
	if slices.Contains(os.Args, "-t") {
		testMode = true
	}

	star = 1
	if slices.Contains(os.Args, "-2") {
		star = 2
	}
}

func main() {
	fmt.Println("--- Day 5: Sunny with a Chance of Asteroids ---")
	fmt.Println()

	for _, program := range parse() {

		if !testMode {
			if star == 1 {
				overrideProgram := make([]int64, len(program))
				copy(overrideProgram, program)
				overrideProgram[1] = 12
				overrideProgram[2] = 2

				fmt.Printf("Program: %v = %d \n", overrideProgram, compute(overrideProgram))
				return
			}

			for left := int64(1); left < 99; left++ {
				for right := int64(1); right < 99; right++ {
					overrideProgram := make([]int64, len(program))
					copy(overrideProgram, program)
					overrideProgram[1] = left
					overrideProgram[2] = right

					fmt.Printf("Program: %v...\n", overrideProgram[0:10])

					if compute(overrideProgram) == 19690720 {
						fmt.Printf("%d%d\n", left, right)
						return
					}
				}

			}
		}

	}
}

func compute(program []int64) int64 {
	idx := 0

	for {
		switch program[idx] {
		case 1:
			left := program[program[idx+1]]
			right := program[program[idx+2]]
			program[program[idx+3]] = left + right

		case 2:
			left := program[program[idx+1]]
			right := program[program[idx+2]]
			program[program[idx+3]] = left * right

		case 99:
			return program[0]

		default:
			log.Fatalf("Invalid opcode %d.", program[idx])
		}

		idx += 4
		if !(idx < len(program)) {
			break
		}
	}

	// ut oh
	log.Fatalf("This is not my beautiful house")
	return -1
}

func parse() [][]int64 {
	path := "input.txt"
	if testMode {
		path = "input.test.txt"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read file because %v.", err)
	}

	result := [][]int64{}

	for _, line := range strings.Split(string(data), "\r\n") {
		if line != "" {
			program := []int64{}

			for _, item := range strings.Split(line, ",") {
				if item != "" {
					i, err := strconv.ParseInt(item, 10, 64)
					if err != nil {
						log.Fatalf("Cannot parse int from value -%v-.", item)
					}

					program = append(program, i)
				}
			}

			result = append(result, program)
		}
	}

	return result
}
