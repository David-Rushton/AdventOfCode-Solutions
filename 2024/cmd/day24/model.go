package main

import (
	"log"
)

type gateOperator int

const (
	operatorAnd gateOperator = iota
	operatorXOr
	operatorOr
)

type gate struct {
	leftIn   *wire
	rightIn  *wire
	Out      *wire
	operator gateOperator
}

func (g *gate) hasValue() bool {
	return g.leftIn.hasValue && g.rightIn.hasValue
}

func (g *gate) getValue() bool {
	if g.hasValue() {
		switch g.operator {
		case operatorAnd:
			return g.leftIn.getValue() && g.rightIn.getValue()
		case operatorXOr:
			return g.leftIn.getValue() != g.rightIn.getValue()
		case operatorOr:
			return g.leftIn.getValue() || g.rightIn.getValue()
		}
	}

	panic("Cannot read value for gate.")
}

type wireUpdate struct {
	wire  *wire
	value bool
}

type wire struct {
	name     string
	value    bool
	hasValue bool
}

func (w *wire) getValue() bool {
	if !w.hasValue {
		log.Fatalf("Cannot get value for wire %s.  Not initialized.", w.name)
	}

	return w.value
}

func (w *wire) setValue(value bool) {
	w.value = value
	w.hasValue = true
}
