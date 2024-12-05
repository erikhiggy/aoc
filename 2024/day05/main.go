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

	// part 1
	var upgradeRules [][]string
	for _, line := range upgradeRuleLines {
		var strs []string
		for _, rule := range strings.Split(line, ",") {
			strs = append(strs, rule)
		}
		upgradeRules = append(upgradeRules, strs)
	}

	var safeUpdates []string
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
		}
	}
	fmt.Println(safeUpdates)
	middleSum := 0
	var middles []int
	// grab the middle number from each unsage update line
	for _, line := range safeUpdates {
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
	fmt.Println(middleSum)
}
