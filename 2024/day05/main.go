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
	//fmt.Println(pageRules)

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

	var unsafeUpdatesList []string
	for i := 0; i < len(unsafeUpdatesMatrix); i++ {
		for j := 0; j < len(unsafeUpdatesMatrix[i])-1; j++ {
			tgtList := pageRulesMap[unsafeUpdatesMatrix[i][j]]
			for k := j + 1; k < len(unsafeUpdatesMatrix[i]); k++ {
				for _, tgt := range tgtList {
					if tgt == unsafeUpdatesMatrix[i][k] {
						unsafeUpdatesMatrix[i][j], unsafeUpdatesMatrix[i][k] = unsafeUpdatesMatrix[i][k], unsafeUpdatesMatrix[i][j]
					}
				}
			}
		}
		unsafeUpdatesList = append(unsafeUpdatesList, strings.Join(unsafeUpdatesMatrix[i], ","))
	}

	// fmt.Println(unsafeUpdatesList)
	part2Sum := getMiddleSum(unsafeUpdatesList)
	fmt.Println(part2Sum)
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
		unsafe := false
		for j := 0; j < len(upgradeRules[i])-1; j++ {
			tgtList := pageRulesMap[upgradeRules[i][j]]
			for k := j + 1; k < len(upgradeRules[i]); k++ {
				for _, tgt := range tgtList {
					if tgt == upgradeRules[i][k] {
						// if we find a match, add the upgrade rule to the list of unsafe updates and break
						unsafe = true
						break
					}
				}
				if unsafe {
					break
				}
			}
			if unsafe {
				break
			}
		}
		if !unsafe {
			safeUpdates = append(safeUpdates, strings.Join(upgradeRules[i], ","))
		} else {
			unsafeUpdates = append(unsafeUpdates, strings.Join(upgradeRules[i], ","))
		}
	}
	return safeUpdates, unsafeUpdates
}
