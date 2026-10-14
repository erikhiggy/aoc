package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return grid
}

type Node struct {
	row int
	col int
	val string
}

type Queue[T any] struct {
	items []T
}

// Enqueue adds an element to the tail
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the head
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	item := q.items[0]

	// Prevent memory leak: erase reference before slicing
	var zero T
	q.items[0] = zero

	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func traverseBreadth(root Node, grid [][]string) ([]Node, map[Node]struct{}) {
	queue := Queue[Node]{}
	var path []Node
	splittersHit := make(map[Node]struct{})
	visited := make(map[Node]bool)

	queue.Enqueue(root)

	// while the stack still has nodes to process
	for queue.Size() > 0 {
		// pop the top of stack
		node, _ := queue.Dequeue()

		// if we have already seen it, skip it
		_, seen := visited[node]
		if seen {
			continue
		}

		v := node.val
		if v == "." {
			v = "|"
		}
		path = append(path, Node{row: node.row, col: node.col, val: v})
		visited[node] = true

		// if we hit a splitter
		if node.val == "^" {
			_, exists := splittersHit[node]
			if !exists {
				splittersHit[node] = struct{}{}
			}
			// push left and right neighbors
			left := Node{row: node.row, col: node.col - 1, val: grid[node.row][node.col-1]}
			right := Node{row: node.row, col: node.col + 1, val: grid[node.row][node.col+1]}
			queue.Enqueue(left)
			queue.Enqueue(right)
		} else if node.row != len(grid)-1 {
			down := Node{row: node.row + 1, col: node.col, val: grid[node.row+1][node.col]}
			queue.Enqueue(down)
		}
	}

	return path, splittersHit
}

func part1(grid [][]string) ([]Node, map[Node]struct{}) {
	// find the root, S
	i, j := find(grid, "S")
	root := Node{
		row: i,
		col: j,
		val: grid[i][j],
	}

	return traverseBreadth(root, grid)
}

func part2(grid [][]string) int {
	beams := make([]int, len(grid[0])+2)
	for _, s := range grid {
		for i, r := range s {
			if r == "S" {
				beams[i]++
			} else if r == "^" && beams[i] > 0 {
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				beams[i] = 0
			}
		}
	}
	tl := 0
	for _, beam := range beams {
		tl += beam
	}

	return tl
}

func main() {
	grid := parseInput("input.txt")
	_, splittersHit := part1(grid)
	timelines := part2(grid)

	fmt.Printf("part 1: %v\n", len(splittersHit))
	fmt.Printf("part 2: %v\n", timelines)
}

func print2d(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func find(grid [][]string, target string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}
