package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 21: Scrambled Letters and Hash ---")
	fmt.Println()

	password := getPassword(aoc.TestMode)
	instructions := aoc.GetInput(21)
	scrambled := scramble(password, instructions)

	fmt.Println()
	fmt.Printf("Scrambled password: %v\n", scrambled)
}

func scramble(s string, instructions []string) string {
	result := []rune(s)

	for i := range instructions {
		elements := strings.Split(instructions[i], " ")

		switch {
		case strings.HasPrefix(instructions[i], "swap position"):
			result = swapPosition(result, aoc.ToInt(elements[2]), aoc.ToInt(elements[5]))

		case strings.HasPrefix(instructions[i], "swap letter"):
			result = swapLetter(result, []rune(elements[2])[0], []rune(elements[5])[0])

		case strings.HasPrefix(instructions[i], "rotate left"):
			result = rotateLeft(result, aoc.ToInt(elements[2]))

		case strings.HasPrefix(instructions[i], "rotate right"):
			result = rotateRight(result, aoc.ToInt(elements[2]))

		case strings.HasPrefix(instructions[i], "rotate based"):
			result = rotateBased(result, []rune(elements[6])[0])

		case strings.HasPrefix(instructions[i], "reverse positions"):
			result = reversePositions(result, aoc.ToInt(elements[2]), aoc.ToInt(elements[4]))

		case strings.HasPrefix(instructions[i], "move position"):
			result = movePosition(result, aoc.ToInt(elements[2]), aoc.ToInt(elements[5]))
		}

		fmt.Printf(" - %v || %v\n", string(result), instructions[i])
	}

	return string(result)
}

func swapPosition(r []rune, x, y int) []rune {
	t := r[y]
	r[y] = r[x]
	r[x] = t

	return r
}

func swapLetter(r []rune, x, y rune) []rune {
	for i := range r {
		switch r[i] {
		case x:
			r[i] = y
		case y:
			r[i] = x
		}
	}

	return r
}

func rotateLeft(r []rune, n int) []rune {
	for range n {
		r = append(r[1:], r[0])
	}

	return r
}

func rotateRight(r []rune, n int) []rune {
	for range n {
		r = append([]rune{r[len(r)-1]}, r[0:len(r)-1]...)
	}

	return r
}

func rotateBased(r []rune, char rune) []rune {
	idx := slices.Index(r, char)
	if idx >= 4 {
		idx++
	}

	return rotateRight(r, idx+1)
}

func reversePositions(r []rune, x, y int) []rune {
	for x < y {
		r = swapPosition(r, x, y)
		x++
		y--
	}

	return r
}

func movePosition(r []rune, x, y int) []rune {
	var tx rune = r[x]
	r = append(r[0:x], r[x+1:]...)

	if y >= len(r) {
		return append(r[0:y], tx)
	}

	tr := make([]rune, len(r))
	copy(tr, r)
	tr = tr[0:y]
	tr = append(tr, tx)
	tr = append(tr, r[y:]...)

	return tr
}

func getPassword(testMode bool) string {
	switch testMode {
	case true:
		return "abcde"
	default:
		return "abcdefgh"
	}
}
