package main

import (
	"fmt"
	"log"
	"math"
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

	fmt.Println("---------")
	fmt.Println(len(wires))
	fmt.Println("---------")

	fmt.Println()
	fmt.Printf("Reading x: % 16d\n", x)
	fmt.Printf("Reading y: % 16d\n", y)
	fmt.Printf("Reading z: % 16d\n", z)
	fmt.Printf("Result:    % 16d\n", z-(x+y))
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
