package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	matrix := readInput()
	part1(matrix)
	product := part2(matrix)
	fmt.Printf("part2: %d\n", product)
}

type Location struct {
	Row int
	Col int
}

func part1(m Matrix) {
	risk := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m.isLowPoint(Location{i, j}) {
				risk += m[i][j] + 1
			}
		}
	}
	fmt.Printf("part1: %d\n", risk)
}

func part2(m Matrix) int {
	var stack []Location
	lowPointCounter := 1
	ids := make([][]int, len(m))
	for i := range m {
		ids[i] = make([]int, len(m[i]))
	}

	basinSize := make(map[int]int)
	
	// mark each initial low point
  // and keep a map of the low points
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m.isLowPoint(Location{i, j}) {
				ids[i][j] = lowPointCounter
				basinSize[lowPointCounter] = 1
				lowPointCounter++
				stack = append(stack, Location{i, j})
			}
		}
	}

	for len(stack) > 0 {
		// grab the top of the stack
		loc := stack[len(stack)-1]
		// pop the element off the stack
		stack = stack[:len(stack)-1]
		id := ids[loc.Row][loc.Col]
		if id == 0 {
			panic("something went terribly wrong! id should not be zero!")
		}

		// BFS all surrounding
		for _, adj := range m.Adj(loc) {
			// check for 9s and if it's not marked
			if m[adj.Row][adj.Col] != 9 && ids[adj.Row][adj.Col] == 0 {
				ids[adj.Row][adj.Col] = id
				basinSize[id]++
				stack = append(stack, adj)
			}
		}
	}

	var basinSizes []int
	for _, size := range basinSize {
		basinSizes = append(basinSizes, size)
	}

	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func findLowPoints(m Matrix) []Location {
	var lowPoints []Location
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m.isLowPoint(Location{i, j}) {
				lowPoints = append(lowPoints, Location{i, j})
			}
		}
	}

	return lowPoints
}

func (m Matrix) isLowPoint(loc Location) bool {
	for _, a := range m.Adj(loc) {
		if m[a.Row][a.Col] <= m[loc.Row][loc.Col] {
			return false
		}
	}
	return true
}

func (m Matrix) Adj(loc Location) []Location {
	var adjPoints []Location

	if loc.Col > 0 {
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col-1})
	}

	if loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row-1, loc.Col})
	}

	if loc.Col < len(m[loc.Row])-1 {
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col+1})
	}

	if loc.Row < len(m)-1 {
		adjPoints = append(adjPoints, Location{loc.Row+1, loc.Col})
	}

	return adjPoints
}

type Matrix [][]int

func readInput() Matrix {
	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var m Matrix
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "")
		nums := make([]int, len(vals))
		for i, v := range vals {
			nums[i], err = strconv.Atoi(v)
			checkErr(err)
		}
		m = append(m, nums)
	}
	return m
}

func Print(m Matrix) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}