package main

type elf struct {
	id       int
	previous *elf
	next     *elf
}

func (e *elf) isWinner() bool {
	return e.id == e.next.id
}

func (e *elf) skip(n int) *elf {
	switch {
	case n < 0:
		return e.previous.skip(n + 1)
	case n > 0:
		return e.next.skip(n - 1)
	default:
		return e
	}
}
