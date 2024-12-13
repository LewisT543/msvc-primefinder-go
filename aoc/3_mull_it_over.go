package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// scan for mul(a int, b int) => accumulate the result of a * b for all mul(a, b)

func SolveMullItOver(input string) string {
	pattern := `mul\(\d+,\d+\)`
	re, _ := regexp.Compile(pattern)
	matches := re.FindAllString(input, -1)

	var sum int64 = 0
	for _, match := range matches {
		fmt.Println(match)
		split := strings.Split(match, ",")

		first := strings.Split(split[0], "(")[1]
		second := split[1][:len(split[1])-1]

		intA, _ := strconv.ParseInt(first, 10, 0)
		intB, _ := strconv.ParseInt(second, 10, 0)

		sum += intA * intB
	}

	return strconv.Itoa(int(sum))
}

const exampleMullItOverInput = "}mul(620,236)where()*@}!&[mul(589,126)]&^]mul(260,42)when()mul(603[%^where() when()$ ?{/^*mul(335,250)>,@!"
const exampleMullItOverSolution = "315204"
