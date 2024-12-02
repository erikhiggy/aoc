package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const sampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func parseInput(sample string) ([][]int, error) {
	var lines []string
	if sample == "" {
		file, err := os.Open("2024/day02/input.txt")
		if err != nil {
			return nil, fmt.Errorf("error opening file: %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}
	} else {
		lines = strings.Split(sample, "\n")
	}

	var result [][]int
	for _, line := range lines {
		values := strings.Fields(line)
		var row []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid number in line: %s", line)
			}
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result, nil
}

func main() {
	in, err := parseInput("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Part 1
	// For each list, we want to detect if it is valid. It can only be valid if the following two constraints are met:
	// 1. The list must be either all increasing or all decreasing.
	// 2. ny two adjacent numbers differ by at least one and at most three.
	validReports := make([]string, len(in))
	// first check if the list is increasing or decreasing
	for idx, list := range in {
		inc, dec := true, true
		for i := 1; i < len(list); i++ {
			if list[i] > list[i-1] {
				dec = false
			} else if list[i] < list[i-1] {
				inc = false
			}
		}

		if !inc && !dec {
			validReports[idx] = "unsafe"
			fmt.Println("invalid due to increasing/decreasing")
			continue
		}
		// next check that the difference between adjacent numbers is at least one and at most three
		for i := 1; i < len(list); i++ {
			if abs(list[i]-list[i-1]) > 3 || abs(list[i]-list[i-1]) < 1 {
				validReports[idx] = "unsafe"
				fmt.Println("invalid due to difference")
				break
			}
		}
		if validReports[idx] == "" {
			validReports[idx] = "safe"
		}
	}

	// sum the number of safe lists
	safeCount := 0
	for _, report := range validReports {
		if report == "safe" {
			safeCount++
		}
	}
	fmt.Println("Safe lists:", safeCount)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
