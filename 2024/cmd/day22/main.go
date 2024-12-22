package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type sequenceKey struct {
	item0 int
	item1 int
	item2 int
	item3 int
}

func main() {
	fmt.Println("--- Day 22: Monkey Market ---")
	fmt.Println()

	secretNumbers := parse(aoc.Input)
	bananaSequences := []map[sequenceKey]int{}
	var secretSum int

	for _, secretNumber := range secretNumbers {
		bananaSequence, nSecret := findNSecret(secretNumber, 2000)
		bananaSequences = append(bananaSequences, bananaSequence)
		secretSum += nSecret
	}

	sequence, bananas := findBestSequences(bananaSequences)

	fmt.Println()
	fmt.Printf("Best Sequence: % 2d,% 2d,% 2d,% 2d\n", sequence.item0, sequence.item1, sequence.item2, sequence.item3)
	fmt.Printf("Total Bananas: %d\n", bananas)
	fmt.Printf("Secret Sum: %d\n", secretSum)
}

func findBestSequences(bananaSequences []map[sequenceKey]int) (sequenceKey, int) {
	sequenceTotals := make(map[sequenceKey]int)

	// Score each sequence.
	for _, bananaSequence := range bananaSequences {
		for key, bananas := range bananaSequence {
			sequenceTotals[key] += bananas
		}
	}

	// find best
	var bestSequence sequenceKey
	var mostBananas int

	for sequence, bananas := range sequenceTotals {
		if bananas > mostBananas {
			mostBananas = bananas
			bestSequence = sequence
		}
	}

	return bestSequence, mostBananas
}

func findNSecret(secretNumber, n int) (bananaSequence map[sequenceKey]int, final int) {
	final = secretNumber
	bananas := make(map[sequenceKey]int)

	sequence := []int{}
	currentResult := 0
	lastResult := secretNumber % 10
	for i := 0; i < n; i++ {
		final = ((final * 64) ^ final) % 16777216
		final = ((final / 32) ^ final) % 16777216
		final = ((final * 2048) ^ final) % 16777216

		currentResult = final % 10
		sequence = append(sequence, currentResult-lastResult)

		if i >= 3 {
			key := sequenceKey{
				sequence[i-3],
				sequence[i-2],
				sequence[i-1],
				sequence[i],
			}
			if _, found := bananas[key]; !found {
				bananas[key] = currentResult
			}
		}

		lastResult = currentResult
	}

	return bananas, final
}

func parse(input []string) []int {
	result := []int{}
	for _, line := range input {
		result = append(result, toInt(line))
	}
	return result
}

func toInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert %s to number.", s)
	}
	return int(num)
}
