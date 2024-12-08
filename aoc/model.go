package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type AOCProblem struct {
	Filename         string
	Title            string
	ShortDescription string
	Day              int
	SolverFn         SolverFn
}

func NewAOCProblem(filename string, shortDescription string, solverFn SolverFn) (*AOCProblem, error) {
	titleParts := strings.Split(filename, "_")
	if len(titleParts) < 2 {
		return nil, fmt.Errorf("invalid filename format")
	}

	dayPart := titleParts[0]
	day, err := strconv.ParseInt(dayPart, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse day from filename: %v", err)
	}

	title, err := formatTitle(titleParts)
	if err != nil {
		return nil, fmt.Errorf("failed to format title: %v", err)
	}

	return &AOCProblem{
		Filename:         filename,
		Title:            title,
		ShortDescription: shortDescription,
		Day:              int(day),
		SolverFn:         solverFn,
	}, nil
}

func formatTitle(filenameParts []string) (string, error) {
	titleParts := filenameParts[1:]

	var builder strings.Builder
	for _, part := range titleParts {
		builder.WriteString(strings.ToUpper(string(part[0])) + strings.ToLower(part[1:]) + " ")
	}
	result := strings.TrimSpace(builder.String())
	return result, nil
}
