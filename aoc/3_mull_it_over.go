package aoc

import "fmt"

func SolveMullItOver(input string) string {
	res := (620 * 236) + (589 * 126) + (260 * 42) + (335 * 250)
	fmt.Println(res)
	return "0"
}

const exampleMullItOverInput = "}mul(620,236)where()*@}!&[mul(589,126)]&^]mul(260,42)when()mul(603[%^where() when()$ ?{/^*mul(335,250)>,@!"
const exampleMullItOverSolution = "315204"
