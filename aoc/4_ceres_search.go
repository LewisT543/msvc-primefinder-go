package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

var directions = []struct{ dx, dy int }{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // U, D, L, R
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // UL, UR, BL, BR
}

func SolveCeresSearch(input string) string {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, len(lines))
	for i, l := range lines {
		split := strings.Split(l, "")
		matrix[i] = split
	}

	xs, err := findAllXs(matrix)
	if err != nil {
		return "0"
	}

	count := 0
	for _, xCoord := range xs {
		for _, dir := range directions {
			if walkGrid(matrix, xCoord.x, xCoord.y, dir.dx, dir.dy, 1, matrix[xCoord.x][xCoord.y]) {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

type coordinate struct {
	x, y int
}

func findAllXs(matrix [][]string) ([]coordinate, error) {
	xs := make([]coordinate, 0)
	for i, row := range matrix {
		for j, v := range row {
			if v == "X" {
				xs = append(xs, coordinate{i, j})
			}
		}
	}
	if len(xs) == 0 {
		return nil, fmt.Errorf("no X's")
	}
	return xs, nil
}

func walkGrid(matrix [][]string, x, y, dx, dy, step int, current string) bool {
	if current == "XMAS" {
		return true
	}
	if step == 4 {
		return false
	}
	newX, newY := x+dx, y+dy
	if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[newX]) {
		return false
	}
	current += matrix[newX][newY]
	return walkGrid(matrix, newX, newY, dx, dy, step+1, current)
}

const exampleCeresSearchInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const exampleCeresSearchSolution = "18"
