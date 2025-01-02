package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

func main() {
	fmt.Println("--- Day 24: Crossed Wires ---")
	fmt.Println()

	wires, gates, initialValues := parse(aoc.Input)
	powerUp(gates, initialValues)
	x, y, z := getReadings(wires)

	fmt.Println()
	fmt.Printf("Reading x: % 16d\n", x)
	fmt.Printf("Reading y: % 16d\n", y)
	fmt.Printf("Reading z: % 16d\n", z)
	fmt.Printf("Result:    % 16d\n", z-(x+y))
	fmt.Printf("Faulty:    %v\n", strings.Join(getFaultyWires(gates), ","))
}

func getFaultyWires(gates map[string][]*gate) []string {
	allGates := map[*gate]bool{}
	for _, connected := range gates {
		for _, current := range connected {
			allGates[current] = true
		}
	}

	faulty := []string{}
	for current := range allGates {
		// if output is z should be XOR
		if current.Out.name[0] == 'z' && current.operator != operatorXOr {
			if current.Out.name != "z45" {
				faulty = append(faulty, current.Out.name)
			}
		}

		// if output is z and inputs are not x,y then not XOR
		if current.Out.name[0] != 'z' {
			if !(current.leftIn.name[0] == 'x' || current.leftIn.name[0] == 'y') {
				if !(current.rightIn.name[1] == 'x' || current.rightIn.name[1] == 'y') {
					if current.operator == operatorXOr {
						faulty = append(faulty, current.Out.name)
					}
				}

			}
		}

		// if XOR with x,y in
		// must be another XOR gate that takes this as in input
		if current.operator == operatorXOr {
			if current.leftIn.name[0] == 'x' || current.leftIn.name[0] == 'y' {
				if current.rightIn.name[0] == 'x' || current.rightIn.name[0] == 'y' {
					if !(current.leftIn.name == "x00" || current.leftIn.name == "y00") {
						if !(current.rightIn.name == "x00" || current.rightIn.name == "y00") {
							if len(gates[current.Out.name]) > 0 {
								var xOrFound bool
								for _, child := range gates[current.Out.name] {
									if child.operator == operatorXOr {
										xOrFound = true
									}
								}

								if !xOrFound {
									faulty = append(faulty, current.Out.name)
								}
							}
						}
					}
				}
			}
		}

		// if AND with x,y in
		// must be another OR gate that takes this as in input
		if current.operator == operatorAnd {
			if current.leftIn.name[0] == 'x' || current.leftIn.name[0] == 'y' {
				if current.rightIn.name[0] == 'x' || current.rightIn.name[0] == 'y' {
					if !(current.leftIn.name == "x00" || current.leftIn.name == "y00") {
						if !(current.rightIn.name == "x00" || current.rightIn.name == "y00") {
							if len(gates[current.Out.name]) > 0 {
								var orFound bool
								for _, child := range gates[current.Out.name] {
									if child.operator == operatorOr {
										orFound = true
									}
								}

								if !orFound {
									faulty = append(faulty, current.Out.name)
								}
							}
						}
					}
				}
			}
		}
	}

	slices.Sort(faulty)
	return faulty
}

func getReadings(wires map[string]*wire) (x, y, z int64) {
	for name := range wires {

		if name[0] == 'x' || name[0] == 'y' || name[0] == 'z' {
			index, err := strconv.ParseInt(name[1:], 10, 64)
			if err != nil {
				log.Fatalf("Cannot convert %s to a number.", err)
			}

			if wires[name].getValue() {
				switch name[0] {
				case 'x':
					x += 1 * int64(math.Pow(2, float64(index)))
				case 'y':
					y += 1 * int64(math.Pow(2, float64(index)))
				case 'z':
					z += 1 * int64(math.Pow(2, float64(index)))
				}
			}
		}
	}

	return x, y, z
}

func powerUp(gates map[string][]*gate, initialValues []wireUpdate) {
	queue := initialValues

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		current.wire.setValue(current.value)

		for _, gate := range gates[current.wire.name] {
			if gate.hasValue() {
				queue = append(queue, wireUpdate{gate.Out, gate.getValue()})
			}
		}
	}
}

func parse(input []string) (wires map[string]*wire, gates map[string][]*gate, initialValues []wireUpdate) {
	wires = map[string]*wire{}
	gates = map[string][]*gate{}
	initialValues = []wireUpdate{}

	addWire := func(name string) {
		if _, exists := wires[name]; !exists {
			wires[name] = &wire{
				name:     name,
				value:    false,
				hasValue: false,
			}
		}
	}

	startupMode := true

	for _, line := range input {
		if line == "" {
			startupMode = false
			continue
		}

		elements := strings.Split(strings.Replace(line, ":", "", -1), " ")

		// Initial values.
		if startupMode {
			addWire(elements[0])
			initialValues = append(initialValues, wireUpdate{
				wires[elements[0]],
				elements[1] == "1"})
			continue
		}

		// Wires.
		for _, name := range []string{elements[0], elements[2], elements[4]} {
			addWire(name)
		}

		var operator gateOperator
		switch elements[1] {
		case "AND":
			operator = operatorAnd
		case "XOR":
			operator = operatorXOr
		case "OR":
			operator = operatorOr
		default:
			log.Fatalf("Invalid operation: %v\n", elements[1])
		}

		// gates
		newGate := &gate{
			leftIn:   wires[elements[0]],
			rightIn:  wires[elements[2]],
			Out:      wires[elements[4]],
			operator: operator,
		}

		for _, wireName := range []string{elements[0], elements[2]} {
			linkedGates, exists := gates[wireName]
			if !exists {
				linkedGates = []*gate{}

			}
			linkedGates = append(linkedGates, newGate)
			gates[wireName] = linkedGates
		}
	}

	return wires, gates, initialValues
}
