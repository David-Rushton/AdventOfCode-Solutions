package main

import "strings"

type sequence map[string]int

func (s sequence) GetLen() int {
	var result int

	for code, count := range s {
		result += len(code) * count
	}

	return result
}

func ToSequence(route string) sequence {
	result := sequence{}

	sections := strings.Split(route, "A")
	sections = sections[0 : len(sections)-1]
	for _, section := range sections {
		result[section+"A"]++
	}

	return result
}
