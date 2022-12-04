package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := readInput()
	// totalFlashes := part1(g)
	step := part2(g)
	// fmt.Printf("part1: %v\n", totalFlashes)
	fmt.Printf("part2: %v\n", step)
}

func part1(g Grid) int {
	t := 0
	for i := 0; i < 100; i++ {
		t += g.Step(i)
	}
	return t
}

func part2(g Grid) int {
	allFlashed := false
	iter := 0
	for !allFlashed {
		flashes := g.Step(iter)
		if flashes == 100 {
			allFlashed = true
		}
		iter++
	}
	return iter
}

func (grid Grid) Step(i int) int {
	var stack []Location
	flashed := make([][]int, len(grid))
	flashes := 0
	for i := range grid {
		flashed[i] = make([]int, len(grid[i]))
	}
	// go through each energy level
	// and increase the level by 1
	// if the level is already a 9, "flash"
	// and increase the energy level of all
	// adjacent (horiz, vert, diag) levels by one
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// increment the number by 1
			// unless 9 then set to 0
			if grid[row][col] == 9 {
				grid[row][col] = 0
				flashed[row][col] = 1
				flashes++
				// add to stack to process later
				stack = append(stack, Location{row, col})
			}

			// this is the normal number case
			if flashed[row][col] != 1 {
				grid[row][col]++
				continue
			}
			// if we hit a 0 (flash!),
			// handle adjacent
			if grid[row][col] == 0 {
				for len(stack) > 0 {
					// grab the location off of the top of the stack
					loc := stack[len(stack)-1]
					// resize the stack
					stack = stack[:len(stack)-1]

					for _, adj := range grid.Adj(loc) {
						// if an adjacent is already marked as flashed, skip
						if flashed[adj.Row][adj.Col] == 1 {
							continue
						}

						// we have hit another flashing scenario
						if grid[adj.Row][adj.Col] == 9 {
							grid[adj.Row][adj.Col] = 0
							flashed[adj.Row][adj.Col] = 1
							flashes++
							stack = append(stack, Location{adj.Row, adj.Col})
						}

						if flashed[adj.Row][adj.Col] != 1 {
							grid[adj.Row][adj.Col]++
						}
					}
				}
			}
		}
	}

	return flashes
}

func (g Grid) Print() {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			fmt.Print(g[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Adj(loc Location) []Location {
	var adjPoints []Location

	if loc.Col > 0 {
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col - 1})
	}

	if loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row - 1, loc.Col})
	}

	if loc.Col < len(g[loc.Row])-1 {
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col + 1})
	}

	if loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row + 1, loc.Col})
	}

	// diag bottom left
	if loc.Col > 0 && loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row + 1, loc.Col - 1})
	}

	// diag bottom right
	if loc.Col < len(g[loc.Row])-1 && loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row + 1, loc.Col + 1})
	}

	// diag top left
	if loc.Col > 0 && loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row - 1, loc.Col - 1})
	}

	// diag top right
	if loc.Col < len(g[loc.Row])-1 && loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row - 1, loc.Col + 1})
	}

	return adjPoints
}

type Location struct {
	Row int
	Col int
}

type Grid [][]int

func readInput() Grid {
	file, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var m Grid
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
