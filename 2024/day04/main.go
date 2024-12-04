package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // down-right
	{1, -1},  // down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}

var xmasPatterns = [][5][2]int{
	{{0, 0}, {1, -1}, {1, 1}, {2, -2}, {2, 2}},     // X-MAS pattern
	{{0, 0}, {-1, -1}, {-1, 1}, {-2, -2}, {-2, 2}}, // Inverted X-MAS pattern
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func searchWord(grid [][]rune, word string, x, y int) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		found := true
		for k := 0; k < wordLen; k++ {
			nx, ny := x+k*dx, y+k*dy
			if !isValid(nx, ny, rows, cols) || grid[nx][ny] != rune(word[k]) {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}
	return count
}

func countOccurrences(grid [][]rune, word string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			count += searchWord(grid, word, i, j)
		}
	}
	return count
}

func searchXMAS(grid [][]rune, x, y int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for _, pattern := range xmasPatterns {
		found := true
		for _, p := range pattern {
			nx, ny := x+p[0], y+p[1]
			if !isValid(nx, ny, rows, cols) || grid[nx][ny] != 'M' && grid[nx][ny] != 'A' && grid[nx][ny] != 'S' {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func countXMASOccurrences(grid [][]rune) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if searchXMAS(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("2024/day04/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	word := "XMAS"
	countPart1 := countOccurrences(grid, word)
	fmt.Println("Number of occurrences of XMAS:", countPart1)

	//countPart2 := countXMASOccurrences(grid)
	//fmt.Println("Number of occurrences of X-MAS:", countPart2)
}
