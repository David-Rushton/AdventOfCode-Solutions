package main

type point struct {
	x int
	y int
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
