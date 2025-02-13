package main

import (
	"fmt"
)

type facilityFactory struct {
	isotopes   map[string]int
	microchips map[int][]string
	generators map[int][]string
	state      uint64
	floorCount int
}

func newFacilityFactory() facilityFactory {
	return facilityFactory{
		isotopes:   map[string]int{},
		microchips: map[int][]string{},
		generators: map[int][]string{},
		state:      0,
		floorCount: 4,
	}
}

func (ff *facilityFactory) addMicrochip(isotope string, floor int) {
	if floor >= ff.floorCount {
		panic(fmt.Sprintf("Cannot add microchip to floor %v.  Floor does not exist.", floor))
	}

	ff.addIsotope(isotope)

	if _, exists := ff.microchips[floor]; !exists {
		ff.microchips[floor] = []string{}
	}

	ff.microchips[floor] = append(ff.microchips[floor], isotope)
}

func (ff *facilityFactory) addGenerator(isotope string, floor int) {
	if floor >= ff.floorCount {
		panic(fmt.Sprintf("Cannot add generator to floor %v.  Floor does not exist.", floor))
	}

	ff.addIsotope(isotope)

	if _, exists := ff.generators[floor]; !exists {
		ff.generators[floor] = []string{}
	}

	ff.generators[floor] = append(ff.generators[floor], isotope)
}

func (ff *facilityFactory) addIsotope(name string) {
	if _, exists := ff.isotopes[name]; !exists {
		ff.isotopes[name] = len(ff.isotopes)
	}
}

func (ff *facilityFactory) build() facility {
	if len(ff.isotopes)*2*ff.floorCount > 64 {
		panic("Not enough memory to store all floors and isotopes")
	}

	// Create facility.
	result := facility{
		floorCount:   ff.floorCount,
		isotopeCount: len(ff.isotopes),
		state:        0,
	}

	// Set default locations.
	result.setCurrentFloor(0)

	for floor, isotopes := range ff.microchips {
		for i := range isotopes {
			result.setMicrochipFloor(ff.isotopes[isotopes[i]], floor)
		}
	}

	for floor, isotopes := range ff.generators {
		for i := range isotopes {
			result.setGeneratorFloor(ff.isotopes[isotopes[i]], floor)
		}
	}

	return result
}
