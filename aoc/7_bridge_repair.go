package aoc

import (
	"strconv"
	"strings"
)

var ops = [2]string{"+", "*"}

func SolveBridgeRepair(input string) string {
	rows := strings.Split(input, "\r\n")
	var sum int64
	for _, row := range rows {
		r := strings.Split(row, ": ")
		target, _ := strconv.ParseInt(r[0], 10, 64)

		numStr := strings.Split(r[1], " ")
		nums := make([]int64, len(numStr))
		for i, num := range numStr {
			nums[i], _ = strconv.ParseInt(num, 10, 64)
		}
		if testPermutations(target, nums) {
			sum += target
		}
	}

	return strconv.Itoa(int(sum))
}

func testPermutations(target int64, nums []int64) bool {
	return evaluate(0, nums, nums[0], target)
}

func evaluate(i int, nums []int64, current int64, target int64) bool {
	if i == len(nums)-1 {
		return current == target
	}
	if evaluate(i+1, nums, current+nums[i+1], target) {
		return true
	}
	if evaluate(i+1, nums, current*nums[i+1], target) {
		return true
	}
	return false
}

const exampleBridgeBuildersInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

const exampleBridgeBuildersSolution = "3749"
