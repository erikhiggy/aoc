package main

import (
	"bufio"
	"errors"
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

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i])
	}
	return d.parent[i]
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
		return true
	}
	return false
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

type JunctionBox struct {
	x, y, z int
}

type Connection struct {
	distance float64
	src, dst int
}

// Euclidian distance between points in 3D space
func euclideanDistance(p1, p2 JunctionBox) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func computeDistances(boxes []JunctionBox) []Connection {
	var connections []Connection
	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			connections = append(connections, Connection{distance: euclideanDistance(boxes[i], boxes[j]), src: i, dst: j})
		}
	}
	sort.Slice(connections, func(a, b int) bool {
		return connections[a].distance < connections[b].distance
	})
	return connections
}

func part1(conns []Connection) int {
	n := len(conns)
	adjList := make(map[int][]int)

	// make 1000 shortest connections
	for i := range 1000 {
		src, dst := conns[i].src, conns[i].dst
		adjList[src] = append(adjList[src], dst)
		adjList[dst] = append(adjList[dst], src)
	}

	visited := make([]bool, n)
	var val []int
	for i := range n {
		if visited[i] != false {
			continue
		}
		visited[i] = true
		var queue Queue[int]
		queue.Enqueue(i)
		count := 1
		for queue.Size() > 0 {
			tail, _ := queue.Dequeue()
			for _, t := range adjList[tail] {
				if visited[t] {
					continue
				}
				visited[t] = true
				queue.Enqueue(t)
				count += 1
			}
		}
		val = append(val, count)
	}
	res := 1
	sort.Ints(val)
	for i := len(val) - 1; i >= len(val)-3; i-- {
		res *= val[i]
	}
	return res
}

func part2(conns []Connection, boxes []JunctionBox) int {
	n := len(boxes)
	dsu := NewDSU(n)

	tmp := n
	res := 0
	for _, c := range conns {
		if dsu.Union(c.src, c.dst) {
			tmp--
			if tmp == 1 {
				x1 := boxes[c.src].x
				x2 := boxes[c.dst].x
				res = x1 * x2
				break
			}
		}
	}
	return res
}

func main() {
	lines := parseInput("input.txt")

	var boxes []JunctionBox
	for _, l := range lines {
		spl := strings.Split(l, ",")
		boxes = append(boxes, JunctionBox{x: mustInt(spl[0]), y: mustInt(spl[1]), z: mustInt(spl[2])})
	}

	connections := computeDistances(boxes)
	part1 := part1(connections)
	part2 := part2(connections, boxes)
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
