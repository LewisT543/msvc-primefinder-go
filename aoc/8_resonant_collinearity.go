package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func SolveResonantCollinearity(input string) string {
	antennas := make(map[rune][]Point)
	rows := strings.Split(input, "\r\n")
	width, height := len(rows[0]), len(rows)

	for _, row := range rows {
		fmt.Println(row)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := rune(rows[y][x])
			if isValidAntenna(char) {
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}
	fmt.Println(antennas)

	antinodeSet := make(map[Point]struct{})
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				antinodes := getAntinodes(positions[i], positions[j])

				for _, antinode := range antinodes {
					if isInBounds(antinode.x, antinode.y, width, height) {
						antinodeSet[antinode] = struct{}{}
					}
				}
			}
		}

		for _, antenna := range positions {
			if _, exists := antinodeSet[antenna]; exists {
				antinodeSet[antenna] = struct{}{}
			}
		}
	}

	fmt.Println(antinodeSet)
	fmt.Println(len(antinodeSet))

	return strconv.Itoa(len(antinodeSet))
}

func isInBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func getAntinodes(antenna1, antenna2 Point) []Point {
	var antinodes []Point

	dx := antenna2.x - antenna1.x
	dy := antenna2.y - antenna1.y

	antinodes = append(antinodes, Point{
		x: antenna1.x - dx,
		y: antenna1.y - dy,
	})

	antinodes = append(antinodes, Point{
		x: antenna2.x + dx,
		y: antenna2.y + dy,
	})
	fmt.Printf("ANTINODES: %v <> POINTS: [%v, %v]\n", antinodes, antenna1, antenna2)
	return antinodes
}

func isValidAntenna(char rune) bool {
	if char == '.' {
		return false
	}
	asciiValue := int(char)
	return (asciiValue >= 48 && asciiValue <= 57) ||
		(asciiValue >= 65 && asciiValue <= 90) ||
		(asciiValue >= 97 && asciiValue <= 122)
}

const exampleResonantCollinearityInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

const exampleResonantCollinearitySolution = "14"
