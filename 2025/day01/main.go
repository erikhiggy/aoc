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

func part1(input []string) int {
	curr := 50
	zeros := 0
	for _, coord := range input {
		dir := string(coord[0])
		amt, err := strconv.Atoi(string(coord[1:]))
		if err != nil {
			log.Fatal("error processing coordinate", err)
		}
		switch dir {
		case "L":
			for i := amt; i > 0; i-- {
				curr -= 1
				if curr == -1 {
					curr = 99
				}
			}
		case "R":
			for i := amt; i > 0; i-- {
				curr += 1
				if curr == 100 {
					curr = 0
				}
			}
		}
		if curr == 0 {
			zeros += 1
		}
	}

	return zeros
}

func part2(input []string) int {
	curr := 50
	zeros := 0
	for _, coord := range input {
		dir := string(coord[0])
		amt, err := strconv.Atoi(string(coord[1:]))
		if err != nil {
			log.Fatal("error processing coordinate", err)
		}
		switch dir {
		case "L":
			for i := amt; i > 0; i-- {
				curr -= 1
				if curr == 0 {
					zeros += 1
				}
				if curr == -1 {
					curr = 99
				}
			}
		case "R":
			for i := amt; i > 0; i-- {
				curr += 1
				if curr == 100 {
					curr = 0
					zeros += 1
				}
			}
		}
	}

	return zeros
}

func main() {
	parsed := parseInput("input.txt")
	part1Sol := part1(parsed)
	part2Sol := part2(parsed)

	fmt.Println("Part 1: ", part1Sol)
	fmt.Println("Part 2: ", part2Sol)
}
