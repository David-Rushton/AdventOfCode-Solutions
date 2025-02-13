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

func TestSetMicrochipFloor(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("plutonium", 1) // bit 4 || isotope id 0
	ff.addMicrochip("hydrogen", 1)  // bit 5 || isotope id 1
	f := ff.build()

	f.setMicrochipFloor(0, 0) // move plutonium to ground floor || bit 0
	f.setMicrochipFloor(1, 3) // move hydrogen to top floor  || bit 13

	expectedBits := []int{0, 15}
	for bit := 0; bit < 13; bit++ {
		if slices.Contains(expectedBits, bit) {
			if !f.getBit(bit) {
				t.Errorf("Expected bit %d to be set", bit)
			}
		}

		if !slices.Contains(expectedBits, bit) {
			if f.getBit(bit) {
				t.Errorf("Expected bit %d to be unset", bit)
			}
		}
	}
}

func TestSetGeneratorFloor(t *testing.T) {
	ff := newFacilityFactory()
	ff.addGenerator("plutonium", 1) // bit 6 || isotope id 0
	ff.addGenerator("hydrogen", 1)  // bit 7  || isotope id 1
	f := ff.build()

	f.setGeneratorFloor(0, 0) // move plutonium to ground floor || bit 0
	f.setGeneratorFloor(1, 3) // move hydrogen to top floor  || bit 13

	expectedBits := []int{2, 14}
	for bit := 0; bit < 13; bit++ {
		if slices.Contains(expectedBits, bit) {
			if !f.getBit(bit) {
				t.Errorf("Expected bit %d to be set", bit)
			}
		}

		if !slices.Contains(expectedBits, bit) {
			if f.getBit(bit) {
				t.Errorf("Expected bit %d to be unset", bit)
			}
		}
	}
}

func TestGetSetFloorRoundTrips(t *testing.T) {
	ff := newFacilityFactory()
	f := ff.build()

	for i := 0; i < f.floorCount; i++ {
		f.setCurrentFloor(i)
		if f.getCurrentFloor() != i {
			t.Errorf("Expected current floor to be %d not %d.", i, f.getCurrentFloor())
		}
	}
}

func TestGetSetBitRoundTrips(t *testing.T) {
	ff := newFacilityFactory()
	f := ff.build()

	testBits := []int{0, 1, 10, 44, 63}
	for _, bit := range testBits {
		f.setBit(bit)
		if !f.getBit(bit) {
			t.Errorf("Expected bit %d to be set.", bit)
		}
	}
}

func TestSetClearBitRoundTrips(t *testing.T) {
	ff := newFacilityFactory()
	f := ff.build()

	testBits := []int{0, 1, 10, 44, 63}
	for _, bit := range testBits {
		f.setBit(bit)
		f.clearBit(bit)
		if f.getBit(bit) {
			t.Errorf("Expected bit %d to be unset.", bit)
		}
	}

	if f.state != 0 {
		t.Error("Expected state to be cleared.")
	}
}

func TestIsValidCorrectlyReturnsTrue(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("hydrogen", 1)
	ff.addGenerator("hydrogen", 1)
	ff.addMicrochip("lithium", 0)
	ff.addGenerator("lithium", 2)
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

func TestStringReturnsExpectedState(t *testing.T) {
	ff := newFacilityFactory()
	ff.addMicrochip("strontium", 1)
	ff.addGenerator("strontium", 1)
	ff.addMicrochip("plutonium", 2)
	ff.addGenerator("plutonium", 2)
	ff.addMicrochip("hydrogen", 0)
	ff.addGenerator("hydrogen", 3)
	f := ff.build()

	// strontium = 0, plutonium = 1 & hydrogen = 2
	expect := `F4 .  .  .  .  .  2G 
F3 .  1M .  .  1G .  
F2 0M .  .  0G .  .  
F1 .  .  2M .  .  .  `
	actual := f.String()

	if actual != expect {
		t.Error("Actual output does not match expected.")
	}
}
