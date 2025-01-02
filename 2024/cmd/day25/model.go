package main

type keyLockBase struct {
	columns []int
}

type key keyLockBase
type lock keyLockBase
