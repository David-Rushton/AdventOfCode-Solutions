package main

type chunk struct {
	count  int
	number int
}

type chunks struct {
	items   []chunk
	current chunk
}

func newChunks() chunks {
	return chunks{
		items: []chunk{},
		current: chunk{
			count:  0,
			number: -1,
		},
	}
}

func (c *chunks) add(next int) {
	if c.current.number == next {
		c.current.count++
		return
	}

	if c.current.count > 0 {
		c.items = append(c.items, c.current)
	}

	c.current = chunk{
		count:  1,
		number: next,
	}
}

func (c *chunks) flush() {
	if c.current.count > 0 {
		c.items = append(c.items, c.current)
	}

	c.current = chunk{
		count:  0,
		number: -1,
	}
}

func (c *chunks) getLen() int {
	var result int

	for _, item := range c.items {
		result += item.count
	}

	return result
}
