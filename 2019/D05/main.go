package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Day 5: Sunny with a Chance of Asteroids ---")
	fmt.Println()

	data, err := os.ReadFile("./input.test.txt")
	if err != nil {
		log.Fatalf("Cannot read file because %v.", err)
	}

	for _, line := range strings.Split(string(data), "\r\n") {
		result := compute(parse(line))
		fmt.Println(result)
	}

}

func parse(input string) []int64 {
	result := []int64{}

	for _, item := range strings.Split(input, ",") {
		i, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse int from value -%v-.", item)
		}

		result = append(result, i)
	}

	return result
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
