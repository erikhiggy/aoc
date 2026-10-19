package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return lines
}

type Point struct {
	x, y int
}

type Pair struct {
	area int
	a, b Point
}

func computePairs(points []Point) []Pair {
	var pairs []Pair
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, Pair{area: computeArea(points[i], points[j]), a: points[i], b: points[j]})
		}
	}
	return pairs
}

func main() {
	lines := parseInput("input.txt")

	points := make([]Point, len(lines))
	for i := range lines {
		spl := strings.Split(lines[i], ",")
		x := mustInt(spl[0])
		y := mustInt(spl[1])
		points[i] = Point{x, y}
	}

	pairs := computePairs(points)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area < pairs[j].area
	})

	fmt.Println("part 1:", pairs[len(pairs)-1].area)
}

func computeArea(p1, p2 Point) int {
	return int((math.Abs(float64(p1.x-p2.x)) + 1) * (math.Abs(float64(p1.y-p2.y)) + 1))
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
