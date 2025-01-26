package main

type tile struct {
	id          int
	cells       [][]string
	connectedTo map[int]bool
}

func (t *tile) addConnection(tileId int) {
	t.connectedTo[tileId] = true
}

func (t *tile) getEdgeIds() []int {
	var edges [][]int

	// top & bottom rows
	for y := 0; y < 10; y += 9 {
		edge := []int{}
		reverseEdge := []int{}

		for x := 0; x < 10; x++ {
			if t.cells[y][x] == "#" {
				edge = append(edge, x)
				reverseEdge = append(reverseEdge, 9-x)
			}
		}

		edges = append(edges, edge)
		edges = append(edges, reverseEdge)
	}

	// left and right columns
	for x := 0; x < 10; x += 9 {
		edge := []int{}
		reverseEdge := []int{}

		for y := 0; y < 10; y++ {
			if t.cells[y][x] == "#" {
				edge = append(edge, y)
				reverseEdge = append(reverseEdge, 9-y)
			}
		}

		edges = append(edges, edge)
		edges = append(edges, reverseEdge)
	}

	// Convert filled in cells to ids
	result := []int{}
	for _, edge := range edges {
		var current int
		for _, loc := range edge {
			current += 1 << loc
		}

		result = append(result, current)
	}

	return result
}
