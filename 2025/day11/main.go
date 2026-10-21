package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	idx := len(s.items) - 1
	top := s.items[idx]
	s.items[idx] = *new(T)
	s.items = s.items[:idx]

	return top, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// moves through data returning the paths
func traverse(adjList map[string][]string, start string) [][]string {
	var stack Stack[string]
	var paths [][]string
	stack.Push(start)
	var path []string
	for !stack.IsEmpty() {
		top, _ := stack.Pop()
		path = append(path, top)
		if top == "out" {
			paths = append(paths, path)
			continue
		}
		for _, l := range adjList[top] {
			stack.Push(l)
		}
	}
	return paths
}

// allPaths returns every path from node to "out" using DFS with memoization.
// memo caches, for each node, all paths from that node to "out".
func allPaths(node string, adj map[string][]string, memo map[string][][]string) [][]string {
	if paths, ok := memo[node]; ok {
		return paths
	}

	if node == "out" {
		memo[node] = [][]string{{"out"}}
		return memo[node]
	}

	var paths [][]string
	for _, neighbor := range adj[node] {
		for _, sub := range allPaths(neighbor, adj, memo) {
			path := make([]string, 0, len(sub)+1)
			path = append(path, node)
			path = append(path, sub...)
			paths = append(paths, path)
		}
	}

	memo[node] = paths
	return paths
}

type state struct {
	node     string
	fft, dac bool
}

// countPaths returns the number of paths from node to "out" that visit
// both "fft" and "dac" (in any order), given whether they were already seen.
func countPaths(node string, fft, dac bool, adj map[string][]string, memo map[state]int) int {
	if node == "fft" {
		fft = true
	}
	if node == "dac" {
		dac = true
	}

	if node == "out" {
		if fft && dac {
			return 1
		}
		return 0
	}

	key := state{node, fft, dac}
	if val, ok := memo[key]; ok {
		return val
	}

	total := 0
	for _, neighbor := range adj[node] {
		total += countPaths(neighbor, fft, dac, adj, memo)
	}

	memo[key] = total
	return total
}

func main() {
	lines := parseInput("input.txt")

	adjList := make(map[string][]string)
	for _, l := range lines {
		colonSplit := strings.Split(l, ":")
		vertex := colonSplit[0]
		edgeSplit := strings.Split(colonSplit[1], " ")
		edgeSplit = edgeSplit[1:]
		adjList[vertex] = edgeSplit
	}

	memo := make(map[string][][]string)
	part1Paths := allPaths("you", adjList, memo)
	part2Memo := make(map[state]int)
	part2 := countPaths("svr", false, false, adjList, part2Memo)

	part1 := len(part1Paths)
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
