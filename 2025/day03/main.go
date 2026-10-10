package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return lines
}

func part1(s []int) int {
	// first pass: find the largest number in s up to len(s)-1 => O(n)
	firstMax, firstMaxIdx := maxWithIndex(s[0 : len(s)-1])

	// second pass: starting at index of firstMax, find the secondMax => O(n)
	secondMax, _ := maxWithIndex(s[firstMaxIdx+1:])

	return (firstMax * 10) + secondMax
}

func main() {
	parsed := parseInput("input.txt")

	part1Sum := 0
	for _, l := range parsed {
		// fmt.Println(l)
		var intList []int
		for _, r := range l {
			intList = append(intList, mustInt(string(r)))
		}
		part1Sum += part1(intList)
	}

	fmt.Println("part 1:", part1Sum)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}

// maxWithIndex returns the largest number in an int slice
// with its index
func maxWithIndex(a []int) (int, int) {
	max := 0
	idx := 0
	for i, x := range a {
		if x > max {
			max = x
			idx = i
		}
	}
	return max, idx
}
