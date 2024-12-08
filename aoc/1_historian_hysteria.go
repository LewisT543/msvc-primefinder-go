package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolveHistorianHysteria(input string) (string, error) {
	var lefts, rights []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lineParts := strings.Split(line, "   ")
		leftInt, _ := strconv.ParseInt(lineParts[0], 10, 0)
		rightInt, _ := strconv.ParseInt(lineParts[1], 10, 0)
		lefts = append(lefts, int(leftInt))
		rights = append(rights, int(rightInt))
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	var finalDif = 0
	for i := 0; i < len(lefts); i++ {
		finalDif += abs(lefts[i] - rights[i])
	}

	return fmt.Sprintf("%d", finalDif), nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const historianHysteriaExampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

const historianHysteriaExampleSolution = "11"
