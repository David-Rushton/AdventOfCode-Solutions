package main

import (
	"encoding/json"
	"fmt"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2015/internal/aoc"
)

// !58042
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

		leafNodes := getLeafNodes(data, aoc.Star == aoc.StarTwo)
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

func getLeafNodes(data map[string]interface{}, skipRed bool) []interface{} {
	var result []interface{}

	for _, value := range data {
		var valueResult []interface{}

		switch valueType := value.(type) {
		case map[string]interface{}:
			valueResult = append(valueResult, getLeafNodes(valueType, skipRed)...)
		case []interface{}:
			var elementResult []interface{}

			for _, element := range valueType {
				switch elementType := element.(type) {
				case map[string]interface{}:
					elementResult = append(elementResult, getLeafNodes(elementType, skipRed)...)
				case []interface{}:
					elementResult = append(elementResult, getLeafNodes(map[string]interface{}{"array": elementType}, skipRed)...)
				default:
					elementResult = append(elementResult, elementType)
				}
			}

			valueResult = append(valueResult, elementResult...)
		case string:
			if skipRed && valueType == "red" {
				return []interface{}{}
			}
		default:
			valueResult = append(valueResult, value)
		}

		result = append(result, valueResult...)
	}

	return result
}
