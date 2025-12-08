package aoc

import (
	"math"
	"strconv"
)

func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
