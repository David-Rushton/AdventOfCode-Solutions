package main

type elf struct {
	id    int
	next  *elf
	value int
}

func (e *elf) setNext(next *elf) {
	e.next = next
}

func (e *elf) isWinner() bool {
	return e.id == e.next.id
}
