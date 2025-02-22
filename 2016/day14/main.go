package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 14: One-Time Pad ---")
	fmt.Println()

	salt := aoc.GetInput(14)[0]
	var i int
	for {
		password := fmt.Sprintf("%v%v", salt, i)
		hash := getMd5(password)

		fmt.Printf(" - Password %v hash == %v\n", password, hash)

		i++
		if i > 10 {
			break
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", -1)
}

func hasTriplet(s string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 2; i < len(s); i++ {

	}

	return false
}

func getMd5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
