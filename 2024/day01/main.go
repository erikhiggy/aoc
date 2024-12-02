package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput() ([]int, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		left, err1 := strconv.Atoi(values[0])
		right, err2 := strconv.Atoi(values[1])

		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %s", line)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return leftList, rightList, nil
}

func main() {
	left, right, err := parseInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort slices and compare left to right
	slices.Sort(left)
	slices.Sort(right)

	listLen := len(left)
	// Part 1
	sum := 0.0
	for i := 0; i < listLen; i++ {
		sum += math.Abs(float64(left[i]) - float64(right[i]))
	}

	fmt.Println("Sum:", int(sum))

	// Part 2
	simScore := 0
	numMap := make(map[int]int)
	for _, r := range right {
		numMap[r] += 1
	}

	for _, l := range left {
		simScore += numMap[l] * l
	}

	fmt.Println("Similarity score:", simScore)
}
