package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput() []string {
	var res []string
	file, err := os.Open("2024/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return res
}

func toMatrix(in []string) [][]string {
	var res [][]string
	for _, line := range in {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		res = append(res, row)
	}
	return res
}

func main() {
	in := parseInput()
	grid := toMatrix(in)
	// Part 1
	var startRow, startCol int
	for i, row := range grid {
		for j, char := range row {
			if char == "^" {
				startRow = i
				startCol = j
				break
			}
		}
	}
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	direction := 0 // Start facing up
	locations := make(map[[2]int]bool)

	row, col := startRow, startCol
	for {
		// Move in the current direction
		newRow := row + directions[direction][0]
		newCol := col + directions[direction][1]

		// Check if we are out of bounds
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
			break
		}

		// Check if we hit a #
		if grid[newRow][newCol] == "#" {
			// Rotate 90 degrees clockwise
			direction = (direction + 1) % 4
		} else {
			// Move to the new position
			row = newRow
			col = newCol
			// save locations
			locations[[2]int{row, col}] = true
		}
	}

	fmt.Println("Part 1:", len(locations))
}
