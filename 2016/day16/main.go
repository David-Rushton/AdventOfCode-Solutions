package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 16: Dragon Checksum ---")
	fmt.Println()

	dataLen := 272
	if aoc.TestMode {
		dataLen = 20
	}

	if aoc.Star == aoc.StarTwo {
		dataLen = 35651584
	}

	data := generateData(aoc.GetInput(16)[0], dataLen)
	checksum := getChecksum(data)

	fmt.Println()
	fmt.Printf("Result: %v\n", checksum)
}

func generateData(seed string, dataLen int) string {
	current := seed
	for {
		if len(current) >= dataLen {
			if len(current) > dataLen {
				current = current[0:dataLen]
			}

			return current
		}

		next := reverse(current)
		next = strings.ReplaceAll(next, "1", "t")
		next = strings.ReplaceAll(next, "0", "1")
		next = strings.ReplaceAll(next, "t", "0")
		current = current + "0" + next

		fmt.Printf(" - Extending data -> %d\n", len(current))
	}
}

func getChecksum(s string) string {
	current := []rune(s)
	for {
		next := []rune(strings.Repeat(" ", len(current)/2))
		iteration := 0
		for i := 0; i < len(current); i += 2 {
			if current[i] == current[i+1] {
				next[iteration] = '1'
			} else {
				next[iteration] = '0'
			}
			iteration++
		}
		current = next

		fmt.Printf(" - Checksum generated: -> %d\n", len(current))

		if len(current)%2 == 1 {
			return string(current)
		}
	}
}

func reverse(s string) string {
	result := []rune(s)

	len := len(s)
	halfLen := len / 2
	for i := 0; i < halfLen; i++ {
		temp := result[i]
		result[i] = result[len-1-i]
		result[len-1-i] = temp

	}

	return string(result)
}
