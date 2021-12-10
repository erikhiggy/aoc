package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := readInput()
	part1Total := part1(lines)
	part2Total := part2(lines)
	fmt.Printf("part1: %d\n", part1Total)
	fmt.Printf("part2: %d\n", part2Total)
}

func part1(lines []string) int {
	totalSum := 0
	// for each line, check for corrupt
	for _, line := range lines {
		totalSum += getIllegalSum(line, false)
	}

	return totalSum
}

func part2(lines []string) int {
	var sums []int
	for _, line := range lines {
		sum := 0
		s := getIllegalSum(line, false)
		if s == 0 { // incomplete line
			// do autocomplete step
			sum += getIllegalSum(line, true)
			sums = append(sums, sum)
		}
	}

	// sort and return the middle sum
	sort.Ints(sums)
	middleIdx := len(sums) / 2

	return sums[middleIdx]
}

// getIllegalSum should return the sum of the illegal characters
// a return of 0 denotes that the line was not corrupt
func getIllegalSum(exp string, part2 bool) int {
	var stack Stack
	sum := 0
	illegalRuneTable := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, e := range exp {
		// if the expression is an opening bracket, append it to the stack
		if e == rune('(') || e == rune('[') || e == rune('{') || e == rune('<') {
			stack = append(stack, e)
			continue
		}

		if stack.isEmpty() {
			return 0
		}

		var check rune
		switch e {
		case ')':
			check = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if check == '{' || check == '[' || check == '<' {
				// illegal character, add to sum based on table
				sum += illegalRuneTable[e]
				return sum
			}
			break
		case ']':
			check = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if check == '{' || check == '(' || check == '<' {
				// illegal character, add to sum based on table
				sum += illegalRuneTable[e]
				return sum
			}
			break
		case '}':
			check = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if check == '(' || check == '[' || check == '<' {
				// illegal character, add to sum based on table
				sum += illegalRuneTable[e]
				return sum
			}
			break
		case '>':
			check = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if check == '(' || check == '[' || check == '{' {
				// illegal character, add to sum based on table
				sum += illegalRuneTable[e]
				return sum
			}
			break
		}
	}

	// for part 2, we want to look sum up all the remaining symbols
	if part2 {
		total := 0
		symMap := map[rune]rune{
			'(': ')',
			'[': ']',
			'{': '}',
			'<': '>',
		}
		part2Map := map[rune]int{
			')': 1,
			']': 2,
			'}': 3,
			'>': 4,
		}
		multiplier := 5
		for i := len(stack) - 1; i >= 0; i-- {
			// multiply the total score by the multiplier, 5
			total *= multiplier
			closingSym := symMap[stack[i]]
			total += part2Map[closingSym]
		}
		return total
	}

	return sum
}

type Stack []rune

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func readInput() []string {
	file, err := os.Open("input.txt")
	checkErr(err)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
