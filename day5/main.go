package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const size = 1000

func main() {
	in := readInput()
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(lines []Line) int {
	// initially mark all spots as not marked
	marked := initGrid()
	m := markLine(lines, marked)
	sum := find2s(m)
	return sum
}

func part2(lines []Line) int {
	return part1(lines)
}

func initGrid() [size][size]int {
	var marked [size][size]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			marked[i][j] = 0
		}
	}
	return marked
}

func markLine(lines []Line, g [size][size]int) [size][size]int {
	for _, l := range lines {
		if l.isHoriz() {
			// figure out the starting and ending points
			startX, endX := l.x1, l.x2
			if l.x2 < startX {
				startX, endX = l.x2, l.x1
			}
			for x := startX; x <= endX; x++ {
				g[x][l.y1]++
			}
		} else if l.isVert() {
			startY, endY := l.y1, l.y2
			if l.y2 < startY {
				startY, endY = l.y2, l.y1
			}

			for y := startY; y <= endY; y++ {
				g[l.x1][y]++
			}
		} else {
			// case for the diagonals
			startX, startY := l.x1, l.y1
			endX, endY := l.x2, l.y2
			if l.x2 < startX {
				startX, startY = l.x2, l.y2
				endX, endY = l.x1, l.y1
			}

			y := startY
			for x := startX; x <= endX; x++ {
				g[x][y]++
				if y < endY {
					y++
				} else {
					y--
				}
			}
		}

}
return g
}

func Filter(vs []Line, f func(Line) bool) []Line {
	vsf := make([]Line, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

type Line struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func readInput() []Line {
	var lines []Line
	numRe := regexp.MustCompile("\\d+")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		points := numRe.FindAllString(scanner.Text(), 4)
		line := Line{
			x1: strToInt(points[0]),
			y1: strToInt(points[1]),
			x2: strToInt(points[2]),
			y2: strToInt(points[3]),
		}
		lines = append(lines, line)
	}

	return lines
}

func (l Line) isHoriz() bool {
	return l.y1 == l.y2
}

func (l Line) isVert() bool {
	return l.x1 == l.x2
}

func strToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func Print(m [size][size]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			fmt.Print(m[j][i])
		}
		fmt.Print("\n")
	}
}

func find2s(m [size][size]int) int {
	sum := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] >= 2 {
				sum++
			}
		}
	}
	return sum
}