package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2025/internal/aoc"
)

type machine struct {
	buttons   buttons
	joltage   joltage
	indicator indicator
}

func (m machine) String() string {
	return fmt.Sprintf("%s %s %s", m.indicator, m.buttons, m.joltage)
}

type indicator []bool

func (i indicator) value() uint32 {
	var result uint32
	for i, v := range i {
		if v {
			result += 1 << i
		}
	}
	return result
}

func (i indicator) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for _, v := range i {
		switch v {
		case true:
			sb.WriteString("#")
		default:
			sb.WriteString(".")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

type buttons []button

func (b buttons) String() string {
	sb := strings.Builder{}
	sb.WriteString("(")
	for _, button := range b {
		if sb.Len() > 1 {
			sb.WriteString(" ")
		}
		sb.WriteString(button.String())
	}
	sb.WriteString(")")
	return sb.String()
}

type button struct {
	id     int
	values []int
}

func (b button) value() uint32 {
	var result uint32
	for _, v := range b.values {
		result += 1 << v
	}
	return result
}

func (b button) String() string {
	sb := strings.Builder{}
	sb.WriteString("(")
	for i, value := range b.values {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.FormatInt(int64(value), 10))
	}
	sb.WriteString(")")
	return sb.String()
}

type joltage []int

func (j joltage) String() string {
	sb := strings.Builder{}
	sb.WriteString("{")
	for i, v := range j {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.FormatInt(int64(v), 10))
	}
	sb.WriteString("}")
	return sb.String()
}

func main() {
	fmt.Println("--- Day 10: Factory ---")
	fmt.Println()

	result := 0
	machines := parse(aoc.Input)
	for _, m := range machines {
		result += solve(m)
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", result)
}

func solve(m machine) int {
	q := []uint32{}
	visited := map[uint32]bool{}

	for _, button := range m.buttons {
		q = append(q, button.value())
	}

	for round := 1; round < 1000; round++ {
		nextQ := []uint32{}

		for _, value := range q {
			if value == m.indicator.value() {
				fmt.Printf(" - New best found for %s in %d presses", m.indicator, round)
				return round
			}

			if visited[value] {
				continue
			}
			visited[value] = true

			for _, button := range m.buttons {
				next := value ^ button.value()
				nextQ = append(nextQ, next)
			}
		}

		q = nextQ
	}

	panic(fmt.Sprintf("cannot solve machine %s", m.String()))
}

func parse(input []string) []machine {
	var result []machine
	var idSeed int

	for _, s := range input {
		elements := strings.Split(s, " ")
		current := machine{}
		for _, element := range elements {
			switch element[0] {
			case '[':
				for _, r := range element[1 : len(element)-1] {
					switch r {
					case '#':
						current.indicator = append(current.indicator, true)
					case '.':
						current.indicator = append(current.indicator, false)
					default:
						panic(fmt.Sprintf("cannot parse indicator %s", string(r)))
					}
				}

			case '(':
				values := strings.Split(strings.ReplaceAll(strings.ReplaceAll(element, "(", ""), ")", ""), ",")
				next := button{id: idSeed}
				idSeed++
				for _, v := range values {
					next.values = append(next.values, aoc.ToInt(v))
				}
				current.buttons = append(current.buttons, next)

			case '{':
				values := strings.Split(strings.ReplaceAll(strings.ReplaceAll(element, "{", ""), "}", ""), ",")
				for _, v := range values {
					current.joltage = append(current.joltage, aoc.ToInt(v))
				}

			default:
				panic("cannot parse input")
			}

		}
		result = append(result, current)
	}

	return result
}
