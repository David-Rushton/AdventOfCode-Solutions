package main

type ipRange struct {
	from  int
	until int // inclusive
}

func (ip *ipRange) count() int {
	return ip.until - ip.from + 1
}
