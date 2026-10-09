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

func main() {
	parsed := parseInput("input.txt")
	ranges := strings.SplitSeq(parsed[0], ",")

	sum := 0
	for r := range ranges {
		startingId := strings.Split(r, "-")[0]
		endingId := strings.Split(r, "-")[1]

		for i := mustInt(startingId); i < mustInt(endingId); i++ {
			numLength := countDigitsLoop(i)
			// we only care about even-lengthed numbers
			if numLength%2 == 0 {
				// numLength == 2 -> 10 == 10^1
				// numLength == 4 -> 100 == 10^2
				// numLength == 6 -> 1000 == 10^3

				// Generalization: i % 10^(numLength / 2)
				// This gives us the second half of the number.
				// The first half can be retrieved though simple division.
				divisor := int(math.Pow10(numLength / 2))
				firstHalf := i / divisor
				secondHalf := i % divisor
				if firstHalf == secondHalf {
					sum += i
				}
			}
		}
	}
	fmt.Println("part 1: ", sum)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
