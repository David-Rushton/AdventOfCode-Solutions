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

	fmt.Println()
	fmt.Println()

	// index by out.
	outs := make(map[string]*gate)
	zs := make(map[string]int)
	zKey := []string{}

	for _, connectedGates := range gates {
		for _, connectedGate := range connectedGates {
			outs[connectedGate.Out.name] = connectedGate

			if strings.HasPrefix(connectedGate.Out.name, "z") {
				zs[connectedGate.Out.name]++

				if !slices.Contains(zKey, connectedGate.Out.name) {
					zKey = append(zKey, connectedGate.Out.name)
				}
			}
		}
	}

	candidates := []string{
		"z06",
		"z07",
		"z08",
		"z09",
		"z10",
		"z11",
		"z12",
		"z13",
	}

	// let's work back from the outputs.
	for outName := range outs {
		queue := []string{outName}
		visited := make(map[string]bool)
		hits := 0

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if visited[current] {
				continue
			}

			fmt.Printf("  %v\n", current)

			visited[current] = true

			for _, n := range []string{outs[current].leftIn.name, outs[current].rightIn.name} {
				if _, exists := outs[n]; exists {
					hits++
					queue = append(queue, outs[current].leftIn.name, outs[current].rightIn.name)
				} else {
					// fmt.Printf("!")
				}
			}
		}

		// fmt.Printf("  %v", current)
		fmt.Printf(" %v == %d\n", outName, hits)
		zs[outName] = hits

		fmt.Println()
		fmt.Println()

		check := []string{
			"bnt",
			"fcd",
			"fjv",
			"gbd",
			"kwb",
			"nmm",
			"z06",
			"z07",
			"z08",
			"z09",
			"z10",
			"z11",
			"z12",
			"z13",
		}

		for _, yyy := range check {
			if _, exists := outs[yyy]; exists {
				fmt.Println(yyy)
			}
		}
	}

	fmt.Println()
	fmt.Println()

	slices.Sort(zKey)

	ex := -2
	for _, k := range zKey {
		if zs[k] != ex || true {
			fmt.Printf("  %s % 3d % 3d\n", k, zs[k], ex)
		}
		ex += 4
	}

	// ?
	fmt.Println()
	fmt.Println()

	for _, k := range candidates {
		if zs[k] != ex {
			fmt.Printf("  %v ( %v, %v ) \n", k, outs[k].leftIn.name, outs[k].rightIn.name)
		}
		ex += 4
	}

}

func someIdea() {
	// z00   0  -2
	// -----------
	// z06   0  22
	// z07  28  26
	// z08  32  30
	// z09  36  34
	// z10  40  38
	// z11  44  42
	// z12  48  46
	// z13  54  50
	// -----------
	// z45  176  178

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
