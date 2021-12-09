package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	One   int = 2
	Four      = 4
	Seven     = 3
	Eight     = 7
)

func main() {
	lines := readInput()
	v := makeArr(lines)
	fmt.Printf("part1: %d\n", part1(v))
	// fmt.Printf("part2: %d\n", part2(v))
}

func part1(arr [][]string) int {
	sum := 0
	for _, line := range arr {
		for i := len(line) - 1; i >= len(line)-4; i-- {
			digit := line[i]
			if isUnique(digit) {
				sum++
			}
		}

	}
	return sum
}

// TODO: Figure this out
// func part2(arr [][]string) int {
// 	sum := 0
// 	digitMap := map[string]string{
// 		"abcdeg": "0",
// 		"acdfg":  "2",
// 		"abcdf":  "3",
// 		"bcdef":  "5",
// 		"bcdefg": "6",
// 		"abcdef": "9",
// 	}

// 	for _, line := range arr {
// 		s := ""
// 		for i := len(line) - 1; i >= len(line)-4; i-- {
// 			nStr := ""
// 			switch len(line[i]) {
// 			case One:
// 				nStr = "1"
// 			case Four:
// 				nStr = "4"
// 			case Seven:
// 				nStr = "7"
// 			case Eight:
// 				nStr = "8"
// 			}

// 			s += digitMap[sorted]
// 		}
// 		// count the running sum
// 		// num, err := strconv.Atoi(s)
// 		// checkErr(err)
// 		// sum += num
// 	}

// 	return sum
// }

func isUnique(s string) bool {
	return len(s) == One || len(s) == Four || len(s) == Seven || len(s) == Eight
}

func makeArr(lines []string) [][]string {
	var v [][]string
	for _, l := range lines {
		s := strings.Split(l, " ")
		v = append(v, s)
	}
	return v
}

func readInput() []string {
	file, err := os.Open("input.txt")
	checkErr(err)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
