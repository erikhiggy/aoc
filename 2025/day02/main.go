package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func countDigitsLoop(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}

	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func isEntirelyRepeating(s string) bool {
	n := len(s)

	// should never happen, but check anyway
	if n < 2 {
		return false
	}

	// double it and remove the first and last char
	modified := s[1:] + s[:n-1]

	return strings.Contains(modified, s)
}

func main() {
	parsed := parseInput("input.txt")
	ranges := strings.SplitSeq(parsed[0], ",")

	part1Sum := 0
	part2Sum := 0
	for r := range ranges {
		startingId := strings.Split(r, "-")[0]
		endingId := strings.Split(r, "-")[1]

		for i := mustInt(startingId); i <= mustInt(endingId); i++ {
			part1Sum = part1(i, part1Sum)
			part2Sum = part2(strconv.Itoa(i), part2Sum)
		}
	}
	fmt.Println("part 1: ", part1Sum)
	fmt.Println("part 2: ", part2Sum)
}

func part1(id int, sum int) int {
	numLength := countDigitsLoop(id)
	// we only care about even-lengthed numbers
	if numLength%2 == 0 {
		// numLength == 2 -> 10 == 10^1
		// numLength == 4 -> 100 == 10^2
		// numLength == 6 -> 1000 == 10^3

		// Generalization: i % 10^(numLength / 2)
		// This gives us the second half of the number.
		// The first half can be retrieved though simple division.
		divisor := int(math.Pow10(numLength / 2))
		firstHalf := id / divisor
		secondHalf := id % divisor
		if firstHalf == secondHalf {
			sum += id
		}
	}
	return sum
}

func part2(id string, sum int) int {
	if isEntirelyRepeating(id) {
		sum += mustInt(id)
	}
	return sum
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
