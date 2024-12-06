package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput() []string {
	var res []string
	file, err := os.Open("2024/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return res
}

func main() {
	// Part 1
	in := parseInput()

	// break the input into the page rules, and the upgrade rules
	var pageRules []string
	var upgradeRuleLines []string
	for _, line := range in {
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			pageRules = append(pageRules, line)
		} else {
			upgradeRuleLines = append(upgradeRuleLines, line)
		}
	}

	// build a reverse map of page rules
	pageRulesMap := make(map[string][]string)
	for _, line := range pageRules {
		parts := strings.Split(line, "|")
		// if the key already exists, append the value to the slice
		if _, ok := pageRulesMap[parts[1]]; ok {
			pageRulesMap[parts[1]] = append(pageRulesMap[parts[1]], parts[0])
		} else {
			// if the key does not exist, create a new slice with the value
			pageRulesMap[parts[1]] = []string{parts[0]}
		}
	}

	var upgradeRules [][]string
	for _, line := range upgradeRuleLines {
		var strs []string
		for _, rule := range strings.Split(line, ",") {
			strs = append(strs, rule)
		}
		upgradeRules = append(upgradeRules, strs)
	}

	safeUpdates, unsafeUpdates := findUpdates(upgradeRules, pageRulesMap)

	// part 1
	part1Sum := getMiddleSum(safeUpdates)
	fmt.Println(part1Sum)

	// part 2
	var unsafeUpdatesMatrix [][]string
	for _, line := range unsafeUpdates {
		var s []string
		for _, rule := range strings.Split(line, ",") {
			s = append(s, rule)
		}
		unsafeUpdatesMatrix = append(unsafeUpdatesMatrix, s)
	}

	// fmt.Println(unsafeUpdatesList)
	fixUnsafeUpdates(unsafeUpdatesMatrix, pageRulesMap)
	unsafeUpdatesList := make([]string, len(unsafeUpdatesMatrix))
	for i, line := range unsafeUpdatesMatrix {
		unsafeUpdatesList[i] = strings.Join(line, ",")
	}
	part2Sum := getMiddleSum(unsafeUpdatesList)
	fmt.Println(part2Sum)
}

func fixUnsafeUpdates(unsafeUpdates [][]string, pageRulesMap map[string][]string) {
	for i := 0; i < len(unsafeUpdates); i++ {
		for j := 0; j < len(unsafeUpdates[i])-1; j++ {
			for k := j + 1; k < len(unsafeUpdates[i]); k++ {
				if isUnsafe(unsafeUpdates[i], j, k, pageRulesMap) {
					swap(unsafeUpdates[i], j, k)
				}
			}
		}
	}
}

func isUnsafe(i []string, j int, k int, rulesMap map[string][]string) bool {
	for _, tgt := range rulesMap[i[j]] {
		if tgt == i[k] {
			return true
		}
	}
	return false
}

func swap(arr []string, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func getMiddleSum(arr []string) int {
	middleSum := 0
	var middles []int
	// grab the middle number from each unsage update line
	for _, line := range arr {
		parts := strings.Split(line, ",")
		asInt, err := strconv.Atoi(parts[len(parts)/2])
		if err != nil {
			log.Fatal(err)
		}
		middles = append(middles, asInt)
	}

	// sum the middle numbers
	for _, middle := range middles {
		middleSum += middle
	}
	return middleSum
}

func findUpdates(upgradeRules [][]string, pageRulesMap map[string][]string) ([]string, []string) {
	var safeUpdates []string
	var unsafeUpdates []string
	for i := 0; i < len(upgradeRules); i++ {
		if isValid(upgradeRules[i], pageRulesMap) {
			safeUpdates = append(safeUpdates, strings.Join(upgradeRules[i], ","))
		} else {
			unsafeUpdates = append(unsafeUpdates, strings.Join(upgradeRules[i], ","))
		}
	}
	return safeUpdates, unsafeUpdates
}

func isValid(update []string, pageRulesMap map[string][]string) bool {
	for j := 0; j < len(update)-1; j++ {
		tgtList := pageRulesMap[update[j]]
		for k := j + 1; k < len(update); k++ {
			for _, tgt := range tgtList {
				if tgt == update[k] {
					return false
				}
			}
		}
	}
	return true
}
