package main

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/iostr"
)

func main() {
	if aoc.Star == aoc.StarOne {
		star1()
		return
	}

	star2()
}

func star1() {
	iostr.Outln("Day 01 *")
	iostr.Outln("--------")

	// parse
	left := []int{}
	right := []int{}

	for _, line := range aoc.Input {
		elements := strings.Split(line, "   ")
		for i, element := range elements {
			value, _ := strconv.ParseInt(element, 10, 64)
			if i == 0 {
				left = append(left, int(value))
			} else {
				right = append(right, int(value))
			}
		}
	}

	// sort
	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		log.Fatalf("Left and right lists are not equal.\n\tLeft: %v.\n\tRight: %v.\n", left, right)
	}

	// calculate answer
	total := 0
	for i := 0; i < len(left); i++ {
		subTotal := int(left[i] - right[i])
		if subTotal < 0 {
			subTotal *= -1
		}

		total += subTotal
	}

	iostr.Outln("--------")
	iostr.Outf("Diff: %d\n", total)
}

func star2() {
	iostr.Outln("Day 01 **")
	iostr.Outln("---------")

	// parse
	left := []int64{}
	right := make(map[int64]int64)

	for _, line := range aoc.Input {
		elements := strings.Split(line, "   ")
		for i, element := range elements {
			value, _ := strconv.ParseInt(element, 10, 64)

			if i == 0 {
				left = append(left, value)
			}

			if i != 0 {
				right[value]++
			}
		}
	}

	// calculate
	total := int64(0)

	for _, multiplicand := range left {
		multiplier, _ := right[multiplicand]

		subTotal := multiplicand * multiplier
		total += subTotal
		iostr.Verbosef("\t%d * %d == %d", multiplicand, multiplier, subTotal)
	}

	iostr.Outln("---------")
	iostr.Outf("Total: %d", total)
}
