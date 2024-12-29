package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePrintQueue(input string) string {
	sections := strings.Split(input, "\r\n\r\n")
	rules := strings.Split(sections[0], "\r\n")
	updates := strings.Split(sections[1], "\r\n")

	rulesMap := make(map[string]*[]string)
	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		if len(splitRule) == 2 {
			if _, exists := rulesMap[splitRule[0]]; !exists {
				rulesMap[splitRule[0]] = &[]string{}
			}
			*rulesMap[splitRule[0]] = append(*rulesMap[splitRule[0]], splitRule[1])
		}
	}

	sum := 0

	for _, update := range updates {
		updatePages := strings.Split(update, ",")
		relevantRules := getRelevantRules(updatePages, rulesMap)
		if updateIsValid(updatePages, relevantRules) {
			sum += getMiddleNumber(updatePages)
		}
	}

	return strconv.Itoa(sum)
}

func getRelevantRules(updatePages []string, rulesMap map[string]*[]string) []string {
	validRules := make([]string, 0)
	updatePagesMap := make(map[string]bool)
	for _, page := range updatePages {
		updatePagesMap[page] = true
	}

	for first, possibleSeconds := range rulesMap {
		if updatePagesMap[first] {
			for _, second := range *possibleSeconds {
				if updatePagesMap[second] {
					validRules = append(validRules, fmt.Sprintf("%s|%s", first, second))
				}
			}
		}
	}
	return validRules
}

func updateIsValid(updatePages []string, rules []string) bool {
	updateMap := make(map[string]int)
	for i, v := range updatePages {
		updateMap[v] = i
	}

	for _, rule := range rules {
		split := strings.Split(rule, "|")
		first, second := split[0], split[1]

		firstInd, firstExists := updateMap[first]
		if !firstExists {
			return false
		}

		secondInd, secondExists := updateMap[second]
		if !secondExists || secondInd <= firstInd {
			return false
		}
	}
	return true
}

func getMiddleNumber(updatePages []string) int {
	middleInd := len(updatePages) / 2
	num, _ := strconv.ParseInt(updatePages[middleInd], 10, 0)
	return int(num)
}

const examplePrintQueueInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

const examplePrintQueueSolution = "143"
