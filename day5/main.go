package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := readInput()
	split := fmtInput(in)
	fmt.Println(split)
	// for _, l := range split {
	// 	fmt.Printf("%q\n", l[0])
	// 	// x1 := string(l[0][0])
	// 	// x2 := string(l[0][1])
	// 	// y1 := string(l[1][0])
	// 	// y2 := string(l[1][1])
	// 	// fmt.Printf("Comparing x1: %s with x2: %s -> %d\n", x1, x2, strings.Compare(string(l[0][0]), string(l[1][0])))
	// 	// fmt.Printf("Comparing y1: %s with y2: %s -> %d\n", y1, y2, strings.Compare(string(l[0][1]), string(l[1][1])))
	// }
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func fmtInput(lines []string) [][]string {
	var list [][]string
	for _, l := range lines {
		splitter := strings.Split(l, " -> ")
		x1y1 := strings.Split(splitter[0], ",")
		fmt.Printf("%q\n", x1y1)
		x2y2 := strings.Split(splitter[1], ",")
		fmt.Printf("%q\n", x2y2)
		list = append(list, splitter)
	}
	return list
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}