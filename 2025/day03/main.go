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

func findTheBiggestVoltage(digits []int, length int) int {
	chosen := ""

	for range length {
		max := 0
		index := 0

		for j := 0; j <= len(digits)+len(chosen)-length; j++ {
			if digits[j] > max {
				max = digits[j]
				index = j
			}
		}

		chosen += strconv.Itoa(max)
		digits = digits[index+1:]
	}

	number := mustInt(chosen)
	return number
}

func main() {
	parsed := parseInput("input.txt")

	var part1Sum, part2Sum int
	for _, l := range parsed {
		var intList []int
		for _, r := range l {
			intList = append(intList, mustInt(string(r)))
		}
		part1Sum += findTheBiggestVoltage(intList, 2)
		part2Sum += findTheBiggestVoltage(intList, 12)
	}

	fmt.Println("part 1:", part1Sum)
	fmt.Println("part 2:", part2Sum)
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
