package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

var (
	seed                string
	enableKeyStretching bool
	passwordCache       map[int]string
)

func init() {
	seed = aoc.GetInput(14)[0]
	enableKeyStretching = aoc.Star == aoc.StarTwo
	passwordCache = map[int]string{}
}

func main() {
	fmt.Println("--- Day 14: One-Time Pad ---")
	fmt.Println()

	result := getPasswords(64)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func getPasswords(n int) int {
	var found int

	for i := 0; ; i++ {
		password := getPassword(i)
		if isMatch, r := hasRepeatedRune(password); isMatch {
			required := strings.Repeat(string(r), 5)
			for j := 1; j < 1001; j++ {
				if strings.Contains(getPassword(i+j), required) {
					found++
					fmt.Printf(" - Found password #%2d %v at %2d \n", found, password, i)

					if found == n {
						return i
					}
				}
			}
		}
	}
}

func getPassword(n int) string {
	stretch := 1
	if enableKeyStretching {
		stretch = 2017
	}

	if _, isCached := passwordCache[n]; !isCached {
		password := fmt.Sprintf("%v%v", seed, n)
		for i := 0; i < stretch; i++ {
			password = getMd5(password)
		}

		passwordCache[n] = password

	}

	return passwordCache[n]
}

func hasRepeatedRune(s string) (bool, rune) {
	var last rune
	var concurrent = 1

	// HACK: s is always hexadecimal.  Adding a non-hex rune to the end ensures we finds repeats at end of string.
	for _, current := range s + "!" {
		switch last == current {
		case true:
			concurrent++

		case false:
			if concurrent >= 3 {
				return true, last
			}

			last = current
			concurrent = 1
		}
	}

	return false, '!'
}

func getMd5(s string) string {
	data := md5.Sum([]byte(s))
	return hex.EncodeToString(data[:])
}
