package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/iostr"
)

func main() {
	iostr.Outln("--- Day 2: Red-Nosed Reports ---")
	iostr.Outln("--------------------------------")

	var safeReports int
	for _, report := range aoc.Input {
		if isSafeReport(toReport(report), aoc.Star == aoc.StarTwo) {
			iostr.Verbosef("Safe report found: %s\n", report)
			safeReports++
		}
	}

	iostr.Outln("--------------------------------")
	iostr.Outf("--- Safe Reports: %d ---", safeReports)
}

func toReport(report string) []int {
	result := []int{}

	for _, element := range strings.Split(report, " ") {
		n, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			log.Fatalln("Cannot convert to int.")
		}

		result = append(result, int(n))
	}

	return result
}

func isSafeReport(report []int, enableProblemDampener bool) bool {
	if enableProblemDampener {
		return isSafeReportWithProblemDampener(report)
	}

	return isSafeReportWithoutProblemDampener(report)
}

func isSafeReportWithProblemDampener(report []int) bool {
	if isSafeReportWithoutProblemDampener(report) {
		return true
	}

	for i := 0; i < len(report); i++ {

		if isSafeReportWithoutProblemDampener(removeAt(report, i)) {
			return true
		}
	}

	return false
}

func isSafeReportWithoutProblemDampener(report []int) bool {
	var increasing int
	var decreasing int

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		absDiff := int(math.Abs(float64(diff)))

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if diff < 0 {
			decreasing++
		}

		if diff > 0 {
			increasing++
		}

		if increasing > 0 && decreasing > 0 {
			return false
		}
	}

	return true
}

func removeAt(s []int, i int) []int {
	result := []int{}

	for j := 0; j < len(s); j++ {
		if j == i {
			continue
		}

		result = append(result, s[j])
	}

	return result
}
