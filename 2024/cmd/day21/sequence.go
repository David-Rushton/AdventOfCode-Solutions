package main

type sequence map[string]int

func (s sequence) GetLen() int {
	var result int

	for code, count := range s {
		result += len(code) * count
	}

	return result
}
