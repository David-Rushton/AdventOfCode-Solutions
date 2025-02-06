package main

import (
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 7: Internet Protocol Version 7 ---")
	fmt.Println()

	ip7Addresses := parse(aoc.GetInput(7))
	var tlsSupportedCount int
	var sslSupportedCount int

	for _, address := range ip7Addresses {
		if address.supportsTLS() {
			tlsSupportedCount++
			fmt.Printf(" - IP7 address %v supports TLS\n", address)
		}

		if address.supportsSSL() {
			sslSupportedCount++
			fmt.Printf(" - IP7 address %v supports SSL\n", address)
		}
	}

	fmt.Println()
	fmt.Printf("Addresses that support TLS: %d\n", tlsSupportedCount)
	fmt.Printf("Addresses that support SSL: %d\n", sslSupportedCount)
}

func parse(input []string) []ip7Address {
	result := []ip7Address{}

	for _, address := range input {
		var buffer string
		var isHypernetSequence bool
		next := ip7Address{sections: []ip7Section{}}

		// HACK: Final closing bracket ensures we flush our buffer.
		for _, r := range address + "[" {
			char := string(r)

			if char == "[" || char == "]" {
				isHypernetSequence = char == "]"
				if len(buffer) > 0 {
					next.sections = append(next.sections, ip7Section{buffer, isHypernetSequence})
					buffer = ""
				}
				continue
			}

			buffer += char
		}

		result = append(result, next)
	}

	return result
}
