package main

type elf struct {
	id       int
	previous *elf
	next     *elf
}

func (e *elf) skip(n int) *elf {
	result := e

	for i := 0; i < n; i++ {
		result = result.next
	}

	return result
}

func (e *elf) isWinner() bool {
	return e.id == e.next.id
}
