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
	g.Print()
	fmt.Println()
	fmt.Println()
	total := 0
	for i := 0; i < 10; i++ {
		total += g.Step()
	}
	fmt.Printf("total: %v\n", total)
}

func (grid Grid) Step() int {
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
			if (grid[row][col] == 9) {
				grid[row][col] = 0
				flashed[row][col] = 1
				flashes++
			}

			// this is the normal number case
			if flashed[row][col] != 1 {
				grid[row][col]++
				continue
			}
			// if we hit a 0 (flash!),
			// handle adjacent
			if grid[row][col] == 0  {
				for _, adj := range grid.Adj(Location{row, col}) {
					// if an adj is 0 then don't increment because it's already flashed
					if grid[adj.Row][adj.Col] == 0 {
						// fmt.Println(adj)
						// fmt.Printf("grid[adj.Row][adj.Col]: %d flashed\n", grid[adj.Row][adj.Col])
						continue
					}
					if grid[adj.Row][adj.Col] == 9 {
						grid[adj.Row][adj.Col] = 0
						flashed[adj.Row][adj.Col] = 1
						flashes++
					} else if grid[adj.Row][adj.Col] != flashed[adj.Row][adj.Col] {
						grid[adj.Row][adj.Col]++
					}
				}
			}
		}
	}

	grid.Print()
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
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col-1})
	}

	if loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row-1, loc.Col})
	}

	if loc.Col < len(g[loc.Row])-1 {
		adjPoints = append(adjPoints, Location{loc.Row, loc.Col+1})
	}

	if loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row+1, loc.Col})
	}

	// diag bottom left
	if loc.Col > 0 && loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row+1, loc.Col-1})
	}

	// diag bottom right
	if loc.Col < len(g[loc.Row])-1 && loc.Row < len(g)-1 {
		adjPoints = append(adjPoints, Location{loc.Row+1, loc.Col+1})
	}

	// diag top left
	if loc.Col > 0 && loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row-1, loc.Col-1})
	}

	// diag top right
	if loc.Col < len(g[loc.Row])-1 && loc.Row > 0 {
		adjPoints = append(adjPoints, Location{loc.Row-1, loc.Col+1})
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