package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

const (
	numberOfIsotopes = 6
	generatorOffset  = numberOfIsotopes
	bitsPerFloor     = numberOfIsotopes * 2
	numberOfFloors   = 4
	floor1BitsStart  = bitsPerFloor * 0
	floor2BitsStart  = bitsPerFloor * 1
	floor3BitsStart  = bitsPerFloor * 2
	floor4BitsStart  = bitsPerFloor * 3
	metaBitsStart    = bitsPerFloor * numberOfFloors
)

func main() {
	fmt.Println("--- Day 11: Radioisotope Thermoelectric Generators ---")
	fmt.Println()

	// Buckle up!
	//
	// The chosen data encoding is a little convoluted.  The entire state of the game is stored
	// within a unsigned 64 bit integer.
	//
	// Floors, chips and generators are encoded by flipping bits.  Or to put it another way, the
	// game state is hashed.
	//
	// Each floor is assigned a range of 12 bits.  The first 6 bits are reserved for microchips.
	// The remaining 6 bits are reserved for generators.  If the bit is true the chip/generator is
	// present.  If not it is not.
	//
	// Chips and generators come in pairs.  Where a pair share a common isotope.  There are up to 6
	// isotopes.  Let's call them a, b, c, d, e and f.  Isotope a is assigned the first chip and
	// generator bit.  Isotope b is assigned the 2nd.  And so on.
	//
	// Finally the player is always located on one of the 4 floors.  The location is encoded in bits
	// 48 through 51.  Only one of these bits is set at a time.
	//
	// Bit Lookup:
	// The table below shows what information is contained within each bit.
	//
	//                      1         2         3         4         5
	// pos	      0123456789012345678901234567890123456789012345678901
	// ---------------------------------------------------------------
	// floor      111111111111222222222222333333333333444444444444
	// isotope    abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef
	// chip       123456      123456      123456      123456
	// generator        123456      123456      123456      123456
	// location                                                   1234
	state := parse(aoc.GetInput(11))
	result := findMinSteps(state)

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func findMinSteps(state uint64) int {
	minSteps := math.MaxInt

	isWin := func(state uint64) bool {
		// check all bits for 4th floor.
		// if all set, then all chips and isotopes are on the 4th floor.
		// this is a win.
		for i := floor4BitsStart; i < floor4BitsStart+bitsPerFloor; i++ {
			if !isBitSet(state, i) {
				return false
			}
		}

		return true
	}

	isValid := func(state uint64) bool {
		// generators fry unshielded chips, on the same floor.
		// chips are shielded if their partner generator is on the same floor.
		// if any chip is fired the state is invalid.
		for i := 0; i < numberOfFloors; i++ {
			floorOffset := bitsPerFloor * i
			var unshieldedChips int
			var generators int
			for j := 0; j < numberOfIsotopes; j++ {
				chipBit := floorOffset + j
				generatorBit := chipBit + generatorOffset

				if isBitSet(state, chipBit) && !isBitSet(state, generatorBit) {
					unshieldedChips++
				}

				if isBitSet(state, generatorBit) {
					generators++
				}

				if unshieldedChips > 0 && generators > 0 {
					return false
				}
			}
		}

		return true
	}

	visited := map[uint64]bool{}
	getPermutations := func(state uint64) []uint64 {
		result := []uint64{}
		floor := getFloor(state)

		for _, floorNumOffset := range []int{1, -1} {
			if nextFloor := floor + floorNumOffset; nextFloor >= 0 && nextFloor <= 3 {
				fromFloorOffset := floor * bitsPerFloor

				// find all chips and generators on this floor.
				setBits := []int{}
				candidateBits := [][]int{}
				for i := fromFloorOffset; i < fromFloorOffset+bitsPerFloor; i++ {
					if isBitSet(state, i) {
						setBits = append(setBits, i)
						candidateBits = append(candidateBits, []int{i})
					}
				}

				// get all unique combinations.
				bitVisited := map[int]bool{}
				for _, i := range setBits {
					bitVisited[i] = true
					for _, j := range setBits {
						if !bitVisited[j] {
							candidateBits = append(candidateBits, []int{i, j})
						}
					}
				}

				// return candidates.
				for _, candidate := range candidateBits {
					next := setFloor(state, nextFloor)

					for _, bit := range candidate {
						next = clearBit(next, bit)
						next = setBit(next, bit+(floorNumOffset*bitsPerFloor))
					}

					if !visited[next] {
						visited[next] = true
						result = append(result, next)
					}
				}
			}
		}

		return result
	}

	var iterations int
	var solve func(steps int, state uint64)
	solve = func(steps int, state uint64) {
		iterations++

		if !isValid(state) {
			return
		}

		if steps > minSteps {
			return
		}

		if isWin(state) {
			if steps < minSteps {
				minSteps = steps
				fmt.Printf(" - Found new best of %d steps\n", steps)
				return
			}
		}

		for _, permutation := range getPermutations(state) {
			solve(steps+1, permutation)
		}
	}

	solve(0, state)

	fmt.Printf(" - Total iterations %d\n", iterations)

	return minSteps
}

func isBitSet(n uint64, pos int) bool {
	return n&(1<<pos) == 1<<pos
}

func setBit(n uint64, pos int) uint64 {
	return n | (1 << pos)
}

func clearBit(n uint64, pos int) uint64 {
	return n &^ (1 << pos)
}

func setFloor(n uint64, floor int) uint64 {
	result := n

	for i := 0; i < numberOfFloors; i++ {
		if i == floor {
			result = setBit(result, metaBitsStart+i)
		} else {
			result = clearBit(result, metaBitsStart+i)
		}
	}

	return result
}

func getFloor(n uint64) int {
	for i := 0; i < numberOfFloors; i++ {
		if isBitSet(n, metaBitsStart+i) {
			return i
		}
	}

	panic("cannot find floor")
}

func print(state uint64) {

}

func parse(input []string) uint64 {
	floorMap := map[string]int{
		"first":  floor1BitsStart,
		"second": floor2BitsStart,
		"third":  floor3BitsStart,
		"fourth": floor4BitsStart,
	}

	ignoreValues := []string{"contains", "nothing", "relevant", "and", "a", "microchip", "generator"}

	var result uint64
	var isotopes int

	idx := map[string]int{}
	nextIdx := 0

	for i := range input {
		elements := strings.Split(strings.ReplaceAll(strings.ReplaceAll(input[i], ",", ""), ".", ""), " ")
		floor := floorMap[elements[1]]

		for _, value := range elements[4:] {
			if !slices.Contains(ignoreValues, value) {
				valueElements := strings.Split(value, "-")

				if _, exists := idx[valueElements[0]]; !exists {
					idx[valueElements[0]] = nextIdx
					nextIdx++
					if isotopes++; isotopes > 6 {
						panic("up to 6 isotopes are supported")
					}
				}
				idx := idx[valueElements[0]]

				// chip
				if len(valueElements) > 1 {
					result = setBit(result, floor+idx)
				}

				// generator
				if len(valueElements) == 1 {
					result = setBit(result, floor+idx+generatorOffset)
				}
			}
		}
	}

	result = setFloor(result, 0)

	return result
}
