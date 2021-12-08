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
}

func part1(lines []Line) int {
	// filter for only horizontal/vertical lines
	filteredIn := Filter(lines, func (l Line) bool {
		return l.isHoriz() || l.isVert()
	})

	// initially mark all spots as not marked
	var marked [size][size]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			marked[i][j] = 0
		}
	}

	m := markLine(filteredIn, marked)

	sum := find2s(m)
	return sum
}

func markLine(filteredIn []Line, g [size][size]int) [size][size]int {
	for _, l := range filteredIn {
		if l.isHoriz() {
			// figure out the starting and ending points
			startX, endX := l.x1, l.x2
			if l.x2 < startX {
				startX, endX = l.x2, l.x1
			}
			for x := startX; x <= endX; x++ {
				g[x][l.y1]++
			}
		}

		if l.isVert() {
			startY, endY := l.y1, l.y2
			if l.y2 < startY {
				startY, endY = l.y2, l.y1
			}

			for y := startY; y <= endY; y++ {
				g[l.x1][y]++
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