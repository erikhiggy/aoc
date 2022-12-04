package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := readInput()
	adjList := makeAdjList(in)
	fmt.Printf("adjList: %v\n", adjList)
	part1(adjList)
}

func part1(adjList AdjacencyList) int {
	start := adjList["start"]
	var stack []string
	stack = append(stack, start[0])

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		// we are at the end, pop off and append to the list
		if adjList[top][0] == "end" {
			stack = append(stack, adjList[top][0])
			break
		}

		stack = append(stack, adjList[top][0])
	}

	var path []string
	path = append(path, "start")
	for _, s := range stack {
		path = append(path, s)
	}

	fmt.Printf("path: %v\n", path)


	return 0
}

func makeAdjList(lines []string) AdjacencyList {
	adjList := make(AdjacencyList)
	for _, line := range lines {
		split := strings.Split(line, "-")
		first := strings.Fields(split[0])[0]
		second := strings.Fields(split[1])[0]

		// if start is second, swap
		if second == "start" {
			temp := second
			second = first
			first = temp
		}
		if first == "end" {
			temp := first
			first = second
			second = temp
		}
		adjList[first] = append(adjList[first], second)
		if first != "start" && second != "end" {
			adjList[second] = append(adjList[second], first)
		}
	}
	return adjList
}

type AdjacencyList map[string][]string

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