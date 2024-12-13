package main

import (
	"fmt"
	"log"
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
	fmt.Println("---  Day 13: Claw Contraption ---")
	fmt.Println()

	prizeBuilder := 0
	if aoc.Star == aoc.StarTwo {
		prizeBuilder = 10000000000000
	}

	clawMachines := parse(aoc.Input, prizeBuilder)
	total := 0

	for _, machine := range clawMachines {
		subTotal := play(machine)
		fmt.Printf("  Result of looking for % 6d x % 6d = %d\n", machine.prize.x, machine.prize.y, subTotal)
		total += subTotal
	}

	fmt.Println()
	fmt.Printf("Result: %d\n", total)
}

func play(machine clawMachine) (tokens int) {
	px := machine.prize.x
	py := machine.prize.y

	ax := machine.buttonA.x
	ay := machine.buttonA.y

	bx := machine.buttonB.x
	by := machine.buttonB.y

	a := (px*by - py*bx) / (ax*by - ay*bx)
	b := (py - a*ay) / by

	tx := a*ax + b*bx
	ty := a*ay + b*by

	if tx == px && ty == py {
		return a*3 + b
	}

	return 0
}

func parse(input []string, prizeBuilder int) []clawMachine {
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
			result[current].prize.x += prizeBuilder
			result[current].prize.y += prizeBuilder
			continue
		}
	}

	return result
}
