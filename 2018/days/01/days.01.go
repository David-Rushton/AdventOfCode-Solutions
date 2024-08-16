package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "./input.txt"
	if os.Args[1] == "test" {
		path = "./input.test.txt"
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	elements := strings.Split(string(content), "\n")
	observed := make(map[int]int)

	total := 0
	for {
		for _, value := range elements {
			if len(value) > 1 {
				symbol := value[:1]
				frequency, _ := strconv.Atoi(strings.Trim(value[1:], "\r\n"))

				if symbol == "+" {
					total += frequency
				}

				if symbol == "-" {
					total -= frequency
				}

				if observed[total] > 0 {
					fmt.Printf("\n\nDuplicate frequency found: %v\n", total)
					return
				}

				observed[total] = 1

				fmt.Printf("\nFrequency %v %v", symbol, frequency)
			}
		}
	}

	fmt.Printf("\n\nFrequency: %v\n", total)
}
