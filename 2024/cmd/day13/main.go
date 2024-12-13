package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type point struct {
	x int
	y int
}

type clawMachine struct {
	buttonA point
	buttonB point
	prize   point
}

func main() {
	fmt.Println("---  Day 13: ??? ---")
	fmt.Println()

	prizeMultiplier := 1
	if aoc.Star == aoc.StarTwo {
		prizeMultiplier = 10000000000000
	}

	clawMachines := parse(aoc.Input, prizeMultiplier)
	total := 0

	for _, machine := range clawMachines {
		subTotal := play(machine)
		fmt.Printf("  Result of looking for % 6d x % 6d = %d\n", machine.prize.x, machine.prize.y, subTotal)
		total += subTotal
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", total)
}

type machineState struct {
	x      int
	y      int
	round  int
	tokens int
}

type stake struct {
	x      int
	y      int
	tokens int
}

func play(machine clawMachine) (tokens int) {
	queue := []machineState{{}}
	stakes := []stake{
		{
			x:      machine.buttonA.x,
			y:      machine.buttonA.y,
			tokens: 3,
		},
		{
			x:      machine.buttonB.x,
			y:      machine.buttonB.y,
			tokens: 1,
		},
	}

	cheapestWin := math.MaxInt
	visited := make(map[machineState]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, stake := range stakes {
			candidate := machineState{
				current.x + stake.x,
				current.y + stake.y,
				current.round + 1,
				current.tokens + stake.tokens,
			}

			if _, found := visited[candidate]; found {
				continue
			}
			visited[candidate]++

			if stake.tokens > cheapestWin {
				continue
			}

			if candidate.x < machine.prize.x && candidate.y < machine.prize.y {
				queue = append(queue, candidate)
			}

			if candidate.x == machine.prize.x && candidate.y == machine.prize.y {
				if candidate.tokens < cheapestWin {
					cheapestWin = candidate.tokens
				}
			}
		}
	}

	if cheapestWin == math.MaxInt {
		cheapestWin = 0
	}

	return cheapestWin
}

func parse(input []string, prizeMultiplier int) []clawMachine {
	result := []clawMachine{}

	getPoint := func(value string) point {
		var x, y int

		positions := strings.Replace(strings.Replace(strings.Replace(value, "X", "", 1), "Y", "", 1), "=", "", 2)
		elements := strings.Split(strings.Split(positions, ":")[1], ", ")
		for i, element := range elements {
			num, err := strconv.ParseInt(strings.Trim(element, " "), 10, 64)
			if err != nil {
				log.Fatalf("Cannot convert %s to a number", element)
			}

			switch i {
			case 0:
				x = int(num)
			case 1:
				y = int(num)
			}
		}

		return point{x, y}
	}

	current := -1
	for _, line := range input {
		if strings.HasPrefix(line, "Button A") {
			result = append(result, clawMachine{})
			current++

			result[current].buttonA = getPoint(line)
			continue
		}

		if strings.HasPrefix(line, "Button B") {
			result[current].buttonB = getPoint(line)
			continue
		}

		if strings.HasPrefix(line, "Prize") {
			result[current].prize = getPoint(line)
			result[current].prize.x *= prizeMultiplier
			result[current].prize.y *= prizeMultiplier
			continue
		}
	}

	return result
}
