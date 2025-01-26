package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

type operations int

const (
	literal operations = iota
	not
	and
	or
	leftShift
	rightShift
)

type operation struct {
	operation operations
	value     uint16
	hasValue  bool
	leftWire  string
	rightWire string
	outWire   string
}

func main() {
	fmt.Println("--- Day 7: Some Assembly Required ---")
	fmt.Println()

	operations := parse(aoc.GetInput(7))
	starOne := runCircuit(operations)
	starTwo := runCircuitWithOverride(operations, "b", starOne)

	fmt.Println()
	fmt.Printf("Result 1: %d\n", starOne)
	fmt.Printf("Result 2: %d\n", starTwo)
}

func runCircuitWithOverride(operations []operation, wire string, value uint16) uint16 {
	newOperations := []operation{}

	for _, current := range operations {
		if current.leftWire == wire {
			newOperations = append(newOperations, operation{
				operation: literal,
				value:     value,
				hasValue:  true,
				leftWire:  "",
				rightWire: "",
				outWire:   wire})
		}

		newOperations = append(newOperations, current)
	}

	return runCircuit(newOperations)
}

func runCircuit(operations []operation) uint16 {
	wires := map[string]uint16{"1": 1}

	canProcess := func(op operation) bool {
		for _, wire := range []string{op.leftWire, op.rightWire} {
			if wire != "" {
				if _, hasValue := wires[wire]; !hasValue {
					return false
				}
			}
		}

		return true
	}

	for len(operations) > 0 {
		fmt.Printf(" - Queue: %d\r", len(operations))

		current := operations[0]
		operations = operations[1:]

		if !canProcess(current) {
			operations = append(operations, current)
			continue
		}

		switch current.operation {
		case literal:
			if current.hasValue {
				wires[current.outWire] = current.value
			} else {
				wires[current.outWire] = wires[current.leftWire]
			}
		case not:
			wires[current.outWire] = ^wires[current.leftWire]
		case and:
			wires[current.outWire] = wires[current.leftWire] & wires[current.rightWire]
		case or:
			wires[current.outWire] = wires[current.leftWire] | wires[current.rightWire]
		case leftShift:
			wires[current.outWire] = wires[current.leftWire] << current.value
		case rightShift:
			wires[current.outWire] = wires[current.leftWire] >> current.value
		}
	}

	fmt.Println()

	return wires["a"]
}

func parse(input []string) []operation {
	result := []operation{}

	for _, line := range input {
		if line == "" {
			continue
		}

		elements := strings.Split(line, " ")
		var next operation

		switch len(elements) {
		case 3:
			next = operation{
				operation: literal,
				leftWire:  "",
				rightWire: "",
				outWire:   elements[2],
			}

			if aoc.IsInt(elements[0]) {
				next.value = uint16(aoc.ToInt(elements[0]))
				next.hasValue = true
			} else {
				next.leftWire = elements[0]
				next.hasValue = false
			}
		case 4:
			next = operation{
				operation: not,
				value:     0,
				hasValue:  false,
				leftWire:  elements[1],
				rightWire: "",
				outWire:   elements[3],
			}
		case 5:
			var operationType operations
			switch elements[1] {
			case "NOT":
				operationType = not
			case "AND":
				operationType = and
			case "OR":
				operationType = or
			case "LSHIFT":
				operationType = leftShift
			case "RSHIFT":
				operationType = rightShift
			}

			next = operation{
				operation: operationType,
				value:     0,
				hasValue:  false,
				leftWire:  elements[0],
				rightWire: elements[2],
				outWire:   elements[4],
			}

			if operationType == leftShift || operationType == rightShift {
				next.rightWire = ""
				next.value = uint16(aoc.ToInt(elements[2]))
				next.hasValue = true
			}
		default:
			log.Fatalf("Cannot parse line: %v.\n", line)
		}

		result = append(result, next)
	}

	return result
}
