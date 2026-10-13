package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func removeRolls(grid [][]string, removed int, pt2 bool) int {
	if removed == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	// up, up-left, down, down-left, left, right, bottom-right, up-right offsets
	dRows := []int{-1, -1, 1, 1, 0, -1, 0, 1}
	dCols := []int{0, -1, 0, -1, -1, 1, 1, 1}

	resGrid := make([][]string, rows)
	xCounter := 0
	for r := range rows {
		resGrid[r] = make([]string, cols)
		for c := range cols {
			if grid[r][c] == "@" {
				// fmt.Printf("'@' found at [%d, %d]\n", r, c)

				wallCounter := 0
				for i := range 8 {
					nextR := r + dRows[i]
					nextC := c + dCols[i]

					if nextR >= 0 && nextR < rows && nextC >= 0 && nextC < cols {
						if grid[nextR][nextC] == "@" {
							// fmt.Printf("Neighbor '@' found at [%d,%d]\n", nextR, nextC)
							wallCounter += 1
						}
					}
				}
				if wallCounter < 4 {
					xCounter += 1
					resGrid[r][c] = "x"
				} else {
					resGrid[r][c] = "@"
				}
			} else {
				resGrid[r][c] = "."
			}
		}
	}

	if pt2 == false {
		return xCounter
	}

	return removed + removeRolls(resGrid, xCounter, true)
}

func main() {
	grid := parseInput("input.txt")

	part1Sum := removeRolls(grid, -1, false)
	part2Sum := removeRolls(grid, -1, true) + 1
	fmt.Println("part 1: ", part1Sum)
	fmt.Println("part 2: ", part2Sum)
}

func print(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
