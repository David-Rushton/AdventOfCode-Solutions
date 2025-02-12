package main

import (
	"slices"
	"testing"
)

func TestFacilityFactoryBuildReturnsExpectedState(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("plutonium", 1)
	ff.addGenerator("plutonium", 2)
	f := ff.build()

	plutoniumChipId := 0
	plutoniumGeneratorId := 1
	firstFloor := 1
	secondFloor := 2
	floorCount := 4
	bitsPerFloor := 2 // 1 each for plutonium chip and generator.
	expectedBits := []int{
		(bitsPerFloor * firstFloor) + plutoniumChipId,
		(bitsPerFloor * secondFloor) + plutoniumGeneratorId,
		// Current location of player (defaults to floor 0)
		(bitsPerFloor * floorCount),
	}

	// State is a 64 bit int.
	for actualBit := 0; actualBit < 64; actualBit++ {
		if slices.Contains(expectedBits, actualBit) {
			if !f.getBit(actualBit) {
				t.Errorf("Expected bit %d to be set", actualBit)
			}
		}

		if !slices.Contains(expectedBits, actualBit) {
			if f.getBit(actualBit) {
				t.Errorf("Expected bit %d to be unset", actualBit)
			}
		}
	}
}

func TestGetSetFloorRoundTrips(t *testing.T) {
	ff := newFacilityFactory()
	f := ff.build()

	// there are always 4 floors
	for i := 0; i < 4; i++ {
		f.setCurrentFloor(i)
		if f.getCurrentFloor() != i {
			t.Errorf("Expected current floor to be %d not %d", i, f.getCurrentFloor())
		}
	}
}

func TestIsValidCorrectlyReturnsTrue(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 3)
	ff.addGenerator("strontium", 3)
	ff.addMicrochip("plutonium", 2)
	ff.addGenerator("plutonium", 2)
	f := ff.build()

	if !f.isValid() {
		t.Error("Expected isValid to return true.")
	}
}

func TestIsValidCorrectlyReturnsFalse(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 2)
	ff.addGenerator("strontium", 3)
	ff.addMicrochip("plutonium", 2)
	ff.addGenerator("plutonium", 2)
	f := ff.build()

	if f.isValid() {
		t.Error("Expected isValid to return false.")
	}
}

func TestIsWinReturnsTrueForWin(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 3)
	ff.addGenerator("strontium", 3)
	ff.addMicrochip("plutonium", 3)
	ff.addGenerator("plutonium", 3)
	f := ff.build()

	if !f.isWin() {
		t.Error("Expected isWin to return true.")
	}
}

func TestIsWinReturnsFalseForNonWin(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 1)
	ff.addGenerator("strontium", 1)
	ff.addMicrochip("plutonium", 2)
	ff.addGenerator("plutonium", 2)
	f := ff.build()

	if f.isWin() {
		t.Error("Expected isWin to return false.")
	}
}

func TestListMovesReturnExpectedStates(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 1)
	ff.addGenerator("strontium", 1)
	ff.addMicrochip("plutonium", 2)
	ff.addGenerator("plutonium", 2)
	f := ff.build()
	f.setCurrentFloor(1)
	moves := f.listMoves()

	// there are 5 valid moves from this state:
	//   - strontium chip down
	//   - strontium generator down
	//   - strontium chip & generator down
	//   - strontium generator up
	//   - strontium chip & generator up
	for len(moves) != 5 {
		t.Errorf("Expected 5 moves not %d", len(moves))
	}
}
