package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- Day 20: Infinite Elves and Infinite Houses ---")
	fmt.Println()

	resultOne := starOne()
	resultTwo := starTwo()

	fmt.Println()
	fmt.Printf("Star One: %d\n", resultOne)
	fmt.Printf("Star Two: %d\n", resultTwo)
}

func starOne() int {
	result := 0
	target := 34_000_000

	for i := 1; ; i++ {
		presents := sumFactors(i)
		if presents >= target {
			result = i
			break
		}

		if i%10000 == 0 {
			fmt.Printf(" - Visited %d houses\n", i)
		}
	}

	return result
}

func sumFactors(n int) int {
	if n == 1 {
		return 1
	}

	result := 1 + n

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == (n / i) {
				result += (i)
			} else {
				result += (i + n/i)
			}
		}
	}

	return result * 10
}

func starTwo() int {
	target := 34_000_000
	houses := make([]int, target)

	for h := 1; h < len(houses); h++ {
		var found int
		for e := h; e < len(houses); e += h {
			houses[e] += h * 11

			found++
			if found >= 50 {
				break
			}
		}

		if h%10000 == 0 {
			fmt.Printf(" - Visited %d houses\n", h)
		}

		if houses[h] >= target {
			return h
		}
	}

	panic("Could not find house")
}
