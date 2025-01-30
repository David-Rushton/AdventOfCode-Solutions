package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- Day 20: Infinite Elves and Infinite Houses ---")
	fmt.Println()

	target := 34_000_000
	result := 0

	for i := 1; ; i++ {
		presents := sumFactors(i)
		if presents >= target {
			result = i
			break
		}

		if i%1000 == 0 {
			fmt.Printf(" - Visited %d houses\n", i)
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func sumFactors(n int) int {
	var result int

	if n == 1 {
		return 1
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == (n / i) {
				result += (i) * 10
			} else {
				result += (i + n/i) * 10

			}
		}
	}

	return result + 10 + (n * 10)
}
