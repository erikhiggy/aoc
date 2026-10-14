package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return grid
}

func calc(nums []int, op string) int {
	total := 0
	for i, n := range nums {
		if op == "*" {
			if i == 0 {
				total = 1
			}
			total *= n
		} else {
			total += n
		}
	}
	return total
}

func part1(grid [][]string) int {
	// first, clean input of spaces
	var cleaned [][]string
	for i := range grid {
		var tmp []string
		for j := range grid[i] {
			if grid[i][j] != "" {
				tmp = append(tmp, grid[i][j])
			}
		}
		cleaned = append(cleaned, tmp)
	}

	rows := len(cleaned)
	cols := len(cleaned[0])

	sum := 0
	for j := range cols {
		var eq []int
		for i := 0; i < rows-1; i++ {
			eq = append(eq, mustInt(cleaned[i][j]))
		}
		sum += calc(eq, cleaned[rows-1][j])
	}
	return sum
}

func part2(grid [][]string) int {
	rows := len(grid)
	cols := len(grid[0])

	sum := 0
	for j := range cols {
		var eq []int
		for i := 0; i < rows-1; i++ {
			if grid[i][j] != "" {
				eq = append(eq, mustInt(grid[i][j]))
			}
		}
		sum += calc(eq, grid[rows-1][j])
	}
	return sum
}

func main() {
	grid := parseInput("input.txt")

	fmt.Println(grid[1][0])

	// part1Sum := part1(grid)
	// fmt.Println("part 1:", part1Sum)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
