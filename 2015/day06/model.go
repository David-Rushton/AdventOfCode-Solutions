package main

import "fmt"

type instructionAction int

func (ia instructionAction) String() string {
	var result string

	switch ia {
	case turnOn:
		result = "turn on"
	case turnOff:
		result = "turn off"
	case toggle:
		result = "toggle"
	}

	return result
}

const (
	turnOn instructionAction = iota
	turnOff
	toggle
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type instruction struct {
	action      instructionAction
	topLeft     point
	bottomRight point
}
