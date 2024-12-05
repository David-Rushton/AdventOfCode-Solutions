package main

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 5: Print Queue ---")
	fmt.Println()

	var starOneTotal int
	var starTwoTotal int

	rulesMap, pageUpdates := parse(aoc.InputRaw)
	for _, pageUpdate := range pageUpdates {
		fmt.Printf("pageUpdate: %v.\n", pageUpdate)

		if subTotal, failedAt := processUpdate(pageUpdate, rulesMap); failedAt == 0 {
			starOneTotal += subTotal
		} else {
			sortedUpdate := sortUpdate(pageUpdate, rulesMap)
			if subTotal, failedAt := processUpdate(sortedUpdate, rulesMap); failedAt == 0 {
				starTwoTotal += subTotal
				continue
			}

			panic("Cannot find middle number :(")
		}
	}

	fmt.Println()
	fmt.Printf("Total Star One: %v\n", starOneTotal)
	fmt.Printf("Total Star Two: %v\n", starTwoTotal)
}

func sortUpdate(pageUpdate []int, rulesMap map[int]big.Int) []int {
	slices.SortFunc(pageUpdate, func(a, b int) int {
		r := rulesMap[b]

		if r.Bit(a) == 1 {
			return -1
		}

		return 1
	})

	return pageUpdate
}

func processUpdate(pageUpdate []int, rulesMap map[int]big.Int) (value, failedAt int) {
	var localRules big.Int

	for i, num := range pageUpdate {
		if rule, found := rulesMap[num]; found {
			localRules.Or(&localRules, &rule)
		}

		if localRules.Bit(num) == 1 {
			fmt.Printf("\tupdate failed: %v.  Rule violation: %v.\n", pageUpdate, num)
			return 0, i
		}
	}

	if len(pageUpdate)%2 == 0 {
		panic("what is the middle number here?")
	}

	middleNumber := len(pageUpdate) / 2
	return pageUpdate[middleNumber], 0
}

type rule struct {
	before int
	after  int
}

func parse(input string) (rulesMap map[int]big.Int, pageUpdates [][]int) {
	rawRules := []rule{}
	pageUpdatesResult := [][]int{}

	// Process raw input
	var pageNumberMode bool
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			pageNumberMode = true
			continue
		}

		if !pageNumberMode {
			elements := strings.Split(line, "|")
			left, _ := strconv.ParseInt(elements[0], 10, 64)
			right, _ := strconv.ParseInt(elements[1], 10, 64)
			newRule := rule{int(left), int(right)}
			rawRules = append(rawRules, newRule)
		}

		if pageNumberMode {
			pageNumbers := []int{}
			for _, item := range strings.Split(line, ",") {
				num, _ := strconv.ParseInt(item, 10, 64)
				pageNumbers = append(pageNumbers, int(num))
			}
			pageUpdatesResult = append(pageUpdatesResult, pageNumbers)
		}
	}

	// Convert rules to return types.
	ruleResult := make(map[int]big.Int)
	for _, r := range rawRules {
		item, ok := ruleResult[r.after]
		if !ok {
			item = big.Int{}
		}

		item.SetBit(&item, r.before, 1)
		ruleResult[r.after] = item
	}

	return ruleResult, pageUpdatesResult
}
