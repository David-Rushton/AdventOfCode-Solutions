package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2017/internal/aoc"
)

type program struct {
	name     string
	weight   int
	parent   *program
	children []*program
}

func (p *program) getTotalWeight() int {
	weight := p.weight

	for _, child := range p.children {
		weight += child.getTotalWeight()
	}

	return weight
}

func main() {
	fmt.Println("--- Day 7: Recursive Circus ---")
	fmt.Println()

	var root = parse(aoc.GetInput(7))
	findUnbalanced(root)

	fmt.Println()
	fmt.Printf("Result: %v\n", root.name)
}

func findUnbalanced(root *program) {
	var printPrograms func(p *program)
	printPrograms = func(p *program) {
		fmt.Printf("- %v %d\n", p.name, p.getTotalWeight())

		for _, child := range p.children {
			printPrograms(child)
		}
	}

	var solve func(p *program) *program
	solve = func(p *program) *program {
		var weights = map[int]int{}
		var byWeights = map[int]*program{}
		var outlier *program

		// Find outlier.
		for _, child := range p.children {
			var totalWeight = child.getTotalWeight()
			weights[totalWeight]++
			byWeights[totalWeight] = child
		}

		for weight, count := range weights {
			if count == 1 {
				outlier = byWeights[weight]
				break
			}
		}

		if outlier == nil {
			return nil
		}

		// If outlier children are balanced investigate them.
		// If not outlier is the problem node.

		subOutlier := solve(outlier)
		if subOutlier == nil {
			fmt.Println(outlier.name)
		}

		return outlier
	}
	solve(root)
}

func parse(input []string) *program {
	var programMap = map[string]*program{}
	var childMap = map[string][]string{}

	// Process input.
	for i := range input {
		// Prepare string.
		var txt = input[i]
		for _, remove := range []string{"(", ")", "->", ","} {
			txt = strings.ReplaceAll(txt, remove, "")
		}

		// Break into sections.
		elements := strings.Split(txt, " ")

		// Create.
		programMap[elements[0]] = &program{
			name:   elements[0],
			weight: aoc.ToInt(elements[1]),
		}

		if len(elements) > 3 {
			childMap[elements[0]] = elements[3:]
		}
	}

	// Convert to return types.
	for parentName, children := range childMap {
		for _, childName := range children {
			programMap[parentName].children = append(
				programMap[parentName].children,
				programMap[childName])

			programMap[childName].parent = programMap[parentName]
		}
	}

	// Find and return root.
	var root *program
	for _, program := range programMap {
		root = program
		for root.parent != nil {
			root = root.parent
		}
		break
	}

	return root
}
