package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	displays := readInput()
	fmt.Printf("part1: %d\n", part1(displays))
	// fmt.Printf("part2: %d\n", part2(v))
}

func part1(displays []Display) int {
	sum := 0
	for _, display := range displays {
		for _, output := range display.output {
			switch len(output) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

// func part2(grid [][]string) int {
// 	return 0
// }

type Display struct {
	patterns []string
	output   []string
}

func readInput() []Display {
	file, err := os.Open("input.txt")
	checkErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var displays []Display
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		display := Display{
			patterns: strings.Fields(line[0]),
			output: strings.Fields(line[1]),
		}
		displays = append(displays, display)
	}

	return displays
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
