package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Day 2: Inventory Management System\n")
	fmt.Print("----------------------------------\n\n")

	total_twos := 0
	total_threes := 0

	for _, v := range GetContentLines() {
		twos, threes := getCounts(v)
		fmt.Println("  ", v, twos, threes)

		if twos > 0 {
			total_twos++
		}

		if threes > 0 {
			total_threes++
		}

	}

	fmt.Printf("\n\nResult: %v * %v = %v\n", total_twos, total_threes, total_twos*total_threes)

	for _, o := range GetContentLines() {
		for _, i := range GetContentLines() {
			if score := getDiffScore(o, i); score == 1 {
				fmt.Printf("Common boxes: %v Vs %v = %v\n", o, i, score)
				return
			}
		}
	}
}

func getDiffScore(left, right string) int {
	result := 0
	for i, c := range left {
		if byte(c) != right[i] {
			result++
		}
	}
	return result
}

func getCounts(s string) (twos, threes int) {
	twos = 0
	threes = 0

	characterCounts := make(map[rune]int)
	for _, c := range s {
		characterCounts[c]++
	}

	for _, v := range characterCounts {
		if v == 2 {
			twos++
		}

		if v == 3 {
			threes++
		}
	}

	return twos, threes
}

func getContent() string {
	file := "live"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}

	content, err := os.ReadFile(fmt.Sprintf("./input.%v.txt", file))
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func GetContentLines() []string {
	var result []string
	for _, v := range strings.Split(getContent(), "\n") {
		if temp := strings.Trim(v, "\r\n"); len(temp) > 0 {
			result = append(result, strings.Trim(v, "\r\n"))
		}
	}
	return result
}
