package main

import (
	"testing"
)

func TestBitFunctionsRoundTrip(t *testing.T) {
	// arrange
	var n uint64
	bits := []int{1, 10, 12, 14, 100}

	for _, i := range bits {
		n = setBit(n, i)
	}

	for _, i := range bits {
		if !isBitSet(n, i) {
			t.Errorf("Expected bit %d to be set", i)
		}
	}
}

func TestIsSetBitReturnsFalseForUnsetBits(t *testing.T) {
	var n uint64

	// set first few even bits.
	for i := 2; i < 20; i += 2 {
		n = setBit(n, i)
	}

	// check odd bits are not set.
	for i := 1; i < 21; i += 2 {
		if isBitSet(n, i) {
			t.Errorf("Expected bit %d to be set to zero", i)
		}
	}
}

func TestClearBitRemovesBits(t *testing.T) {
	// arrange
	var n uint64
	bits := []int{1, 10, 12, 14, 100}

	for _, i := range bits {
		n = setBit(n, i)
		n = clearBit(n, i)
	}

	if n != 0 {
		t.Errorf("Expected all bits to be clear but n has value of %v", n)
	}
}

func TestFloorEncoding(t *testing.T) {
	var n uint64

	for i := 0; i < 4; i++ {
		n = setFloor(n, i)
		if actual := getFloor(n); actual != i {
			t.Errorf("Set/get floor does not roundtrip.. Expected %d.  Actual: %d", i, actual)
		}
	}
}
