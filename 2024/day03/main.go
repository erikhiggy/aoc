package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const SampleInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func parseInput(sample string) (string, error) {
	// If sample is given, use it instead of the file
	if sample != "" {
		return sample, nil
	}

	// Open the file
	file, err := os.Open("2024/day03/input.txt")
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var in string
	for scanner.Scan() {
		in += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return in, nil
}

func extractAndCompute(s string) int {
	// Extract the integers from the string
	var num1AsInt, num2AsInt int
	nums := strings.Split(s[4:len(s)-1], ",")
	num1, num2 := nums[0], nums[1]
	num1AsInt, err := strconv.Atoi(num1)
	if err != nil {
		println("Error:", err)
		return 0
	}
	num2AsInt, err = strconv.Atoi(num2)
	if err != nil {
		println("Error:", err)
		return 0
	}

	// Compute the product
	return num1AsInt * num2AsInt
}

func main() {
	in, err := parseInput("")
	if err != nil {
		println("Error:", err)
		return
	}

	// Part 1
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(in, -1)

	sum := 0
	for _, match := range matches {
		// Extract the integers from the match
		sum += extractAndCompute(match)
	}
	fmt.Println("Part 1:", sum)

	// Part 2
	newPattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	re = regexp.MustCompile(newPattern)
	matches = re.FindAllString(in, -1)

	shouldCompute := true
	sum = 0
	for _, match := range matches {
		// if the match is "don't()", keep skipping matches until we find "do()"
		if match == "don't()" {
			shouldCompute = false
			continue
		}

		// if the match is "do()", start computing again
		if match == "do()" {
			shouldCompute = true
			continue
		}

		// if we should compute, extract the integers from the match
		if shouldCompute {
			sum += extractAndCompute(match)
		}
	}
	fmt.Println("Part 2:", sum)
}
