package main

import (
	"fmt"
	"slices"
	"strings"
)

// Facility stores the puzzle state within a unsigned 64 bit integer.  This allows for fast cloning
// of state.  At the expense of added complexity.  The below describes how bits are encoded within
// the state field.
//
// Floors, chips and generators are encoded by flipping bits.  Or to put it another way, the
// game state is hashed.
//
// Each floor is assigned a range of bits.  The range is dynamic.  It is equal to the total number of
// isotopes multiplied by 2.  Each isotope has a microchip and a generator.  The first series bits
// are reserved for microchips.  The remaining bits are reserved for generators.  If the bit is true
// the chip/generator is present.  If not it is not.
//
// Chips and generators come in pairs.  Where a pair share a common isotope.  There are up to 6
// isotopes.  Let's call them a, b, c, d, e and f.  Isotope a is assigned the first chip and
// generator bit.  Isotope b is assigned the 2nd.  And so on.  The offset between is a chip and its
// isotope is equal to the number of isotopes.
//
// Finally the player is always located on one of the 4 floors.  The location is encoded in metadata
// section.  This appears after the final floor.  Only one of these bits is set at a time.
//
// Bit Lookup:
// The table below shows what information is contained within each bit.  The table assumes 6 isotopes.
// Fewer isotopes would result in fewer bits required.
//
//	1         2         3         4         5
//
// pos	      0123456789012345678901234567890123456789012345678901
// ---------------------------------------------------------------
// floor      111111111111222222222222333333333333444444444444
// isotope    abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef
// chip       123456      123456      123456      123456
// generator        123456      123456      123456      123456
// location                                                   1234
type facility struct {
	floorCount   int
	isotopeCount int
	state        uint64
}

func (f *facility) copy() facility {
	return facility{
		f.floorCount,
		f.isotopeCount,
		f.state,
	}
}

// True if the current state solves the puzzle.
// The puzzle is solved when all chips and generators are on the top floor.
func (f *facility) isWin() bool {
	fromBit := f.getFloorBitOffset(f.floorCount - 1)
	toBit := fromBit + f.getBitsPerFloor()
	for i := fromBit; i < toBit; i++ {
		if !f.getBit(i) {
			return false
		}
	}

	return true
}

// State is valid if there are no unshielded microchips sharing a floor with a generator.
// These chips would be fried, and the puzzle lost.
func (f *facility) isValid() bool {
	for i := 0; i < f.floorCount; i++ {

		var unshieldedMicrochips int
		var generatorCount int
		for j := 0; j < f.isotopeCount; j++ {
			microchipBit := f.getMicrochipBitIndex(j, i)
			generatorBit := f.getGeneratorBitIndex(j, i)

			if f.getBit(microchipBit) && !f.getBit(generatorBit) {
				unshieldedMicrochips++
			}

			if f.getBit(generatorBit) {
				generatorCount++
			}

			if unshieldedMicrochips > 0 && generatorCount > 0 {
				return false
			}
		}
	}

	return true
}

func (f *facility) getBitsPerFloor() int {
	// 2 bits reserved for each isotopes.
	// 1 for its microchip, the other for its generator.
	return f.isotopeCount * 2
}

func (f *facility) getMicrochipBitIndex(isotopeId int, floor int) int {
	if floor >= f.floorCount {
		panic(fmt.Sprintf("Cannot get bit index for floor %d.  Floor does not exist.", floor))
	}

	return (f.getBitsPerFloor() * floor) + isotopeId
}

func (f *facility) getGeneratorBitIndex(isotopeId int, floor int) int {
	if floor >= f.floorCount {
		panic(fmt.Sprintf("Cannot get bit index for floor %d.  Floor does not exist.", floor))
	}

	return (f.getBitsPerFloor() * floor) + f.isotopeCount + isotopeId
}

func (f *facility) getFloorBitOffset(floor int) int {
	if floor >= f.floorCount {
		panic(fmt.Sprintf("Cannot get bit offset for floor %d.  Floor does not exist.", floor))
	}

	return f.getBitsPerFloor() * floor
}

func (f *facility) getMetadataBitOffset() int {
	return f.floorCount * f.getBitsPerFloor()
}

func (f *facility) setMicrochipFloor(isotopeId, floor int) {
	for i := 0; i < f.floorCount; i++ {
		bitIndex := f.getMicrochipBitIndex(isotopeId, i)
		if i == floor {
			f.setBit(bitIndex)
		} else {
			f.clearBit(bitIndex)
		}
	}
}

func (f *facility) setGeneratorFloor(isotopeId, floor int) {
	for i := 0; i < f.floorCount; i++ {
		bitIndex := f.getGeneratorBitIndex(isotopeId, i)
		if i == floor {
			f.setBit(bitIndex)
		} else {
			f.clearBit(bitIndex)
		}
	}
}

func (f *facility) getCurrentFloor() int {
	for i := 0; i < f.floorCount; i++ {
		if f.getBit(f.getMetadataBitOffset() + i) {
			return i
		}
	}

	panic("Facility state is corrupted.  Could not read current floor.")
}

func (f *facility) setCurrentFloor(floor int) {
	for i := 0; i < f.floorCount; i++ {
		if i == floor {
			f.setBit(f.getMetadataBitOffset() + i)
		} else {
			f.clearBit(f.getMetadataBitOffset() + i)
		}
	}
}

func (f *facility) getBit(pos int) bool {
	return f.state&(1<<pos) == (1 << pos)
}

func (f *facility) setBit(pos int) {
	f.state = f.state | (1 << pos)
}

func (f *facility) clearBit(pos int) {
	f.state = f.state &^ (1 << pos)
}

func (f *facility) listMoves() []facility {
	candidateBits := []int{}
	result := []facility{}

	for _, floorOffset := range []int{1, -1} {
		nextFloor := f.getCurrentFloor() + floorOffset
		if nextFloor >= 0 && nextFloor < f.floorCount {
			nextFloorOffset := floorOffset * f.getBitsPerFloor()

			// get bits on floor
			for i := f.getFloorBitOffset(f.getCurrentFloor()); i <= f.getFloorBitOffset(nextFloor); i++ {
				if f.getBit(i) {
					candidateBits = append(candidateBits, i)
				}
			}

			// append unique combos
			visited := map[int]bool{}
			for _, i := range candidateBits {
				visited[i] = true
				for _, j := range candidateBits {
					if !visited[j] {
						candidate := f.copy()
						candidate.clearBit(i)
						candidate.clearBit(j)
						candidate.setBit(i + nextFloorOffset)
						candidate.setBit(j + nextFloorOffset)
						if candidate.isValid() {
							result = append(result, candidate)
						}
					}
				}
			}

			// append singles
			for _, i := range candidateBits {
				candidate := f.copy()
				candidate.clearBit(i)
				candidate.setBit(i + nextFloorOffset)
				if candidate.isValid() {
					result = append(result, candidate)
				}
			}
		}
	}

	return result
}

func (f *facility) String() string {
	floors := []string{"F1 ", "F2 ", "F3 ", "F4 "}

	for i := 0; i < f.floorCount; i++ {
		for j := 0; j < f.getBitsPerFloor(); j++ {
			bit := (i * f.getBitsPerFloor()) + j
			if f.getBit(bit) {
				if j < f.isotopeCount {
					floors[i] += fmt.Sprintf("%vM ", j)
				} else {
					floors[i] += fmt.Sprintf("%vG ", j-f.isotopeCount)
				}
			} else {
				floors[i] += ".  "
			}
		}
	}

	slices.Reverse(floors)
	return strings.Join(floors, "\n")
}
