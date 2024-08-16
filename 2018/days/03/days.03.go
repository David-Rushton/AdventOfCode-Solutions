package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type CellKey struct {
	X, Y int
}

type Cell struct {
	Id     int
	Left   int
	Top    int
	Width  int
	Height int
}

func main() {
	fmt.Print("---------------------------------\n")
	fmt.Print("Day 3: No Matter How You Slice It\n")
	fmt.Print("---------------------------------\n\n")

	cellKeys := make(map[CellKey]int)
	cells := []Cell{}
	for _, line := range GetContentLines() {
		cell := ParseCell(line)
		cells = append(cells, cell)
		fmt.Printf("  Processing: %v\n", cell)

		for _, key := range GetCellKeys(cell) {
			cellKeys[key]++
		}
	}

	if slices.Contains(os.Args, "--plot") {
		PlotCells(cellKeys)
	}

	var winner Cell
	for _, cell := range cells {
		score := 0
		cellCount := 0
		for _, key := range GetCellKeys(cell) {
			score += cellKeys[key]
			cellCount++
		}

		if score == cellCount {
			winner = cell
		}
	}

	result := 0
	for _, k := range cellKeys {
		if k > 1 {
			result++
		}
	}

	fmt.Print("\n---------------------------------\n")
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Winner: %v\n", winner.Id)
	fmt.Print("---------------------------------\n")

}

func PlotCells(cellKeys map[CellKey]int) {
	fmt.Println()
	for y := 0; y < 1000; y++ {
		fmt.Print("  ")
		for x := 0; x < 1000; x++ {
			value := strconv.Itoa(cellKeys[CellKey{X: x, Y: y}])
			if value == "0" {
				value = "."
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func GetCellKeys(cell Cell) []CellKey {
	result := []CellKey{}

	for x := cell.Left; x < cell.Left+cell.Width; x++ {
		for y := cell.Top; y < cell.Top+cell.Height; y++ {
			result = append(result, CellKey{X: x, Y: y})
		}
	}

	return result
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"Cell { Id = %v, Left = %v, Top = %v, Width = %v, Height = %v }",
		c.Id,
		c.Left,
		c.Top,
		c.Width,
		c.Height)
}

func ParseCell(s string) Cell {
	// format: <#>ID <@> top,left: width<x>height
	// ints: 0 (48) to 9 (57)
	var cell Cell
	var buffer []rune
	fieldsPopulated := 0

	for _, v := range fmt.Sprintf("%v-", s) {
		if v >= 48 && v <= 57 {
			buffer = append(buffer, v)
		} else {
			if len(buffer) > 0 {
				value, err := strconv.Atoi(string(buffer))
				if err != nil {
					log.Fatal("Cannot convert buffer to int")
				}

				switch fieldsPopulated {
				case 0:
					cell.Id = value
				case 1:
					cell.Left = value
				case 2:
					cell.Top = value
				case 3:
					cell.Width = value
				case 4:
					cell.Height = value
				default:
					log.Fatal("Too many cell fields")
				}

				fieldsPopulated++
				buffer = []rune{}
			}
		}
	}

	return cell
}

func getContent() string {
	file := "live"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}

	content, err := os.ReadFile(fmt.Sprintf("./input.%v.txt", file))
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func GetContentLines() []string {
	var result []string
	for _, v := range strings.Split(getContent(), "\n") {
		if temp := strings.Trim(v, "\r\n"); len(temp) > 0 {
			result = append(result, strings.Trim(v, "\r\n"))
		}
	}
	return result
}
