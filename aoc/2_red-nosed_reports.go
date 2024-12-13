package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

// Find the safe reports - a report is a line. A report is safe if it is A). ascending or descending, B). 0 < dif(a - b...) < 3 for each level where [a, b, c...]

func SolveRedNosedReports(input string) string {
	reports := strings.Split(input, "\r\n")
	counter := 0
	for _, r := range reports {
		reportStrs := strings.Split(r, " ")
		report := make([]int, len(reportStrs))

		for i, l := range reportStrs {
			level, err := strconv.ParseInt(l, 10, 0)
			if err != nil {
				fmt.Printf("error parsing int: %v", err)
			}
			var levelInt = int(level)
			report[i] = levelInt
		}

		if isSafe(report) {
			fmt.Printf("report is safe: %v\n", report)
			counter++
		}
	}

	return strconv.Itoa(counter)
}

func isSafe(report []int) bool {
	current := report[0]
	rest := report[1:]
	isAscending := current < rest[0]

	for _, v := range rest {
		d := dif(current, v)
		if d < 1 || d > 3 {
			fmt.Printf("not gradual report: %v - with dif: %d\n", report, d)
			return false
		}
		if isAscending {
			if v < current {
				fmt.Printf("ascending report is not ascending: %v\n", report)
				return false
			}
		} else {
			if v > current {
				fmt.Printf("descending report is not descending: %v\n", report)
				return false
			}
		}
		current = v
	}
	return true
}

func dif(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

const redNosedReportsExampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

const redNosedReportsExampleSolution = "2"
