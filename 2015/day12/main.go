package main

import (
	"encoding/json"
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

func main() {
	fmt.Println("--- Day 12: JSAbacusFramework.io ---")
	fmt.Println()

	for _, input := range aoc.GetInput(12) {
		result := 0

		var data map[string]interface{}
		if err := json.Unmarshal([]byte(input), &data); err != nil {
			// Maybe an array?
			var arrayData []interface{}
			if err := json.Unmarshal([]byte(input), &arrayData); err != nil {
				panic(fmt.Sprintf("Cannot parse JSON because: %v", err))
			}

			data = make(map[string]interface{})
			data["some-key"] = arrayData
		}

		leafNodes := getLeafNodes(data)
		for _, node := range leafNodes {
			switch nodeType := node.(type) {
			case float64:
				result += int(nodeType)
			}
		}

		if len(input) > 10 {
			fmt.Printf("  - %v... -> %d\n", input[0:10], result)
		} else {
			fmt.Printf("  - %v -> %d\n", input, result)
		}

	}

	fmt.Println()
	fmt.Println()
}

func getLeafNodes(data map[string]interface{}) []interface{} {
	var result []interface{}

	for _, value := range data {
		switch valueType := value.(type) {
		case map[string]interface{}:
			result = append(result, getLeafNodes(valueType)...)
		case []interface{}:
			for _, element := range valueType {
				switch elementType := element.(type) {
				case map[string]interface{}:
					result = append(result, getLeafNodes(elementType)...)
				case []interface{}:
					result = append(result, getLeafNodes(map[string]interface{}{"array": elementType})...)
				default:
					result = append(result, elementType)
				}
			}
		default:
			result = append(result, value)
		}
	}

	return result
}
