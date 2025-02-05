package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 5: How About a Nice Game of Chess? ---")
	fmt.Println()

	doorId := aoc.GetInput(5)[0]

	var password1 string
	var password2Len int
	password2 := [8]string{"_", "_", "_", "_", "_", "_", "_", "_"}
	for i := int64(0); len(password1) < 8 || password2Len < 8; i++ {
		candidate := toMD5(doorId + strconv.FormatInt(i, 10))

		if strings.HasPrefix(candidate, "00000") {
			// Password 1.
			if len(password1) < 8 {
				password1 += candidate[5:6]
				fmt.Printf(
					" - Hash: %v at index %d extends password 1: %v\n",
					candidate,
					i,
					password1)
			}

			// Password 2.
			if password2Len < 8 {
				idx, err := strconv.ParseInt(candidate[5:6], 10, 64)
				if err == nil {
					if idx < 8 {
						if password2[idx] == "_" {
							password2Len++
							password2[idx] = candidate[6:7]
							fmt.Printf(
								" - Hash: %v at index %d extends password 2: %v\n",
								candidate,
								i,
								strings.Join(password2[:], ""))
						}
					}
				}
			}
		}
	}

	fmt.Println()
	fmt.Printf("Password One: %v\n", password1)
	fmt.Printf("Password Two: %v\n", strings.Join(password2[:], ""))
}

func toMD5(value string) string {
	data := md5.Sum([]byte(value))
	return hex.EncodeToString(data[:])
}
