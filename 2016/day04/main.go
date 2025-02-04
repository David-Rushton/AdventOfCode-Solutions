package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

type room struct {
	name      string
	sectorId  int
	checksum  string
	validated bool
}

func main() {
	fmt.Println("--- Day 4: Security Through Obscurity ---")
	fmt.Println()

	rooms := parse(aoc.GetInput(4))
	sectorSum, northPoleSector := solve(rooms)

	fmt.Println()
	fmt.Printf("Valid sector Sum: %d\n", sectorSum)
	fmt.Printf("North Pole Object Storage Sector: %d\n", northPoleSector)
}

func solve(rooms []room) (sectorSum, northPoleSector int) {
	sectorSum = 0
	northPoleSector = 0

	for _, room := range rooms {
		// Count occurrence of each character in the name.
		characterCount := map[rune]int{}
		for _, r := range room.name {
			if r == '-' {
				continue
			}

			characterCount[r]++
		}

		// Create count to character(s) lookup.
		countMap := map[int]string{}
		countMapKeys := []int{}
		for k, v := range characterCount {
			if !slices.Contains(countMapKeys, v) {
				countMapKeys = append(countMapKeys, v)
			}
			countMap[v] += string(k)
		}

		// Get top five by count, tie-breaking by alphabetical order.
		slices.Sort(countMapKeys)
		slices.Reverse(countMapKeys)

		var checksum string
		for _, countKey := range countMapKeys {
			checksum += alphaSort(countMap[countKey])

			if len(checksum) >= 5 {
				break
			}
		}

		if len(checksum) > 5 {
			checksum = checksum[0:5]
		}

		// Is checksum valid?
		if room.checksum == checksum {
			sectorSum += room.sectorId

			decryptedName := decrypt(room.name, room.sectorId)
			fmt.Printf(
				" - Valid room `%v` found in section %v\n",
				decryptedName,
				room.sectorId)

			if decryptedName == "northpole object storage" {
				northPoleSector = room.sectorId
				fmt.Println(" - ❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
			}

		}
	}

	return sectorSum, northPoleSector
}

func decrypt(value string, shift int) string {
	var result string

	for _, r := range value {
		if r == '-' {
			result += " "
		} else {
			result += string(rune(97 + ((int(r) - 97 + shift) % 26)))
		}
	}

	return result
}

func alphaSort(value string) string {
	characters := []rune{}
	for _, r := range value {
		characters = append(characters, r)
	}

	slices.Sort(characters)

	var result string
	for _, r := range characters {
		result += string(r)
	}

	return result
}

func parse(input []string) []room {
	rooms := []room{}

	for i := range input {
		sectorStart := strings.LastIndex(input[i], "-")
		checksumStart := strings.Index(input[i], "[")
		rooms = append(rooms, room{
			name:     input[i][0:sectorStart],
			sectorId: aoc.ToInt(input[i][sectorStart+1 : checksumStart]),
			checksum: input[i][checksumStart+1 : len(input[i])-1]})
	}

	return rooms
}
