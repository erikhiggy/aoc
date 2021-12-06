package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, boards := readInput()
	part1(nums, boards)
	part2(nums, boards)
}

type boardNum struct {
	num int
	visited bool
}

type board struct {
	nums [5][5]boardNum
}

func part1(numsDrawn []int, boards []*board) {
	for _, numDrawn := range numsDrawn {
		for _, b := range boards {
			b.visit(numDrawn)
			if b.bingoed() {
				fmt.Printf("bingo! ==> %d", b.unvisitedSum()*numDrawn)
				return
			}
		}
	}
}

func part2(numsDrawn []int, boards []*board) {
	for _, n := range numsDrawn {
		var havntWon []*board
		for _, b := range boards {
			b.visit(n)
			if !b.bingoed() {
				havntWon = append(havntWon, b)
			} else if len(boards) == 1 {
				fmt.Println("found the last winning board! ==> %d", b.unvisitedSum()*n)
				return
			}
		}
		boards = havntWon
	}
}

func (b *board) unvisitedSum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.nums[i][j].visited {
				sum += b.nums[i][j].num
			}
		}
	}

	return sum
}

func (b *board) visit(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.nums[i][j].num == n {
				b.nums[i][j].visited = true
			}
		}
	}
}

func (b *board) bingoed() bool {
	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; j < 5; j++ {
			if !b.nums[i][j].visited {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; j < 5; j++ {
			if !b.nums[j][i].visited {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	return false
}

func makeBoard(lines []string) *board {
	var b board
	for i := 0; i < 5; i++ {
		for j, numAsStr := range strings.Fields(lines[i]) {
			n, err := strconv.Atoi(numAsStr)
			checkErr(err)
			b.nums[i][j] = boardNum{n, false}
		}
	}

	return &b
}

func readInput() ([]int, []*board) {
	file, err := os.Open("input.txt")
	checkErr(err)

	var lines []string
	var numsDrawn []int
	var boards []*board
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for _, n := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(n)
		checkErr(err)
		numsDrawn = append(numsDrawn, num)
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = []string{}
		} else {
			lines = append(lines, scanner.Text())
			if len(lines) == 5 {
				board := makeBoard(lines)
				boards = append(boards, board)
			}
		}
	}
	return numsDrawn, boards
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}