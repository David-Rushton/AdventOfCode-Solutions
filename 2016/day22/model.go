package main

import "math"

type point struct {
	x int
	y int
}

func (p *point) isConnected(other point) bool {
	return (p.x == other.x && math.Abs(float64(p.y-other.y)) == 1) ||
		(p.y == other.y && math.Abs(float64(p.x-other.x)) == 1)
}

type node struct {
	address   point
	size      int
	used      int
	available int
}

type connectedNodes struct {
	from  node
	to    node
	hash1 int64
	hash2 int64
}

func (cn *connectedNodes) move() (left, right node) {
	left = node{
		address:   cn.from.address,
		size:      cn.from.size,
		used:      0,
		available: cn.from.size,
	}

	right = node{
		address:   cn.to.address,
		size:      cn.to.size,
		used:      cn.to.used + cn.from.used,
		available: cn.to.size - cn.to.used - cn.from.used,
	}

	return
}
