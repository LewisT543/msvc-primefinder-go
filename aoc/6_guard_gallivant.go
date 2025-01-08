package aoc

import (
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type Direction struct {
	name string
	dx   int
	dy   int
}

var (
	Up    = Direction{"Up", -1, 0}
	Right = Direction{"Right", 0, 1}
	Down  = Direction{"Down", 1, 0}
	Left  = Direction{"Left", 0, -1}
)

var directions2 = []Direction{Up, Right, Down, Left}

func (d Direction) turnRight() Direction {
	for i, dir := range directions2 {
		if d == dir {
			return directions2[(i+1)%len(directions2)]
		}
	}
	panic("unknown direction")
}

func SolveGuardGallivant(input string) string {
	rows := strings.Split(input, "\r\n")
	grid := make([][]string, len(rows))

	start := position{0, 0}

	for i, row := range rows {
		grid[i] = strings.Split(row, "")
		if j := strings.Index(row, "^"); j != -1 {
			start = position{i, j}
		}
	}

	if start.x == -1 {
		panic("starting position not found")
	}

	visitedPositions := walkGridGG(grid, start, Up)

	return strconv.Itoa(len(visitedPositions))
}

func walkGridGG(grid [][]string, start position, d Direction) map[position]bool {
	visitedPositions := make(map[position]bool)
	current := start

	for {
		visitedPositions[current] = true
		nextX, nextY := current.x+d.dx, current.y+d.dy
		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[nextX]) {
			break
		}
		if grid[nextX][nextY] == "#" {
			d = d.turnRight()
		} else {
			current = position{nextX, nextY}
		}
	}

	return visitedPositions
}

const exampleGuardGallivantInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const exampleGuardGallivantSolution = "41"
