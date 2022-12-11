package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	priorityMap := make(map[rune]int, 52)
	for i, letter := range letters {
		priorityMap[letter] = i + 1
	}
	input := readInput()

	// sum := 0
	// part 1
	// for _, rucksack := range input {
	// 	priority := priorityMap[findDupePart1(rucksack)]
	// 	sum += priority
	// }

	// part 2
	sum := 0
	rucksackGroups := make([][]string, len(input) / 3)
	idx := 0
	for i, rucksack := range input {
		if i != 0 && i % 3 == 0 {
			idx += 1
		}
		rucksackGroups[idx] = append(rucksackGroups[idx], rucksack)
	}

	for _, rucksackGroup := range rucksackGroups {
		priority := priorityMap[findDupePart2(rucksackGroup)]
		sum += priority
	}

	fmt.Println(sum)
}

func findDupePart2(rucksacks []string) rune {
	firstMap := make(map[rune]int)
	secondMap := make(map[rune]int)
	var letter rune
	for _, r := range rucksacks[0] {
		if firstMap[r] == 0 {
			firstMap[r] = 1
		}
	}

	for _, r := range rucksacks[1] {
		if secondMap[r] == 0 {
			secondMap[r] = 1
		}
	}

	for _, r := range rucksacks[2] {
		if firstMap[rune(r)] == 1 && secondMap[rune(r)] == 1 {
			letter = rune(r)
			break
		}
	}
	return letter
}

func findDupePart1(rucksack string) rune {
	mappy := make(map[rune]int)
	var letter rune
	for i, r := range rucksack {
		if i == len(rucksack) / 2 {
			break
		}

		if mappy[r] == 0 {
			mappy[r] = 1
		}
	}

	for r := len(rucksack) / 2; r < len(rucksack); r++ {
		if mappy[rune(rucksack[r])] == 1 {
			letter = rune(rucksack[r])
			break
		}
	}

	return letter
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}